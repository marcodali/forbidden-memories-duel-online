package models

import (
	"fmt"
	"path/filepath"
	"runtime"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var singletonForCardModel sync.Once

func initializeCardTestSuite() {
	singletonForCardModel.Do(func() {
		_, filename, _, _ := runtime.Caller(0)
		fmt.Println("This setup code executes only one time for the file", filepath.Base(filename))
		CleanRegistry()
	})
}

func InitializeCardRegistryWithFakeYAMLData() {
	initializeCardTestSuite()
	data := `
- id: 1001
  name: "Test Card"
  description: "A test card"
  baseAttack: 800
  baseDefense: 400
  level: 4
  type: "Warrior"
  guardianStars: ["Mars", "Jupiter"]
  rarity: "NORMAL"
- id: 2002
  name: "Fake Malevolent Nuzzler"
  description: "A fake equip card"
  type: "Equip"
  rarity: "NORMAL"
  equipRules:
    validTargetIDs: [33]
    bonus: 700
`
	GetCardRegistry().LoadCardsfromYAML([]byte(data))
}

func TestLoadCardsfromYAMLWithInvalidData(t *testing.T) {
	initializeCardTestSuite()
	invalidData := `
- id: 5000
  name: "Invalid YAML Card because there is no closing quote here
  description: "A card with invalid YAML format"
  baseAttack: 800
  baseDefense: 400
  level: 4
  type: "Warrior"
  guardianStars: ["Mars", "Jupiter"]
  rarity: "NORMAL"
`
	err := GetCardRegistry().LoadCardsfromYAML([]byte(invalidData))
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected error trying to load cards from YAML data")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)
}

func TestGetCard(t *testing.T) {
	initializeCardTestSuite()
	InitializeCardRegistryWithFakeYAMLData()
	cardTemplate := GetCardRegistry().GetCard(1001) // Test Card
	assert.NotNil(t, cardTemplate)
	assert.Equal(t, 1001, cardTemplate.ID)
}

func TestNewCardInstance(t *testing.T) {
	initializeCardTestSuite()
	InitializeCardRegistryWithFakeYAMLData()
	playableCard, err := NewCardInstance(1001) // Test Card
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
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no card template found for the given ID")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)
}

func TestSingletonCardRegistry(t *testing.T) {
	initializeCardTestSuite()
	registryA := GetCardRegistry()
	registryB := GetCardRegistry()
	assert.Equal(t, registryA, registryB)
	assert.EqualValues(t, registryA, registryB)
}

func TestEquipCard(t *testing.T) {
	initializeCardTestSuite()
	InitializeCardRegistryWithFakeYAMLData()
	equipCard := GetCardRegistry().GetCard(2002) // Fake Malevolent Nuzzler

	assert.NotNil(t, equipCard)
	assert.Equal(t, equipCard.ID, 2002)
	assert.Equal(t, equipCard.Type, TypeEquip)
	assert.NotNil(t, equipCard.EquipRules)
	assert.Equal(t, 1, len(equipCard.EquipRules.ValidTargetIDs))
	assert.Equal(t, 33, equipCard.EquipRules.ValidTargetIDs[0])
	assert.Equal(t, 700, equipCard.EquipRules.Bonus)
}
