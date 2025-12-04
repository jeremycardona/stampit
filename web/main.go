package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type App struct {
	// depencies here
	logger *slog.Logger
}

// Middleware to log all requests
func (app *App) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler
		next.ServeHTTP(w, r)

		// Log after request is handled
		app.logger.Info("Request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"remote_addr", r.RemoteAddr,
			"duration_ms", time.Since(start).Milliseconds(),
		)
	})
}

func (app *App) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, requested: " + r.URL.Path + "\n"))
}

func (app *App) claim(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Claim rewards " + r.URL.Path + "\n"))
}

func (app *App) rewards(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("View rewards " + r.URL.Path + "\n"))
}

func setupLogger(env string) *slog.Logger {
	var handler slog.Handler

	if env == "production" {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	} else {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	}

	return slog.New(handler)
}

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	logger := setupLogger(env)
	app := &App{
		logger: logger,
	}
	r := mux.NewRouter()
	r.Use(app.loggingMiddleware)
	r.HandleFunc("/", app.home)
	r.HandleFunc("/claim", app.claim)
	r.HandleFunc("/rewards", app.rewards)

	logger.Info("Server starting",
		"port", 8080,
		"environmnet", env,
	)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		logger.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
