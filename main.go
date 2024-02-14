package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MarkTBSS/go-hexagonalArchitecture/handler"
	"github.com/MarkTBSS/go-hexagonalArchitecture/repository"
	"github.com/MarkTBSS/go-hexagonalArchitecture/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432 // Default port
	user     = "postgres"
	password = "pass123"
	dbname   = "banking"
)

func main() {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		log.Fatal(err)
	}
	customerRepositoryDB := repository.NewCustomRepositoryDB(db)
	customerRepositoryMock := repository.NewCustomRepositoryMock()
	_ = customerRepositoryDB
	customerService := service.NewCustomerService(customerRepositoryMock)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)
	http.ListenAndServe(":8000", router)
}
