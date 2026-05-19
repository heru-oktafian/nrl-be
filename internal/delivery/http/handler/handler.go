package handler

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/heru-oktafian/nrl-be/internal/delivery/http/middleware"
	"github.com/heru-oktafian/nrl-be/internal/domain/entity"
	"github.com/heru-oktafian/nrl-be/internal/domain/repository"
	"github.com/heru-oktafian/nrl-be/pkg/response"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	db                  *pgxpool.Pool
	profileRepo         *repository.ProfileRepository
	experienceRepo      *repository.ExperienceRepository
	skillRepo           *repository.SkillRepository
	projectRepo         *repository.ProjectRepository
	socialLinkRepo      *repository.SocialLinkRepository
	toolRepo            *repository.ToolRepository
	contactMessageRepo  *repository.ContactMessageRepository
	serviceRepo         *repository.ServiceRepository
	adminUserRepo       *repository.AdminUserRepository
}

func New(db *pgxpool.Pool) *Handler {
	return &Handler{
		db:                  db,
		profileRepo:         repository.NewProfileRepository(db),
		experienceRepo:      repository.NewExperienceRepository(db),
		skillRepo:           repository.NewSkillRepository(db),
		projectRepo:         repository.NewProjectRepository(db),
		socialLinkRepo:      repository.NewSocialLinkRepository(db),
		toolRepo:            repository.NewToolRepository(db),
		contactMessageRepo:  repository.NewContactMessageRepository(db),
		serviceRepo:         repository.NewServiceRepository(db),
		adminUserRepo:       repository.NewAdminUserRepository(db),
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

// GetContactMessages returns all contact messages
func (h *Handler) GetContactMessages(c *fiber.Ctx) error {
	messages, err := h.contactMessageRepo.GetAll(c.Context())
	if err != nil {
		return response.InternalError(c, "Failed to fetch messages")
	}
	if messages == nil {
		messages = []entity.ContactMessage{}
	}
	return response.Success(c, messages)
}

// GetServices returns all services
func (h *Handler) GetServices(c *fiber.Ctx) error {
	services, err := h.serviceRepo.GetAll(c.Context())
	if err != nil {
		return response.InternalError(c, "Failed to fetch services")
	}
	if services == nil {
		services = []entity.Service{}
	}
	return response.Success(c, services)
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

	services, _ := h.serviceRepo.GetAll(c.Context())
	if services == nil {
		services = []entity.Service{}
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
		"services":     services,
	})
}

// ============ ADMIN HANDLERS ============

// LoginRequest represents the login request body
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginAdmin handles admin authentication
func (h *Handler) LoginAdmin(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	if req.Username == "" || req.Password == "" {
		return response.BadRequest(c, "Username and password are required")
	}

	user, err := h.adminUserRepo.GetByUsername(c.Context(), req.Username)
	if err != nil {
		return response.Unauthorized(c, "Invalid credentials")
	}

	// For simplicity, we're using plain text password check
	// In production, use bcrypt: bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if user.Password != req.Password {
		return response.Unauthorized(c, "Invalid credentials")
	}

	token, err := middleware.GenerateToken(user.ID, user.Username)
	if err != nil {
		return response.InternalError(c, "Failed to generate token")
	}

	return response.Success(c, fiber.Map{
		"token": token,
		"user": fiber.Map{
			"id":       user.ID,
			"username": user.Username,
			"name":     user.Name,
			"email":    user.Email,
		},
	})
}

// GetAuthMe returns the current authenticated user
func (h *Handler) GetAuthMe(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	user, err := h.adminUserRepo.GetByID(c.Context(), userID)
	if err != nil {
		return response.NotFound(c, "User not found")
	}

	return response.Success(c, fiber.Map{
		"id":       user.ID,
		"username": user.Username,
		"name":     user.Name,
		"email":    user.Email,
	})
}

