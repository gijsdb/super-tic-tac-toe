package middlewares

// func PrepContext(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		ctx := r.Context()
// 		reqId := rand.Int63()

// 		ctx = context.WithValue(r.Context(), "reqId", reqId)
// 		ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
// 		defer cancel()
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }
