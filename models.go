package alfapay

// OrderStatus represents the status of an order.
type OrderStatus int

const (
	OrderStatusRegistered       OrderStatus = 0 // Order registered, not paid
	OrderStatusPreAuthorized    OrderStatus = 1 // Pre-authorized amount held
	OrderStatusFullyAuthorized  OrderStatus = 2 // Full authorization completed
	OrderStatusCancelled        OrderStatus = 3 // Authorization cancelled
	OrderStatusRefunded         OrderStatus = 4 // Refund completed
	OrderStatusACSAuthorization OrderStatus = 5 // ACS authorization initiated
	OrderStatusDeclined         OrderStatus = 6 // Authorization declined
)

// TaxType represents the tax type for fiscal operations.
type TaxType int

const (
	TaxTypeNoVAT      TaxType = 0 // Without VAT
	TaxTypeVAT0       TaxType = 1 // VAT 0%
	TaxTypeVAT10      TaxType = 2 // VAT 10%
	TaxTypeVAT20      TaxType = 3 // VAT 20%
	TaxTypeVAT10_110  TaxType = 4 // VAT 10/110
	TaxTypeVAT20_120  TaxType = 5 // VAT 20/120
	TaxTypeVAT5       TaxType = 7 // VAT 5%
	TaxTypeVAT7       TaxType = 8 // VAT 7%
	TaxTypeVAT5_105   TaxType = 9 // VAT 5/105
	TaxTypeVAT7_107   TaxType = 10 // VAT 7/107
)

// TaxSystem represents the taxation system.
type TaxSystem int

const (
	TaxSystemGeneral         TaxSystem = 0 // General
	TaxSystemSimplifiedIncome TaxSystem = 1 // Simplified, income
	TaxSystemSimplifiedDiff  TaxSystem = 2 // Simplified, income minus expenses
	TaxSystemENVD            TaxSystem = 3 // ENVD
	TaxSystemESN             TaxSystem = 4 // ESN (agricultural)
	TaxSystemPatent          TaxSystem = 5 // Patent
)

