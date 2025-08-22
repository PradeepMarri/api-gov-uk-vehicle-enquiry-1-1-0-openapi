package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Vehicle represents the Vehicle schema from the OpenAPI specification
type Vehicle struct {
	Colour string `json:"colour,omitempty"` // Vehicle colour
	Motstatus string `json:"motStatus,omitempty"` // MOT Status of the vehicle
	Dateoflastv5cissued string `json:"dateOfLastV5CIssued,omitempty"` // Date of last V5C issued
	Registrationnumber string `json:"registrationNumber"` // Registration number of the vehicle
	Taxstatus string `json:"taxStatus,omitempty"` // Tax status of the vehicle
	Co2emissions int `json:"co2Emissions,omitempty"` // Carbon Dioxide emissions in grams per kilometre
	Monthoffirstdvlaregistration string `json:"monthOfFirstDvlaRegistration,omitempty"` // Month of First DVLA Registration
	Eurostatus string `json:"euroStatus,omitempty"` // Euro Status (Dealer / Customer Provided (new vehicles))
	Realdrivingemissions string `json:"realDrivingEmissions,omitempty"` // Real Driving Emissions value
	Revenueweight int `json:"revenueWeight,omitempty"` // Revenue weight in kilograms
	Fueltype string `json:"fuelType,omitempty"` // Fuel type (Method of Propulsion)
	Markedforexport bool `json:"markedForExport,omitempty"` // True only if vehicle has been export marked
	Monthoffirstregistration string `json:"monthOfFirstRegistration,omitempty"` // Month of First Registration
	Wheelplan string `json:"wheelplan,omitempty"` // Vehicle wheel plan
	Artenddate string `json:"artEndDate,omitempty"` // Additional Rate of Tax End Date, format: YYYY-MM-DD
	Enginecapacity int `json:"engineCapacity,omitempty"` // Engine capacity in cubic centimetres
	Taxduedate string `json:"taxDueDate,omitempty"` // Date of tax liablity, Used in calculating licence information presented to user
	Motexpirydate string `json:"motExpiryDate,omitempty"` // Mot Expiry Date
	MakeField string `json:"make,omitempty"` // Vehicle make
	Yearofmanufacture int `json:"yearOfManufacture,omitempty"` // Year of Manufacture
	Typeapproval string `json:"typeApproval,omitempty"` // Vehicle Type Approval Category
}

// VehicleRequest represents the VehicleRequest schema from the OpenAPI specification
type VehicleRequest struct {
	Registrationnumber string `json:"registrationNumber,omitempty"`
}

// ErrorResponse represents the ErrorResponse schema from the OpenAPI specification
type ErrorResponse struct {
	Errors []Errors `json:"errors,omitempty"`
}

// Errors represents the Errors schema from the OpenAPI specification
type Errors struct {
	Detail string `json:"detail,omitempty"` // A meaningful description of the error which has occurred
	Status string `json:"status,omitempty"`
	Title string `json:"title"` // Error title
	Code string `json:"code,omitempty"` // DVLA reference code
}
