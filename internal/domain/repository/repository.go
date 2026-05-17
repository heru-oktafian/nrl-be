package repository

import (
	"context"

	"github.com/heru-oktafian/nrl-be/internal/domain/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProfileRepository struct {
	db *pgxpool.Pool
}

func NewProfileRepository(db *pgxpool.Pool) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) Get(ctx context.Context) (*entity.Profile, error) {
	var p entity.Profile
	err := r.db.QueryRow(ctx, `
		SELECT id, name, title, bio, email, phone, location, image_url, created_at, updated_at
		FROM profiles WHERE id = 1
	`).Scan(&p.ID, &p.Name, &p.Title, &p.Bio, &p.Email, &p.Phone, &p.Location, &p.ImageURL, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProfileRepository) Update(ctx context.Context, p *entity.Profile) error {
	_, err := r.db.Exec(ctx, `
		UPDATE profiles SET name=$1, title=$2, bio=$3, email=$4, phone=$5, location=$6, image_url=$7, updated_at=CURRENT_TIMESTAMP
		WHERE id=1
	`, p.Name, p.Title, p.Bio, p.Email, p.Phone, p.Location, p.ImageURL)
	return err
}

type ExperienceRepository struct {
	db *pgxpool.Pool
}

func NewExperienceRepository(db *pgxpool.Pool) *ExperienceRepository {
	return &ExperienceRepository{db: db}
}

func (r *ExperienceRepository) GetAll(ctx context.Context) ([]entity.Experience, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, profile_id, title, company, location, start_date, end_date, current, description, created_at, updated_at
		FROM experiences ORDER BY start_date DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var experiences []entity.Experience
	for rows.Next() {
		var e entity.Experience
		if err := rows.Scan(&e.ID, &e.ProfileID, &e.Title, &e.Company, &e.Location, &e.StartDate, &e.EndDate, &e.Current, &e.Description, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, err
		}
		experiences = append(experiences, e)
	}
	return experiences, nil
}

type SkillRepository struct {
	db *pgxpool.Pool
}

func NewSkillRepository(db *pgxpool.Pool) *SkillRepository {
	return &SkillRepository{db: db}
}

func (r *SkillRepository) GetAll(ctx context.Context) ([]entity.Skill, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, name, level, category, icon, created_at, updated_at
		FROM skills ORDER BY level DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var skills []entity.Skill
	for rows.Next() {
		var s entity.Skill
		if err := rows.Scan(&s.ID, &s.Name, &s.Level, &s.Category, &s.Icon, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}
		skills = append(skills, s)
	}
	return skills, nil
}

type ProjectRepository struct {
	db *pgxpool.Pool
}

func NewProjectRepository(db *pgxpool.Pool) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) GetAll(ctx context.Context) ([]entity.Project, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, title, description, tech_stack, image_url, demo_url, repo_url, featured, created_at, updated_at
		FROM projects ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []entity.Project
	for rows.Next() {
		var p entity.Project
		if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.TechStack, &p.ImageURL, &p.DemoURL, &p.RepoURL, &p.Featured, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	return projects, nil
}

type SocialLinkRepository struct {
	db *pgxpool.Pool
}

func NewSocialLinkRepository(db *pgxpool.Pool) *SocialLinkRepository {
	return &SocialLinkRepository{db: db}
}

func (r *SocialLinkRepository) GetAll(ctx context.Context) ([]entity.SocialLink, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, platform, url, icon, created_at, updated_at
		FROM social_links ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []entity.SocialLink
	for rows.Next() {
		var s entity.SocialLink
		if err := rows.Scan(&s.ID, &s.Platform, &s.URL, &s.Icon, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}
		links = append(links, s)
	}
	return links, nil
}

type ToolRepository struct {
	db *pgxpool.Pool
}

func NewToolRepository(db *pgxpool.Pool) *ToolRepository {
	return &ToolRepository{db: db}
}

func (r *ToolRepository) GetAll(ctx context.Context) ([]entity.Tool, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, name, icon, url, created_at, updated_at
		FROM tools ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tools []entity.Tool
	for rows.Next() {
		var t entity.Tool
		if err := rows.Scan(&t.ID, &t.Name, &t.Icon, &t.URL, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tools = append(tools, t)
	}
	return tools, nil
}

type ContactMessageRepository struct {
	db *pgxpool.Pool
}

func NewContactMessageRepository(db *pgxpool.Pool) *ContactMessageRepository {
	return &ContactMessageRepository{db: db}
}

func (r *ContactMessageRepository) Create(ctx context.Context, msg *entity.ContactMessage) error {
	return r.db.QueryRow(ctx, `
		INSERT INTO contact_messages (name, email, subject, message)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`, msg.Name, msg.Email, msg.Subject, msg.Message).Scan(&msg.ID, &msg.CreatedAt)
}

func (r *ContactMessageRepository) GetAll(ctx context.Context) ([]entity.ContactMessage, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, name, email, subject, message, read, created_at
		FROM contact_messages ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []entity.ContactMessage
	for rows.Next() {
		var m entity.ContactMessage
		if err := rows.Scan(&m.ID, &m.Name, &m.Email, &m.Subject, &m.Message, &m.Read, &m.CreatedAt); err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}
	return messages, nil
}