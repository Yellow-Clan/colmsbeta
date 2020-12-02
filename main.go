package main

import (
	//"fmt"
	"log"
	"net/http"
	"os"

	"github.com/202lp1/colms/cfig"
	"github.com/202lp1/colms/controllers"
	"github.com/202lp1/colms/models"
	"github.com/gorilla/mux" //gin
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "Hello World!")
	//})
	cfig.DB, err = connectDB()
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	log.Printf("db is connected: %v", cfig.DB)

	// Migrate the schema
	cfig.DB.AutoMigrate(&models.Empleado{})
	cfig.DB.AutoMigrate(&models.Alumno{})
	cfig.DB.AutoMigrate(&models.Matricula{})
	cfig.DB.AutoMigrate(&models.Docente{})
	//cfig.DB.Create(&models.Empleado{Name: "Juan", City: "Juliaca"})

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home).Methods("GET")

	r.HandleFunc("/item/index", controllers.ItemList).Methods("GET")

	r.HandleFunc("/employee/index", controllers.EmployeeList).Methods("GET")
	r.HandleFunc("/employee/form", controllers.EmployeeForm).Methods("GET", "POST")
	r.HandleFunc("/employee/delete", controllers.EmployeeDel).Methods("GET")

	r.HandleFunc("/alumno/index", controllers.AlumnoList).Methods("GET")
	r.HandleFunc("/alumno/form", controllers.AlumnoForm).Methods("GET", "POST")
	r.HandleFunc("/alumno/delete", controllers.AlumnoDel).Methods("GET")

	r.HandleFunc("/matricula/index", controllers.MatriculaList).Methods("GET")
	r.HandleFunc("/matricula/form", controllers.MatriculaForm).Methods("GET", "POST")
	r.HandleFunc("/matricula/delete", controllers.MatriculaDel).Methods("GET")

	r.HandleFunc("/docente/index", controllers.DocenteList).Methods("GET")
	r.HandleFunc("/docente/form", controllers.DocenteForm).Methods("GET", "POST")
	r.HandleFunc("/docente/delete", controllers.DocenteDel).Methods("GET")

	//http.ListenAndServe(":80", r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("port: %v", port)
	http.ListenAndServe(":"+port, r)

}

func connectDBmysql() (c *gorm.DB, err error) {
	dsn := "docker:docker@tcp(mysql-db:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "docker:docker@tcp(localhost:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return conn, err
}

func connectDB() (c *gorm.DB, err error) {
	////dsn := "docker:docker@tcp(mysql-db:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "docker:docker@tcp(localhost:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	//conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dsn := "user=fcliujwibyxsdc password=a8b47332106e2c8ed9cdfdf4383d68f68f734790c714516102e201ec49e643d7 host=ec2-174-129-199-54.compute-1.amazonaws.com dbname=d35n7dqvk2ndfr port=5432 sslmode=require TimeZone=Asia/Shanghai"
	//dsn := "user=postgres password=postgres2 dbname=users_test host=localhost port=5435 sslmode=disable TimeZone=Asia/Shanghai"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return conn, err
}
