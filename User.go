package main

type User struct {
	Nik           string `json:"nik"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Full_name     string `json:"full_name"`
	Position      string `json:"position"`
	Email         string `json:"email"`
	Hired_date    string `json:"hired_date"`
	Resign_date   string `json:"resign_date"`
	Unitkerja_id  string `json:"unitkerja_id"`
	Unitkerja     string `json:"unitkerja"`
	Manager_id    string `json:"manager_id"`
	Status        string `json:"status"`
	Employee_type string `json:"employee_type"`
	Person_grade  string `json:"person_grade"`
	Job_grade     string `json:"job_grade"`
	Position_id   string `json:"position_id"`
	Divisi        string `json:"divisi"`
	Divisi_id     string `json:"divisi_id"`
	Flag          string `json:"flag"`
}
