package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/nrl-be/internal/domain/entity"
	"github.com/heru-oktafian/nrl-be/internal/domain/repository"
	"github.com/heru-oktafian/nrl-be/pkg/response"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	db                *pgxpool.Pool
	profileRepo       *repository.ProfileRepository
	experienceRepo    *repository.ExperienceRepository
	skillRepo         *repository.SkillRepository
	projectRepo       *repository.ProjectRepository
	socialLinkRepo    *repository.SocialLinkRepository
	toolRepo          *repository.ToolRepository
	contactMessageRepo *repository.ContactMessageRepository
}

func New(db *pgxpool.Pool) *Handler {
	return &Handler{
		db:                db,
		profileRepo:       repository.NewProfileRepository(db),
		experienceRepo:    repository.NewExperienceRepository(db),
		skillRepo:         repository.NewSkillRepository(db),
		projectRepo:       repository.NewProjectRepository(db),
		socialLinkRepo:    repository.NewSocialLinkRepository(db),
		toolRepo:          repository.NewToolRepository(db),
		contactMessageRepo: repository.NewContactMessageRepository(db),
	}
}

// GetProfile returns the profile data
func (h *Handler) GetProfile(c *fiber.Ctx) error {
	profile, err := h.profileRepo.Get(c.Context())
	if err != nil {
		return response.NotFound(c, "Profile not found")
	}
	// Convert to JSON-friendly map
	return response.Success(c, fiber.Map{
		"id":         profile.ID,
		"name":       profile.Name,
		"title":      profile.Title,
		"bio":        profile.Bio,
		"email":      profile.Email,
		"phone":      profile.GetPhone(),
		"location":   profile.GetLocation(),
		"image_url":  profile.GetImageURL(),
		"created_at": profile.CreatedAt,
		"updated_at": profile.UpdatedAt,
	})
}

// GetExperiences returns all experiences
func (h *Handler) GetExperiences(c *fiber.Ctx) error {
	experiences, err := h.experienceRepo.GetAll(c.Context())
	if err != nil {
		return response.InternalError(c, "Failed to fetch experiences")
	}
	if experiences == nil {
		experiences = []entity.Experience{}
	}
	return response.Success(c, experiences)
}

// GetSkills returns all skills
func (h *Handler) GetSkills(c *fiber.Ctx) error {
	skills, err := h.skillRepo.GetAll(c.Context())
	if err != nil {
		return response.InternalError(c, "Failed to fetch skills")
	}
	if skills == nil {
		skills = []entity.Skill{}
	}
	return response.Success(c, skills)
}

// GetProjects returns all projects
func (h *Handler) GetProjects(c *fiber.Ctx) error {
	projects, err := h.projectRepo.GetAll(c.Context())
	if err != nil {
		log.Printf("GetProjects error: %v", err)
		return response.InternalError(c, "Failed to fetch projects")
	}
	if projects == nil {
		return response.Success(c, []interface{}{})
	}

	// Convert to JSON-friendly format
	var result []fiber.Map
	for _, p := range projects {
		result = append(result, fiber.Map{
			"id":          p.ID,
			"title":       p.Title,
			"description": p.Description,
			"tech_stack":  p.GetTechStack(),
			"image_url":   p.GetImageURL(),
			"demo_url":    p.GetDemoURL(),
			"repo_url":    p.GetRepoURL(),
			"featured":    p.Featured,
			"created_at":  p.CreatedAt,
			"updated_at":  p.UpdatedAt,
		})
	}
	return response.Success(c, result)
}

// GetSocialLinks returns all social links
func (h *Handler) GetSocialLinks(c *fiber.Ctx) error {
	links, err := h.socialLinkRepo.GetAll(c.Context())
	if err != nil {
		return response.InternalError(c, "Failed to fetch social links")
	}
	if links == nil {
		links = []entity.SocialLink{}
	}
	return response.Success(c, links)
}

// GetTools returns all tools
func (h *Handler) GetTools(c *fiber.Ctx) error {
	tools, err := h.toolRepo.GetAll(c.Context())
	if err != nil {
		return response.InternalError(c, "Failed to fetch tools")
	}
	if tools == nil {
		tools = []entity.Tool{}
	}
	return response.Success(c, tools)
}

// ContactMessage represents the incoming contact form data
type ContactMessage struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

// CreateContactMessage handles contact form submissions
func (h *Handler) CreateContactMessage(c *fiber.Ctx) error {
	var msg ContactMessage
	if err := c.BodyParser(&msg); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	if msg.Name == "" || msg.Email == "" || msg.Message == "" {
		return response.BadRequest(c, "Name, email, and message are required")
	}

	contactMsg := &entity.ContactMessage{
		Name:    msg.Name,
		Email:   msg.Email,
		Subject: msg.Subject,
		Message: msg.Message,
	}

	if err := h.contactMessageRepo.Create(c.Context(), contactMsg); err != nil {
		return response.InternalError(c, "Failed to save message")
	}

	return response.SuccessWithMessage(c, "Message sent successfully", contactMsg)
}

// GetPortfolio returns all portfolio data combined
func (h *Handler) GetPortfolio(c *fiber.Ctx) error {
	profile, err := h.profileRepo.Get(c.Context())
	if err != nil {
		profile = &entity.Profile{}
	}

	experiences, _ := h.experienceRepo.GetAll(c.Context())
	if experiences == nil {
		experiences = []entity.Experience{}
	}

	skills, _ := h.skillRepo.GetAll(c.Context())
	if skills == nil {
		skills = []entity.Skill{}
	}

	projects, _ := h.projectRepo.GetAll(c.Context())
	var projectList []fiber.Map
	if projects != nil {
		for _, p := range projects {
			projectList = append(projectList, fiber.Map{
				"id":          p.ID,
				"title":       p.Title,
				"description": p.Description,
				"tech_stack":  p.GetTechStack(),
				"image_url":   p.GetImageURL(),
				"demo_url":    p.GetDemoURL(),
				"repo_url":    p.GetRepoURL(),
				"featured":    p.Featured,
				"created_at":  p.CreatedAt,
				"updated_at":  p.UpdatedAt,
			})
		}
	}
	if projectList == nil {
		projectList = []fiber.Map{}
	}

	socialLinks, _ := h.socialLinkRepo.GetAll(c.Context())
	if socialLinks == nil {
		socialLinks = []entity.SocialLink{}
	}

	tools, _ := h.toolRepo.GetAll(c.Context())
	if tools == nil {
		tools = []entity.Tool{}
	}

	// Convert profile to JSON-friendly map
	return response.Success(c, fiber.Map{
		"profile": fiber.Map{
			"id":         profile.ID,
			"name":       profile.Name,
			"title":      profile.Title,
			"bio":        profile.Bio,
			"email":      profile.Email,
			"phone":      profile.GetPhone(),
			"location":   profile.GetLocation(),
			"image_url":  profile.GetImageURL(),
			"created_at": profile.CreatedAt,
			"updated_at": profile.UpdatedAt,
		},
		"experiences":  experiences,
		"skills":       skills,
		"projects":     projectList,
		"social_links": socialLinks,
		"tools":        tools,
	})
}