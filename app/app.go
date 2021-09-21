package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gautampgit/banking/domain"
	"github.com/gautampgit/banking/service"
	"github.com/gorilla/mux"
)

func getDBClient() *sql.DB {
	dbUser := "gautam"        //os.Getenv("USERNAME")
	dbPasswrd := "Laxmi@1971" //os.Getenv("PASSWORD")
	dbAddrress := "localhost" //os.Getenv("DB_ADDRESS")
	dbPort := "3306"          //os.Getenv("DB_PASSWORD")
	dbName := "banking"       //os.Getenv("DB_NAME")
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswrd, dbAddrress, dbPort, dbName)
	conn, err := sql.Open("mysql", datasource)
	if err != nil {
		log.Fatal("Unable to coonect to the DB")
	}
	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)
	//defer conn.Close()
	return conn
}

//
func Start() {
	//creating a new request multiplexer
	router := mux.NewRouter()

	//wiring

	dbClient := getDBClient()
	customerRepositoryDB := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccountRepositoryDb(dbClient)
	//ch := CustomHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomHandler{service: service.NewCustomerService(customerRepositoryDB)}
	ah := AccountHandler{service: service.NewAccountService(accountRepositoryDB)}
	//routing to endpoints
	router.HandleFunc("/customers", ch.getAllCustomers)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getCustomerById)
	router.HandleFunc("/customer/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	// := os.Getenv("SERVER_ADDRESS")
	//port := os.Getenv("SERVERPORT")

	log.Fatal(http.ListenAndServe(":8080" /*fmt.Sprintf("%s:%s", address, port)*/, router))

}
