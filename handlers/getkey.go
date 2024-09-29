package handlers

import "net/http"
import "./storage"


func GetKey(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "Key is required", http.StatusBadRequest)
			return
		}

		val ,err := db.Get(key)
		if err == storage.ErrNotFound {
			http.Error(w, "Key not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, "Internal Server Error unavle to get error from surver", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(val)


	})
}