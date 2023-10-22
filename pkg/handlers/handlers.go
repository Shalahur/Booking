package handlers

import (
	"Booking/pkg/config"
	"Booking/pkg/models"
	"Booking/pkg/render"
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

	render.RenderTemplate(writer, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(writer http.ResponseWriter, request *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello there"

	remoteIp := m.App.Session.GetString(request.Context(), "remote_ip")
	stringMap["remoteIp"] = remoteIp
	render.RenderTemplate(writer, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Generals render room page
func (m *Repository) Generals(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "generals.page.tmpl", &models.TemplateData{})
}

// majors render room page
func (m *Repository) Majors(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "majors.page.tmpl", &models.TemplateData{})
}

// Availability render room page
func (m *Repository) Availability(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "search-availability.page.tmpl", &models.TemplateData{})
}

// Contact render room page
func (m *Repository) Contact(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "contact.page.tmpl", &models.TemplateData{})
}

// Reservation render room page
func (m *Repository) Reservation(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "make-reservation.page.tmpl", &models.TemplateData{})
}