// BaseResponse contains common response fields.
type BaseResponse struct {
	ErrorCode    string `json:"errorCode,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	UserMessage  string `json:"userMessage,omitempty"`
}

// IsSuccess returns true if the response indicates success.
func (r BaseResponse) IsSuccess() bool {
	return r.ErrorCode == "" || r.ErrorCode == "0"
}

// RegisterOrderRequest represents a request to register a new order.
type RegisterOrderRequest struct {
	OrderNumber          string                 `json:"orderNumber"`
	Amount               int64                  `json:"amount"`
	ReturnURL            string                 `json:"returnUrl"`
	FailURL              string                 `json:"failUrl,omitempty"`
	Description          string                 `json:"description,omitempty"`
	Language             string                 `json:"language,omitempty"`
	PageView             string                 `json:"pageView,omitempty"`
	ClientID             string                 `json:"clientId,omitempty"`
	MerchantLogin        string                 `json:"merchantLogin,omitempty"`
	JSONParams           map[string]string      `json:"jsonParams,omitempty"`
	SessionTimeoutSecs   int                    `json:"sessionTimeoutSecs,omitempty"`
	ExpirationDate       string                 `json:"expirationDate,omitempty"`
	BindingID            string                 `json:"bindingId,omitempty"`
	Features             string                 `json:"features,omitempty"`
	Email                string                 `json:"email,omitempty"`
	Phone                string                 `json:"phone,omitempty"`
	Currency             string                 `json:"currency,omitempty"`
	OrderBundle          *OrderBundle           `json:"orderBundle,omitempty"`
	TaxSystem            *TaxSystem             `json:"taxSystem,omitempty"`
	AdditionalParameters map[string]string      `json:"additionalParameters,omitempty"`
	DynamicCallbackURL   string                 `json:"dynamicCallbackUrl,omitempty"`
	FeeInput             int64                  `json:"feeInput,omitempty"`
}

// RegisterOrderResponse represents the response from order registration.
type RegisterOrderResponse struct {
	BaseResponse
	OrderID string `json:"orderId,omitempty"`
	FormURL string `json:"formUrl,omitempty"`
}

// OrderBundle represents the shopping cart.
type OrderBundle struct {
	OrderCreationDate string            `json:"orderCreationDate,omitempty"`
	CustomerDetails   *CustomerDetails  `json:"customerDetails,omitempty"`
	CartItems         *CartItems        `json:"cartItems,omitempty"`
}

// CustomerDetails represents customer information.
type CustomerDetails struct {
	Email       string        `json:"email,omitempty"`
	Phone       string        `json:"phone,omitempty"`
	Contact     string        `json:"contact,omitempty"`
	DeliveryInfo *DeliveryInfo `json:"deliveryInfo,omitempty"`
}

// DeliveryInfo represents delivery information.
type DeliveryInfo struct {
	DeliveryType string `json:"deliveryType,omitempty"`
	Country      string `json:"country,omitempty"`
	City         string `json:"city,omitempty"`
	PostAddress  string `json:"postAddress,omitempty"`
}

// CartItems represents shopping cart items.
type CartItems struct {
	Items []Item `json:"items"`
}

// Item represents a single cart item.
type Item struct {
	PositionID      int         `json:"positionId"`
	Name            string      `json:"name"`
	Quantity        *Quantity   `json:"quantity"`
	ItemAmount      int64       `json:"itemAmount"`
	ItemCode        string      `json:"itemCode,omitempty"`
	ItemPrice       int64       `json:"itemPrice,omitempty"`
	Tax             *Tax        `json:"tax,omitempty"`
	ItemDetails     *ItemDetails `json:"itemDetails,omitempty"`
	ItemAttributes  *ItemAttributes `json:"itemAttributes,omitempty"`
	AgentInterest   *AgentInterest `json:"agentInterest,omitempty"`
}

// Quantity represents item quantity.
type Quantity struct {
	Value   float64 `json:"value"`
	Measure string  `json:"measure"`
}

// Tax represents tax information.
type Tax struct {
	TaxType TaxType `json:"taxType"`
	TaxSum  int64   `json:"taxSum,omitempty"`
}

// ItemDetails represents additional item details.
type ItemDetails struct {
	ItemDetailsParams []ItemDetailsParam `json:"itemDetailsParams,omitempty"`
}

// ItemDetailsParam represents a single item detail parameter.
type ItemDetailsParam struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

// ItemAttributes represents item attributes for fiscal operations.
type ItemAttributes struct {
	Attributes []ItemAttribute `json:"attributes,omitempty"`
}

// ItemAttribute represents a single item attribute.
type ItemAttribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// AgentInterest represents agent commission information.
type AgentInterest struct {
	InterestType string  `json:"interestType,omitempty"`
	Interest     float64 `json:"interest,omitempty"`
}

// GetOrderStatusRequest represents a request to get order status.
type GetOrderStatusRequest struct {
	OrderID      string `json:"orderId,omitempty"`
	OrderNumber  string `json:"orderNumber,omitempty"`
	Language     string `json:"language,omitempty"`
	MerchantLogin string `json:"merchantLogin,omitempty"`
}

// GetOrderStatusExtendedResponse represents the extended order status response.
type GetOrderStatusExtendedResponse struct {
	BaseResponse
	OrderNumber            string              `json:"orderNumber,omitempty"`
	OrderStatus            OrderStatus         `json:"orderStatus"`
	ActionCode             int                 `json:"actionCode,omitempty"`
	ActionCodeDescription  string              `json:"actionCodeDescription,omitempty"`
	Amount                 int64               `json:"amount,omitempty"`
	Currency               string              `json:"currency,omitempty"`
	Date                   int64               `json:"date,omitempty"`
	OrderDescription       string              `json:"orderDescription,omitempty"`
	IP                     string              `json:"ip,omitempty"`
	AuthDateTime           int64               `json:"authDateTime,omitempty"`
	AuthRefNum             string              `json:"authRefNum,omitempty"`
	TerminalID             string              `json:"terminalId,omitempty"`
	DepositedDate          int64               `json:"depositedDate,omitempty"`
	RefundedDate           int64               `json:"refundedDate,omitempty"`
	ReversedDate           int64               `json:"reversedDate,omitempty"`
	PaymentWay             string              `json:"paymentWay,omitempty"`
	Chargeback             bool                `json:"chargeback,omitempty"`
	CardAuthInfo           *CardAuthInfo       `json:"cardAuthInfo,omitempty"`
	BindingInfo            *CardBindingInfo    `json:"bindingInfo,omitempty"`
	PaymentAmountInfo      *PaymentAmountInfo  `json:"paymentAmountInfo,omitempty"`
	BankInfo               *BankInfo           `json:"bankInfo,omitempty"`
	PayerData              *PayerData          `json:"payerData,omitempty"`
	Refunds                []Refund            `json:"refunds,omitempty"`
	MerchantOrderParams    []OrderAddendum     `json:"merchantOrderParams,omitempty"`
	Attributes             []OrderAddendum     `json:"attributes,omitempty"`
	TransactionAttributes  []OrderAddendum     `json:"transactionAttributes,omitempty"`
}

// CardAuthInfo represents card authentication information.
type CardAuthInfo struct {
	MaskedPan       string `json:"maskedPan,omitempty"`
	Expiration      string `json:"expiration,omitempty"`
	CardholderName  string `json:"cardholderName,omitempty"`
	ApprovalCode    string `json:"approvalCode,omitempty"`
	Pan             string `json:"pan,omitempty"`
	SecureAuthInfo  *SecureAuthInfo `json:"secureAuthInfo,omitempty"`
}

// SecureAuthInfo represents 3D Secure authentication information.
type SecureAuthInfo struct {
	Eci         int    `json:"eci,omitempty"`
	ThreeDSInfo *ThreeDSInfo `json:"threeDSInfo,omitempty"`
}

// ThreeDSInfo represents 3D Secure specific information.
type ThreeDSInfo struct {
	Xid string `json:"xid,omitempty"`
}

// CardBindingInfo represents card binding information.
type CardBindingInfo struct {
	BindingID    string `json:"bindingId,omitempty"`
	ClientID     string `json:"clientId,omitempty"`
	AuthDateTime string `json:"authDateTime,omitempty"`
	TerminalID   string `json:"terminalId,omitempty"`
}

// PaymentAmountInfo represents payment amount details.
type PaymentAmountInfo struct {
	ApprovedAmount  int64  `json:"approvedAmount,omitempty"`
	DepositedAmount int64  `json:"depositedAmount,omitempty"`
	RefundedAmount  int64  `json:"refundedAmount,omitempty"`
	PaymentState    string `json:"paymentState,omitempty"`
	FeeAmount       int64  `json:"feeAmount,omitempty"`
}

// BankInfo represents bank information.
type BankInfo struct {
	BankName      string `json:"bankName,omitempty"`
	BankCountryCode string `json:"bankCountryCode,omitempty"`
	BankCountryName string `json:"bankCountryName,omitempty"`
}

// PayerData represents payer information.
type PayerData struct {
	Email string `json:"email,omitempty"`
	Phone string `json:"phone,omitempty"`
}

// Refund represents a refund operation.
type Refund struct {
	RefundID      string `json:"refundId,omitempty"`
	RefundDate    int64  `json:"refundDate,omitempty"`
	RefundAmount  int64  `json:"refundAmount,omitempty"`
	RefundItems   []Item `json:"refundItems,omitempty"`
}

// OrderAddendum represents additional order parameters.
type OrderAddendum struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// DepositRequest represents a request to complete a payment.
type DepositRequest struct {
	OrderID       string `json:"orderId"`
	Amount        int64  `json:"amount"`
	Language      string `json:"language,omitempty"`
	JSONParams    string `json:"jsonParams,omitempty"`
	DepositItems  string `json:"depositItems,omitempty"`
	DepositType   int64  `json:"depositType,omitempty"`
	Currency      string `json:"currency,omitempty"`
}

// RefundRequest represents a request to refund a payment.
type RefundRequest struct {
	OrderID      string `json:"orderId"`
	Amount       int64  `json:"amount"`
	Language     string `json:"language,omitempty"`
	JSONParams   string `json:"jsonParams,omitempty"`
	RefundItems  string `json:"refundItems,omitempty"`
}

// ReverseRequest represents a request to reverse a payment.
type ReverseRequest struct {
	OrderID      string `json:"orderId"`
	Language     string `json:"language,omitempty"`
	JSONParams   string `json:"jsonParams,omitempty"`
	Amount       int64  `json:"amount,omitempty"`
}

// BindingRequest represents a request to activate a binding.
type BindingRequest struct {
	BindingID string `json:"bindingId"`
	Language  string `json:"language,omitempty"`
}

// GetBindingsRequest represents a request to get bindings.
type GetBindingsRequest struct {
	ClientID    string `json:"clientId"`
	BindingType string `json:"bindingType,omitempty"`
	ShowExpired string `json:"showExpired,omitempty"`
	Language    string `json:"language,omitempty"`
}

// GetBindingsResponse represents the response with bindings list.
type GetBindingsResponse struct {
	BaseResponse
	Bindings []Binding `json:"bindings,omitempty"`
}

// Binding represents a card binding.
type Binding struct {
	BindingID         string `json:"bindingId,omitempty"`
	MaskedPan         string `json:"maskedPan,omitempty"`
	ExpiryDate        string `json:"expiryDate,omitempty"`
	ClientID          string `json:"clientId,omitempty"`
	BindingCategory   string `json:"bindingCategory,omitempty"`
	IsExpired         bool   `json:"isExpired,omitempty"`
	CardholderName    string `json:"cardholderName,omitempty"`
	PaymentSystem     string `json:"paymentSystem,omitempty"`
	CreatedDate       int64  `json:"createdDate,omitempty"`
	LastUsedDate      int64  `json:"lastUsedDate,omitempty"`
}

// ExtendBindingRequest represents a request to extend binding expiry.
type ExtendBindingRequest struct {
	BindingID string `json:"bindingId"`
	NewExpiry string `json:"newExpiry"` // Format: YYYYMM
	Language  string `json:"language,omitempty"`
}

// UnbindRequest represents a request to deactivate a binding.
type UnbindRequest struct {
	BindingID string `json:"bindingId"`
	Language  string `json:"language,omitempty"`
}

// RecurrentPaymentRequest represents a request for recurrent payment.
type RecurrentPaymentRequest struct {
	OrderNumber          string                 `json:"orderNumber"`
	BindingID            string                 `json:"bindingId"`
	Amount               int64                  `json:"amount"`
	Currency             string                 `json:"currency,omitempty"`
	Language             string                 `json:"language,omitempty"`
	Description          string                 `json:"description,omitempty"`
	ClientID             string                 `json:"clientId,omitempty"`
	PreAuth              bool                   `json:"preAuth,omitempty"`
	AdditionalParameters map[string]string      `json:"additionalParameters,omitempty"`
	OrderBundle          *OrderBundle           `json:"orderBundle,omitempty"`
	TaxSystem            *TaxSystem             `json:"taxSystem,omitempty"`
	DynamicCallbackURL   string                 `json:"dynamicCallbackUrl,omitempty"`
	FeeInput             int64                  `json:"feeInput,omitempty"`
	AutocompletionDate   string                 `json:"autocompletionDate,omitempty"`
	AutoReverseDate      string                 `json:"autoReverseDate,omitempty"`
}

// RecurrentPaymentResponse represents the recurrent payment response.
type RecurrentPaymentResponse struct {
	Success     bool                            `json:"success"`
	Data        *RecurrentPaymentData           `json:"data,omitempty"`
	Error       *RecurrentPaymentError          `json:"error,omitempty"`
	OrderStatus *GetOrderStatusExtendedResponse `json:"orderStatus,omitempty"`
}

// RecurrentPaymentData represents successful recurrent payment data.
type RecurrentPaymentData struct {
	OrderID     string `json:"orderId,omitempty"`
	OrderNumber string `json:"orderNumber,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
}

