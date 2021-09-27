package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/mike-wb/udemy-go/bookings-app/internal/config"
	"github.com/mike-wb/udemy-go/bookings-app/internal/forms"
	"github.com/mike-wb/udemy-go/bookings-app/internal/models"
	"github.com/mike-wb/udemy-go/bookings-app/internal/render"
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
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
		Page: "Home",
		Path: "home",
	})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// send data to the template
	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		Page: "About",
		Path: "about",
	})
}

// Contact is the handler for the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	// send data to the template
	render.RenderTemplate(w, r, "contact.page.tmpl", &models.TemplateData{
		Page: "Contact Us",
		Path: "contact",
	})
}

// Rooms is the handler to show the list of rooms
func (m *Repository) Rooms(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "rooms.page.tmpl", &models.TemplateData{
		Page: "Our Rooms",
		Path: "rooms",
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
		render.RenderTemplate(w, r, "room.page.tmpl", &models.TemplateData{
			Page: room.Name,
			Path: "rooms",
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
	render.RenderTemplate(w, r, "search-availability.page.tmpl", &models.TemplateData{
		Page: "Search Availability",
		Path: "search-availability",
	})
}

// PostAvailability renders the response for the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	end := r.FormValue("end")
	log.Printf("Start: %s - End: %s\n", start, end)
	w.Write([]byte(fmt.Sprintf("Posted to search availability - Start: %s - End: %s\n", start, end)))
}

// JSON Response structure containing availability for a single room
type jsonresponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON provides the JSON response for the room search availability page
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	end := r.FormValue("end")
	room := r.FormValue("room")
	log.Printf("Start: %s - End: %s - Room: %s\n", start, end, room)
	resp := jsonresponse{
		OK:      true,
		Message: "Available!",
	}
	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Reservation is the handler for the make-reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	// send data to the template
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Page: "Make Reservation",
		Path: "search-availability",
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Phone:     r.Form.Get("phone"),
		Email:     r.Form.Get("email"),
	}

	form := forms.New(r.PostForm)
	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 2)
	form.MinLength("last_name", 2)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation.page.tmpl", &models.TemplateData{
			Page: "Make Reservation",
			Path: "search-availability",
			Form: form,
			Data: data,
		})
		return
	}
	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther)
}

// ReservationSummary handles the rendering of the Reservation Summary page
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)

	if !ok {
		log.Println("cannot get item from session")
		m.App.Session.Put(r.Context(), "error", "Can't get reservation from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation
	render.RenderTemplate(w, r, "reservation-summary.page.tmpl", &models.TemplateData{
		Page: "Reservation Summary",
		Path: "search-availability",
		Data: data,
	})
}
