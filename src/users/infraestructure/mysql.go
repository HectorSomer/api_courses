package infraestructure

import (
	"event_driven/src/config"
	"event_driven/src/users/domain/entities"
	"log"
    "fmt"
	_"github.com/go-sql-driver/mysql"
	entitiesLogin "event_driven/src/users/application/entities"
)

type MySql struct {
	conn *config.Conn_MySQL
}

func NewMySql() *MySql{
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySql{conn: conn}
}

func(mysql *MySql) CreateUser(user entities.User) (*entities.User, error) {
	query := "INSERT INTO user (username, password, email, role) VALUES (?, ?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, user.Username, user.Password, user.Email, user.Role)
	if err != nil {
		fmt.Println("Error al insertar el usuario:", err)
		return nil, err
	}
    if result != nil {
		rowsAfected, _ := result.RowsAffected()
		if rowsAfected == 1 {
			lastInsertId, err :=  result.LastInsertId()
           if err != nil {
			    fmt.Println("Error al obtener el ID del usuario insertado:", err)
				return nil, err
        }else {
			log.Printf("[MySQL] - Ninguna fila fue afectada.")
		}
		user.ID = int(lastInsertId)
	}
	}else{
		log.Printf("[MySQL] - Ha habido un error en la consulta (ningún resultado).")
	}
	return &user, nil
}

func (mysql *MySql) GetUser(id int) (*entities.User, error) {
	query := "SELECT * FROM user WHERE idUser = ?"
	rows, err := mysql.conn.FetchRows(query, id)
	if err != nil {
		return nil, fmt.Errorf("error al obtener el usuario: %w", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, fmt.Errorf("usuario con ID %d no encontrado", id)
	}
	var user entities.User
	err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Role)
	if err != nil {
		return nil, fmt.Errorf("error al leer los datos del usuario: %w", err)
	}
	return &user, nil
}

func(mysql *MySql) Login(gmail string, password string) (*entitiesLogin.UserLogin, error) {
	query := "SELECT idUser, username, email, role FROM user WHERE email = ? AND password = ?"
	rows, err := mysql.conn.FetchRows(query, gmail, password)
	if err != nil {
		return nil, fmt.Errorf("error al obtener el usuario: %w", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, fmt.Errorf("usuario o contraseña%s incorrecto", gmail)
	}
	var user entitiesLogin.UserLogin
	err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Role)
	if err != nil {
		return nil, fmt.Errorf("error al leer los datos del usuario: %w", err)
	}
	return &user, nil
}