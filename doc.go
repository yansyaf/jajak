// Package main jajak API.
//
// Jajak API Documentation
//
//
//     Schemes: http
//     Host: 127.0.0.1:8071
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

// swagger:route GET /ping ping GetPing
// 	responses:
//		200: PingResponse
//		500: ErrorResponse

// swagger:route GET /polls polls GetPolls
// 	responses:
//		200: PollResponse
//		500: ErrorResponse

// swagger:route GET /polls/{ID} polls GetPollById
// 	responses:
//		200: PollResponse
//		500: ErrorResponse

// swagger:parameters GetPollById
type PollPathParameter struct {
	// In: path
	// Required: true
	ID string `json:"id"`
}

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
	// In: body
	Body struct {
		Message string `json:"message"`
	}
}

// Poll response, return list of polling
// swagger:response PollResponse
type PollResponse struct {
	// In: body
	Body struct {
		ID      string   `json:"id"`
		Title   string   `json:"title"`
		Creator string   `json:"creator"`
		Items   []string `json:"items"`
	}
}