// RecurrentPaymentError represents a recurrent payment error.
type RecurrentPaymentError struct {
	Code        int    `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Message     string `json:"message,omitempty"`
}

// PaymentOrderBindingRequest represents a request for payment with binding.
type PaymentOrderBindingRequest struct {
	MDOrder   string `json:"mdOrder"`
	BindingID string `json:"bindingId"`
	CVC       string `json:"cvc,omitempty"`
	Language  string `json:"language,omitempty"`
	IP        string `json:"ip,omitempty"`
	Email     string `json:"email,omitempty"`
}

// ApplePayPaymentRequest represents a request for Apple Pay payment.
type ApplePayPaymentRequest struct {
	Merchant             string                 `json:"merchant"`
	OrderNumber          string                 `json:"orderNumber"`
	PaymentToken         string                 `json:"paymentToken"`
	Description          string                 `json:"description,omitempty"`
	Language             string                 `json:"language,omitempty"`
	AdditionalParameters map[string]string      `json:"additionalParameters,omitempty"`
	ClientID             string                 `json:"clientId,omitempty"`
	PreAuth              bool                   `json:"preAuth,omitempty"`
	Email                string                 `json:"email,omitempty"`
	Phone                string                 `json:"phone,omitempty"`
	FailURL              string                 `json:"failUrl,omitempty"`
	PostAddress          string                 `json:"postAddress,omitempty"`
	Amount               int64                  `json:"amount,omitempty"`
	CurrencyCode         string                 `json:"currencyCode,omitempty"`
	IP                   string                 `json:"ip,omitempty"`
	ReturnURL            string                 `json:"returnUrl,omitempty"`
	OrderBundle          *OrderBundle           `json:"orderBundle,omitempty"`
}

// ApplePayPaymentResponse represents the Apple Pay payment response.
type ApplePayPaymentResponse struct {
	Success     bool                            `json:"success"`
	Data        *ApplePayData                   `json:"data,omitempty"`
	Error       *ApplePayError                  `json:"error,omitempty"`
	OrderStatus *GetOrderStatusExtendedResponse `json:"orderStatus,omitempty"`
}

// ApplePayData represents successful Apple Pay payment data.
type ApplePayData struct {
	OrderID string `json:"orderId,omitempty"`
}

// ApplePayError represents an Apple Pay error.
type ApplePayError struct {
	Code        int    `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Message     string `json:"message,omitempty"`
}

