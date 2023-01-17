package appstore

import (
	"context"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	current := New()

	expected := &IAPClient{
		ProductionURL: ProductionURL,
		SandboxURL:    SandboxURL,
		Client:        &http.Client{Timeout: 30 * time.Second},
	}

	if !reflect.DeepEqual(current, expected) {
		t.Errorf("object %v is not equal to %v", current, expected)
	}
}

func TestNewWithClient(t *testing.T) {
	client := &http.Client{Timeout: 30 * time.Second}
	current := NewWithClient(client)

	expected := &IAPClient{
		ProductionURL: ProductionURL,
		SandboxURL:    SandboxURL,
		Client:        client,
	}

	if !reflect.DeepEqual(current, expected) {
		t.Errorf("object %v is not equal to %v", current, expected)
	}
}

func TestVerifyTimeout(t *testing.T) {
	client := &IAPClient{
		ProductionURL: ProductionURL,
		SandboxURL:    SandboxURL,
		Client:        &http.Client{Timeout: time.Millisecond},
	}

	req := IAPValidationRequest{
		ReceiptData: "receipt data",
	}

	var res *IAPValidationResponse
	ctx := context.Background()
	err := client.VerifyReceipt(ctx, req, res)

	if err == nil {
		t.Errorf("error should be occurred because of timeout")
	}

	t.Log(err)
}

func TestVerifyWithCancel(t *testing.T) {
	client := New()

	req := IAPValidationRequest{
		ReceiptData: "receipt data",
	}

	res := &IAPValidationResponse{}
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(10 * time.Millisecond)
		cancel()
	}()

	err := client.VerifyReceipt(ctx, req, res)

	if err == nil {
		t.Errorf("error should be occurred because of context cancel")
	}

	t.Log(err)
}

func TestResponse(t *testing.T) {
	req := IAPValidationRequest{
		ReceiptData: "receipt data",
	}
	res := &IAPValidationResponse{}

	type testCase struct {
		name          string
		testServer    *httptest.Server
		sandboxServer *httptest.Server
		expected      *IAPValidationResponse
	}

	testCases := []testCase{
		{
			name:          "VerifySandboxReceipt",
			testServer:    httptest.NewServer(serverWithResponse(http.StatusOK, `{"status": 21007}`)),
			sandboxServer: httptest.NewServer(serverWithResponse(http.StatusOK, `{"status": 0}`)),
			expected: &IAPValidationResponse{
				Status: 0,
			},
		},
		{
			name:          "VerifyBadPayload",
			testServer:    httptest.NewServer(serverWithResponse(http.StatusOK, `{"status": 21002}`)),
			sandboxServer: httptest.NewServer(serverWithResponse(http.StatusOK, `{"status": 21002}`)),
			expected: &IAPValidationResponse{
				Status: 21002,
			},
		},
	}

	client := New()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer tc.testServer.Close()

			client.SandboxURL = tc.sandboxServer.URL
			client.ProductionURL = tc.testServer.URL

			ctx := context.Background()
			err := client.VerifyReceipt(ctx, req, res)

			if err != nil {
				t.Errorf("%s", err)
			}
			if !reflect.DeepEqual(res, tc.expected) {
				t.Errorf("object %v is not equal to %v", res, tc.expected)
			}
		})
	}
}

func serverWithResponse(statusCode int, response string) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			rw.Header().Set("Content-Type", "application/json")
			rw.WriteHeader(statusCode)
			rw.Write([]byte(response))
		} else {
			rw.Write([]byte(`unsupported request`))
		}
	})
}
