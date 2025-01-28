package models

import (
	"testing"

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
- id: 2
  name: "Malevolent Nuzzler"
  description: "An equip card"
  baseAttack: 0
  baseDefense: 0
  level: 0
  mainType: "Magic"
  subType: "Equip"
  guardianStars: ["None", "None"]
  rarity: "NORMAL"
  isFusion: false
  isFusionMaterial: false
  isMagic: true
  isEquip: true
  isTrap: false
  isRitual: false
  image: "malevolent_nuzzler.png"
  equipRules:
    validTargetIDs: [1]
    bonus: 700
`
	registry := GetCardRegistry()
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
	registry := GetCardRegistry()
	err := registry.LoadCards([]byte(invalidData))
	if err == nil {
		t.Fatal("Expected an error when loading invalid YAML data, but got none")
	}
	assert.Contains(t, err.Error(), "error unmarshalling YAML data")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)
}

func TestGetCard(t *testing.T) {
	registry := GetCardRegistry()
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
	playableCard, err := NewCardInstance(1)
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
	_, err = NewCardInstance(999)
	if err == nil {
		t.Fatal("Expected an error when creating a card instance with invalid ID, but got none")
	}
	assert.Contains(t, err.Error(), "no card template found for the given ID")
}

func TestSingletonCardRegistry(t *testing.T) {
	registryA := GetCardRegistry()
	registryB := GetCardRegistry()
	if registryA != registryB {
		t.Error("Expected both registry instances to be the same because we are using a singleton")
	}
}

func TestEquipCard(t *testing.T) {
	registry := GetCardRegistry()
	equipCard := registry.GetCard(2) // Malevolent Nuzzler

	if equipCard == nil {
		t.Fatal("Expected to find equip card with ID 2, but got nil")
	}

	if !equipCard.IsEquip {
		t.Error("Expected card to be an equip card")
	}

	if equipCard.EquipRules == nil {
		t.Fatal("Expected equip card to have EquipRules")
	}

	if len(equipCard.EquipRules.ValidTargetIDs) == 0 {
		t.Error("Expected equip card to have valid target IDs")
	}

	if equipCard.EquipRules.Bonus != 700 {
		t.Errorf("Expected bonus to be 700, got %d", equipCard.EquipRules.Bonus)
	}
}

func TestCardInstanceStateChange(t *testing.T) {
	playableCard, err := NewCardInstance(1)
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
