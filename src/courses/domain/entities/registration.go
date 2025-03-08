package entities

type Registration struct {
	IDCourse     int `json:"idCourse"`
	IDUser int `json:"idUser"`// pending, confirmed, cancelled, expired
}