// GooglePayRequest represents a request for Google Pay payment.
type GooglePayRequest struct {
	Merchant             string                 `json:"merchant"`
	OrderNumber          string                 `json:"orderNumber"`
	PaymentToken         string                 `json:"paymentToken"`
	Amount               int64                  `json:"amount"`
	IP                   string                 `json:"ip"`
	ReturnURL            string                 `json:"returnUrl"`
	Description          string                 `json:"description,omitempty"`
	Language             string                 `json:"language,omitempty"`
	AdditionalParameters map[string]string      `json:"additionalParameters,omitempty"`
	ClientID             string                 `json:"clientId,omitempty"`
	PreAuth              bool                   `json:"preAuth,omitempty"`
	CurrencyCode         string                 `json:"currencyCode,omitempty"`
	Email                string                 `json:"email,omitempty"`
	Phone                string                 `json:"phone,omitempty"`
	FailURL              string                 `json:"failUrl,omitempty"`
	PostAddress          string                 `json:"postAddress,omitempty"`
	OrderBundle          *OrderBundle           `json:"orderBundle,omitempty"`
}

// GooglePayResponse represents the Google Pay payment response.
type GooglePayResponse struct {
	Success     bool                            `json:"success"`
	Data        *GooglePayData                  `json:"data,omitempty"`
	Error       *GooglePayError                 `json:"error,omitempty"`
	OrderStatus *GetOrderStatusExtendedResponse `json:"orderStatus,omitempty"`
}

