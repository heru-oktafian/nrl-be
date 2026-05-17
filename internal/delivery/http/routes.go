package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/nrl-be/internal/delivery/http/handler"
	"github.com/jackc/pgx/v5/pgxpool"
)

// SetupRoutes configures all API routes
func SetupRoutes(app *fiber.App, db *pgxpool.Pool) {
	h := handler.New(db)

	// API v1 group
	api := app.Group("/api/v1")

	// Public portfolio routes
	api.Get("/portfolio", h.GetPortfolio)
	api.Get("/profile", h.GetProfile)
	api.Get("/experiences", h.GetExperiences)
	api.Get("/skills", h.GetSkills)
	api.Get("/projects", h.GetProjects)
	api.Get("/social-links", h.GetSocialLinks)
	api.Get("/tools", h.GetTools)

	// Contact form submission
	api.Post("/contact", h.CreateContactMessage)
}
