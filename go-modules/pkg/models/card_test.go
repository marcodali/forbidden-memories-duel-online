package models_test

import (
	"testing"

	"github.com/marcodali/forbidden-memories-duel-online/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestLoadCards(t *testing.T) {
	data := `
- id: 1
  name: "Test Card"
  description: "A test card"
  baseAttack: 800
  baseDefense: 400
  level: 4
  mainType: "Warrior"
  subType: "Normal"
  guardianStars: ["Mars", "Jupiter"]
  rarity: "NORMAL"
  isFusion: false
  isFusionMaterial: false
  isMagic: false
  isEquip: false
  isTrap: false
  isRitual: false
  image: "test_card.png"
`
	registry := models.GetCardRegistry()
	err := registry.LoadCards([]byte(data))
	if err != nil {
		t.Fatalf("Failed to load cards: %v", err)
	}
}

func TestLoadCardsWithInvalidData(t *testing.T) {
	invalidData := `
- id: 1
  name: "Invalid YAML Card because there is no closing quote here
  description: "A card with invalid YAML format"
  baseAttack: 800
  baseDefense: 400
  level: 4
  mainType: "Warrior"
  subType: "Normal"
  guardianStars: ["Mars", "Jupiter"]
  rarity: "NORMAL"
  isFusion: false
  isFusionMaterial: false
  isMagic: false
  isEquip: false
  isTrap: false
  isRitual: false
  image: "invalid_card.png"
`
	registry := models.GetCardRegistry()
	err := registry.LoadCards([]byte(invalidData))
	if err == nil {
		t.Fatal("Expected an error when loading invalid YAML data, but got none")
	}
	assert.Contains(t, err.Error(), "error unmarshalling YAML data")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)
}

func TestGetCard(t *testing.T) {
	registry := models.GetCardRegistry()
	cardTemplate := registry.GetCard(1)
	if cardTemplate == nil {
		t.Fatal("Expected to find card with ID 1, but got nil")
	}
	if cardTemplate.ID != 1 {
		t.Errorf("Expected card ID to be 1, got %d", cardTemplate.ID)
	}
	if cardTemplate.Name != "Test Card" {
		t.Errorf("Expected card name to be 'Test Card', got %s", cardTemplate.Name)
	}
}

func TestNewCardInstance(t *testing.T) {
	playableCard, err := models.NewCardInstance(1)
	if err != nil {
		t.Fatalf("Failed to create playable card instance because: %v", err)
	}
	if playableCard == nil {
		t.Fatal("Expected playable card instance to be created, but got nil")
	}
	if playableCard.CurrentAttack != 800 {
		t.Errorf("Expected CurrentAttack to be 800, got %d", playableCard.CurrentAttack)
	}
	if playableCard.CurrentDefense != 400 {
		t.Errorf("Expected CurrentDefense to be 400, got %d", playableCard.CurrentDefense)
	}
	if playableCard.IsInAttackMode {
		t.Error("Expected IsInAttackMode to be false by default")
	}

	// Test with invalid card ID
	_, err = models.NewCardInstance(999)
	if err == nil {
		t.Fatal("Expected an error when creating a card instance with invalid ID, but got none")
	}
	assert.Contains(t, err.Error(), "no card template found for the given ID")
}

func TestSingletonCardRegistry(t *testing.T) {
	registryA := models.GetCardRegistry()
	registryB := models.GetCardRegistry()
	if registryA != registryB {
		t.Error("Expected both registry instances to be the same because we are using a singleton")
	}
}

func TestCardInstanceStateChange(t *testing.T) {
	playableCard, err := models.NewCardInstance(1)
	if err != nil {
		t.Fatalf("Failed to create playable card instance because: %v", err)
	}

	// Simulate changing state
	playableCard.IsInAttackMode = true
	if !playableCard.IsInAttackMode {
		t.Error("Expected IsInAttackMode to be true after change")
	}

	playableCard.IsInAttackMode = false
	if playableCard.IsInAttackMode {
		t.Error("Expected IsInAttackMode to be false after change")
	}
}
