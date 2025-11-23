package alfapay

import "context"

// ApplePayService handles Apple Pay payment operations.
type ApplePayService struct {
	client *Client
}

// Payment performs a payment using Apple Pay token.
func (s *ApplePayService) Payment(ctx context.Context, req *ApplePayPaymentRequest) (*ApplePayPaymentResponse, error) {
	var resp ApplePayPaymentResponse
	err := s.client.doJSONRequestNoAuth(ctx, "/applepay/payment.do", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
