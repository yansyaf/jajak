// Package main jajak API.
//
// Jajak API Documentation
//
//
//     Schemes: http
//     Host: 128.199.91.172:8071
//     BasePath: /
//     Version: 0.0.1
//     Contact: Artiko W <artikow@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta
package main

import (
	"github.com/toshim45/jajak/survey"
)

// swagger:route GET /ping ping GetPing
// 	responses:
//		200: PingResponse
//		500: ErrorResponse

// swagger:route GET /surveys surveys GetSurveys
// 	responses:
//		200: SurveyResponse
//		500: ErrorResponse

// swagger:route GET /surveys/{SurveyID} surveys GetSurveyById
// 	responses:
//		200: SurveyResponse
//		500: ErrorResponse

// swagger:route POST /surveys surveys StoreSurvey
// 	responses:
//		200: SurveyResponse
//		500: ErrorResponse

// swagger:route POST /surveys/{SurveyID}/poll surveys StorePoll
// 	responses:
//		201: NoResponse
//		500: ErrorResponse

// swagger:parameters GetSurveyById
type SurveyPathParameter struct {
	// In: path
	// Required: true
	SurveyID string
}

// swagger:parameters StoreSurvey
type StoreSurveyParameter struct {
	// In: body
	Body survey.Survey
}

// swagger:parameters StorePoll
type StorePollParameter struct {
	// In: path
	// Required: true
	SurveyID string
	// In: body
	Body survey.Poll
}

// Standard error response returned when got Exception in apps
// swagger:response ErrorResponse
type ErrorResponse struct {
	// In: body
	Body struct {
		Error string `json: "error"`
	}
}

// swagger:response NoResponse
type NoResponse struct{}

// Standard ping response, return inputted message if provided
// swagger:response PingResponse
type PingResponse struct {
	// In: body
	Body struct {
		Message string `json:"message"`
	}
}

// Poll response, return list of polling
// swagger:response SurveyResponse
type SurveyResponse struct {
	// In: body
	Body survey.Survey
}
