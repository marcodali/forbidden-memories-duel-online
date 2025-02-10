package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// represents the authentication method used by the player
type AuthProvider string

const (
	Google   AuthProvider = "GOOGLE"
	Facebook AuthProvider = "FACEBOOK"
	Apple    AuthProvider = "APPLE"
)

var verifiedProviders = []AuthProvider{Google, Facebook, Apple}

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

var verifiedCountries = []string{Canada, USA, Mexico, Colombia, Brazil, Chile, Peru, Aregentina}

type Player struct {
	ID           string
	Username     string
	Country      string
	RegisteredAt time.Time
	LastLogin    time.Time
	AuthProvider AuthProvider
	IsOnline     bool
	IsDueling    bool
	LifePoints   int
	TotalDuels   int
	WinCount     int
	LossCount    int
}

func NewPlayer(username string) (*Player, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	now := time.Now()
	return &Player{
		ID:           generateUUID(),
		Username:     username,
		RegisteredAt: now,
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
	// verify if the country is in the list of verified countries
	for _, c := range verifiedCountries {
		if c == country {
			p.Country = country
			return nil
		}
	}
	return fmt.Errorf("invalid country %q: expected one of [%v]", country, verifiedCountries)
}

func (p *Player) SetAuthProvider(authProvider AuthProvider) error {
	// verify if the auth provider is in the list of verified providers
	for _, provider := range verifiedProviders {
		if provider == authProvider {
			p.AuthProvider = authProvider
			return nil
		}
	}
	return fmt.Errorf("invalid auth provider %q: expected one of [%v]", authProvider, verifiedProviders)
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
