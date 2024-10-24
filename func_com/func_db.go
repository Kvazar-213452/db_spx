package func_com

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Check_db(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile(Data_db)

	if err != nil {
		http.Error(w, "erorr", http.StatusInternalServerError)
		return
	}

	var jsonData map[string]string

	json.Unmarshal(data, &jsonData)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(jsonData)
}

func Check_server(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "good")
}

func Add_db(w http.ResponseWriter, r *http.Request) {
	key1 := r.URL.Query().Get("key")
	key2 := r.URL.Query().Get("key1")

	data, err := ioutil.ReadFile(Data_db)
	if err != nil {
		http.Error(w, "erorr", http.StatusInternalServerError)
		return
	}

	var jsonData map[string]string

	json.Unmarshal(data, &jsonData)

	jsonData[key1] = key2

	updatedData, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		http.Error(w, "Помилка при формуванні JSON", http.StatusInternalServerError)
		return
	}

	ioutil.WriteFile(Data_db, updatedData, 0644)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "good")
}

func Del_data_db(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	data, err := ioutil.ReadFile(Data_db)
	if err != nil {
		http.Error(w, "erorr", http.StatusInternalServerError)
		return
	}

	var jsonData map[string]interface{}

	json.Unmarshal(data, &jsonData)

	delete(jsonData, key)

	updatedData, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		http.Error(w, "erorr", http.StatusInternalServerError)
		return
	}

	ioutil.WriteFile(Data_db, updatedData, 0644)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "good")
}

func Del_all_data_db(w http.ResponseWriter, r *http.Request) {
	emptyData := make(map[string]string)

	updatedData, err := json.MarshalIndent(emptyData, "", "    ")
	if err != nil {
		http.Error(w, "Помилка при формуванні JSON", http.StatusInternalServerError)
		return
	}

	ioutil.WriteFile(Data_db, updatedData, 0644)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "good")
}

func Off(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}
