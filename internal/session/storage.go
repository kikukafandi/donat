package session

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// SessionFile defines the name of the persistent file stored in the user's home directory.
var SessionFile = ".donat_session.json"

// DonatSession represents the persistent state required to maintain the user's identity.
type DonatSession struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

// getFilePath resolves the absolute path to the session file based on the OS user's home directory.
func getFilePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, SessionFile)
}

// Save writes the current email and session token to disk.
func Save(email, token string) error {
	data := DonatSession{Email: email, Token: token}
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	// 0644 permission: Read/Write for owner, Read-only for others.
	return os.WriteFile(getFilePath(), file, 0644)
}

// Load reads and unmarshals the session file from disk.
func Load() (*DonatSession, error) {
	data, err := os.ReadFile(getFilePath())
	if err != nil {
		return nil, err
	}
	var session DonatSession
	if err := json.Unmarshal(data, &session); err != nil {
		return nil, err
	}
	return &session, nil
}

// Clear removes the session file, effectively logging the user out.
func Clear() error {
	return os.Remove(getFilePath())
}