// UpdateProfile updates the profile data
func (h *Handler) UpdateProfile(c *fiber.Ctx) error {
	var req struct {
		Name     string `json:"name"`
		Title    string `json:"title"`
		Bio      string `json:"bio"`
		About    string `json:"about"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Location string `json:"location"`
		ImageURL string `json:"image_url"`
	}

	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	profile, err := h.profileRepo.Get(c.Context())
	if err != nil {
		// Create if not exists
		profile = &entity.Profile{
			Name: req.Name,
			Title: req.Title,
			Bio: req.Bio,
			Email: req.Email,
		}
		if err := h.profileRepo.Create(c.Context(), profile); err != nil {
			return response.InternalError(c, "Failed to create profile")
		}
		return response.SuccessWithMessage(c, "Profile created", profile)
	}

	profile.Name = req.Name
	profile.Title = req.Title
	profile.Bio = req.Bio
	profile.Email = req.Email

	if req.About != "" {
		profile.About.Valid = true
		profile.About.String = req.About
	}
	if req.Phone != "" {
		profile.Phone.Valid = true
		profile.Phone.String = req.Phone
	}
	if req.Location != "" {
		profile.Location.Valid = true
		profile.Location.String = req.Location
	}
	if req.ImageURL != "" {
		profile.ImageURL.Valid = true
		profile.ImageURL.String = req.ImageURL
	}

	if err := h.profileRepo.Update(c.Context(), profile); err != nil {
		return response.InternalError(c, "Failed to update profile")
	}

	return response.SuccessWithMessage(c, "Profile updated", profile)
}

// GetAdminProfile returns the profile for admin
func (h *Handler) GetAdminProfile(c *fiber.Ctx) error {
	profile, err := h.profileRepo.Get(c.Context())
	if err != nil {
		return response.NotFound(c, "Profile not found")
	}
	return response.Success(c, fiber.Map{
		"id":         profile.ID,
		"name":       profile.Name,
		"title":      profile.Title,
		"bio":        profile.Bio,
		"about":      profile.GetAbout(),
		"email":      profile.Email,
		"phone":      profile.GetPhone(),
		"location":   profile.GetLocation(),
		"image_url":  profile.GetImageURL(),
		"created_at": profile.CreatedAt,
		"updated_at": profile.UpdatedAt,
	})
}

// CreateProfile creates a new profile
func (h *Handler) CreateProfile(c *fiber.Ctx) error {
	var req struct {
		Name     string `json:"name"`
		Title    string `json:"title"`
		Bio      string `json:"bio"`
		About    string `json:"about"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Location string `json:"location"`
		ImageURL string `json:"image_url"`
	}

	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	profile := &entity.Profile{
		Name:  req.Name,
		Title: req.Title,
		Bio:   req.Bio,
		Email: req.Email,
	}

	if req.About != "" {
		profile.About.Valid = true
		profile.About.String = req.About
	}
	if req.Phone != "" {
		profile.Phone.Valid = true
		profile.Phone.String = req.Phone
	}
	if req.Location != "" {
		profile.Location.Valid = true
		profile.Location.String = req.Location
	}
	if req.ImageURL != "" {
		profile.ImageURL.Valid = true
		profile.ImageURL.String = req.ImageURL
	}

	if err := h.profileRepo.Create(c.Context(), profile); err != nil {
		return response.InternalError(c, "Failed to create profile")
	}

	return response.SuccessWithMessage(c, "Profile created", profile)
}

// DeleteProfile deletes a profile
func (h *Handler) DeleteProfile(c *fiber.Ctx) error {
	if err := h.profileRepo.Delete(c.Context(), 1); err != nil {
		return response.InternalError(c, "Failed to delete profile")
	}
	return response.SuccessWithMessage(c, "Profile deleted", nil)
}

