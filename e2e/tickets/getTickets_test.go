package tickets_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/KevinCarDev/ticket/db"
	"github.com/KevinCarDev/ticket/models"
	"github.com/KevinCarDev/ticket/routes"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestDB() {
	database, _ := gorm.Open(postgres.Open(db.DSN), &gorm.Config{})
	database.AutoMigrate(&models.Ticket{})
	db.DB = database
}

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/tickets/{id}", routes.GetTicketHandler).Methods("GET")
	r.HandleFunc("/tickets", routes.PostTicketHandler).Methods("POST")
	return r
}

func TestGetTicketHandler_E2E(t *testing.T) {
	setupTestDB()
	router := setupRouter()

	// Crear un ticket de prueba
	ticket := models.Ticket{Cliente: "Juan", Origen: "A", Destino: "B", Price: 10}
	db.DB.Create(&ticket)

	// Hacer la petición GET
	req := httptest.NewRequest("GET", "/tickets/"+strconv.Itoa(int(ticket.ID)), nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("esperado status 200, obtuviste %d", rr.Code)
	}

	var respTicket models.Ticket
	json.NewDecoder(rr.Body).Decode(&respTicket)
	if respTicket.Cliente != "Juan" {
		t.Errorf("esperado Cliente 'Juan', obtuviste '%s'", respTicket.Cliente)
	}
}

func TestPostTicketHandler_E2E(t *testing.T) {
	setupTestDB()
	router := setupRouter()

	body := `{"Cliente":"Ana","Origen":"X","Destino":"Y","Price":20}`
	req := httptest.NewRequest("POST", "/tickets", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("esperado status 200, obtuviste %d", rr.Code)
	}

	var respTicket models.Ticket
	json.NewDecoder(rr.Body).Decode(&respTicket)
	if respTicket.Cliente != "Ana" {
		t.Errorf("esperado Cliente 'Ana', obtuviste '%s'", respTicket.Cliente)
	}
}
