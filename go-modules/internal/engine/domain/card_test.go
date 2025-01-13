package domain

import (
	"testing"

	"github.com/marcodali/forbidden-memories-duel-online/pkg/models"
)

func TestLoadCards(t *testing.T) {
	data := `
- id: 1
  name: "Test Card"
  description: "A test card"
  baseAttack: 1000
  baseDefense: 1000
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

	card := registry.GetCard(1)
	if card == nil {
		t.Fatal("Expected card to be loaded, but got nil")
	}
	if card.Name != "Test Card" {
		t.Errorf("Expected card name to be 'Test Card', got %s", card.Name)
	}
}

func TestGetCard(t *testing.T) {
	registry := models.GetCardRegistry()
	card := registry.GetCard(1)
	if card == nil {
		t.Fatal("Expected to find card with ID 1, but got nil")
	}
	if card.ID != 1 {
		t.Errorf("Expected card ID to be 1, got %d", card.ID)
	}
}

func TestNewCardInstance(t *testing.T) {
	instance := models.NewCardInstance(1)
	if instance == nil {
		t.Fatal("Expected card instance to be created, but got nil")
	}
	if instance.CurrentAttack != 1000 {
		t.Errorf("Expected CurrentAttack to be 1000, got %d", instance.CurrentAttack)
	}
	if instance.CurrentDefense != 1000 {
		t.Errorf("Expected CurrentDefense to be 1000, got %d", instance.CurrentDefense)
	}
}

func TestSingletonCardRegistry(t *testing.T) {
	registry1 := models.GetCardRegistry()
	registry2 := models.GetCardRegistry()
	if registry1 != registry2 {
		t.Error("Expected both registry instances to be the same")
	}
}

func TestCardInstanceAttributes(t *testing.T) {
	instance := models.NewCardInstance(1)
	if instance == nil {
		t.Fatal("Expected card instance to be created, but got nil")
	}
	if instance.IsInAttackMode {
		t.Error("Expected IsInAttackMode to be false by default")
	}
}

func TestCardInstanceStateChange(t *testing.T) {
	instance := models.NewCardInstance(1)
	if instance == nil {
		t.Fatal("Expected card instance to be created, but got nil")
	}

	// Simulate changing state
	instance.IsInAttackMode = true
	if !instance.IsInAttackMode {
		t.Error("Expected IsInAttackMode to be true after change")
	}

	instance.IsInAttackMode = false
	if instance.IsInAttackMode {
		t.Error("Expected IsInAttackMode to be false after change")
	}
}
