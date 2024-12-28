package handlers

import (
	"fmt"
	"net/http"

	"github.com/AtharvBhujbal/bookings/pkg/config"
	"github.com/AtharvBhujbal/bookings/pkg/models"
	"github.com/AtharvBhujbal/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// Create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RequestURI
	fmt.Printf("Remote IP: %s\n", remoteIP)
	fmt.Printf("Remote IP type: %T\n", remoteIP)
	// m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["Test"] = "Hello Again."

	// remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	// stringMap["remote_ip"] = remoteIp

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
