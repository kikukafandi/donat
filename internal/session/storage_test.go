package session

import (
	"os"
	"testing"
)

func TestSaveAndLoad(t *testing.T) {
	// 1. Setup: Gunakan file temporary untuk testing
	originalFile := SessionFile
	SessionFile = "test_donat_session.json"

	// Bersihkan file test setelah selesai
	defer func() {
		os.Remove(getFilePath())
		SessionFile = originalFile // Balikin ke semula
	}()

	// 2. Data Dummy
	email := "test@sharklasers.com"
	token := "abc-123-xyz"

	// 3. Test SAVE
	err := Save(email, token)
	if err != nil {
		t.Fatalf("Gagal menyimpan session: %v", err)
	}

	// Cek apakah file benar-benar terbuat
	if _, err := os.Stat(getFilePath()); os.IsNotExist(err) {
		t.Error("File session tidak ditemukan setelah Save()")
	}

	// 4. Test LOAD
	loadedSession, err := Load()
	if err != nil {
		t.Fatalf("Gagal me-load session: %v", err)
	}

	// 5. Validasi Data
	if loadedSession.Email != email {
		t.Errorf("Email salah. Harapan: %s, Dapat: %s", email, loadedSession.Email)
	}
	if loadedSession.Token != token {
		t.Errorf("Token salah. Harapan: %s, Dapat: %s", token, loadedSession.Token)
	}
}

func TestClear(t *testing.T) {
	// Setup
	SessionFile = "test_donat_clear.json"
	Save("delete@me.com", "token")

	// Test CLEAR
	err := Clear()
	if err != nil {
		t.Fatalf("Gagal menghapus session: %v", err)
	}

	// Validasi file harusnya hilang
	if _, err := os.Stat(getFilePath()); !os.IsNotExist(err) {
		t.Error("File session masih ada padahal sudah di-Clear()")
	}
}
