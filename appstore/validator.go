package appstore

import (
	"bytes"
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
	VerifyReceipt(request IAPValidationRequest) error
}

// IAPClients implements Client
type IAPClient struct {
	URL    string
	Client *http.Client
}

// Send a receipt to the App Store for verification.
func (c *IAPClient) verifyReceipt(URL string, data []byte) error {
	r, err := http.NewRequest("POST", URL, bytes.NewBuffer(data))

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

	return c.parseResponse(res, data)
}

// Send a receipt to the App Store for verification.
func (c *IAPClient) VerifyReceipt(request IAPValidationRequest) error {
	data, err := json.Marshal(request)

	if err != nil {
		return err
	}

	return c.verifyReceipt(c.URL, data)
}

// Parse response from the App Store.
//
// The method reads data from the response's body and decodes it to the IAPValidationResponse.
//
// Also, the method checks the status code of the response and if it is equal to 21007 sends a new request with a sandbox URL.
func (c *IAPClient) parseResponse(res *http.Response, data []byte) error {
	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	var s Status
	if err := json.Unmarshal(body, &s); err != nil {
		return err
	}

	if s.Status == 21007 {
		return c.verifyReceipt(SandboxURL, data)
	}

	var receipt IAPValidationResponse
	if err := json.Unmarshal(body, &receipt); err != nil {
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
