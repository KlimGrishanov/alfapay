package alfapay

import "context"

// MirPayService handles MIR Pay payment operations.
type MirPayService struct {
	client *Client
}

// Payment performs a payment using MIR Pay.
func (s *MirPayService) Payment(ctx context.Context, req *MirPayPaymentRequest) (*MirPayResponse, error) {
	var resp MirPayResponse
	err := s.client.doJSONRequestNoAuth(ctx, "/mir/payment.do", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DirectPayment performs a direct payment using MIR Pay without pre-registration.
func (s *MirPayService) DirectPayment(ctx context.Context, req *MirPayPaymentRequest) (*MirPayResponse, error) {
	var resp MirPayResponse
	err := s.client.doJSONRequestNoAuth(ctx, "/mir/paymentDirect.do", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
