package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type RequestBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type JsonMessage2 struct {
	Code    string `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	// router := gin.Default()
	http.HandleFunc("/", HelloServer)
	// router.POST("iamAPI/v1/aralia/setUser", setUserAralia)
	http.HandleFunc("/iamAPI/v1/aralia/setUser", setUserAralia)
	// router.POST("iamAPI/v1/hcis/setUser", setUserHcis)
	http.HandleFunc("/iamAPI/v1/hcis/setUser", setUserHcis)
	// router.Run("0.0.0.0:8081")
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", 8081), nil)
}

//	func setUserAralia(c *gin.Context) {
//		SetUser(c, "aralia")
//	}
func setUserAralia(w http.ResponseWriter, r *http.Request) {
	SetUser(w, r, "aralia")
}

//	func setUserHcis(c *gin.Context) {
//		SetUser(c, "hcis")
//	}
func setUserHcis(w http.ResponseWriter, r *http.Request) {
	SetUser(w, r, "hcis")
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var requestBody RequestBody
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			jsonResponse, err := json.Marshal(&JsonMessage2{"400", "failed", "Please check body request"})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
			// http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		reAlphaNum := regexp.MustCompile(`^[a-zA-Z0-9',. ]+$`)
		if !reAlphaNum.Match([]byte(requestBody.Name)) {
			// response := JsonMessage2{
			// 	Message: "Check Name",
			// 	Code:    "404",
			// 	Status:  "Failed",
			// }
			jsonResponse, err := json.Marshal(&JsonMessage2{"500", "failed", err.Error()})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
			return
		} else {
			jsonResponse, err := json.Marshal(requestBody)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
		}

		// fmt.Fprintf(w, "Received request body: %+v", requestBody)
	} else {
		response := map[string]string{"message": "Hello, World!"}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}
}