// GooglePayData represents successful Google Pay payment data.
type GooglePayData struct {
	OrderID string `json:"orderId,omitempty"`
}

// GooglePayError represents a Google Pay error.
type GooglePayError struct {
	Code        int    `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Message     string `json:"message,omitempty"`
}

// SamsungPayPaymentRequest represents a request for Samsung Pay payment.
type SamsungPayPaymentRequest struct {
	Merchant             string                 `json:"merchant"`
	OrderNumber          string                 `json:"orderNumber"`
	PaymentToken         string                 `json:"paymentToken"`
	Amount               int64                  `json:"amount,omitempty"`
	IP                   string                 `json:"ip,omitempty"`
	ReturnURL            string                 `json:"returnUrl,omitempty"`
	Description          string                 `json:"description,omitempty"`
	Language             string                 `json:"language,omitempty"`
	AdditionalParameters map[string]string      `json:"additionalParameters,omitempty"`
	ClientID             string                 `json:"clientId,omitempty"`
	PreAuth              bool                   `json:"preAuth,omitempty"`
	CurrencyCode         string                 `json:"currencyCode,omitempty"`
	Email                string                 `json:"email,omitempty"`
	Phone                string                 `json:"phone,omitempty"`
	FailURL              string                 `json:"failUrl,omitempty"`
	OrderBundle          *OrderBundle           `json:"orderBundle,omitempty"`
}

// SamsungPayPaymentResponse represents the Samsung Pay payment response.
type SamsungPayPaymentResponse struct {
	Success     bool                            `json:"success"`
	Data        *SamsungPayData                 `json:"data,omitempty"`
	Error       *SamsungPayError                `json:"error,omitempty"`
	OrderStatus *GetOrderStatusExtendedResponse `json:"orderStatus,omitempty"`
}

// SamsungPayData represents successful Samsung Pay payment data.
type SamsungPayData struct {
	OrderID string `json:"orderId,omitempty"`
}

// SamsungPayError represents a Samsung Pay error.
type SamsungPayError struct {
	Code        int    `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Message     string `json:"message,omitempty"`
}

// MirPayPaymentRequest represents a request for MIR Pay payment.
type MirPayPaymentRequest struct {
	Merchant             string                 `json:"merchant"`
	OrderNumber          string                 `json:"orderNumber"`
	Amount               int64                  `json:"amount"`
	ReturnURL            string                 `json:"returnUrl"`
	Description          string                 `json:"description,omitempty"`
	Language             string                 `json:"language,omitempty"`
	AdditionalParameters map[string]string      `json:"additionalParameters,omitempty"`
	ClientID             string                 `json:"clientId,omitempty"`
	PreAuth              bool                   `json:"preAuth,omitempty"`
	CurrencyCode         string                 `json:"currencyCode,omitempty"`
	Email                string                 `json:"email,omitempty"`
	Phone                string                 `json:"phone,omitempty"`
	FailURL              string                 `json:"failUrl,omitempty"`
	OrderBundle          *OrderBundle           `json:"orderBundle,omitempty"`
}

