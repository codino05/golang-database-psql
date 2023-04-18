package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "dicoding"
	dbname   = "belajar_golang_database"
)

var (
	db  *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("successfully connected database")

	//CreateEmployee()
	GetEmployee()
	// UpdateEmployee()
	// DeleteEmployee()
}

type Employee struct {
	Id        int
	Full_name string
	Email     string
	Age       int
	Division  string
}

func CreateEmployee() {
	var employee = Employee{}

	sqlStatement := `insert into employees (full_name, email, age, division)
	values($1, $2, $3, $4) returning * `

	err = db.QueryRow(sqlStatement, "yespapa", "yespapa@gmail.com", 23, "BackEnd Golang").
		Scan(&employee.Id, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("new employee data : %+v\n", employee)
}

func GetEmployee() {
	var result = []Employee{}

	sqlStatement := `select *from employees`
	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}
		err = rows.Scan(&employee.Id, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

		if err != nil {
			panic(err)
		}

		result = append(result, employee)
	}
	fmt.Println("employee datas : ", result)
}

func UpdateEmployee() {
	sqlStatement := `update employees set full_name = $2, email = $3, division= $4, age = $5 where id = $1;`
	res, err := db.Exec(sqlStatement, 1, "budi", "budi@gmail.com", "BackEnd Java", 30)
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("update data amount : ", count)
}

func DeleteEmployee() {
	sqlStatement := `delete from employees where id = $1;`

	res, err := db.Exec(sqlStatement, 4)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("deleted data amount : ", count)
}
