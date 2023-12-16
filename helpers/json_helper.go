package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Helper function to decode JSON request payload
// @summary Decode JSON request payload
// @param r:body Request body containing JSON data
// @param v:body Request payload structure to decode into
// @router /decodejson [post]
func DecodeJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

// Helper function to render JSON response
// @summary Render JSON response
// @param w:body http.ResponseWriter Response writer
// @param status:body int HTTP status code
// @param v:body interface{} Data to be encoded and sent as JSON response
// @router /renderjson [post]
func RenderJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		fmt.Println("Error encoding JSON response:", err)
	}
}

// Helper function to render error response
// @summary Render error response
// @param w:body http.ResponseWriter Response writer
// @param status:body int HTTP status code
// @param message:body string Error message to be included in the response
// @router /rendererror [post]
func RenderError(w http.ResponseWriter, status int, message string) {
	RenderJSON(w, status, map[string]string{"error": message})
}
