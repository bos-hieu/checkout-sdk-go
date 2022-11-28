package customers

import "github.com/bos-hieu/checkout-sdk-go"

type (
	// Request -
	Request struct {
		*Customer
	}
	// Customer -
	Customer struct {
		Email   string `json:"email,omitempty"`
		Name    string `json:"name,omitempty"`
		Default string `json:"default,omitempty"`
	}
)

type (
	// Response -
	Response struct {
		StatusResponse *checkout.StatusResponse `json:"api_response,omitempty"`
	}
)
