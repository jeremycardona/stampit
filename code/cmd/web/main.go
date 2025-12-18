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
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	r := mux.NewRouter()
	r.Use(app.loggingMiddleware)
	r.HandleFunc("/", app.home).Methods("GET")
	r.HandleFunc("/redeem", app.redeemPost).Methods("POST")
	r.HandleFunc("/claim", app.claimPost).Methods("POST")
	r.HandleFunc("/points", app.points).Methods("GET")
	r.HandleFunc("/points", app.pointsPost).Methods("POST")
	r.HandleFunc("/rewards", app.rewards).Methods("GET")
	r.HandleFunc("/rewards", app.rewardsPost).Methods("POST")
	r.HandleFunc("/history", app.customerHistory).Methods("GET")
	r.HandleFunc("/transactions", app.employeeHistory).Methods("GET")
	r.HandleFunc("/offer", app.offers).Methods("GET")
	r.HandleFunc("/offer", app.offerPost).Methods("POST")
	r.HandleFunc("/offer", app.offerPut).Methods("PUT")
	r.HandleFunc("/offer", app.offerDelete).Methods("DELETE")
	r.HandleFunc("/search", app.search).Methods("GET")
	r.HandleFunc("/search", app.searchPost).Methods("POST")
	r.HandleFunc("/purchase", app.purchase).Methods("GET")
	r.HandleFunc("/purchase", app.purchasePost).Methods("POST")
	r.HandleFunc("/login", app.loginPost).Methods("POST")
	r.HandleFunc("/signup", app.signupPost).Methods("POST")
	r.HandleFunc("/member", app.memberPost).Methods("POST")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer)).Methods("GET")

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
