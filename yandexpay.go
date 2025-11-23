package alfapay

import "context"

// YandexPayService handles Yandex Pay payment operations.
type YandexPayService struct {
	client *Client
}

// Payment performs a payment using Yandex Pay token.
func (s *YandexPayService) Payment(ctx context.Context, req *YandexPayRequest) (*YandexPayResponse, error) {
	var resp YandexPayResponse
	err := s.client.doJSONRequestNoAuth(ctx, "/yandex/payment.do", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DirectPayment performs a direct payment using Yandex Pay without pre-registration.
func (s *YandexPayService) DirectPayment(ctx context.Context, req *YandexPayRequest) (*YandexPayResponse, error) {
	var resp YandexPayResponse
	err := s.client.doJSONRequestNoAuth(ctx, "/yandex/paymentDirect.do", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// InstantPayment performs instant payment with Yandex Pay (register + pay).
func (s *YandexPayService) InstantPayment(ctx context.Context, req *YandexPayRequest) (*YandexPayResponse, error) {
	var resp YandexPayResponse
	err := s.client.doJSONRequestNoAuth(ctx, "/yandex/instantPayment.do", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
