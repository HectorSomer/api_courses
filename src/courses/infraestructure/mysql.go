package infraestructure

import (
	"event_driven/src/config"
	"event_driven/src/courses/domain/entities"
	"event_driven/src/users/infraestructure"
	"fmt"
	"log"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

func NewMySql() *MySQL{
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}
func (mysql *MySQL) CreateCourse(course entities.Course) (*entities.Course, error) {
	query := "INSERT INTO course (name, description, teacher, idUser) VALUES (?, ?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, course.Name, course.Description, course.Teacher, course.IDUser)
	if err != nil {
		fmt.Println("Error in insert course: ", err)
		return nil, err
	}
	if result != nil {
		rowsAfected, _ := result.RowsAffected()
		if rowsAfected == 1 {
			lastInsertId, err := result.LastInsertId()
			if err != nil {
				fmt.Println("Error in get ID of course inserted: ", err)
				return nil, err
			}else {
				log.Printf("[MySQL] - Ninguna fila fue afectada.")
			}
			course.ID = int(lastInsertId)
     }
}else{
	log.Printf("[MySQL] - Ha habido un error en la consulta (ningún resultado).")
}
	return &course, nil
}

func(mysql *MySQL) GetCourse(id int) (*entities.Course, error) {
     query := "SELECT * FROM course WHERE idCourse = ?"
	 rows, err := mysql.conn.FetchRows(query, id)
	 if err != nil {
		 fmt.Println("Error in get course: ", err)
		 return nil, err
	 }
	 defer rows.Close()
	 if !rows.Next() {
		 return nil, fmt.Errorf("course with ID %d not found", id)
	 }
	 var course entities.Course
	 err = rows.Scan(&course.ID, &course.Name, &course.Description, &course.Teacher, &course.IDUser)
	 if err != nil {
		 fmt.Println("Error in scan course: ", err)
		 return nil, err
	 }
	 return &course, nil
}
func (mysql *MySQL) RegistrateToCourse(registration entities.Registration) (*entities.RegistrationInfo, error) {
	query := "INSERT INTO registration (course_id, user_id) VALUES (?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, registration.IDCourse, registration.IDUser)
	if err != nil {
		fmt.Println("Error in insert registration: ", err)
		return nil, err
	}
	var messageSend entities.RegistrationInfo
	if result != nil {
		rowsAfected, _ := result.RowsAffected()
		if rowsAfected == 1 {
			fmt.Println("Se supone que todo va bien")
	   }
    student, err := infraestructure.NewMySql().GetUser(registration.IDUser)
	if err != nil{
		fmt.Println("Error in get user: ", err)
		return nil, err
	}
	courseFind, err := mysql.GetCourse(registration.IDCourse)
	if err != nil {
		fmt.Println("Error in get course: ", err)
        return nil, err
    }
	if student != nil && courseFind != nil {
		messageSend.IDCourse = registration.IDCourse
		messageSend.IDUserStudent = registration.IDUser
		messageSend.UserStudent = student.Username
		messageSend.IDUserTeacher = courseFind.IDUser
	}

}else{
	log.Printf("[MySQL] - Ha habido un error en la consulta (ningún resultado).")
}
    return &messageSend, nil
}

func(mysq *MySQL) GetCourses() (*[]entities.Course, error){
	query := "SELECT * FROM course"
    rows, err := mysq.conn.FetchRows(query)
    if err != nil {
        fmt.Println("Error in get courses: ", err)
        return nil, err
    }
    defer rows.Close()
    var courses []entities.Course
    for rows.Next() {
        var course entities.Course
        err = rows.Scan(&course.ID, &course.Name, &course.Description, &course.Teacher, &course.IDUser)
        if err != nil {
            fmt.Println("Error in scan course: ", err)
            return nil, err
        }
        courses = append(courses, course)
    }
    return &courses, nil
}