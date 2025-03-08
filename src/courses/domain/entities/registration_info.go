package entities

type RegistrationInfo struct {
	IDCourse     int `json:"idCourse"`
	IDUserStudent int `json:"idUser"`
    IDUserTeacher int `json:"idUserTeacher"`
	UserStudent string `json:"userStudent"`
}