package alfapay

import (
	"context"
	"net/url"
	"strconv"
)

// RefundService handles refund operations.
type RefundService struct {
	client *Client
}

// Refund performs a refund for a completed payment.
// Amount must be less than or equal to the deposited amount.
func (s *RefundService) Refund(ctx context.Context, req *RefundRequest) (*BaseResponse, error) {
	params := url.Values{}
	params.Set("orderId", req.OrderID)
	params.Set("amount", strconv.FormatInt(req.Amount, 10))

	if req.Language != "" {
		params.Set("language", req.Language)
	}
	if req.JSONParams != "" {
		params.Set("jsonParams", req.JSONParams)
	}
	if req.RefundItems != "" {
		params.Set("refundItems", req.RefundItems)
	}

	var resp BaseResponse
	err := s.client.doFormRequest(ctx, "/rest/refund.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// InstantRefund performs an instant refund without requiring an orderId.
// This is used for quick refunds when you have the necessary payment details.
func (s *RefundService) InstantRefund(ctx context.Context, orderID string, amount int64) (*BaseResponse, error) {
	params := url.Values{}
	params.Set("orderId", orderID)
	params.Set("amount", strconv.FormatInt(amount, 10))

	var resp BaseResponse
	err := s.client.doFormRequest(ctx, "/rest/instantRefund.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
