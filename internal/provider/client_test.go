package provider

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGenerateEmail(t *testing.T) {
	// 1. Mock Server (Pura-pura jadi GuerrillaMail)
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Cek apakah request benar
		if r.URL.Query().Get("f") != "get_email_address" {
			t.Errorf("Parameter 'f' salah. Dapat: %s", r.URL.Query().Get("f"))
		}

		// Berikan respon palsu (JSON)
		response := `{"email_addr": "mock@donat.com", "sid_token": "mock-token-123"}`
		fmt.Fprintln(w, response)
	}))
	defer mockServer.Close()

	// 2. Inisialisasi Client dengan URL Mock
	client := NewClient()
	client.BaseURL = mockServer.URL // Override URL asli ke server palsu

	// 3. Eksekusi
	email, token, err := client.GenerateEmail()

	// 4. Validasi
	if err != nil {
		t.Fatalf("Error tidak diharapkan: %v", err)
	}
	if email != "mock@donat.com" {
		t.Errorf("Email salah. Dapat: %s", email)
	}
	if token != "mock-token-123" {
		t.Errorf("Token salah. Dapat: %s", token)
	}
}

func TestGetMessages_ErrorWithoutSession(t *testing.T) {
	client := NewClient()
	// Kita sengaja TIDAK set token

	_, err := client.GetMessages()
	if err == nil {
		t.Error("Harusnya error jika session belum ada, tapi malah sukses")
	}
}
