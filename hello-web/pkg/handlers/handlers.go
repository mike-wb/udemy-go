package handlers

import (
	"net/http"

	"github.com/mike-wb/udemy-go/pkg/config"
	"github.com/mike-wb/udemy-go/pkg/models"
	"github.com/mike-wb/udemy-go/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic here
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!"

	// send the data to the template
	render.RenderTemplates(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
