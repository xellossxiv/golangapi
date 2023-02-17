package controller

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"teststart/model"
	"teststart/tools"

	"github.com/gin-gonic/gin"
)

func SetUser(c *gin.Context, app string) {

	newDataUser := model.User{}
	var insertColKey, insertValue, updateValue, targetTable string
	db, err := tools.ConnectMysql()
	defer db.Close()

	//Check DB Connection is available
	err = db.Ping()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, &model.JsonMessage{"500", "failed", err.Error()})
		return
	}

	//Check if IP is whitelisted in DB
	if !tools.CheckClientIP(c.ClientIP(), db) {
		c.IndentedJSON(http.StatusUnauthorized, &model.JsonMessage{"401", "Failed", "Unauthorized"})
		return
	}

	//Check Request Body is NOT NULL
	if err := c.BindJSON(&newDataUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4002", "failed", "Empty Request Body"})
		return
	}

	reNik := regexp.MustCompile("^[a-zA-Z0-9]+$")
	reAlphaNum := regexp.MustCompile(`^[a-zA-Z0-9',. ]+$`)
	reEmail := regexp.MustCompile("^[a-zA-Z0-9@,. ]+$")
	reDate := regexp.MustCompile(`^[0-9\-]+$`)
	reNum := regexp.MustCompile("^[0-9]+$")

	if !reNik.Match([]byte(newDataUser.Nik)) {
		c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4003", "failed", "Recheck NIK"})
		return
	}
	if !reAlphaNum.Match([]byte(newDataUser.First_name)) && newDataUser.First_name != "" {
		c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4003", "failed", "Recheck First Name"})
		return
	}
	if !reAlphaNum.Match([]byte(newDataUser.Last_name)) && newDataUser.Last_name != "" {
		c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4003", "failed", "Recheck Last Name"})
		return
	}
	if !reAlphaNum.Match([]byte(newDataUser.Full_name)) {
		c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4003", "failed", "Recheck Full Name"})
		return
	}
	if !reAlphaNum.Match([]byte(newDataUser.Position)) {
		c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4003", "failed", "Recheck Position"})
		return
	}
	if !reEmail.Match([]byte(newDataUser.Email)) {
		c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4003", "failed", "Recheck Email"})
		return
	}
	if !reDate.Match([]byte(newDataUser.Hired_date)) {
		c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4003", "failed", "Recheck Hired Date"})
		return
	}
	if !reDate.Match([]byte(newDataUser.Resign_date)) && newDataUser.Resign_date != "" {
		c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4003", "failed", "Recheck Resign Date"})
		return
	}
	if !reNum.Match([]byte(newDataUser.Unitkerja_id)) {
		c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4003", "failed", "Recheck Unit Kerja ID"})
		return
	}
	if !reAlphaNum.Match([]byte(newDataUser.Unitkerja)) {
		c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4003", "failed", "Recheck Unit Kerja"})
		return
	}
	if !reNik.Match([]byte(newDataUser.Manager_id)) {
		c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4003", "failed", "Recheck Manager ID"})
		return
	}

	if strings.ToLower(newDataUser.Status) == "inactive" {
		newDataUser.Status = "true"
	} else if strings.ToLower(newDataUser.Status) == "active" {
		newDataUser.Status = ""
	} else {
		c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4003", "failed", "Recheck Status"})
		return
	}

	if app == "hcis" {
		if !reNum.Match([]byte(newDataUser.Position_id)) {
			c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4003", "failed", "Recheck Position ID"})
			return
		}
		if !reAlphaNum.Match([]byte(newDataUser.Employee_type)) {
			c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4003", "failed", "Recheck Employee Type"})
			return
		}
		if !reNum.Match([]byte(newDataUser.Person_grade)) {
			c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4003", "failed", "Recheck Person Grade"})
			return
		}
		if !reNum.Match([]byte(newDataUser.Job_grade)) {
			c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4003", "failed", "Recheck Job Grade"})
			return
		}
		if !reNum.Match([]byte(newDataUser.Divisi_id)) {
			c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4003", "failed", "Recheck Divisi ID"})
			return
		}
		if !reAlphaNum.Match([]byte(newDataUser.Divisi)) {
			c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4003", "failed", "Recheck Divisi"})
			return
		}
		if !reAlphaNum.Match([]byte(newDataUser.Flag)) && newDataUser.Flag != "" {
			c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4003", "failed", "Recheck Flag"})
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
		c.IndentedJSON(http.StatusBadRequest, &model.JsonMessage{"4004", "failed", err.Error()})
		panic(err.Error())
	}

	defer upsert.Close()

	c.IndentedJSON(http.StatusAccepted, &model.JsonMessage{"200", "Success", "User Processed to Sailpoint"})
	fmt.Println("Aplikasi : " + app)
	fmt.Println("EOL")

}
