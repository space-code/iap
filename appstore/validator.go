package appstore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	// The endpoint to sandbox environment.
	SandboxURL string = "https://sandbox.itunes.apple.com/verifyReceipt"
	// The endpoint to production environment.
	ProductionURL string = "https://buy.itunes.apple.com/verifyReceipt"
	// The content type definition.
	ContentType string = "application/json"
)

type Client interface {
	VerifyReceipt(ctx context.Context, request IAPValidationRequest, response *IAPValidationResponse) error
}

// IAPClients implements Client
type IAPClient struct {
	URL    string
	Client *http.Client
}

// Send a receipt to the App Store for verification.
func (c *IAPClient) verifyReceipt(ctx context.Context, URL string, data []byte, response *IAPValidationResponse) error {
	r, err := http.NewRequest("POST", URL, bytes.NewBuffer(data))
	r = r.WithContext(ctx)

	if err != nil {
		return err
	}

	r.Header.Add("Content-Type", ContentType)

	res, err := c.Client.Do(r)

	if err != nil {
		return err
	}

	if res.StatusCode >= 500 {
		return fmt.Errorf("Received http status code %d from the App Store", res.StatusCode)
	}

	defer r.Body.Close()

	return c.parseResponse(ctx, res, data, response)
}

// Send a receipt to the App Store for verification.
func (c *IAPClient) VerifyReceipt(ctx context.Context, request IAPValidationRequest, response *IAPValidationResponse) error {
	data, err := json.Marshal(request)

	if err != nil {
		return err
	}

	return c.verifyReceipt(ctx, c.URL, data, response)
}

// Parse response from the App Store.
//
// The method reads data from the response's body and decodes it to the IAPValidationResponse.
//
// Also, the method checks the status code of the response and if it is equal to 21007 sends a new request with a sandbox URL.
func (c *IAPClient) parseResponse(ctx context.Context, res *http.Response, data []byte, result *IAPValidationResponse) error {
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	var s Status
	if err := json.Unmarshal(body, &s); err != nil {
		return err
	}

	if s.Status == 21007 {
		return c.verifyReceipt(ctx, SandboxURL, data, result)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return err
	}

	return nil
}

// Create new instance of IAPClient.
func NewWithClient(c *http.Client) *IAPClient {
	client := &IAPClient{
		URL:    ProductionURL,
		Client: c,
	}
	return client
}

// Create new instance of IAPClient.
func New() *IAPClient {
	client := &IAPClient{
		URL:    ProductionURL,
		Client: &http.Client{Timeout: 30 * time.Second},
	}
	return client
}
