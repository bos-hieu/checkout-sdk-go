package payments

import "github.com/bos-hieu/checkout-sdk-go"

// VoidsRequest ...
type VoidsRequest struct {
	Reference string            `json:"reference,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

// VoidsResponse ...
type VoidsResponse struct {
	StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
	Accepted       *Accepted                `json:"accepted,omitempty"`
}
