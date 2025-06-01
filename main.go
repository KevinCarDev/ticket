package main

import (
	"fmt"
	"net/http"

	"github.com/KevinCarDev/ticket/db"
	"github.com/KevinCarDev/ticket/helpers"
	"github.com/KevinCarDev/ticket/models"
	"github.com/KevinCarDev/ticket/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBconnect()
	db.DB.AutoMigrate(models.Ticket{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/tickets", routes.GetTicketsHandler).Methods("GET")
	r.HandleFunc("/tickets/{id}", routes.GetTicketHandler).Methods("GET")
	r.HandleFunc("/tickets", routes.PostTicketHandler).Methods("POST")
	r.HandleFunc("/tickets/{id}", routes.DeleteTicketHandler).Methods("DELETE")

	input := "Oferta Especial para usuarios"

	fmt.Println("Vocales:", helpers.ContarVocales(input))         // 7
	fmt.Println("Consonantes:", helpers.ContarConsonantes(input)) // 6
	fmt.Println("Iniciales:", helpers.Iniciales(input))

	http.ListenAndServe(":8080", r)
}
