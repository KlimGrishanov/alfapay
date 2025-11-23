package alfapay

import (
	"context"
	"net/url"
	"strconv"
)

// StatusService handles order status operations.
type StatusService struct {
	client *Client
}

// GetExtended retrieves extended order status information.
// Either orderID or orderNumber must be provided.
func (s *StatusService) GetExtended(ctx context.Context, req *GetOrderStatusRequest) (*GetOrderStatusExtendedResponse, error) {
	params := url.Values{}
	if req.OrderID != "" {
		params.Set("orderId", req.OrderID)
	}
	if req.OrderNumber != "" {
		params.Set("orderNumber", req.OrderNumber)
	}
	if req.Language != "" {
		params.Set("language", req.Language)
	}
	if req.MerchantLogin != "" {
		params.Set("merchantLogin", req.MerchantLogin)
	}

	var resp GetOrderStatusExtendedResponse
	err := s.client.doFormRequest(ctx, "/rest/getOrderStatusExtended.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetByOrderID retrieves extended order status by order ID.
func (s *StatusService) GetByOrderID(ctx context.Context, orderID string) (*GetOrderStatusExtendedResponse, error) {
	return s.GetExtended(ctx, &GetOrderStatusRequest{OrderID: orderID})
}

// GetByOrderNumber retrieves extended order status by order number.
func (s *StatusService) GetByOrderNumber(ctx context.Context, orderNumber string) (*GetOrderStatusExtendedResponse, error) {
	return s.GetExtended(ctx, &GetOrderStatusRequest{OrderNumber: orderNumber})
}

// GetLastOrders retrieves orders for a date range.
// Date format: yyyyMMddHHmmss
func (s *StatusService) GetLastOrders(ctx context.Context, req *GetLastOrdersRequest) (*GetLastOrdersResponse, error) {
	params := url.Values{}
	params.Set("from", req.FromDate)
	params.Set("to", req.ToDate)

	if req.Page > 0 {
		params.Set("page", strconv.Itoa(req.Page))
	}
	if req.Size > 0 {
		params.Set("size", strconv.Itoa(req.Size))
	}
	if req.TransactionStates != "" {
		params.Set("transactionStates", req.TransactionStates)
	}
	if req.Merchants != "" {
		params.Set("merchants", req.Merchants)
	}
	if req.Language != "" {
		params.Set("language", req.Language)
	}

	var resp GetLastOrdersResponse
	err := s.client.doFormRequest(ctx, "/rest/getLastOrdersForMerchants.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// VerifyEnrollment checks if a card is enrolled in 3D Secure.
func (s *StatusService) VerifyEnrollment(ctx context.Context, req *VerifyEnrollmentRequest) (*VerifyEnrollmentResponse, error) {
	params := url.Values{}
	params.Set("pan", req.PAN)

	if req.Language != "" {
		params.Set("language", req.Language)
	}

	var resp VerifyEnrollmentResponse
	err := s.client.doFormRequest(ctx, "/rest/verifyEnrollment.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
