package cors

import "net/http"

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
  (*w).Header().Set("Access-Control-Allow-Headers", "X-PINGOTHER, Accept, Authorization, Content-Type, X-CSRF-Token")
}