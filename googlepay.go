package alfapay

import "context"

// GooglePayService handles Google Pay payment operations.
type GooglePayService struct {
	client *Client
}

// Payment performs a payment using Google Pay token.
func (s *GooglePayService) Payment(ctx context.Context, req *GooglePayRequest) (*GooglePayResponse, error) {
	var resp GooglePayResponse
	err := s.client.doJSONRequestNoAuth(ctx, "/google/payment.do", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
