package alfapay

import (
	"context"
	"net/url"
	"strconv"
)

// PaymentService handles payment operations.
type PaymentService struct {
	client *Client
}

// Deposit completes a pre-authorized payment.
// Amount can be less than or equal to the pre-authorized amount, but not less than 1 ruble.
func (s *PaymentService) Deposit(ctx context.Context, req *DepositRequest) (*BaseResponse, error) {
	params := url.Values{}
	params.Set("orderId", req.OrderID)
	params.Set("amount", strconv.FormatInt(req.Amount, 10))

	if req.Language != "" {
		params.Set("language", req.Language)
	}
	if req.JSONParams != "" {
		params.Set("jsonParams", req.JSONParams)
	}
	if req.DepositItems != "" {
		params.Set("depositItems", req.DepositItems)
	}
	if req.DepositType > 0 {
		params.Set("depositType", strconv.FormatInt(req.DepositType, 10))
	}
	if req.Currency != "" {
		params.Set("currency", req.Currency)
	}

	var resp BaseResponse
	err := s.client.doFormRequest(ctx, "/rest/deposit.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Reverse cancels an authorized payment (before settlement).
func (s *PaymentService) Reverse(ctx context.Context, req *ReverseRequest) (*BaseResponse, error) {
	params := url.Values{}
	params.Set("orderId", req.OrderID)

	if req.Language != "" {
		params.Set("language", req.Language)
	}
	if req.JSONParams != "" {
		params.Set("jsonParams", req.JSONParams)
	}
	if req.Amount > 0 {
		params.Set("amount", strconv.FormatInt(req.Amount, 10))
	}

	var resp BaseResponse
	err := s.client.doFormRequest(ctx, "/rest/reverse.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PayWithBinding performs payment using a saved card binding.
func (s *PaymentService) PayWithBinding(ctx context.Context, req *PaymentOrderBindingRequest) (*PaymentFormResult, error) {
	params := url.Values{}
	params.Set("mdOrder", req.MDOrder)
	params.Set("bindingId", req.BindingID)

	if req.CVC != "" {
		params.Set("cvc", req.CVC)
	}
	if req.Language != "" {
		params.Set("language", req.Language)
	}
	if req.IP != "" {
		params.Set("ip", req.IP)
	}
	if req.Email != "" {
		params.Set("email", req.Email)
	}

	var resp PaymentFormResult
	err := s.client.doFormRequest(ctx, "/rest/paymentOrderBinding.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Instant registers an order and initiates payment in a single request.
func (s *PaymentService) Instant(ctx context.Context, req *InstantPaymentRequest) (*InstantPaymentResponse, error) {
	params := url.Values{}
	params.Set("orderNumber", req.OrderNumber)
	params.Set("amount", strconv.FormatInt(req.Amount, 10))
	params.Set("returnUrl", req.ReturnURL)

	if req.FailURL != "" {
		params.Set("failUrl", req.FailURL)
	}
	if req.Description != "" {
		params.Set("description", req.Description)
	}
	if req.Language != "" {
		params.Set("language", req.Language)
	}
	if req.Email != "" {
		params.Set("email", req.Email)
	}
	if req.Phone != "" {
		params.Set("phone", req.Phone)
	}
	if req.Currency != "" {
		params.Set("currency", req.Currency)
	}
	if req.BindingID != "" {
		params.Set("bindingId", req.BindingID)
	}
	if req.CVC != "" {
		params.Set("cvc", req.CVC)
	}
	if req.IP != "" {
		params.Set("ip", req.IP)
	}

	var resp InstantPaymentResponse
	err := s.client.doFormRequest(ctx, "/rest/instantPayment.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Recurrent performs a recurrent (recurring) payment using a binding.
func (s *PaymentService) Recurrent(ctx context.Context, req *RecurrentPaymentRequest) (*RecurrentPaymentResponse, error) {
	// Add credentials to the request body for JSON API
	type recurrentReqWithAuth struct {
		*RecurrentPaymentRequest
		UserName string `json:"userName"`
		Password string `json:"password"`
	}

	reqBody := &recurrentReqWithAuth{
		RecurrentPaymentRequest: req,
		UserName:                s.client.userName,
		Password:                s.client.password,
	}

	var resp RecurrentPaymentResponse
	// Use direct JSON request without query auth params
	err := s.client.doJSONRequestNoAuth(ctx, "/recurrentPayment.do", reqBody, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Finish3DS completes a 3D Secure payment after customer authentication.
func (s *PaymentService) Finish3DS(ctx context.Context, req *Finish3DSPaymentRequest) (*PaymentFormResult, error) {
	params := url.Values{}
	params.Set("mdOrder", req.MDOrder)
	params.Set("paRes", req.PaRes)

	var resp PaymentFormResult
	err := s.client.doFormRequest(ctx, "/rest/finish3dsPayment.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
