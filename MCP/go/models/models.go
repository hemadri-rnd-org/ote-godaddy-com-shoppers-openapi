package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// SubaccountCreate represents the SubaccountCreate schema from the OpenAPI specification
type SubaccountCreate struct {
	Marketid string `json:"marketId,omitempty"`
	Namefirst string `json:"nameFirst"`
	Namelast string `json:"nameLast"`
	Password string `json:"password"`
	Email string `json:"email"`
	Externalid int `json:"externalId,omitempty"`
}

// ShopperUpdate represents the ShopperUpdate schema from the OpenAPI specification
type ShopperUpdate struct {
	Externalid int `json:"externalId,omitempty"`
	Marketid string `json:"marketId,omitempty"`
	Namefirst string `json:"nameFirst,omitempty"`
	Namelast string `json:"nameLast,omitempty"`
	Email string `json:"email,omitempty"`
}

// ErrorField represents the ErrorField schema from the OpenAPI specification
type ErrorField struct {
	Code string `json:"code"` // Short identifier for the error, suitable for indicating the specific error within client code
	Message string `json:"message,omitempty"` // Human-readable, English description of the problem with the contents of the field
	Path string `json:"path"` // <ul> <li style='margin-left: 12px;'>JSONPath referring to a field containing an error</li> <strong style='margin-left: 12px;'>OR</strong> <li style='margin-left: 12px;'>JSONPath referring to a field that refers to an object containing an error, with more detail in `pathRelated`</li> </ul>
	Pathrelated string `json:"pathRelated,omitempty"` // JSONPath referring to a field containing an error, which is referenced by `path`
}

// PasswordError represents the PasswordError schema from the OpenAPI specification
type PasswordError struct {
	Code string `json:"code,omitempty"` // Short identifier for the error, suitable for indicating the specific error within client code
	Message string `json:"message,omitempty"` // Human-readable, English description of the error
	TypeField string `json:"type,omitempty"` // Response type, always 'error'
}

// Shopper represents the Shopper schema from the OpenAPI specification
type Shopper struct {
	Shopperid string `json:"shopperId"`
	Customerid string `json:"customerId,omitempty"` // Identifier for the Customer record associated with this Shopper record. This is an alternate identifier that some systems use to identify an individual shopper record
	Email string `json:"email"`
	Externalid int `json:"externalId,omitempty"`
	Marketid string `json:"marketId"`
	Namefirst string `json:"nameFirst"`
	Namelast string `json:"nameLast"`
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Code string `json:"code"` // Short identifier for the error, suitable for indicating the specific error within client code
	Fields []ErrorField `json:"fields,omitempty"` // List of the specific fields, and the errors found with their contents
	Message string `json:"message,omitempty"` // Human-readable, English description of the error
}

// ErrorLimit represents the ErrorLimit schema from the OpenAPI specification
type ErrorLimit struct {
	Fields []ErrorField `json:"fields,omitempty"` // List of the specific fields, and the errors found with their contents
	Message string `json:"message,omitempty"` // Human-readable, English description of the error
	Retryaftersec int `json:"retryAfterSec"` // Number of seconds to wait before attempting a similar request
	Code string `json:"code"` // Short identifier for the error, suitable for indicating the specific error within client code
}

// Secret represents the Secret schema from the OpenAPI specification
type Secret struct {
	Secret string `json:"secret,omitempty"` // The secret value used to set a subaccount's password
}

// ShopperId represents the ShopperId schema from the OpenAPI specification
type ShopperId struct {
	Customerid string `json:"customerId,omitempty"` // Identifier for the Customer record associated with this Shopper record. This is an alternate identifier that some systems use to identify an individual shopper record
	Shopperid string `json:"shopperId"`
}

// ShopperStatus represents the ShopperStatus schema from the OpenAPI specification
type ShopperStatus struct {
	Billingstate string `json:"billingState,omitempty"` // Indicates the billing state of the Shopper.<br />ABANDONED: The shopper has not been billed in at least 10 years and has no active subscriptions.<br />INACTIVE: The shopper has been billed within the last 10 years but has no active subscriptions.<br />ACTIVE: The shopper has at least one active subscription.
}
