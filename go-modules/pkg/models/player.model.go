package models

import (
	"errors"
	"fmt"
	"time"

	"slices"

	"github.com/google/uuid"
)

// represents the authentication method used by the player
type AuthProvider string

const (
	Google   AuthProvider = "GOOGLE"
	Facebook AuthProvider = "FACEBOOK"
	Apple    AuthProvider = "APPLE"
)

var validAuthProviders = []AuthProvider{Google, Facebook, Apple}

const (
	Canada     = "CA"
	USA        = "US"
	Mexico     = "MX"
	Colombia   = "CO"
	Brazil     = "BR"
	Chile      = "CL"
	Peru       = "PE"
	Aregentina = "AR"
)

var validSignUpCountries = []string{Canada, USA, Mexico, Colombia, Brazil, Chile, Peru, Aregentina}

type Player struct {
	ID                    string
	Username              string
	Country               string
	WhenSignedUp          time.Time
	LastLogin             time.Time
	AuthProvider          AuthProvider
	IsOnline              bool
	IsDueling             bool
	RemainingTurnsToAtack int // 0 means no waiting is needed and the player can atack now!
	LifePoints            int
	TotalDuels            int
	WinCount              int
	LossCount             int
}

func NewPlayer(username string) (*Player, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	now := time.Now()
	return &Player{
		ID:           generateUUID(),
		Username:     username,
		WhenSignedUp: now,
		LastLogin:    now,
		IsOnline:     true,
		IsDueling:    false,
		TotalDuels:   0,
		WinCount:     0,
		LossCount:    0,
	}, nil
}

func generateUUID() string {
	return uuid.New().String()
}

func (p *Player) SetCountry(country string) error {
	// verify if the country is in the list of valid sign-up countries
	if slices.Contains(validSignUpCountries, country) {
		p.Country = country
		return nil
	}
	return fmt.Errorf("invalid country %q: expected one of [%v]", country, validSignUpCountries)
}

func (p *Player) SetAuthProvider(authProvider AuthProvider) error {
	// verify if the auth provider is in the list of valid auth providers
	if slices.Contains(validAuthProviders, authProvider) {
		p.AuthProvider = authProvider
		return nil
	}
	return fmt.Errorf("invalid auth provider %q: expected one of [%v]", authProvider, validAuthProviders)
}

func (p *Player) UpdateLastLogin() {
	p.LastLogin = time.Now()
}

// calculates the player's win rate percentage
func (p *Player) GetWinRate() float64 {
	if p.TotalDuels == 0 {
		return 0.0
	}
	return float64(p.WinCount) / float64(p.TotalDuels) * 100
}
