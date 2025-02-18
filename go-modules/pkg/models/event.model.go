package models

import (
	"fmt"
	"time"
)

type EventType string

const (
	EventDeckShuffled                    EventType = "DECK_SHUFFLED"
	EventOneCardDroppedOrDestroyed       EventType = "ONE_CARD_DROPPED_OR_DESTROYED"
	EventBulkCardDestruction             EventType = "BULK_CARD_DESTRUCTION"
	EventCardFusionFailed                EventType = "CARD_FUSION_FAILED"
	EventCardFused                       EventType = "CARD_FUSED"
	EventMonsterBattle                   EventType = "MONSTER_BATTLE"
	EventSacrificeCardsForRitual         EventType = "SACRIFICE_CARDS_FOR_RITUAL"
	EventOneCardStateAndPositionChanged  EventType = "ONE_CARD_STATE_AND_POSITION_CHANGED"
	EventBulkCardStateAndPositionChanged EventType = "BULK_CARD_STATE_AND_POSITION_CHANGED"
	EventGetOutOfCards                   EventType = "GET_OUT_OF_CARDS"
	EventOneCardPointsUpdate             EventType = "ONE_CARD_POINTS_UPDATE"
	EventBulkCardPointsUpdate            EventType = "BULK_CARD_POINTS_UPDATE"
	EventDirectDamageToLifePoints        EventType = "DIRECT_DAMAGE_TO_LIFE_POINTS"
	EventPlayerLifePointsUpdate          EventType = "PLAYER_LIFE_POINTS_UPDATE"
	EventTrapActivated                   EventType = "TRAP_ACTIVATED"
	EventChangeFieldLand                 EventType = "CHANGE_FIELD_LAND"
	EventEquipCardAttached               EventType = "EQUIP_CARD_ATTACHED"
	EventGuardianStarChange              EventType = "GUARDIAN_STAR_CHANGE"
	EventMagicCardActivated              EventType = "MAGIC_CARD_ACTIVATED"
	EventPlayerWins                      EventType = "PLAYER_WINS"
	EventPlayerLoses                     EventType = "PLAYER_LOSES"
	EventTurnPhaseChange                 EventType = "TURN_PHASE_CHANGE"
	EventProhibitOpponentToAtack         EventType = "PROHIBIT_OPPONENT_TO_ATACK"
)

var validEventTypes = map[EventType]bool{
	EventDeckShuffled:                    true,
	EventOneCardDroppedOrDestroyed:       true,
	EventBulkCardDestruction:             true,
	EventCardFusionFailed:                true,
	EventCardFused:                       true,
	EventMonsterBattle:                   true,
	EventSacrificeCardsForRitual:         true,
	EventOneCardStateAndPositionChanged:  true,
	EventBulkCardStateAndPositionChanged: true,
	EventGetOutOfCards:                   true,
	EventOneCardPointsUpdate:             true,
	EventBulkCardPointsUpdate:            true,
	EventDirectDamageToLifePoints:        true,
	EventPlayerLifePointsUpdate:          true,
	EventTrapActivated:                   true,
	EventChangeFieldLand:                 true,
	EventEquipCardAttached:               true,
	EventGuardianStarChange:              true,
	EventMagicCardActivated:              true,
	EventPlayerWins:                      true,
	EventPlayerLoses:                     true,
	EventTurnPhaseChange:                 true,
	EventProhibitOpponentToAtack:         true,
}

type Event struct {
	Type      EventType
	Timestamp time.Time
	Data      map[string]any // Flexible data storage
}

func NewEvent(eventType EventType, data map[string]any) (*Event, error) {
	if validEventTypes[eventType] {
		return &Event{
			Type:      eventType,
			Timestamp: time.Now(),
			Data:      data,
		}, nil
	}
	return nil, fmt.Errorf("invalid event type %q: expected one of [%v]", eventType, validEventTypes)
}
