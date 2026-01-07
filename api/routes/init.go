package routes

import (
	"net/http"

	"github.com/YahiaJouini/careflow/api/middleware"
	"github.com/gorilla/mux"
)

func InitializeRoutes() *mux.Router {
	// main router
	router := mux.NewRouter()
	// auth router
	authRouter := router.PathPrefix("/auth").Subrouter()
	// user router
	meRouter := router.PathPrefix("/me").Subrouter()
	meRouter.Use(middleware.AuthMiddleware(middleware.All))
	InitUserRoutes(meRouter)

	// admin router
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middleware.AuthMiddleware(middleware.Admin))

	// patient router
	patientRouter := router.PathPrefix("/patient").Subrouter()
	patientRouter.Use(middleware.AuthMiddleware(middleware.Patient))

	// doctor router
	doctorRouter := router.PathPrefix("/doctor").Subrouter()
	doctorRouter.Use(middleware.AuthMiddleware(middleware.Doctor))

	// public router
	publicRouter := router.PathPrefix("/public").Subrouter()

	// init routers
	InitAuthRoutes(authRouter)

	InitAdminRoutes(adminRouter)
	InitPatientRoutes(patientRouter)
	InitPublicRoutes(publicRouter)
	InitDoctorRoutes(doctorRouter)
	return router
}

func UseSecurityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set security headers
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
