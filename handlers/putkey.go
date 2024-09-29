package handlers

import (
	"io"
	"net/http"
	"../storage"
)

func PutKey(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")

		if key == "" {
			http.Error(w, "Key is required", http.StatusBadRequest)
			return
		}

		defer r.Body.Close()

		val, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Internal Server Error unable to read body", http.StatusInternalServerError)
			return
		}

		if db.Set(key, val); err != nil {

			http.Error(w, "Internal Server Error unable to set value", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	})
}