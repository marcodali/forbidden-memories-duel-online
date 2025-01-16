package models

import (
	"errors"
	"time"
)

// AuthProvider represents the authentication method used by the player
type AuthProvider string

const (
	Google   AuthProvider = "GOOGLE"
	Facebook AuthProvider = "FACEBOOK"
	Apple    AuthProvider = "APPLE"
)

// Player represents a player in the game
type Player struct {
	ID           string
	Username     string
	Country      string
	RegisteredAt time.Time
	LastLogin    time.Time
	AuthProvider AuthProvider
	IsOnline     bool
	IsInDuel     bool
	TotalDuels   int
	WinCount     int
	LossCount    int
}

// NewPlayer creates a new player instance
func NewPlayer(id string, username string, country string, authProvider AuthProvider) (*Player, error) {
	if id == "" {
		return nil, errors.New("player id cannot be empty")
	}
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}
	if country == "" {
		return nil, errors.New("country cannot be empty")
	}

	now := time.Now()
	return &Player{
		ID:           id,
		Username:     username,
		Country:      country,
		RegisteredAt: now,
		LastLogin:    now,
		AuthProvider: authProvider,
		IsOnline:     true,
		IsInDuel:     false,
		TotalDuels:   0,
		WinCount:     0,
		LossCount:    0,
	}, nil
}

// UpdateLastLogin updates the player's last login timestamp
func (p *Player) UpdateLastLogin() {
	p.LastLogin = time.Now()
}

// GetWinRate calculates the player's win rate percentage
func (p *Player) GetWinRate() float64 {
	if p.TotalDuels == 0 {
		return 0.0
	}
	return float64(p.WinCount) / float64(p.TotalDuels) * 100
}
