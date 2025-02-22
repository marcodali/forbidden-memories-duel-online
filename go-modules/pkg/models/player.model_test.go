package models

import (
	"fmt"
	"path/filepath"
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var singletonForPlayerModel sync.Once

func initializePlayerTestSuite() {
	singletonForPlayerModel.Do(func() {
		_, filename, _, _ := runtime.Caller(0)
		fmt.Println("This setup code executes only one time for the file", filepath.Base(filename))
		CleanRegistry()
	})
}

func TestNewPlayer(t *testing.T) {
	initializePlayerTestSuite()
	tests := []struct {
		name         string
		username     string
		country      string
		authProvider AuthProvider
		wantErr      bool
	}{
		{
			name:         "Valid player creation with Google",
			username:     "TestPlayer",
			country:      Mexico,
			authProvider: Google,
			wantErr:      false,
		},
		{
			name:         "Valid player creation with Facebook",
			username:     "TestPlayer",
			country:      USA,
			authProvider: Facebook,
			wantErr:      false,
		},
		{
			name:         "Empty username",
			username:     "",
			country:      Colombia,
			authProvider: Apple,
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			player, err := NewPlayer(tt.username)

			if tt.wantErr {
				assert.Nil(t, player)
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "username cannot be empty")
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, player)
				assert.Equal(t, tt.username, player.Username)
				assert.True(t, player.IsOnline)
				assert.False(t, player.IsDueling)
				assert.Equal(t, 0, player.TotalDuels)
				assert.Equal(t, 0, player.WinCount)
				assert.Equal(t, 0, player.LossCount)

				// Verify timestamps are initialized
				assert.False(t, player.WhenSignedUp.IsZero(), "should not be January 1, year 1")
				assert.False(t, player.LastLogin.IsZero(), "should not be January 1, year 1")

				// call country setter method
				err = player.SetCountry(tt.country)
				assert.NoError(t, err)
				assert.Equal(t, tt.country, player.Country)

				// call auth provider setter method
				err = player.SetAuthProvider(tt.authProvider)
				assert.NoError(t, err)
				assert.Equal(t, tt.authProvider, player.AuthProvider)
			}
		})
	}
}

func TestInvalidCountry(t *testing.T) {
	initializePlayerTestSuite()
	test := struct {
		username string
		country  string
	}{
		username: "TestPlayer",
		country:  "",
	}

	player, _ := NewPlayer(test.username)
	err := player.SetCountry(test.country)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid country")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)
}

func TestInvalidAuthProvider(t *testing.T) {
	initializePlayerTestSuite()
	test := struct {
		username     string
		authProvider AuthProvider
	}{
		username:     "TestPlayer",
		authProvider: "",
	}

	player, _ := NewPlayer(test.username)
	err := player.SetAuthProvider(test.authProvider)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid auth provider")

	// to see this error message, run the test with -v flag
	t.Logf("Error: %v", err)
}

func TestUpdateLastLogin(t *testing.T) {
	initializePlayerTestSuite()
	player, _ := NewPlayer("TestPlayer")

	// Simulate time passing
	time.Sleep(1 * time.Nanosecond)

	oldLogin := player.LastLogin
	player.UpdateLastLogin()

	assert.True(t, player.LastLogin.After(oldLogin))
}

func TestGetWinRate(t *testing.T) {
	initializePlayerTestSuite()
	player, _ := NewPlayer("TestPlayer")

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
