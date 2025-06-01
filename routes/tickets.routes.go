package routes

import (
	"encoding/json"
	"net/http"

	"github.com/KevinCarDev/ticket/db"
	"github.com/KevinCarDev/ticket/models"
	"github.com/gorilla/mux"
)

func GetTicketsHandler(w http.ResponseWriter, r *http.Request) {
	var tickets []models.Ticket
	db.DB.Find(&tickets)
	json.NewEncoder(w).Encode(&tickets)
}

func GetTicketHandler(w http.ResponseWriter, r *http.Request) {
	var ticket models.Ticket
	params := mux.Vars(r)
	db.DB.First(&ticket, params["id"])

	if ticket.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Ticket not found"))
		return
	}

	json.NewEncoder(w).Encode(&ticket)

}
func PostTicketHandler(w http.ResponseWriter, r *http.Request) {
	var ticket models.Ticket
	json.NewDecoder(r.Body).Decode(&ticket)

	createdTicket := db.DB.Create(&ticket)
	err := createdTicket.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&ticket)
}
func DeleteTicketHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var ticket models.Ticket
	db.DB.First(&ticket, params["id"])

	if ticket.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Ticket not found"))
		return
	}

	db.DB.Unscoped().Delete(&ticket)
	w.WriteHeader(http.StatusOK)

}
