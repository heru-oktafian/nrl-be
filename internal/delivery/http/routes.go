package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/nrl-be/internal/delivery/http/handler"
	"github.com/heru-oktafian/nrl-be/internal/delivery/http/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

// SetupRoutes configures all API routes
func SetupRoutes(app *fiber.App, db *pgxpool.Pool) {
	h := handler.New(db)

	// API v1 group
	api := app.Group("/api/v1")

	// Health check
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"service": "nrl-be",
		})
	})

	// Public portfolio routes (for frontend)
	public := api.Group("/public")
	public.Get("/portfolio", h.GetPortfolio)
	public.Get("/profile", h.GetProfile)
	public.Get("/experiences", h.GetExperiences)
	public.Get("/skills", h.GetSkills)
	public.Get("/projects", h.GetProjects)
	public.Get("/social-links", h.GetSocialLinks)
	public.Get("/tools", h.GetTools)
	public.Get("/services", h.GetServices)
	public.Post("/contact", h.CreateContactMessage)
	public.Get("/contact-messages", h.GetContactMessages)

	// Admin authentication routes (no auth required)
	adminAuth := api.Group("/admin/auth")
	adminAuth.Post("/login", h.LoginAdmin)
	adminAuth.Post("/logout", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"success": true, "message": "Logged out successfully"})
	})
	adminAuth.Get("/me", middleware.JWTAuth(), h.GetAuthMe)

	// Admin routes (JWT auth required)
	admin := api.Group("/admin", middleware.JWTAuth())
	
	// Profile
	admin.Get("/profile", h.GetAdminProfile)
	admin.Put("/profile", h.UpsertProfile)
	
	// Projects
	admin.Get("/projects", h.GetProjects)
	admin.Get("/projects/:id", h.GetProject)
	admin.Post("/projects", h.CreateProject)
	admin.Put("/projects/:id", h.UpdateProject)
	admin.Delete("/projects/:id", h.DeleteProject)
	
	// Skills
	admin.Get("/skills", h.GetSkills)
	admin.Get("/skills/:id", h.GetSkill)
	admin.Post("/skills", h.CreateSkill)
	admin.Put("/skills/:id", h.UpdateSkill)
	admin.Delete("/skills/:id", h.DeleteSkill)
	
	// Experiences
	admin.Get("/experiences", h.GetExperiences)
	admin.Get("/experiences/:id", h.GetExperience)
	admin.Post("/experiences", h.CreateExperience)
	admin.Put("/experiences/:id", h.UpdateExperience)
	admin.Delete("/experiences/:id", h.DeleteExperience)
	
	// Social Links
	admin.Get("/social-links", h.GetSocialLinks)
	admin.Get("/social-links/:id", h.GetSocialLink)
	admin.Post("/social-links", h.CreateSocialLink)
	admin.Put("/social-links/:id", h.UpdateSocialLink)
	admin.Delete("/social-links/:id", h.DeleteSocialLink)
	
	// Tools
	admin.Get("/tools", h.GetTools)
	admin.Get("/tools/:id", h.GetTool)
	admin.Post("/tools", h.CreateTool)
	admin.Put("/tools/:id", h.UpdateTool)
	admin.Delete("/tools/:id", h.DeleteTool)
	
	// Services
	admin.Get("/services", h.GetServices)
	admin.Get("/services/:id", h.GetService)
	admin.Post("/services", h.CreateService)
	admin.Put("/services/:id", h.UpdateService)
	admin.Delete("/services/:id", h.DeleteService)
	
	// Contact Messages
	admin.Get("/contact-messages", h.GetContactMessages)
	admin.Get("/contact-messages/:id", h.GetContactMessage)
	admin.Put("/contact-messages/:id", h.UpdateContactMessage)
	admin.Delete("/contact-messages/:id", h.DeleteContactMessage)
	
	// Dashboard counts
	admin.Get("/counts", h.GetCounts)
}