// Package alfapay provides a Go client for the Alfa Payments API.
package alfapay

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// DefaultBaseURL is the default Alfa Payments API base URL.
	DefaultBaseURL = "https://alfa.rbsuat.com/payment"
	// DefaultTimeout is the default HTTP client timeout.
	DefaultTimeout = 30 * time.Second
)

// Client is the Alfa Payments API client.
type Client struct {
	baseURL    string
	httpClient *http.Client
	userName   string
	password   string

	// Services
	Orders     *OrderService
	Status     *StatusService
	Bindings   *BindingService
	Payments   *PaymentService
	Refunds    *RefundService
	SBP        *SBPService
	ApplePay   *ApplePayService
	GooglePay  *GooglePayService
	SamsungPay *SamsungPayService
	MirPay     *MirPayService
	YandexPay  *YandexPayService
}

// ClientOption is a function that configures the client.
type ClientOption func(*Client)

// WithBaseURL sets a custom base URL for the client.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.baseURL = strings.TrimRight(baseURL, "/")
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithTimeout sets a custom timeout for the HTTP client.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.httpClient.Timeout = timeout
	}
}

// NewClient creates a new Alfa Payments API client.
func NewClient(userName, password string, opts ...ClientOption) *Client {
	c := &Client{
		baseURL:  DefaultBaseURL,
		userName: userName,
		password: password,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	// Initialize services
	c.Orders = &OrderService{client: c}
	c.Status = &StatusService{client: c}
	c.Bindings = &BindingService{client: c}
	c.Payments = &PaymentService{client: c}
	c.Refunds = &RefundService{client: c}
	c.SBP = &SBPService{client: c}
	c.ApplePay = &ApplePayService{client: c}
	c.GooglePay = &GooglePayService{client: c}
	c.SamsungPay = &SamsungPayService{client: c}
	c.MirPay = &MirPayService{client: c}
	c.YandexPay = &YandexPayService{client: c}

	return c
}

// doRequest performs an HTTP request and decodes the response.
func (c *Client) doRequest(ctx context.Context, method, endpoint string, query url.Values, body interface{}, result interface{}) error {
	// Add authentication to query params
	if query == nil {
		query = url.Values{}
	}
	query.Set("userName", c.userName)
	query.Set("password", c.password)

	fullURL := fmt.Sprintf("%s%s", c.baseURL, endpoint)
	if len(query) > 0 {
		fullURL = fmt.Sprintf("%s?%s", fullURL, query.Encode())
	}

	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, fullURL, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(respBody),
		}
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// doJSONRequest performs a JSON POST request.
func (c *Client) doJSONRequest(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	return c.doRequest(ctx, http.MethodPost, endpoint, nil, body, result)
}

// doFormRequest performs a form POST request with query parameters.
func (c *Client) doFormRequest(ctx context.Context, endpoint string, params url.Values, result interface{}) error {
	return c.doRequest(ctx, http.MethodPost, endpoint, params, nil, result)
}

// doJSONRequestNoAuth performs a JSON POST request without query auth params.
// Used for endpoints where auth is passed in the request body.
func (c *Client) doJSONRequestNoAuth(ctx context.Context, endpoint string, body interface{}, result interface{}) error {
	fullURL := fmt.Sprintf("%s%s", c.baseURL, endpoint)

	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fullURL, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    string(respBody),
		}
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// APIError represents an API error response.
type APIError struct {
	StatusCode int
	Message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error (status %d): %s", e.StatusCode, e.Message)
}
