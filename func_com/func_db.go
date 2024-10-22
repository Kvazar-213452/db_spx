package func_com

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Check_db(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("main.json")

	if err != nil {
		http.Error(w, "erorr", http.StatusInternalServerError)
		return
	}

	var jsonData map[string]string

	json.Unmarshal(data, &jsonData)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(jsonData)
}
