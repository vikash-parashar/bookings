package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vikash-parashar/bookings/internal/config"
	"github.com/vikash-parashar/bookings/internal/forms"
	"github.com/vikash-parashar/bookings/internal/models"
	"github.com/vikash-parashar/bookings/internal/render"
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

// Home renders the search Home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr

	m.App.SessionManager.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home", &models.TemplateData{})
}

// About renders the search About page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, r, "about", &models.TemplateData{StringMap: stringMap})
}

// Generals renders the search GeneralsRoom page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, r, "generals", &models.TemplateData{StringMap: stringMap})
}

// Majors renders the search Majors page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, r, "majors", &models.TemplateData{StringMap: stringMap})
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, r, "search-availability", &models.TemplateData{StringMap: stringMap})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("%s,%s", start, end)))
}

// jsonResponse
type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles request for availability and send JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	res := jsonResponse{
		true,
		"Available",
	}

	out, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact renders the search Contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	render.RenderTemplate(w, r, "contact", &models.TemplateData{StringMap: stringMap})
}

// Reservation renders the search Reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]any)
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "make-reservation", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// PostReservation handlers the posting of a  Reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}
	form := forms.New(r.PostForm)

	// form.Has("first_name", r)
	form.MinLength("first_name", 3, r)
	form.Required("first_name", "last_name", "email")
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]any)
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation", &models.TemplateData{
			Form: form,
			Data: data,
		})
	}

	// getting values of reservation from sessions to display on reservation-summery page
	m.App.SessionManager.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summery", http.StatusSeeOther)

}

func (m *Repository) ReservationSummery(w http.ResponseWriter, r *http.Request) {
	//TODO:
	reservation, ok := m.App.SessionManager.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("can not get item from session")
		m.App.SessionManager.Put(r.Context(), "error", "can not get reservation from session manager")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	data := make(map[string]any)
	data["reservation"] = reservation
	render.RenderTemplate(w, r, "reservation-summery", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
	m.App.SessionManager.Remove(r.Context(), "reservation")
}

//FIXME: do't forget to import the latest templates from git repo by ts
