package handlers

import (
	"net/http"
	"path"

	"github.com/mike-wb/udemy-go/bookings-app/pkg/config"
	"github.com/mike-wb/udemy-go/bookings-app/pkg/models"
	"github.com/mike-wb/udemy-go/bookings-app/pkg/render"
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

// Rooms is the handler to show the list of rooms
func (m *Repository) Rooms(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "rooms.page.tmpl", &models.TemplateData{
		Page: r.URL.Path,
		Data: map[string]interface{}{
			"rooms": models.Rooms,
		},
	})
}

// Room is the handler for an individual room page
func (m *Repository) Room(w http.ResponseWriter, r *http.Request) {
	// send data to the template
	path := path.Base(path.Clean(r.URL.Path))

	room, ok := models.Rooms[path]

	if ok {
		render.RenderTemplate(w, "room.page.tmpl", &models.TemplateData{
			Page: r.URL.Path,
			Data: map[string]interface{}{
				"room": room,
			},
		})
	} else {
		http.Redirect(w, r, "/rooms", http.StatusTemporaryRedirect)
	}
}

// Availability renders the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	// send data to the template
	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{
		Page: r.URL.Path,
	})
}

// PostAvailability renders the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	// send data to the template
	render.RenderTemplate(w, "search-availability.page.tmpl", &models.TemplateData{
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
