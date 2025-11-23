package alfapay

import (
	"context"
	"net/url"
)

// BindingService handles card binding operations.
type BindingService struct {
	client *Client
}

// GetBindings retrieves all bindings for a client.
func (s *BindingService) GetBindings(ctx context.Context, req *GetBindingsRequest) (*GetBindingsResponse, error) {
	params := url.Values{}
	params.Set("clientId", req.ClientID)

	if req.BindingType != "" {
		params.Set("bindingType", req.BindingType)
	}
	if req.ShowExpired != "" {
		params.Set("showExpired", req.ShowExpired)
	}
	if req.Language != "" {
		params.Set("language", req.Language)
	}

	var resp GetBindingsResponse
	err := s.client.doFormRequest(ctx, "/rest/getBindings.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAllBindings retrieves all bindings including duplicates for a client.
func (s *BindingService) GetAllBindings(ctx context.Context, req *GetBindingsRequest) (*GetBindingsResponse, error) {
	params := url.Values{}
	params.Set("clientId", req.ClientID)

	if req.BindingType != "" {
		params.Set("bindingType", req.BindingType)
	}
	if req.ShowExpired != "" {
		params.Set("showExpired", req.ShowExpired)
	}
	if req.Language != "" {
		params.Set("language", req.Language)
	}

	var resp GetBindingsResponse
	err := s.client.doFormRequest(ctx, "/rest/getAllBindings.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Activate activates a deactivated binding.
func (s *BindingService) Activate(ctx context.Context, req *BindingRequest) (*BaseResponse, error) {
	params := url.Values{}
	params.Set("bindingId", req.BindingID)

	if req.Language != "" {
		params.Set("language", req.Language)
	}

	var resp BaseResponse
	err := s.client.doFormRequest(ctx, "/rest/bindCard.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deactivate deactivates an active binding.
func (s *BindingService) Deactivate(ctx context.Context, req *UnbindRequest) (*BaseResponse, error) {
	params := url.Values{}
	params.Set("bindingId", req.BindingID)

	if req.Language != "" {
		params.Set("language", req.Language)
	}

	var resp BaseResponse
	err := s.client.doFormRequest(ctx, "/rest/unBindCard.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Extend extends the expiry date of a binding.
// NewExpiry format: YYYYMM
func (s *BindingService) Extend(ctx context.Context, req *ExtendBindingRequest) (*BaseResponse, error) {
	params := url.Values{}
	params.Set("bindingId", req.BindingID)
	params.Set("newExpiry", req.NewExpiry)

	if req.Language != "" {
		params.Set("language", req.Language)
	}

	var resp BaseResponse
	err := s.client.doFormRequest(ctx, "/rest/extendBinding.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetByCardOrID retrieves bindings by card or binding ID.
func (s *BindingService) GetByCardOrID(ctx context.Context, bindingID, pan string) (*GetBindingsResponse, error) {
	params := url.Values{}
	if bindingID != "" {
		params.Set("bindingId", bindingID)
	}
	if pan != "" {
		params.Set("pan", pan)
	}

	var resp GetBindingsResponse
	err := s.client.doFormRequest(ctx, "/rest/getBindingsByCardOrId.do", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