// MirPayResponse represents the MIR Pay payment response.
type MirPayResponse struct {
	Success     bool                            `json:"success"`
	Data        *MirPayData                     `json:"data,omitempty"`
	Error       *MirPayError                    `json:"error,omitempty"`
	OrderStatus *GetOrderStatusExtendedResponse `json:"orderStatus,omitempty"`
}

// MirPayData represents successful MIR Pay payment data.
type MirPayData struct {
	OrderID   string `json:"orderId,omitempty"`
	FormURL   string `json:"formUrl,omitempty"`
	Deeplink  string `json:"deeplink,omitempty"`
}

// MirPayError represents a MIR Pay error.
type MirPayError struct {
	Code        int    `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Message     string `json:"message,omitempty"`
}

// YandexPayRequest represents a request for Yandex Pay payment.
type YandexPayRequest struct {
	Merchant             string                 `json:"merchant"`
	OrderNumber          string                 `json:"orderNumber"`
	PaymentToken         string                 `json:"paymentToken"`
	Amount               int64                  `json:"amount"`
	IP                   string                 `json:"ip,omitempty"`
	ReturnURL            string                 `json:"returnUrl"`
	Description          string                 `json:"description,omitempty"`
	Language             string                 `json:"language,omitempty"`
	AdditionalParameters map[string]string      `json:"additionalParameters,omitempty"`
	ClientID             string                 `json:"clientId,omitempty"`
	PreAuth              bool                   `json:"preAuth,omitempty"`
	CurrencyCode         string                 `json:"currencyCode,omitempty"`
	Email                string                 `json:"email,omitempty"`
	Phone                string                 `json:"phone,omitempty"`
	FailURL              string                 `json:"failUrl,omitempty"`
	OrderBundle          *OrderBundle           `json:"orderBundle,omitempty"`
}

// YandexPayResponse represents the Yandex Pay payment response.
type YandexPayResponse struct {
	Success bool               `json:"success"`
	Data    *YandexPayData     `json:"data,omitempty"`
	Error   *YandexPayError    `json:"error,omitempty"`
}

// YandexPayData represents successful Yandex Pay payment data.
type YandexPayData struct {
	OrderID   string `json:"orderId,omitempty"`
	Redirect  string `json:"redirect,omitempty"`
	AcsURL    string `json:"acsUrl,omitempty"`
	PaReq     string `json:"paReq,omitempty"`
}

// YandexPayError represents a Yandex Pay error.
type YandexPayError struct {
	Code        int    `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Message     string `json:"message,omitempty"`
}

// SBP QR Code types

// SBPGetQRRequest represents a request to get SBP QR code.
type SBPGetQRRequest struct {
	MDOrder     string `json:"mdOrder,omitempty"`
	QRHeight    int    `json:"qrHeight,omitempty"`
	QRWidth     int    `json:"qrWidth,omitempty"`
	QRFormat    string `json:"qrFormat,omitempty"` // image, matrix
}

// SBPGetQRResponse represents the SBP QR code response.
type SBPGetQRResponse struct {
	BaseResponse
	QRImage   string `json:"qrImage,omitempty"`   // Base64 encoded image
	Payload   string `json:"payload,omitempty"`   // QR code payload
	QRURL     string `json:"qrUrl,omitempty"`
}

// SBPQRStatusRequest represents a request to get SBP QR status.
type SBPQRStatusRequest struct {
	MDOrder string `json:"mdOrder"`
}

// SBPQRStatusResponse represents the SBP QR status response.
type SBPQRStatusResponse struct {
	BaseResponse
	OrderID     string      `json:"orderId,omitempty"`
	OrderStatus OrderStatus `json:"orderStatus"`
}

// SBPBindRequest represents a request to create SBP binding.
type SBPBindRequest struct {
	MDOrder string `json:"mdOrder"`
}

// SBPBindResponse represents the SBP binding response.
type SBPBindResponse struct {
	BaseResponse
	BindingID string `json:"bindingId,omitempty"`
}

// SBPUnbindRequest represents a request to remove SBP binding.
type SBPUnbindRequest struct {
	BindingID string `json:"bindingId"`
}

