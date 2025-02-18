package models

import "time"

type EventType string

const (
	EventOneCardDroppedOrDestroyed EventType = "ONE_CARD_DROPPED_OR_DESTROYED"
	EventPlayerLifePointsUpdate    EventType = "PLAYER_LIFE_POINTS_UPDATE"
	EventCardFusionFailed          EventType = "CARD_FUSION_FAILED"
	EventCardFused                 EventType = "CARD_FUSED"
	EventMonsterBattle             EventType = "MONSTER_BATTLE"
	EventSacrificeOfferedForRitual EventType = "SACRIFICE_OFFERED_FOR_RITUAL"
	EventCardPositionChanged       EventType = "CARD_POSITION_CHANGED"
	EventGetOutOfCards             EventType = "GET_OUT_OF_CARDS"
	EventDirectDamageToLifePoints  EventType = "DIRECT_DAMAGE_TO_LIFE_POINTS"
	EventTrapActivated             EventType = "TRAP_ACTIVATED"
	EventBulkCardDestruction       EventType = "BULK_CARD_DESTRUCTION"
	EventPlayerWins                EventType = "PLAYER_WINS"
	EventPlayerLoses               EventType = "PLAYER_LOSES"
	EventChangeFieldLand           EventType = "CHANGE_FIELD_LAND"
	EventOneCardPointsUpdate       EventType = "ONE_CARD_POINTS_UPDATE"
	EventBulkCardPointsUpdate      EventType = "BULK_CARD_POINTS_UPDATE"
)

type Event struct {
	Type      EventType
	Timestamp time.Time
	Data      map[string]any // Flexible data storage
}

func NewEvent(eventType EventType, data map[string]interface{}) *Event {
	return &Event{
		Type:      eventType,
		Timestamp: time.Now(),
		Data:      data,
	}
}
