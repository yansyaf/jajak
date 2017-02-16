// Package main jajak API.
//
// Jajak API Documentation
//
//
//     Schemes: http
//     Host: 127.0.0.1:8071
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
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

// swagger:route GET /ping ping GetPing
// 	responses:
//		200: PingResponse
//		500: ErrorResponse

// Standard error response returned when got Exception in apps
// swagger:response ErrorResponse
type ErrorResponse struct {
	// in: body
	Body struct {
		Error string `json: "error"`
	}
}

// Standard ping response, return inputted message if provided
// swagger:response PingResponse
type PingResponse struct {
	// in: body
	Body struct {
		Message string `json: "message"`
	}
}
