package handlers

import (
	"Booking/internal/config"
	"Booking/internal/models"
	"Booking/internal/render"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// TODO: Repository concept does not understand. need to read more\
// Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(writer http.ResponseWriter, request *http.Request) {
	remoteIp := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(writer, request, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(writer http.ResponseWriter, request *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello there"

	remoteIp := m.App.Session.GetString(request.Context(), "remote_ip")
	stringMap["remoteIp"] = remoteIp
	render.RenderTemplate(writer, request, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Generals render room page
func (m *Repository) Generals(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "generals.page.tmpl", &models.TemplateData{})
}

// majors render room page
func (m *Repository) Majors(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "majors.page.tmpl", &models.TemplateData{})
}

// Availability render room page
func (m *Repository) Availability(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability post form to check room availability
func (m *Repository) PostAvailability(writer http.ResponseWriter, request *http.Request) {
	start := request.Form.Get("start")
	end := request.Form.Get("end")

	writer.Write([]byte(fmt.Sprintf("Start date is %s and end Date is %s", start, end)))
}

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJson post form and send json
func (m *Repository) AvailabilityJson(writer http.ResponseWriter, request *http.Request) {
	resp := jsonResponse{
		Ok:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}
	log.Println(out)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(out)
}

// Contact render room page
func (m *Repository) Contact(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "contact.page.tmpl", &models.TemplateData{})
}

// Reservation render room page
func (m *Repository) Reservation(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, request, "make-reservation.page.tmpl", &models.TemplateData{})
}