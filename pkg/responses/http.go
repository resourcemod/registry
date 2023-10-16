// Package responses contains default http responses structs
package responses

type Unauthorized struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
