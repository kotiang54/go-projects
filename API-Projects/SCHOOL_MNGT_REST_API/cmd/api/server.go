package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "school_management_api/internal/api/middlewares"
	"school_management_api/internal/api/router"
	"school_management_api/pkg/utils"
)

func main() {
	// Main entry of the api

	port := 3000
	cert := "cert.pem"
	key := "key.pem"

	router := router.Router()

	fmt.Println("Server is running on port:", port)

	// Make HTTP 1.1 with TLS server
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS13,
	}

	// rate limiting middleware can be added here
	// rl := mw.NewRateLimiter(5, time.Minute)

	// HPP middleware options configuration
	// hppOptions := mw.HPPOptions{
	// 	CheckQuery:                 true,
	// 	CheckBody:                  true,
	// 	CheckBodyOnlyForContenType: "application/x-www-form-urlencoded",
	// 	Whitelist:                  []string{"sortBy", "sortOrder", "first_name", "last_name", "class"},
	// }

	// Recommended middleware order (from outermost to innermost)
	// secureMux := mw.Cors( // 1. CORS: Handle cross-origin and preflight requests first
	// 	mw.Hpp(hppOptions)( // 2. HPP: Sanitize query/body params before any logic uses them
	// 		rl.Middleware( // 3. Rate Limiting: Block abusive clients early, before expensive work
	// 			mw.SecurityHeaders( // 4. Security Headers: Set headers for all responses
	// 				mw.ResponseTime( // 5. Response Time: Measure as much as possible
	// 					mw.Compression( // 6. Compression: Compress the final response
	// 						mux, // 7. Your main router/handler
	// 					),
	// 				),
	// 			),
	// 		),
	// 	),
	// )

	// Using helper function to apply middlewares
	secureMux := utils.ApplyMiddlewares(router,
		// mw.Compression,     // 6. Compression: Compress the final response
		// mw.ResponseTime,    // 5. Response Time: Measure as much as possible
		mw.SecurityHeaders, // 4. Security Headers: Set headers for all responses
		// rl.Middleware,      // 3. Rate Limiting: Block abusive clients early, before expensive work
		// mw.Hpp(hppOptions), // 2. HPP: Sanitize query/body params before any logic uses them
		// mw.Cors,            // 1. CORS: Handle cross-origin and preflight requests first
	)

	// Create a custom server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      secureMux,
		TLSConfig:    tlsConfig,
		TLSNextProto: map[string]func(*http.Server, *tls.Conn, http.Handler){},
	}

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Error starting the server:", err)
	}
}
