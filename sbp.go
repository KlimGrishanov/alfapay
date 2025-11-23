package alfapay

import (
	"context"
	"net/url"
	"strconv"
)

// SBPService handles SBP (Система быстрых платежей / Fast Payment System) operations.
type SBPService struct {
	client *Client
}

// GetQR retrieves a QR code for SBP payment.
func (s *SBPService) GetQR(ctx context.Context, req *SBPGetQRRequest) (*SBPGetQRResponse, error) {
	params := url.Values{}
	if req.MDOrder != "" {
		params.Set("mdOrder", req.MDOrder)
	}
	if req.QRHeight > 0 {
		params.Set("qrHeight", strconv.Itoa(req.QRHeight))
	}
	if req.QRWidth > 0 {
		params.Set("qrWidth", strconv.Itoa(req.QRWidth))
	}
	if req.QRFormat != "" {
		params.Set("qrFormat", req.QRFormat)
	}

	var resp SBPGetQRResponse
	err := s.client.doFormRequest(ctx, "/rest/sbp/c2b/qr/dynamic/get.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetQRStatus retrieves the status of an SBP QR payment.
func (s *SBPService) GetQRStatus(ctx context.Context, mdOrder string) (*SBPQRStatusResponse, error) {
	params := url.Values{}
	params.Set("mdOrder", mdOrder)

	var resp SBPQRStatusResponse
	err := s.client.doFormRequest(ctx, "/rest/sbp/c2b/qr/status.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RejectQR rejects/cancels an SBP QR payment.
func (s *SBPService) RejectQR(ctx context.Context, mdOrder string) (*BaseResponse, error) {
	params := url.Values{}
	params.Set("mdOrder", mdOrder)

	var resp BaseResponse
	err := s.client.doFormRequest(ctx, "/rest/sbp/c2b/qr/dynamic/reject.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Bind creates an SBP binding for recurring payments.
func (s *SBPService) Bind(ctx context.Context, mdOrder string) (*SBPBindResponse, error) {
	params := url.Values{}
	params.Set("mdOrder", mdOrder)

	var resp SBPBindResponse
	err := s.client.doFormRequest(ctx, "/rest/sbp/c2b/bind.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Unbind removes an SBP binding.
func (s *SBPService) Unbind(ctx context.Context, bindingID string) (*BaseResponse, error) {
	params := url.Values{}
	params.Set("bindingId", bindingID)

	var resp BaseResponse
	err := s.client.doFormRequest(ctx, "/rest/sbp/c2b/unBind.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBindings retrieves all SBP bindings for a client.
func (s *SBPService) GetBindings(ctx context.Context, clientID string) (*SBPGetBindingsResponse, error) {
	params := url.Values{}
	params.Set("clientId", clientID)

	var resp SBPGetBindingsResponse
	err := s.client.doFormRequest(ctx, "/rest/sbp/c2b/getBindings.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SBP B2B (Business to Business) operations

// B2BGetPayload retrieves payload for B2B SBP payment.
func (s *SBPService) B2BGetPayload(ctx context.Context, orderID string) (*SBPB2BPayloadResponse, error) {
	params := url.Values{}
	params.Set("orderId", orderID)

	var resp SBPB2BPayloadResponse
	err := s.client.doFormRequest(ctx, "/rest/sbp/b2b/getPayload.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// B2BPerform performs a B2B SBP payment.
func (s *SBPService) B2BPerform(ctx context.Context, req *SBPB2BPerformRequest) (*SBPB2BPerformResponse, error) {
	var resp SBPB2BPerformResponse
	err := s.client.doJSONRequest(ctx, "/rest/sbp/b2b/perform.do", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SBP B2C (Business to Consumer) payout operations

// B2CPerformPayout performs a B2C SBP payout.
func (s *SBPService) B2CPerformPayout(ctx context.Context, req *SBPB2CPayoutRequest) (*SBPB2CPayoutResponse, error) {
	var resp SBPB2CPayoutResponse
	err := s.client.doJSONRequest(ctx, "/rest/sbp/b2c/performPayout.do", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// B2CCheckPayout checks the status of a B2C SBP payout.
func (s *SBPService) B2CCheckPayout(ctx context.Context, orderID string) (*SBPB2CCheckPayoutResponse, error) {
	params := url.Values{}
	params.Set("orderId", orderID)

	var resp SBPB2CCheckPayoutResponse
	err := s.client.doFormRequest(ctx, "/rest/sbp/b2c/checkPayout.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// B2CGetPayoutStatus retrieves the status of a B2C SBP payout.
func (s *SBPService) B2CGetPayoutStatus(ctx context.Context, orderID string) (*SBPB2CPayoutStatusResponse, error) {
	params := url.Values{}
	params.Set("orderId", orderID)

	var resp SBPB2CPayoutStatusResponse
	err := s.client.doFormRequest(ctx, "/rest/sbp/b2c/getPayoutStatus.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
