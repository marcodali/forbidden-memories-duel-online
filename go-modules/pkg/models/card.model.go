package models

import (
	"fmt"
	"sync"

	"gopkg.in/yaml.v3"
)

// represents the rarity level of a card
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

type TypeCard string

const (
	TypeEquip        TypeCard = "Equip"
	TypeMagic        TypeCard = "Magic"
	TypeRitual       TypeCard = "Ritual"
	TypeTrap         TypeCard = "Trap"
	TypeZombie       TypeCard = "Zombie"
	TypeSpellCaster  TypeCard = "SpellCaster"
	TypeFiend        TypeCard = "Fiend"
	TypePlant        TypeCard = "Plant"
	TypeRock         TypeCard = "Rock"
	TypeInsect       TypeCard = "Inscect"
	TypeAqua         TypeCard = "Aqua"
	TypeFairy        TypeCard = "Faity"
	TypeMachine      TypeCard = "Machine"
	TypeWarrior      TypeCard = "Warrior"
	TypeBeast        TypeCard = "Beast"
	TypeReptile      TypeCard = "Repitle"
	TypePyro         TypeCard = "Pyro"
	TypeDinasour     TypeCard = "Dinasour"
	TypeDragon       TypeCard = "Dragon"
	TypeThunder      TypeCard = "Thunder"
	TypeWingedBeast  TypeCard = "Winged Beast"
	TypeBeastWarrior TypeCard = "Beast-Warrior"
	TypeFish         TypeCard = "Fish"
	TypeSeaSerpent   TypeCard = "Sea Serpent"
)

var validMonsterTypes = map[TypeCard]bool{
	TypeZombie:       true,
	TypeSpellCaster:  true,
	TypeFiend:        true,
	TypePlant:        true,
	TypeRock:         true,
	TypeInsect:       true,
	TypeAqua:         true,
	TypeFairy:        true,
	TypeMachine:      true,
	TypeWarrior:      true,
	TypeBeast:        true,
	TypeReptile:      true,
	TypePyro:         true,
	TypeDinasour:     true,
	TypeDragon:       true,
	TypeThunder:      true,
	TypeWingedBeast:  true,
	TypeBeastWarrior: true,
	TypeFish:         true,
	TypeSeaSerpent:   true,
}

// defines the rules for equip cards
type EquipRules struct {
	ValidTargetIDs []int `yaml:"validTargetIDs"` // List of card IDs that can be equipped
	Bonus          int   `yaml:"bonus"`          // Bonus applied to both ATK and DEF
}

type RitualRules struct {
	Material         []int `yaml:"materialIDs"`
	IsRitualResult   bool  `yaml:"isRitualResult"`
	IsRitualMaterial bool  `yaml:"isRitualMaterial"`
}

// contains the immutable properties of a card
type CardTemplate struct {
	ID            int          `yaml:"id"`
	Name          string       `yaml:"name"`
	Description   string       `yaml:"description"`
	BaseAttack    int          `yaml:"baseAttack"`
	BaseDefense   int          `yaml:"baseDefense"`
	Level         int          `yaml:"level"`
	Type          TypeCard     `yaml:"type"`
	GuardianStars []string     `yaml:"guardianStars"` // slices initialize to nil instead of {"", ""}
	Rarity        Rarity       `yaml:"rarity"`
	EquipRules    *EquipRules  `yaml:"equipRules,omitempty"`
	RitualRules   *RitualRules `yaml:"ritualRules,omitempty"`
}

// represents a card in play
type CardInstance struct {
	Template       *CardTemplate
	IsInAttackMode bool
	CurrentAttack  int
	CurrentDefense int
}

// is the global registry of card templates
type CardRegistry struct {
	templates map[int]*CardTemplate
}

var (
	registry *CardRegistry
	once     sync.Once
)

// returns the singleton instance of the card registry
func GetCardRegistry() *CardRegistry {
	once.Do(func() {
		registry = &CardRegistry{
			templates: make(map[int]*CardTemplate),
		}
	})
	return registry
}

func (r *CardRegistry) LoadCardsfromYAML(data []byte) error {
	var templates []*CardTemplate
	if err := yaml.Unmarshal(data, &templates); err != nil {
		return fmt.Errorf("unexpected error trying to load cards from YAML data: %w", err)
	}

	for _, template := range templates {
		r.templates[template.ID] = template
	}
	return nil
}

// returns a card template by ID
func (r *CardRegistry) GetCard(id int) *CardTemplate {
	return r.templates[id]
}

// creates a new instance of a card in play
func NewCardInstance(templateID int) (*CardInstance, error) {
	template := GetCardRegistry().GetCard(templateID)
	if template == nil {
		return nil, fmt.Errorf("no card template found for the given ID: %d", templateID)
	}

	return &CardInstance{
		Template:       template,
		CurrentAttack:  template.BaseAttack,
		CurrentDefense: template.BaseDefense,
		IsInAttackMode: false,
	}, nil
}