// SBPGetBindingsResponse represents the SBP bindings list response.
type SBPGetBindingsResponse struct {
	BaseResponse
	Bindings []SBPBinding `json:"bindings,omitempty"`
}

// SBPBinding represents an SBP binding.
type SBPBinding struct {
	BindingID   string `json:"bindingId,omitempty"`
	BankName    string `json:"bankName,omitempty"`
	MaskedPhone string `json:"maskedPhone,omitempty"`
	CreatedDate int64  `json:"createdDate,omitempty"`
}

// InstantPaymentRequest represents a request for instant payment (register + pay).
type InstantPaymentRequest struct {
	OrderNumber          string                 `json:"orderNumber"`
	Amount               int64                  `json:"amount"`
	ReturnURL            string                 `json:"returnUrl"`
	FailURL              string                 `json:"failUrl,omitempty"`
	Description          string                 `json:"description,omitempty"`
	Language             string                 `json:"language,omitempty"`
	Email                string                 `json:"email,omitempty"`
	Phone                string                 `json:"phone,omitempty"`
	Currency             string                 `json:"currency,omitempty"`
	BindingID            string                 `json:"bindingId,omitempty"`
	CVC                  string                 `json:"cvc,omitempty"`
	IP                   string                 `json:"ip,omitempty"`
	AdditionalParameters map[string]string      `json:"additionalParameters,omitempty"`
	OrderBundle          *OrderBundle           `json:"orderBundle,omitempty"`
}

// InstantPaymentResponse represents the instant payment response.
type InstantPaymentResponse struct {
	BaseResponse
	OrderID   string                          `json:"orderId,omitempty"`
	FormURL   string                          `json:"formUrl,omitempty"`
	AcsURL    string                          `json:"acsUrl,omitempty"`
	PaReq     string                          `json:"paReq,omitempty"`
	TermURL   string                          `json:"termUrl,omitempty"`
	Redirect  string                          `json:"redirect,omitempty"`
	Info      string                          `json:"info,omitempty"`
}

// DeclineRequest represents a request to decline an order.
type DeclineRequest struct {
	OrderID       string `json:"orderId,omitempty"`
	OrderNumber   string `json:"orderNumber,omitempty"`
	MerchantLogin string `json:"merchantLogin,omitempty"`
	Language      string `json:"language,omitempty"`
}

// VerifyEnrollmentRequest represents a request to verify 3DS enrollment.
type VerifyEnrollmentRequest struct {
	PAN       string `json:"pan"`
	Language  string `json:"language,omitempty"`
}

// VerifyEnrollmentResponse represents the 3DS enrollment verification response.
type VerifyEnrollmentResponse struct {
	BaseResponse
	Enrolled   string `json:"enrolled,omitempty"` // Y, N, U
	EmitterName string `json:"emitterName,omitempty"`
	EmitterCountryCode string `json:"emitterCountryCode,omitempty"`
}

// AddParamsRequest represents a request to add additional parameters to an order.
type AddParamsRequest struct {
	OrderID  string            `json:"orderId"`
	Params   map[string]string `json:"params"`
	Language string            `json:"language,omitempty"`
}

// Finish3DSPaymentRequest represents a request to finish 3DS payment.
type Finish3DSPaymentRequest struct {
	MDOrder string `json:"mdOrder"`
	PaRes   string `json:"paRes"`
}

// PaymentFormResult represents the result of a payment form.
type PaymentFormResult struct {
	BaseResponse
	Redirect   string `json:"redirect,omitempty"`
	AcsURL     string `json:"acsUrl,omitempty"`
	PaReq      string `json:"paReq,omitempty"`
	TermURL    string `json:"termUrl,omitempty"`
	OrderID    string `json:"orderId,omitempty"`
	Info       string `json:"info,omitempty"`
}

// GetLastOrdersRequest represents a request to get last orders for merchants.
type GetLastOrdersRequest struct {
	FromDate      string `json:"from"` // Format: yyyyMMddHHmmss
	ToDate        string `json:"to"`   // Format: yyyyMMddHHmmss
	Page          int    `json:"page,omitempty"`
	Size          int    `json:"size,omitempty"`
	TransactionStates string `json:"transactionStates,omitempty"`
	Merchants     string `json:"merchants,omitempty"`
	Language      string `json:"language,omitempty"`
}

// GetLastOrdersResponse represents the last orders response.
type GetLastOrdersResponse struct {
	BaseResponse
	TotalCount int                              `json:"totalCount,omitempty"`
	Page       int                              `json:"page,omitempty"`
	PageSize   int                              `json:"pageSize,omitempty"`
	Orders     []GetOrderStatusExtendedResponse `json:"orderStatuses,omitempty"`
}

