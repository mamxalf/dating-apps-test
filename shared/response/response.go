package response

import (
	"dating-apps/shared/failure"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Failure struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Base struct {
	Data    *any    `json:"data,omitempty"`
	Error   *string `json:"error,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NoContent sends a response without any content
func NoContent(w http.ResponseWriter) {
	respond(w, http.StatusNoContent, nil)
}

// WithMessage sends a response with a simple text message
func WithMessage(w http.ResponseWriter, code int, message string) {
	respond(w, code, Base{Message: &message})
}

// WithJSON sends a response containing a JSON object
func WithJSON(w http.ResponseWriter, code int, jsonPayload any) {
	respond(w, code, Base{Data: &jsonPayload})
}

// WithError sends a response with an error message
func WithError(w http.ResponseWriter, err error) {
	code := failure.GetCode(err)
	errMsg := err.Error()
	respond(w, code, Base{Error: &errMsg})
}

// WithPreparingShutdown sends a default response for when the server is preparing to shut down
func WithPreparingShutdown(w http.ResponseWriter) {
	WithMessage(w, http.StatusServiceUnavailable, "SERVER PREPARING TO SHUT DOWN")
}

// WithUnhealthy sends a default response for when the server is unhealthy
func WithUnhealthy(w http.ResponseWriter) {
	WithMessage(w, http.StatusServiceUnavailable, "SERVER UNHEALTHY")
}

func (e *Failure) Error() string {
	return fmt.Sprintf("%s: %s", http.StatusText(e.Code), e.Message)
}

func respond(w http.ResponseWriter, code int, payload any) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed Error!")
	}
}
