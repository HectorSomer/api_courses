package entities 

    type Course struct {
		ID          int    `json:"idCourse"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Teacher     string `json:"teacher"`
		IDUser 	int    `json:"idUser"`
	}
