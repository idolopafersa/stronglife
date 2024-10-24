package security

import (
	structmodels "backendgo/StructModels"
	"encoding/json"
	"fmt"
	"net/http"
)

func DecodeAndVerifyToken(w http.ResponseWriter, r *http.Request, routine structmodels.NewRoutine) (string, error) {
	if err := json.NewDecoder(r.Body).Decode(&routine); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		fmt.Println(err)

	}

	cookie, err := r.Cookie("token")
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Unauthorized: No valid cookie", http.StatusUnauthorized)

	}
}
