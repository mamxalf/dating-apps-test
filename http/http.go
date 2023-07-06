package http

import (
	"dating-apps/configs"
	"dating-apps/docs"
	router "dating-apps/http/routers"
	"dating-apps/infras"
	"dating-apps/shared/logger"
	"dating-apps/shared/response"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog/log"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"strings"
)

// HTTP is the HTTP server.
type HTTP struct {
	Config *configs.Config
	DB     *infras.PostgresConn
	Router router.Router
	mux    *chi.Mux
}

// ProvideHTTP is the provider for HTTP.
func ProvideHTTP(db *infras.PostgresConn, config *configs.Config, router router.Router) *HTTP {
	return &HTTP{
		DB:     db,
		Config: config,
		Router: router,
	}
}

// SetupAndServe sets up the server and gets it up and running.
func (h *HTTP) SetupAndServe() {
	h.mux = chi.NewRouter()
	h.setupMiddleware()
	h.setupSwaggerDocs()
	h.setupRoutes()

	h.logServerInfo()

	log.Info().Str("port", h.Config.Server.Port).Msg("Starting up HTTP server.")

	err := http.ListenAndServe(":"+h.Config.Server.Port, h.mux)
	if err != nil {
		log.Err(err)
	}
}

func (h *HTTP) setupSwaggerDocs() {
	if h.Config.Server.Env == "development" {
		docs.SwaggerInfo.Title = h.Config.App.Name
		docs.SwaggerInfo.Version = h.Config.App.Revision
		swaggerURL := fmt.Sprintf("%s/swagger/doc.json", h.Config.App.URL)
		h.mux.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(swaggerURL)))
		log.Info().Str("url", swaggerURL).Msg("Swagger documentation enabled.")
	}
}

func (h *HTTP) setupRoutes() {
	h.mux.Get("/health", h.HealthCheck)
	h.Router.SetupRoutes(h.mux)
}

func (h *HTTP) setupMiddleware() {
	h.mux.Use(middleware.Logger)
	h.mux.Use(middleware.Recoverer)
	h.setupCORS()
}

func (h *HTTP) logServerInfo() {
	h.logCORSConfigInfo()
}

func (h *HTTP) logCORSConfigInfo() {
	corsConfig := h.Config.App.CORS
	corsHeaderInfo := "CORS Header"
	if corsConfig.Enable {
		log.Info().Msg("CORS Headers and Handlers are enabled.")
		log.Info().Str(corsHeaderInfo, fmt.Sprintf("Access-Control-Allow-Credentials: %t", corsConfig.AllowCredentials)).Msg("")
		log.Info().Str(corsHeaderInfo, fmt.Sprintf("Access-Control-Allow-Headers: %s", strings.Join(corsConfig.AllowedHeaders, ", "))).Msg("")
		log.Info().Str(corsHeaderInfo, fmt.Sprintf("Access-Control-Allow-Methods: %s", strings.Join(corsConfig.AllowedMethods, ", "))).Msg("")
		log.Info().Str(corsHeaderInfo, fmt.Sprintf("Access-Control-Allow-Origin: %s", strings.Join(corsConfig.AllowedOrigins, ", "))).Msg("")
		log.Info().Str(corsHeaderInfo, fmt.Sprintf("Access-Control-Max-Age: %d", corsConfig.MaxAgeSeconds)).Msg("")
	} else {
		log.Info().Msg("CORS Headers are disabled.")
	}
}

func (h *HTTP) setupCORS() {
	corsConfig := h.Config.App.CORS
	if corsConfig.Enable {
		h.mux.Use(cors.Handler(cors.Options{
			AllowCredentials: corsConfig.AllowCredentials,
			AllowedHeaders:   corsConfig.AllowedHeaders,
			AllowedMethods:   corsConfig.AllowedMethods,
			AllowedOrigins:   corsConfig.AllowedOrigins,
			MaxAge:           corsConfig.MaxAgeSeconds,
		}))
	}
}

// HealthCheck performs a health check on the server. Usually required by
// Kubernetes to check if the service is healthy.
// @Summary Health Check
// @Description Health Check Endpoint
// @Tags service
// @Produce json
// @Accept json
// @Success 200 {object} response.Base
// @Failure 503 {object} response.Base
// @Router /health [get]
func (h *HTTP) HealthCheck(w http.ResponseWriter, r *http.Request) {
	if err := h.DB.Read.Ping(); err != nil {
		logger.ErrorWithStack(err)
		response.WithUnhealthy(w)
		return
	}
	response.WithMessage(w, http.StatusOK, "OK")
}
