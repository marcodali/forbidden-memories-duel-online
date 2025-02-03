package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadCardsfromYAML(t *testing.T) {
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
    validTargetIDs: [33]
    bonus: 700
`
	registry := GetCardRegistry()
	err := registry.LoadCardsfromYAML([]byte(data))
	assert.NoError(t, err)
}

func TestLoadCardsfromYAMLWithInvalidData(t *testing.T) {
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
	err := registry.LoadCardsfromYAML([]byte(invalidData))
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "unexpected error trying to load cards from YAML data")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)
}

func TestGetCard(t *testing.T) {
	registry := GetCardRegistry()
	cardTemplate := registry.GetCard(1)
	assert.NotNil(t, cardTemplate)
	assert.Equal(t, 1, cardTemplate.ID)
}

func TestNewCardInstance(t *testing.T) {
	playableCard, err := NewCardInstance(1)
	assert.NoError(t, err)
	assert.NotNil(t, playableCard)
	assert.Equal(t, 800, playableCard.CurrentAttack)
	assert.Equal(t, 400, playableCard.CurrentDefense)
	assert.False(t, playableCard.IsInAttackMode)

	// Simulate changing state
	playableCard.IsInAttackMode = true
	assert.True(t, playableCard.IsInAttackMode)

	// Test with invalid card ID
	_, err = NewCardInstance(999)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "no card template found for the given ID")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)
}

func TestSingletonCardRegistry(t *testing.T) {
	registryA := GetCardRegistry()
	registryB := GetCardRegistry()
	assert.Equal(t, registryA, registryB)
	assert.EqualValues(t, registryA, registryB)
}

func TestEquipCard(t *testing.T) {
	registry := GetCardRegistry()
	equipCard := registry.GetCard(2) // Malevolent Nuzzler

	assert.NotNil(t, equipCard)
	assert.True(t, equipCard.IsEquip)
	assert.NotNil(t, equipCard.EquipRules)
	assert.Equal(t, 1, len(equipCard.EquipRules.ValidTargetIDs))
	assert.Equal(t, 33, equipCard.EquipRules.ValidTargetIDs[0])
	assert.Equal(t, 700, equipCard.EquipRules.Bonus)
}
