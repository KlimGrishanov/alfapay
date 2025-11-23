package alfapay

import "context"

// SamsungPayService handles Samsung Pay payment operations.
type SamsungPayService struct {
	client *Client
}

// Payment performs a payment using Samsung Pay token.
func (s *SamsungPayService) Payment(ctx context.Context, req *SamsungPayPaymentRequest) (*SamsungPayPaymentResponse, error) {
	var resp SamsungPayPaymentResponse
	err := s.client.doJSONRequestNoAuth(ctx, "/samsung/payment.do", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DirectPayment performs a direct payment using Samsung Pay without pre-registration.
func (s *SamsungPayService) DirectPayment(ctx context.Context, req *SamsungPayPaymentRequest) (*SamsungPayPaymentResponse, error) {
	var resp SamsungPayPaymentResponse
	err := s.client.doJSONRequestNoAuth(ctx, "/samsung/paymentDirect.do", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
