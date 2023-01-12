package appstore

import (
	"errors"
	"testing"
)

func TestHandleErro(t *testing.T) {
	tests := []struct {
		name string
		in   int
		out  error
	}{
		{
			name: "status 0",
			in:   0,
			out:  nil,
		},
		{
			name: "status 21000",
			in:   21000,
			out:  ErrInvalidJSON,
		},
		{
			name: "status 21002",
			in:   21002,
			out:  ErrInvalidReceipt,
		},
		{
			name: "status 21003",
			in:   21003,
			out:  ErrReceiptNotAuthenticated,
		},
		{
			name: "status 21004",
			in:   21004,
			out:  ErrInvalidSharedSecret,
		},
		{
			name: "status 21005",
			in:   21005,
			out:  ErrServiceUnavailable,
		},
		{
			name: "status 21006",
			in:   21006,
			out:  ErrSubscriptionExpired,
		},
		{
			name: "status 21007",
			in:   21007,
			out:  ErrTestEnvironment,
		},
		{
			name: "status 21008",
			in:   21008,
			out:  ErrProductionEnvironment,
		},
		{
			name: "status 21009",
			in:   21009,
			out:  ErrInternalError,
		},
		{
			name: "status 21010",
			in:   21010,
			out:  ErrReceiptUnauthorized,
		},
	}

	for _, d := range tests {
		t.Run(d.name, func(t *testing.T) {
			out := HandleError(d.in)

			if !errors.Is(out, d.out) {
				t.Errorf("input: %d\ngot: %v\nwant: %v\n", d.in, out, d.out)
			}
		})
	}
}