// SBP B2B models

// SBPB2BPayloadResponse represents the B2B SBP payload response.
type SBPB2BPayloadResponse struct {
	BaseResponse
	Payload  string `json:"payload,omitempty"`
	QRURL    string `json:"qrUrl,omitempty"`
	OrderID  string `json:"orderId,omitempty"`
}

// SBPB2BPerformRequest represents a B2B SBP payment request.
type SBPB2BPerformRequest struct {
	OrderID         string              `json:"orderId"`
	Amount          int64               `json:"amount"`
	Currency        string              `json:"currency,omitempty"`
	SenderParams    *SBPB2BSenderParams `json:"senderParams,omitempty"`
	RecipientParams *SBPB2BRecipientParams `json:"recipientParams,omitempty"`
}

// SBPB2BSenderParams represents B2B SBP sender parameters.
type SBPB2BSenderParams struct {
	BankID      string `json:"bankId,omitempty"`
	Account     string `json:"account,omitempty"`
	LegalName   string `json:"legalName,omitempty"`
	INN         string `json:"inn,omitempty"`
	KPP         string `json:"kpp,omitempty"`
}

// SBPB2BRecipientParams represents B2B SBP recipient parameters.
type SBPB2BRecipientParams struct {
	BankID      string `json:"bankId,omitempty"`
	Account     string `json:"account,omitempty"`
	LegalName   string `json:"legalName,omitempty"`
	INN         string `json:"inn,omitempty"`
	KPP         string `json:"kpp,omitempty"`
}

// SBPB2BPerformResponse represents the B2B SBP payment response.
type SBPB2BPerformResponse struct {
	BaseResponse
	OrderID     string `json:"orderId,omitempty"`
	OrderStatus string `json:"orderStatus,omitempty"`
}

// SBP B2C models

// SBPB2CPayoutRequest represents a B2C SBP payout request.
type SBPB2CPayoutRequest struct {
	OrderNumber     string                  `json:"orderNumber"`
	Amount          int64                   `json:"amount"`
	Currency        string                  `json:"currency,omitempty"`
	Purpose         string                  `json:"purpose,omitempty"`
	RecipientParams *SBPB2CRecipientParams  `json:"recipientParams,omitempty"`
	SenderParams    *SBPB2CSenderParams     `json:"senderParams,omitempty"`
}

// SBPB2CRecipientParams represents B2C SBP recipient parameters.
type SBPB2CRecipientParams struct {
	BankID   string `json:"bankId,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Account  string `json:"account,omitempty"`
	Name     *SBPB2CRecipientName `json:"name,omitempty"`
}

// SBPB2CRecipientName represents B2C SBP recipient name.
type SBPB2CRecipientName struct {
	FirstName  string `json:"firstName,omitempty"`
	MiddleName string `json:"middleName,omitempty"`
	LastName   string `json:"lastName,omitempty"`
}

// SBPB2CSenderParams represents B2C SBP sender parameters.
type SBPB2CSenderParams struct {
	BankID    string `json:"bankId,omitempty"`
	Account   string `json:"account,omitempty"`
	LegalName string `json:"legalName,omitempty"`
	INN       string `json:"inn,omitempty"`
	KPP       string `json:"kpp,omitempty"`
}

// SBPB2CPayoutResponse represents the B2C SBP payout response.
type SBPB2CPayoutResponse struct {
	BaseResponse
	OrderID     string `json:"orderId,omitempty"`
	OrderStatus string `json:"orderStatus,omitempty"`
}

// SBPB2CCheckPayoutResponse represents the B2C SBP payout check response.
type SBPB2CCheckPayoutResponse struct {
	BaseResponse
	OrderID     string `json:"orderId,omitempty"`
	OrderStatus string `json:"orderStatus,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
}

// SBPB2CPayoutStatusResponse represents the B2C SBP payout status response.
type SBPB2CPayoutStatusResponse struct {
	BaseResponse
	OrderID     string `json:"orderId,omitempty"`
	OrderStatus string `json:"orderStatus,omitempty"`
	Amount      int64  `json:"amount,omitempty"`
	StatusInfo  *SBPB2CStatusInfo `json:"statusInfo,omitempty"`
}

// SBPB2CStatusInfo represents B2C SBP status information.
type SBPB2CStatusInfo struct {
	Status      string `json:"status,omitempty"`
	Description string `json:"description,omitempty"`
}
