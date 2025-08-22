package main

import (
	"github.com/vehicle-enquiry-api/mcp-server/config"
	"github.com/vehicle-enquiry-api/mcp-server/models"
	tools_vehicle "github.com/vehicle-enquiry-api/mcp-server/tools/vehicle"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_vehicle.CreateGetvehicledetailsbyregistrationnumberTool(cfg),
	}
}
