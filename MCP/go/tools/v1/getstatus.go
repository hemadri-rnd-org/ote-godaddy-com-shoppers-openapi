package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mcp-server/mcp-server/config"
	"github.com/mcp-server/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func GetstatusHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		shopperIdVal, ok := args["shopperId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: shopperId"), nil
		}
		shopperId, ok := shopperIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: shopperId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["auditClientIp"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("auditClientIp=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v1/shoppers/%s/status%s", cfg.BaseURL, shopperId, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.ShopperStatus
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateGetstatusTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v1_shoppers_shopperId_status",
		mcp.WithDescription("Get details for the specified Shopper"),
		mcp.WithString("shopperId", mcp.Required(), mcp.Description("The ID of the shopper to retrieve. Must agree with the shopper id on the token or header, if present")),
		mcp.WithString("auditClientIp", mcp.Required(), mcp.Description("The client IP of the user who originated the request leading to this call.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    GetstatusHandler(cfg),
	}
}
