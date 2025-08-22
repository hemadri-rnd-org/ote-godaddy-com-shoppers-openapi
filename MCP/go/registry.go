package main

import (
	"github.com/mcp-server/mcp-server/config"
	"github.com/mcp-server/mcp-server/models"
	tools_v1 "github.com/mcp-server/mcp-server/tools/v1"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_v1.CreateChangepasswordTool(cfg),
		tools_v1.CreateGetstatusTool(cfg),
		tools_v1.CreateCreatesubaccountTool(cfg),
		tools_v1.CreateDeleteTool(cfg),
		tools_v1.CreateGetTool(cfg),
		tools_v1.CreateUpdateTool(cfg),
	}
}
