package products

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetAllProducts implements Products.
func (h ProductHandlers) GetAllProducts(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]byte{})
	}
}
