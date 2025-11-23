package alfapay

import (
	"context"
	"net/url"
	"strconv"
)

// OrderService handles order registration and management operations.
type OrderService struct {
	client *Client
}

// Register registers a new single-stage order.
// The returnURL is required - it's where the customer will be redirected after payment.
func (s *OrderService) Register(ctx context.Context, req *RegisterOrderRequest) (*RegisterOrderResponse, error) {
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
	if req.PageView != "" {
		params.Set("pageView", req.PageView)
	}
	if req.ClientID != "" {
		params.Set("clientId", req.ClientID)
	}
	if req.MerchantLogin != "" {
		params.Set("merchantLogin", req.MerchantLogin)
	}
	if req.SessionTimeoutSecs > 0 {
		params.Set("sessionTimeoutSecs", strconv.Itoa(req.SessionTimeoutSecs))
	}
	if req.ExpirationDate != "" {
		params.Set("expirationDate", req.ExpirationDate)
	}
	if req.BindingID != "" {
		params.Set("bindingId", req.BindingID)
	}
	if req.Features != "" {
		params.Set("features", req.Features)
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
	if req.DynamicCallbackURL != "" {
		params.Set("dynamicCallbackUrl", req.DynamicCallbackURL)
	}
	if req.FeeInput > 0 {
		params.Set("feeInput", strconv.FormatInt(req.FeeInput, 10))
	}
	if req.TaxSystem != nil {
		params.Set("taxSystem", strconv.Itoa(int(*req.TaxSystem)))
	}

	var resp RegisterOrderResponse
	err := s.client.doFormRequest(ctx, "/rest/register.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RegisterPreAuth registers a new two-stage (pre-authorized) order.
// Pre-authorization holds funds on customer's account without actual charge.
// Use Deposit to complete the payment later.
func (s *OrderService) RegisterPreAuth(ctx context.Context, req *RegisterOrderRequest) (*RegisterOrderResponse, error) {
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
	if req.PageView != "" {
		params.Set("pageView", req.PageView)
	}
	if req.ClientID != "" {
		params.Set("clientId", req.ClientID)
	}
	if req.MerchantLogin != "" {
		params.Set("merchantLogin", req.MerchantLogin)
	}
	if req.SessionTimeoutSecs > 0 {
		params.Set("sessionTimeoutSecs", strconv.Itoa(req.SessionTimeoutSecs))
	}
	if req.ExpirationDate != "" {
		params.Set("expirationDate", req.ExpirationDate)
	}
	if req.BindingID != "" {
		params.Set("bindingId", req.BindingID)
	}
	if req.Features != "" {
		params.Set("features", req.Features)
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
	if req.DynamicCallbackURL != "" {
		params.Set("dynamicCallbackUrl", req.DynamicCallbackURL)
	}
	if req.FeeInput > 0 {
		params.Set("feeInput", strconv.FormatInt(req.FeeInput, 10))
	}
	if req.TaxSystem != nil {
		params.Set("taxSystem", strconv.Itoa(int(*req.TaxSystem)))
	}

	var resp RegisterOrderResponse
	err := s.client.doFormRequest(ctx, "/rest/registerPreAuth.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Decline cancels an unpaid order.
func (s *OrderService) Decline(ctx context.Context, req *DeclineRequest) (*BaseResponse, error) {
	params := url.Values{}
	if req.OrderID != "" {
		params.Set("orderId", req.OrderID)
	}
	if req.OrderNumber != "" {
		params.Set("orderNumber", req.OrderNumber)
	}
	if req.MerchantLogin != "" {
		params.Set("merchantLogin", req.MerchantLogin)
	}
	if req.Language != "" {
		params.Set("language", req.Language)
	}

	var resp BaseResponse
	err := s.client.doFormRequest(ctx, "/rest/decline.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddParams adds additional parameters to an existing order.
func (s *OrderService) AddParams(ctx context.Context, req *AddParamsRequest) (*BaseResponse, error) {
	params := url.Values{}
	params.Set("orderId", req.OrderID)

	// Serialize params as JSON
	if len(req.Params) > 0 {
		paramsJSON := "{"
		first := true
		for k, v := range req.Params {
			if !first {
				paramsJSON += ","
			}
			paramsJSON += `"` + k + `":"` + v + `"`
			first = false
		}
		paramsJSON += "}"
		params.Set("params", paramsJSON)
	}

	if req.Language != "" {
		params.Set("language", req.Language)
	}

	var resp BaseResponse
	err := s.client.doFormRequest(ctx, "/rest/addParams.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
