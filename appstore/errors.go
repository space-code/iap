package appstore

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidJSON             = errors.New("The App Store could not read the JSON object you provided.")
	ErrInvalidReceipt          = errors.New("The data in the receipt-data property was malformed or missing.")
	ErrReceiptNotAuthenticated = errors.New("The receipt could not be authenticated.")
	ErrInvalidSharedSecret     = errors.New("The shared secret you provided does not match the shared secret on file for your account.")
	ErrServiceUnavailable      = errors.New("The receipt server is not currently available.")
	ErrSubscriptionExpired     = errors.New("This receipt is valid but the subscription has expired. When this status code is returned to your server, the receipt data is also decoded and returned as part of the response.")
	ErrTestEnvironment         = errors.New("This receipt is from the test environment, but it was sent to the production environment for verification. Send it to the test environment instead.")
	ErrProductionEnvironment   = errors.New("This receipt is from the production environment, but it was sent to the test environment for verification. Send it to the production environment instead.")
	ErrInternalError           = errors.New("Internal data access error.")
	ErrReceiptUnauthorized     = errors.New("This receipt could not be authorized. Treat this the same as if a purchase was never made.")
	ErrUnknown                 = errors.New("Unknown error")
)

// Handle App Store errors from response status code.
func HandleError(status int) error {
	var e error

	switch status {
	case 0:
		return nil
	case 21000:
		e = ErrInvalidJSON
	case 21002:
		e = ErrInvalidReceipt
	case 21003:
		e = ErrReceiptNotAuthenticated
	case 21004:
		e = ErrInvalidSharedSecret
	case 21005:
		e = ErrServiceUnavailable
	case 21006:
		e = ErrSubscriptionExpired
	case 21007:
		e = ErrTestEnvironment
	case 21008:
		e = ErrProductionEnvironment
	case 21009:
		e = ErrInternalError
	case 21010:
		e = ErrReceiptUnauthorized
	default:
		e = ErrUnknown
	}

	return fmt.Errorf("status %d: %w", status, e)
}
