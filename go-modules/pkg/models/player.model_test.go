package models

import (
    "testing"
    "time"
    "github.com/stretchr/testify/assert"
)

func TestNewPlayer(t *testing.T) {
    tests := []struct {
        name         string
        id          string
        username    string
        country     string
        authProvider AuthProvider
        wantErr     bool
    }{
        {
            name:         "Valid player creation with Google",
            id:          "player-1",
            username:    "TestPlayer",
            country:     "MX",
            authProvider: Google,
            wantErr:     false,
        },
        {
            name:         "Valid player creation with Facebook",
            id:          "player-2",
            username:    "TestPlayer2",
            country:     "US",
            authProvider: Facebook,
            wantErr:     false,
        },
        {
            name:         "Empty ID",
            id:          "",
            username:    "TestPlayer",
            country:     "MX",
            authProvider: Google,
            wantErr:     true,
        },
        {
            name:         "Empty username",
            id:          "player-1",
            username:    "",
            country:     "MX",
            authProvider: Google,
            wantErr:     true,
        },
        {
            name:         "Empty country",
            id:          "player-1",
            username:    "TestPlayer",
            country:     "",
            authProvider: Google,
            wantErr:     true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            player, err := NewPlayer(tt.id, tt.username, tt.country, tt.authProvider)
            
            if tt.wantErr {
                assert.Error(t, err)
                assert.Nil(t, player)
            } else {
                assert.NoError(t, err)
                assert.NotNil(t, player)
                assert.Equal(t, tt.id, player.ID)
                assert.Equal(t, tt.username, player.Username)
                assert.Equal(t, tt.country, player.Country)
                assert.Equal(t, tt.authProvider, player.AuthProvider)
                assert.True(t, player.IsOnline)
                assert.False(t, player.IsInDuel)
                assert.Equal(t, 0, player.TotalDuels)
                assert.Equal(t, 0, player.WinCount)
                assert.Equal(t, 0, player.LossCount)
                
                // Verify timestamps are recent
                assert.WithinDuration(t, time.Now(), player.RegisteredAt, 2*time.Second)
                assert.WithinDuration(t, time.Now(), player.LastLogin, 2*time.Second)
            }
        })
    }
}

func TestUpdateLastLogin(t *testing.T) {
    player, _ := NewPlayer("player-1", "TestPlayer", "MX", Google)
    
    // Simulate time passing
    time.Sleep(time.Millisecond * 100)
    
    oldLogin := player.LastLogin
    player.UpdateLastLogin()
    
    assert.True(t, player.LastLogin.After(oldLogin))
}

func TestGetWinRate(t *testing.T) {
    player, _ := NewPlayer("player-1", "TestPlayer", "MX", Google)

    tests := []struct {
        name       string
        totalDuels int
        wins       int
        losses     int
        expected   float64
    }{
        {
            name:       "No duels played",
            totalDuels: 0,
            wins:       0,
            losses:     0,
            expected:   0.0,
        },
        {
            name:       "50% win rate",
            totalDuels: 10,
            wins:       5,
            losses:     5,
            expected:   50.0,
        },
        {
            name:       "100% win rate",
            totalDuels: 5,
            wins:       5,
            losses:     0,
            expected:   100.0,
        },
        {
            name:       "0% win rate",
            totalDuels: 3,
            wins:       0,
            losses:     3,
            expected:   0.0,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            player.TotalDuels = tt.totalDuels
            player.WinCount = tt.wins
            player.LossCount = tt.losses
            
            assert.Equal(t, tt.expected, player.GetWinRate())
        })
    }
}
