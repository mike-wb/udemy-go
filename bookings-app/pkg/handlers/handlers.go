package handlers

import (
	"net/http"

	"github.com/tsawler/bookings-app/pkg/config"
	"github.com/tsawler/bookings-app/pkg/models"
	"github.com/tsawler/bookings-app/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		Page: r.URL.Path,
	})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// send data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		Page: r.URL.Path,
	})
}

// Contact is the handler for the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	// send data to the template
	render.RenderTemplate(w, "contact.page.tmpl", &models.TemplateData{
		Page: r.URL.Path,
	})
}

// Room is the handler for the rooms page
func (m *Repository) Room(w http.ResponseWriter, r *http.Request) {
	// send data to the template
	if r.URL.Path == "/rooms/generals-quarters" {
		render.RenderTemplate(w, "generals.page.tmpl", &models.TemplateData{
			Page: r.URL.Path,
		})
	} else if r.URL.Path == "/rooms/majors-suite" {
		render.RenderTemplate(w, "majors.page.tmpl", &models.TemplateData{
			Page: r.URL.Path,
		})
	} else {
		Repo.Home(w, r)
	}
}

// Reservation is the handler for the reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	// send data to the template
	render.RenderTemplate(w, "reservation.page.tmpl", &models.TemplateData{
		Page: r.URL.Path,
	})
}

// MakeReservation is the handler for the make-reservation page
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	// send data to the template
	render.RenderTemplate(w, "make-reservation.page.tmpl", &models.TemplateData{
		Page: r.URL.Path,
	})
}
