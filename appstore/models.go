package appstore

type IAPValidationRequest struct {
	ReceiptData            string `json:"receipt-data"`
	Password               string `json:"password,omitempty"`
	ExcludeOldTransactions bool   `json:"exclude-old-transactions"`
}

type Status struct {
	Status int `json:"status"`
}

type IAPValidationResponse struct {
	Environment        string               `json:"environment"`
	IsRetryable        bool                 `json:"is_retryable,omitempty"`
	LatestReceipt      []byte               `json:"latest_receipt,omitempty"`
	LatestReceiptInfo  []LatestReceiptInfo  `json:"latest_receipt_info,omitempty"`
	PendingRenewalInfo []PendingRenewalInfo `json:"pending_renewal_info,omitempty"`
	Receipt            Receipt              `json:"receipt"`
	Status             int                  `json:"status"`
}

type LatestReceiptInfo struct {
	AppAccountToken             string `json:"app_account_token,omitempty"`
	CancellationDate            string `json:"cancellation_date,omitempty"`
	CancellationDateMs          string `json:"cancellation_date_ms,omitempty"`
	CancellationDatePst         string `json:"cancellation_date_pst,omitempty"`
	CancellationReason          string `json:"cancellation_reason,omitempty"`
	ExpiresDate                 string `json:"expires_date,omitempty"`
	ExpiresDateMs               string `json:"expires_date_ms,omitempty"`
	ExpiresDatePst              string `json:"expires_date_pst,omitempty"`
	InAppOwnershipType          string `json:"in_app_ownership_type,omitempty"`
	IsInIntroOfferPeriod        string `json:"is_in_intro_offer_period,omitempty"`
	IsTrialPeriod               string `json:"is_trial_period"`
	IsUpgraded                  string `json:"is_upgraded,omitempty"`
	OfferCodeRefName            string `json:"offer_code_ref_name"`
	OriginalPurchaseDate        string `json:"original_purchase_date"`
	OriginalPurchaseDateMs      string `json:"original_purchase_date_ms"`
	OriginalPurchaseDatePst     string `json:"original_purchase_date_pst"`
	OriginalTransactionId       string `json:"original_transaction_id"`
	ProductId                   string `json:"product_id"`
	PromotionalOfferId          string `json:"promotional_offer_id"`
	PurchaseDate                string `json:"purchase_date"`
	PurchaseDateMs              string `json:"purchase_date_ms"`
	PurchaseDatePst             string `json:"purchase_date_pst"`
	Quantity                    string `json:"quantity"`
	SubscriptionGroupIdentifier string `json:"subscription_group_identifier"`
	WebOrderLineItemId          string `json:"web_order_line_item_id,omitempty"`
	TransactionId               string `json:"transaction_id"`
}

type PendingRenewalInfo struct {
	AutoRenewProductId        string `json:"auto_renew_product_id"`
	AutoRenewStatus           string `json:"auto_renew_status"`
	ExpirationIntent          string `json:"expiration_intent,omitempty"`
	GracePeriodExpiresDate    string `json:"grace_period_expires_date,omitempty"`
	GracePeriodExpiresDateMs  string `json:"grace_period_expires_date_ms,omitempty"`
	GracePeriodExpiresDatePst string `json:"grace_period_expires_date_pst,omitempty"`
	IsInBillingRetryPeriod    string `json:"is_in_billing_retry_period,omitempty"`
	OfferCodeRefName          string `json:"offer_code_ref_name,omitempty"`
	OriginalTransactionId     string `json:"original_transaction_id"`
	PriceConsentStatus        string `json:"price_consent_status"`
	ProductId                 string `json:"product_id"`
	PromotionalOfferId        string `json:"promotional_offer_id"`
	PriceIncreaseStatus       string `json:"price_increase_status"`
}

type Receipt struct {
	AdamId                     int64   `json:"adam_id"`
	AppItemId                  int64   `json:"app_item_id"`
	ApplicationVersion         string  `json:"application_version"`
	BundleId                   string  `json:"bundle_id"`
	DownloadId                 int64   `json:"download_id"`
	ExpirationDate             string  `json:"expiration_date"`
	ExpirationDateMs           string  `json:"expiration_date_ms"`
	ExpirationDatePst          string  `json:"expiration_date_pst"`
	InApp                      []InApp `json:"in_app"`
	OriginalApplicationVersion string  `json:"original_application_version"`
	OriginalPurchaseDate       string  `json:"original_purchase_date"`
	OriginalPurchaseDateMs     string  `json:"original_purchase_date_ms"`
	OriginalPurchaseDatePst    string  `json:"original_purchase_date_pst"`
	PreorderDate               string  `json:"preorder_date"`
	PreorderDateMs             string  `json:"preorder_date_ms"`
	PreorderDatePst            string  `json:"preorder_date_pst"`
	ReceiptCreationDate        string  `json:"receipt_creation_date"`
	ReceiptCreationDateMs      string  `json:"receipt_creation_date_ms"`
	ReceiptCreationDatePst     string  `json:"receipt_creation_date_pst"`
	ReceiptType                string  `json:"receipt_type"`
	RequestDate                string  `json:"request_date"`
	RequestDateMs              string  `json:"request_date_ms"`
	RequestDatePst             string  `json:"request_date_pst"`
	VersionExternalIdentifier  int64   `json:"version_external_identifier,omitempty"`
}

type InApp struct {
	CancellationDate        string `json:"cancellation_date,omitempty"`
	CancellationDateMs      string `json:"cancellation_date_ms,omitempty"`
	CancellationDatePst     string `json:"cancellation_date_pst,omitempty"`
	CancellationReason      string `json:"cancellation_reason"`
	ExpiresDate             string `json:"expires_date,omitempty"`
	ExpiresDateMs           string `json:"expires_date_ms,omitempty"`
	ExpiresDatePst          string `json:"expires_date_pst,omitempty"`
	IsInIntroOfferPeriod    string `json:"is_in_intro_offer_period"`
	IsTrialPeriod           string `json:"is_trial_period"`
	OriginalPurchaseDate    string `json:"original_purchase_date"`
	OriginalPurchaseDateMs  string `json:"original_purchase_date_ms"`
	OriginalPurchaseDatePst string `json:"original_purchase_date_pst"`
	OriginalTransactionId   string `json:"original_transaction_id"`
	ProductId               string `json:"product_id"`
	PromotionalOfferId      string `json:"promotional_offer_id"`
	PurchaseDate            string `json:"purchase_date"`
	PurchaseDateMs          string `json:"purchase_date_ms"`
	PurchaseDatePst         string `json:"purchase_date_pst"`
	Quantity                string `json:"quantity"`
	TransactionId           string `json:"transaction_id"`
	WebOrderLineItemId      string `json:"web_order_line_item_id,omitempty"`
}
