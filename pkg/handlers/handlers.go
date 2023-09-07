package handlers

import (
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

	render.RenderTemplate(w, "home", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, "about", &models.TemplateData{StringMap: stringMap})
}
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, "about", &models.TemplateData{StringMap: stringMap})
}
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, "about", &models.TemplateData{StringMap: stringMap})
}
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, "about", &models.TemplateData{StringMap: stringMap})
}
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, "about", &models.TemplateData{StringMap: stringMap})
}
