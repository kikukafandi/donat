package provider

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// Client handles interaction with the GuerrillaMail API.
// It maintains the session token required for sequential requests.
type Client struct {
	BaseURL      string
	HTTPClient   *http.Client
	SessionToken string
}

// MailMessage represents a summary of an email in the inbox list.
type MailMessage struct {
	ID      int    `json:"mail_id"`
	From    string `json:"mail_from"`
	Subject string `json:"mail_subject"`
	Date    string `json:"mail_date"`
}

// FullMessage represents the complete email content including body.
type FullMessage struct {
	ID       int    `json:"mail_id"`
	From     string `json:"mail_from"`
	Subject  string `json:"mail_subject"`
	Body     string `json:"mail_body"`    // Raw HTML body
	TextBody string `json:"mail_excerpt"` // Text snippet
}

// internal struct for authentication response parsing.
type guerrillaAuthResponse struct {
	Email    string `json:"email_addr"`
	SIDToken string `json:"sid_token"`
}

// internal struct for list response parsing.
type guerrillaListResponse struct {
	List []MailMessage `json:"list"`
}

// NewClient initializes a new client with a predefined timeout.
func NewClient() *Client {
	return &Client{
		BaseURL:    "https://api.guerrillamail.com/ajax.php",
		HTTPClient: &http.Client{Timeout: 20 * time.Second},
	}
}

// SetToken updates the client's session token.
// This is typically used when loading a saved session from disk.
func (c *Client) SetToken(token string) {
	c.SessionToken = token
}

// doRequest performs a GET request with specific headers to mimic a browser.
// This helps avoid 403 Forbidden errors from the API provider.
func (c *Client) doRequest(params url.Values) (*http.Response, error) {
	reqURL := fmt.Sprintf("%s?%s", c.BaseURL, params.Encode())
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}
	// Use a standard User-Agent to prevent bot detection.
	req.Header.Set("User-Agent", "Mozilla/5.0 (DonatCLI/1.0)")
	return c.HTTPClient.Do(req)
}

// GenerateEmail requests a new disposable email address.
// Returns the email address and the session token required for future requests.
func (c *Client) GenerateEmail() (string, string, error) {
	params := url.Values{}
	params.Add("f", "get_email_address")

	resp, err := c.doRequest(params)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	var result guerrillaAuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", "", fmt.Errorf("failed to parse response: %w", err)
	}

	c.SessionToken = result.SIDToken
	return result.Email, result.SIDToken, nil
}

// GetMessages retrieves the list of emails currently in the inbox.
// Requires a valid SessionToken to be set.
func (c *Client) GetMessages() ([]MailMessage, error) {
	if c.SessionToken == "" {
		return nil, fmt.Errorf("session not initialized (run 'donat bake' first)")
	}
	params := url.Values{}
	params.Add("f", "get_email_list")
	params.Add("offset", "0")
	params.Add("sid_token", c.SessionToken)

	resp, err := c.doRequest(params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result guerrillaListResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		// Guerrilla API might return malformed JSON for empty lists or errors.
		// We treat decode errors as an empty list to avoid crashing the flow.
		return nil, nil
	}
	return result.List, nil
}

// ReadMessage fetches the full content of a specific email by ID.
func (c *Client) ReadMessage(id int) (*FullMessage, error) {
	if c.SessionToken == "" {
		return nil, fmt.Errorf("session not initialized")
	}
	params := url.Values{}
	params.Add("f", "fetch_email")
	params.Add("email_id", strconv.Itoa(id))
	params.Add("sid_token", c.SessionToken)

	resp, err := c.doRequest(params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var msg FullMessage
	if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
		return nil, fmt.Errorf("failed to parse message body: %w", err)
	}
	return &msg, nil
}
