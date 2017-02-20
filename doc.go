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

// Poll response, return list of polling
// swagger:response PollResponse
type PollResponse struct {
	// in: body
	Body struct {
		Title   string   `db:"title" json:"title"`
		Creator string   `db:"creator" json:"creator"`
		Items   []string `db:"items" json:"items"`
	}
}
