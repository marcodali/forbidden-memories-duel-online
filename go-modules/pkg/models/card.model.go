package models

import (
	"fmt"
	"sync"

	"gopkg.in/yaml.v3"
)

// Rarity represents the rarity level of a card
type Rarity string

const (
	RarityNormal     Rarity = "NORMAL"
	RarityRare       Rarity = "RARE"
	RaritySuperRare  Rarity = "SUPER_RARE"
	RarityUltraRare  Rarity = "ULTRA_RARE"
	RaritySecretRare Rarity = "SECRET_RARE"
	RarityUltimate   Rarity = "ULTIMATE_RARE"
	RarityGhost      Rarity = "GHOST_RARE"
)

// EquipRules defines the rules for equip cards
type EquipRules struct {
	ValidTargetIDs []int `yaml:"validTargetIDs"` // List of card IDs that can be equipped
	Bonus          int   `yaml:"bonus"`          // Bonus applied to both ATK and DEF
}

// CardTemplate contains the immutable properties of a card
type CardTemplate struct {
	ID               int         `yaml:"id"`
	Name             string      `yaml:"name"`
	Description      string      `yaml:"description"`
	BaseAttack       int         `yaml:"baseAttack"`
	BaseDefense      int         `yaml:"baseDefense"`
	Level            int         `yaml:"level"`
	MainType         string      `yaml:"mainType"`
	SubType          string      `yaml:"subType"`
	GuardianStars    [2]string   `yaml:"guardianStars"`
	Rarity           Rarity      `yaml:"rarity"`
	IsFusion         bool        `yaml:"isFusion"`
	IsFusionMaterial bool        `yaml:"isFusionMaterial"`
	IsMagic          bool        `yaml:"isMagic"`
	IsEquip          bool        `yaml:"isEquip"`
	IsTrap           bool        `yaml:"isTrap"`
	IsRitual         bool        `yaml:"isRitual"`
	Image            string      `yaml:"image"`
	EquipRules       *EquipRules `yaml:"equipRules,omitempty"`
}

// CardInstance represents a card in play
type CardInstance struct {
	Template       *CardTemplate
	IsInAttackMode bool
	CurrentAttack  int
	CurrentDefense int
}

// CardRegistry is the global registry of card templates
type CardRegistry struct {
	templates map[int]*CardTemplate
}

var (
	registry *CardRegistry
	once     sync.Once
)

// GetCardRegistry returns the singleton instance of the card registry
func GetCardRegistry() *CardRegistry {
	once.Do(func() {
		registry = &CardRegistry{
			templates: make(map[int]*CardTemplate),
		}
	})
	return registry
}

// LoadCards loads card templates from YAML data
func (r *CardRegistry) LoadCards(data []byte) error {
	var templates []*CardTemplate
	if err := yaml.Unmarshal(data, &templates); err != nil {
		return fmt.Errorf("error unmarshalling YAML data: %w", err)
	}

	for _, template := range templates {
		r.templates[template.ID] = template
	}
	return nil
}

// GetCard returns a card template by ID
func (r *CardRegistry) GetCard(id int) *CardTemplate {
	return r.templates[id]
}

// NewCardInstance creates a new instance of a card in play
func NewCardInstance(templateID int) (*CardInstance, error) {
	template := GetCardRegistry().GetCard(templateID)
	if template == nil {
		return nil, fmt.Errorf("no card template found for the given ID")
	}

	return &CardInstance{
		Template:       template,
		CurrentAttack:  template.BaseAttack,
		CurrentDefense: template.BaseDefense,
		IsInAttackMode: false,
	}, nil
}
