package handlers

import (
	"fmt"
	"net/http"

	"github.com/vikash-parashar/bookings/pkg/config"
	"github.com/vikash-parashar/bookings/pkg/models"
	"github.com/vikash-parashar/bookings/pkg/render"
)

// Repo the Repository used by handlers
var Repo *Repository

// Repository is the type Repository
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	m.App.SessionManager.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, r, "about", &models.TemplateData{StringMap: stringMap})
}
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, r, "generals", &models.TemplateData{StringMap: stringMap})
}
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, r, "majors", &models.TemplateData{StringMap: stringMap})
}
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, r, "search-availability", &models.TemplateData{StringMap: stringMap})
}
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("%s,%s", start, end)))
}
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, r, "contact", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, r, "make-reservation", &models.TemplateData{StringMap: stringMap})
}
