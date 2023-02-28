package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strings"
)

func SetUser(w http.ResponseWriter, r *http.Request, app string) {

	newDataUser := User{}
	var insertColKey, insertValue, updateValue, targetTable string
	db, err := ConnectMysql()
	defer db.Close()

	//Check DB Connection is available
	err = db.Ping()
	if err != nil {
		// c.IndentedJSON(http.StatusInternalServerError, &JsonMessage{"500", "failed", err.Error()})
		jsonResponse, err := json.Marshal(&JsonMessage{"500", "failed", err.Error()})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}
	// fullAddr := r.RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	fullAddr := net.ParseIP(ip).String()
	// fmt.Println(fullAddr)
	// ip := strings.Split(fullAddr, ":")[0]
	//Check if IP is whitelisted in DB
	if !CheckClientIP(fullAddr, db) {
		// c.IndentedJSON(http.StatusUnauthorized, &JsonMessage{"401", "Failed", "Unauthorized"})
		jsonResponse, err := json.Marshal(&JsonMessage{"401", "Failed", "Unauthorized"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&newDataUser)
	if err != nil {
		jsonResponse, err := json.Marshal(&JsonMessage{"4002", "failed", "Please check body request"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		// http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	//Check Request Body is NOT NULL
	// if err := c.BindJSON(&newDataUser); err != nil {
	// 	c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4002", "failed", "Empty Request Body"})
	// 	return
	// }

	reNik := regexp.MustCompile("^[a-zA-Z0-9]+$")
	reAlphaNum := regexp.MustCompile(`^[a-zA-Z0-9',. ]+$`)
	reEmail := regexp.MustCompile("^[a-zA-Z0-9@,. ]+$")
	reDate := regexp.MustCompile(`^[0-9\-]+$`)
	reNum := regexp.MustCompile("^[0-9]+$")

	if !reNik.Match([]byte(newDataUser.Nik)) {
		// c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4003", "failed", "Recheck NIK"})
		jsonResponse, err := json.Marshal(&JsonMessage{"4003", "failed", "Recheck NIK"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}
	if !reAlphaNum.Match([]byte(newDataUser.First_name)) && newDataUser.First_name != "" {
		// c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4003", "failed", "Recheck First Name"})
		jsonResponse, err := json.Marshal(&JsonMessage{"4003", "failed", "Recheck First Name"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}
	if !reAlphaNum.Match([]byte(newDataUser.Last_name)) && newDataUser.Last_name != "" {
		// c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4003", "failed", "Recheck Last Name"})
		jsonResponse, err := json.Marshal(&JsonMessage{"4003", "failed", "Recheck Last Name"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}
	if !reAlphaNum.Match([]byte(newDataUser.Full_name)) {
		// c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4003", "failed", "Recheck Full Name"})
		jsonResponse, err := json.Marshal(&JsonMessage{"4003", "failed", "Recheck Full Name"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}
	if !reAlphaNum.Match([]byte(newDataUser.Position)) {
		// c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4003", "failed", "Recheck Position"})
		jsonResponse, err := json.Marshal(&JsonMessage{"4003", "failed", "Recheck Position"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}
	if !reEmail.Match([]byte(newDataUser.Email)) {
		// c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4003", "failed", "Recheck Email"})
		jsonResponse, err := json.Marshal(&JsonMessage{"4003", "failed", "Recheck Email"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}
	if !reDate.Match([]byte(newDataUser.Hired_date)) {
		// c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4003", "failed", "Recheck Hired Date"})
		jsonResponse, err := json.Marshal(&JsonMessage{"4003", "failed", "Recheck Hired Date"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}
	if !reDate.Match([]byte(newDataUser.Resign_date)) && newDataUser.Resign_date != "" {
		// c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4003", "failed", "Recheck Resign Date"})
		jsonResponse, err := json.Marshal(&JsonMessage{"4003", "failed", "Recheck Resign Date"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}
	if !reNum.Match([]byte(newDataUser.Unitkerja_id)) {
		// c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4003", "failed", "Recheck Unit Kerja ID"})
		jsonResponse, err := json.Marshal(&JsonMessage{"4003", "failed", "Recheck Unit Kerja ID"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}
	if !reAlphaNum.Match([]byte(newDataUser.Unitkerja)) {
		// c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4003", "failed", "Recheck Unit Kerja"})
		jsonResponse, err := json.Marshal(&JsonMessage{"4003", "failed", "Recheck Unit Kerja"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}
	if !reNik.Match([]byte(newDataUser.Manager_id)) {
		// c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4003", "failed", "Recheck Manager ID"})
		jsonResponse, err := json.Marshal(&JsonMessage{"4003", "failed", "Recheck Manager ID"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	if strings.ToLower(newDataUser.Status) == "inactive" {
		newDataUser.Status = "true"
	} else if strings.ToLower(newDataUser.Status) == "active" {
		newDataUser.Status = ""
	} else {
		// c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4003", "failed", "Recheck Status"})
		jsonResponse, err := json.Marshal(&JsonMessage{"4003", "failed", "Recheck Status"})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	if app == "hcis" {
		if !reNum.Match([]byte(newDataUser.Position_id)) {
			// c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4003", "failed", "Recheck Position ID"})
			jsonResponse, err := json.Marshal(&JsonMessage{"4003", "failed", "Recheck Position ID"})
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
			return
		}
		if !reAlphaNum.Match([]byte(newDataUser.Employee_type)) {
			// c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4003", "failed", "Recheck Employee Type"})
			jsonResponse, err := json.Marshal(&JsonMessage{"4003", "failed", "Recheck Employee Type"})
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
			return
		}
		if !reNum.Match([]byte(newDataUser.Person_grade)) {
			// c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4003", "failed", "Recheck Person Grade"})
			jsonResponse, err := json.Marshal(&JsonMessage{"4003", "failed", "Recheck Person Grade"})
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
			return
		}
		if !reNum.Match([]byte(newDataUser.Job_grade)) {
			// c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4003", "failed", "Recheck Job Grade"})
			jsonResponse, err := json.Marshal(&JsonMessage{"4003", "failed", "Recheck Job Grade"})
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
			return
		}
		if !reNum.Match([]byte(newDataUser.Divisi_id)) {
			// c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4003", "failed", "Recheck Divisi ID"})
			jsonResponse, err := json.Marshal(&JsonMessage{"4003", "failed", "Recheck Divisi ID"})
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
			return
		}
		if !reAlphaNum.Match([]byte(newDataUser.Divisi)) {
			// c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4003", "failed", "Recheck Divisi"})
			jsonResponse, err := json.Marshal(&JsonMessage{"4003", "failed", "Recheck Divisi"})
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
			return
		}
		if !reAlphaNum.Match([]byte(newDataUser.Flag)) && newDataUser.Flag != "" {
			// c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4003", "failed", "Recheck Flag"})
			jsonResponse, err := json.Marshal(&JsonMessage{"4003", "failed", "Recheck Flag"})
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
			return
		}

		insertColKey = "nik, first_name, last_name, full_name, position, inactive, " +
			"email, hired_date, resign_date, unitkerja_id, unitkerja, " +
			"manager_id, employee_type, person_grade, job_grade, position_id, " +
			"divisi_id, divisi, flag"
		insertValue = fmt.Sprintf("'%s','%s','%s','%s','%s','%s',   '%s','%s','%s','%s','%s',   '%s','%s','%s','%s','%s',   '%s','%s','%s'",
			newDataUser.Nik, newDataUser.First_name, newDataUser.Last_name, newDataUser.Full_name, newDataUser.Position, newDataUser.Status,
			newDataUser.Email, newDataUser.Hired_date, newDataUser.Resign_date, newDataUser.Unitkerja_id, newDataUser.Unitkerja,
			newDataUser.Manager_id, newDataUser.Employee_type, newDataUser.Person_grade, newDataUser.Job_grade, newDataUser.Position_id,
			newDataUser.Divisi_id, newDataUser.Divisi, newDataUser.Flag)
		updateValue = fmt.Sprintf("first_name='%s', last_name='%s', full_name='%s', position='%s', inactive='%s',"+
			"email='%s', hired_date='%s', resign_date='%s', unitkerja_id='%s', unitkerja='%s',"+
			"manager_id='%s', employee_type='%s', person_grade='%s', job_grade='%s', position_id='%s', "+
			"divisi_id='%s', divisi='%s', flag='%s', statusUpdate='1'",
			newDataUser.First_name, newDataUser.Last_name, newDataUser.Full_name, newDataUser.Position, newDataUser.Status,
			newDataUser.Email, newDataUser.Hired_date, newDataUser.Resign_date, newDataUser.Unitkerja_id, newDataUser.Unitkerja,
			newDataUser.Manager_id, newDataUser.Employee_type, newDataUser.Person_grade, newDataUser.Job_grade, newDataUser.Position_id,
			newDataUser.Divisi_id, newDataUser.Divisi, newDataUser.Flag)
		targetTable = "tableproses_hcis"

	} else {
		insertColKey = "nik, first_name, last_name, full_name, position, inactive, email, " +
			"hired_date, resign_date, unitkerja_id, unitkerja, manager_id"
		insertValue = fmt.Sprintf("'%s','%s','%s','%s','%s','%s','%s',   '%s','%s','%s','%s','%s'",
			newDataUser.Nik, newDataUser.First_name, newDataUser.Last_name, newDataUser.Full_name, newDataUser.Position, newDataUser.Status, newDataUser.Email,
			newDataUser.Hired_date, newDataUser.Resign_date, newDataUser.Unitkerja_id, newDataUser.Unitkerja, newDataUser.Manager_id)
		updateValue = fmt.Sprintf("first_name='%s', last_name='%s', full_name='%s',`position`='%s',inactive='%s',email='%s',"+
			"hired_date='%s',resign_date='%s',unitkerja_id='%s',unitkerja='%s',manager_id='%s',statusUpdate='1'",
			newDataUser.First_name, newDataUser.Last_name, newDataUser.Full_name, newDataUser.Position, newDataUser.Status, newDataUser.Email,
			newDataUser.Hired_date, newDataUser.Resign_date, newDataUser.Unitkerja_id, newDataUser.Unitkerja, newDataUser.Manager_id)
		targetTable = "tableproses_aralia"
	}
	upsertQuery := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) on duplicate key update %s;", targetTable, insertColKey, insertValue, updateValue)
	// fmt.Println(upsertQuery)
	upsert, err := db.Query(upsertQuery)
	if err != nil {
		// c.IndentedJSON(http.StatusBadRequest, &JsonMessage{"4004", "failed", err.Error()})
		jsonResponse, err := json.Marshal(&JsonMessage{"4004", "failed", err.Error()})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		panic(err.Error())
	}

	defer upsert.Close()

	// c.IndentedJSON(http.StatusAccepted, &JsonMessage{"200", "Success", "User Processed to Sailpoint"})
	jsonResponse, err := json.Marshal(&JsonMessage{"200", "Success", "User Processed to Sailpoint"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
	fmt.Println("Aplikasi : " + app)
	fmt.Println("EOL")

}