// Admin Profile CRUD (using existing profile, id=1)
func (h *Handler) UpsertProfile(c *fiber.Ctx) error {
	var req struct {
		Name     string `json:"name"`
		Title    string `json:"title"`
		Bio      string `json:"bio"`
		About    string `json:"about"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Location string `json:"location"`
		ImageURL string `json:"image_url"`
	}

	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	profile, err := h.profileRepo.Get(c.Context())
	if err != nil {
		return response.NotFound(c, "Profile not found")
	}

	profile.Name = req.Name
	profile.Title = req.Title
	profile.Bio = req.Bio
	profile.Email = req.Email

	if req.About != "" {
		profile.About.Valid = true
		profile.About.String = req.About
	}
	if req.Phone != "" {
		profile.Phone.Valid = true
		profile.Phone.String = req.Phone
	}
	if req.Location != "" {
		profile.Location.Valid = true
		profile.Location.String = req.Location
	}
	if req.ImageURL != "" {
		profile.ImageURL.Valid = true
		profile.ImageURL.String = req.ImageURL
	}

	if err := h.profileRepo.Update(c.Context(), profile); err != nil {
		return response.InternalError(c, "Failed to update profile")
	}

	return response.SuccessWithMessage(c, "Profile updated", fiber.Map{
		"id":         profile.ID,
		"name":       profile.Name,
		"title":      profile.Title,
		"bio":        profile.Bio,
		"about":      profile.GetAbout(),
		"email":      profile.Email,
		"phone":      profile.GetPhone(),
		"location":   profile.GetLocation(),
		"image_url":  profile.GetImageURL(),
		"created_at": profile.CreatedAt,
		"updated_at": profile.UpdatedAt,
	})
}

// Experience Admin CRUD
func (h *Handler) CreateExperience(c *fiber.Ctx) error {
	var req struct {
		Title       string `json:"title"`
		Company     string `json:"company"`
		Location    string `json:"location"`
		StartDate   string `json:"start_date"`
		EndDate     string `json:"end_date"`
		Current     bool   `json:"current"`
		Description string `json:"description"`
	}

	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	exp := &entity.Experience{
		ProfileID:   1,
		Title:       req.Title,
		Company:     req.Company,
		Location:    req.Location,
		Current:     req.Current,
		Description: req.Description,
	}

	// Parse dates
	if req.StartDate != "" {
		t, err := time.Parse("2006-01-02", req.StartDate)
		if err == nil {
			exp.StartDate = &t
		}
	}
	if req.EndDate != "" && !req.Current {
		t, err := time.Parse("2006-01-02", req.EndDate)
		if err == nil {
			exp.EndDate = &t
		}
	}

	if err := h.experienceRepo.Create(c.Context(), exp); err != nil {
		return response.InternalError(c, "Failed to create experience")
	}

	return response.SuccessWithMessage(c, "Experience created", exp)
}

func (h *Handler) UpdateExperience(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	var req struct {
		Title       string `json:"title"`
		Company     string `json:"company"`
		Location    string `json:"location"`
		StartDate   string `json:"start_date"`
		EndDate     string `json:"end_date"`
		Current     bool   `json:"current"`
		Description string `json:"description"`
	}

	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	exp, err := h.experienceRepo.GetByID(c.Context(), id)
	if err != nil {
		return response.NotFound(c, "Experience not found")
	}

	exp.Title = req.Title
	exp.Company = req.Company
	exp.Location = req.Location
	exp.Current = req.Current
	exp.Description = req.Description

	if req.StartDate != "" {
		t, err := time.Parse("2006-01-02", req.StartDate)
		if err == nil {
			exp.StartDate = &t
		}
	}
	if req.EndDate != "" && !req.Current {
		t, err := time.Parse("2006-01-02", req.EndDate)
		if err == nil {
			exp.EndDate = &t
		}
	} else {
		exp.EndDate = nil
	}

	if err := h.experienceRepo.Update(c.Context(), exp); err != nil {
		return response.InternalError(c, "Failed to update experience")
	}

	return response.SuccessWithMessage(c, "Experience updated", exp)
}

func (h *Handler) DeleteExperience(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	if err := h.experienceRepo.Delete(c.Context(), id); err != nil {
		return response.InternalError(c, "Failed to delete experience")
	}

	return response.SuccessWithMessage(c, "Experience deleted", nil)
}

func (h *Handler) GetExperience(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	exp, err := h.experienceRepo.GetByID(c.Context(), id)
	if err != nil {
		return response.NotFound(c, "Experience not found")
	}

	return response.Success(c, exp)
}

// Skill Admin CRUD
func (h *Handler) CreateSkill(c *fiber.Ctx) error {
	var req struct {
		Name     string `json:"name"`
		Level    int    `json:"level"`
		Category string `json:"category"`
		Icon     string `json:"icon"`
	}

	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	skill := &entity.Skill{
		Name:     req.Name,
		Level:    req.Level,
		Category: req.Category,
		Icon:     req.Icon,
	}

	if err := h.skillRepo.Create(c.Context(), skill); err != nil {
		return response.InternalError(c, "Failed to create skill")
	}

	return response.SuccessWithMessage(c, "Skill created", skill)
}

func (h *Handler) UpdateSkill(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	var req struct {
		Name     string `json:"name"`
		Level    int    `json:"level"`
		Category string `json:"category"`
		Icon     string `json:"icon"`
	}

	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	skill, err := h.skillRepo.GetByID(c.Context(), id)
	if err != nil {
		return response.NotFound(c, "Skill not found")
	}

	skill.Name = req.Name
	skill.Level = req.Level
	skill.Category = req.Category
	skill.Icon = req.Icon

	if err := h.skillRepo.Update(c.Context(), skill); err != nil {
		return response.InternalError(c, "Failed to update skill")
	}

	return response.SuccessWithMessage(c, "Skill updated", skill)
}

func (h *Handler) DeleteSkill(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	if err := h.skillRepo.Delete(c.Context(), id); err != nil {
		return response.InternalError(c, "Failed to delete skill")
	}

	return response.SuccessWithMessage(c, "Skill deleted", nil)
}

func (h *Handler) GetSkill(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	skill, err := h.skillRepo.GetByID(c.Context(), id)
	if err != nil {
		return response.NotFound(c, "Skill not found")
	}

	return response.Success(c, skill)
}

// Project Admin CRUD
func (h *Handler) CreateProject(c *fiber.Ctx) error {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		TechStack   string `json:"tech_stack"`
		ImageURL    string `json:"image_url"`
		DemoURL     string `json:"demo_url"`
		RepoURL     string `json:"repo_url"`
		Featured    bool   `json:"featured"`
	}

	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	project := &entity.Project{
		Title:       req.Title,
		Description: req.Description,
		Featured:    req.Featured,
	}

	if req.TechStack != "" {
		project.TechStack.Valid = true
		project.TechStack.String = req.TechStack
	}
	if req.ImageURL != "" {
		project.ImageURL.Valid = true
		project.ImageURL.String = req.ImageURL
	}
	if req.DemoURL != "" {
		project.DemoURL.Valid = true
		project.DemoURL.String = req.DemoURL
	}
	if req.RepoURL != "" {
		project.RepoURL.Valid = true
		project.RepoURL.String = req.RepoURL
	}

	if err := h.projectRepo.Create(c.Context(), project); err != nil {
		return response.InternalError(c, "Failed to create project")
	}

	return response.SuccessWithMessage(c, "Project created", project)
}

func (h *Handler) UpdateProject(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		TechStack   string `json:"tech_stack"`
		ImageURL    string `json:"image_url"`
		DemoURL     string `json:"demo_url"`
		RepoURL     string `json:"repo_url"`
		Featured    bool   `json:"featured"`
	}

	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	project, err := h.projectRepo.GetByID(c.Context(), id)
	if err != nil {
		return response.NotFound(c, "Project not found")
	}

	project.Title = req.Title
	project.Description = req.Description
	project.Featured = req.Featured

	if req.TechStack != "" {
		project.TechStack.Valid = true
		project.TechStack.String = req.TechStack
	}
	if req.ImageURL != "" {
		project.ImageURL.Valid = true
		project.ImageURL.String = req.ImageURL
	}
	if req.DemoURL != "" {
		project.DemoURL.Valid = true
		project.DemoURL.String = req.DemoURL
	}
	if req.RepoURL != "" {
		project.RepoURL.Valid = true
		project.RepoURL.String = req.RepoURL
	}

	if err := h.projectRepo.Update(c.Context(), project); err != nil {
		return response.InternalError(c, "Failed to update project")
	}

	return response.SuccessWithMessage(c, "Project updated", project)
}

func (h *Handler) DeleteProject(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	if err := h.projectRepo.Delete(c.Context(), id); err != nil {
		return response.InternalError(c, "Failed to delete project")
	}

	return response.SuccessWithMessage(c, "Project deleted", nil)
}

func (h *Handler) GetProject(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	project, err := h.projectRepo.GetByID(c.Context(), id)
	if err != nil {
		return response.NotFound(c, "Project not found")
	}

	return response.Success(c, fiber.Map{
		"id":          project.ID,
		"title":       project.Title,
		"description": project.Description,
		"tech_stack":  project.GetTechStack(),
		"image_url":   project.GetImageURL(),
		"demo_url":    project.GetDemoURL(),
		"repo_url":    project.GetRepoURL(),
		"featured":    project.Featured,
		"created_at":  project.CreatedAt,
		"updated_at":  project.UpdatedAt,
	})
}

// SocialLink Admin CRUD
func (h *Handler) CreateSocialLink(c *fiber.Ctx) error {
	var req struct {
		Platform string `json:"platform"`
		URL      string `json:"url"`
		Icon     string `json:"icon"`
	}

	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	link := &entity.SocialLink{
		Platform: req.Platform,
		URL:      req.URL,
		Icon:     req.Icon,
	}

	if err := h.socialLinkRepo.Create(c.Context(), link); err != nil {
		return response.InternalError(c, "Failed to create social link")
	}

	return response.SuccessWithMessage(c, "Social link created", link)
}

func (h *Handler) UpdateSocialLink(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	var req struct {
		Platform string `json:"platform"`
		URL      string `json:"url"`
		Icon     string `json:"icon"`
	}

	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	link, err := h.socialLinkRepo.GetByID(c.Context(), id)
	if err != nil {
		return response.NotFound(c, "Social link not found")
	}

	link.Platform = req.Platform
	link.URL = req.URL
	link.Icon = req.Icon

	if err := h.socialLinkRepo.Update(c.Context(), link); err != nil {
		return response.InternalError(c, "Failed to update social link")
	}

	return response.SuccessWithMessage(c, "Social link updated", link)
}

func (h *Handler) DeleteSocialLink(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	if err := h.socialLinkRepo.Delete(c.Context(), id); err != nil {
		return response.InternalError(c, "Failed to delete social link")
	}

	return response.SuccessWithMessage(c, "Social link deleted", nil)
}

func (h *Handler) GetSocialLink(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	link, err := h.socialLinkRepo.GetByID(c.Context(), id)
	if err != nil {
		return response.NotFound(c, "Social link not found")
	}

	return response.Success(c, link)
}

// Tool Admin CRUD
func (h *Handler) CreateTool(c *fiber.Ctx) error {
	var req struct {
		Name string `json:"name"`
		Icon string `json:"icon"`
		URL  string `json:"url"`
	}

	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	tool := &entity.Tool{
		Name: req.Name,
		Icon: req.Icon,
		URL:  req.URL,
	}

	if err := h.toolRepo.Create(c.Context(), tool); err != nil {
		return response.InternalError(c, "Failed to create tool")
	}

	return response.SuccessWithMessage(c, "Tool created", tool)
}

func (h *Handler) UpdateTool(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	var req struct {
		Name string `json:"name"`
		Icon string `json:"icon"`
		URL  string `json:"url"`
	}

	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	tool, err := h.toolRepo.GetByID(c.Context(), id)
	if err != nil {
		return response.NotFound(c, "Tool not found")
	}

	tool.Name = req.Name
	tool.Icon = req.Icon
	tool.URL = req.URL

	if err := h.toolRepo.Update(c.Context(), tool); err != nil {
		return response.InternalError(c, "Failed to update tool")
	}

	return response.SuccessWithMessage(c, "Tool updated", tool)
}

func (h *Handler) DeleteTool(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	if err := h.toolRepo.Delete(c.Context(), id); err != nil {
		return response.InternalError(c, "Failed to delete tool")
	}

	return response.SuccessWithMessage(c, "Tool deleted", nil)
}

func (h *Handler) GetTool(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	tool, err := h.toolRepo.GetByID(c.Context(), id)
	if err != nil {
		return response.NotFound(c, "Tool not found")
	}

	return response.Success(c, tool)
}

// Service Admin CRUD
func (h *Handler) CreateService(c *fiber.Ctx) error {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
		Order       int    `json:"order"`
	}

	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	service := &entity.Service{
		Title:       req.Title,
		Description: req.Description,
		Icon:        req.Icon,
		Order:       req.Order,
	}

	if err := h.serviceRepo.Create(c.Context(), service); err != nil {
		return response.InternalError(c, "Failed to create service")
	}

	return response.SuccessWithMessage(c, "Service created", service)
}

func (h *Handler) UpdateService(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
		Order       int    `json:"order"`
	}

	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	service, err := h.serviceRepo.GetByID(c.Context(), id)
	if err != nil {
		return response.NotFound(c, "Service not found")
	}

	service.Title = req.Title
	service.Description = req.Description
	service.Icon = req.Icon
	service.Order = req.Order

	if err := h.serviceRepo.Update(c.Context(), service); err != nil {
		return response.InternalError(c, "Failed to update service")
	}

	return response.SuccessWithMessage(c, "Service updated", service)
}

func (h *Handler) DeleteService(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	if err := h.serviceRepo.Delete(c.Context(), id); err != nil {
		return response.InternalError(c, "Failed to delete service")
	}

	return response.SuccessWithMessage(c, "Service deleted", nil)
}

func (h *Handler) GetService(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	service, err := h.serviceRepo.GetByID(c.Context(), id)
	if err != nil {
		return response.NotFound(c, "Service not found")
	}

	return response.Success(c, service)
}

// ContactMessage Admin CRUD
func (h *Handler) UpdateContactMessage(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	var req struct {
		Read bool `json:"read"`
	}

	if err := c.BodyParser(&req); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	if err := h.contactMessageRepo.UpdateRead(c.Context(), id, req.Read); err != nil {
		return response.InternalError(c, "Failed to update contact message")
	}

	return response.SuccessWithMessage(c, "Contact message updated", nil)
}

func (h *Handler) DeleteContactMessage(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	if err := h.contactMessageRepo.Delete(c.Context(), id); err != nil {
		return response.InternalError(c, "Failed to delete contact message")
	}

	return response.SuccessWithMessage(c, "Contact message deleted", nil)
}

func (h *Handler) GetContactMessage(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return response.BadRequest(c, "Invalid ID")
	}

	msg, err := h.contactMessageRepo.GetByID(c.Context(), id)
	if err != nil {
		return response.NotFound(c, "Contact message not found")
	}

	return response.Success(c, msg)
}

// GetCounts returns counts for dashboard
func (h *Handler) GetCounts(c *fiber.Ctx) error {
	projects, _ := h.projectRepo.GetAll(c.Context())
	experiences, _ := h.experienceRepo.GetAll(c.Context())
	skills, _ := h.skillRepo.GetAll(c.Context())
	messages, _ := h.contactMessageRepo.GetAll(c.Context())

	unreadCount := 0
	for _, m := range messages {
		if !m.Read {
			unreadCount++
		}
	}

	return response.Success(c, fiber.Map{
		"projects":      len(projects),
		"experiences":   len(experiences),
		"skills":        len(skills),
		"messages":      len(messages),
		"unread":        unreadCount,
	})
}