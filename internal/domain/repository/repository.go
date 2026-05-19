package repository

import (
	"context"

	"github.com/heru-oktafian/nrl-be/internal/domain/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AdminUserRepository struct {
	db *pgxpool.Pool
}

func NewAdminUserRepository(db *pgxpool.Pool) *AdminUserRepository {
	return &AdminUserRepository{db: db}
}

func (r *AdminUserRepository) GetByUsername(ctx context.Context, username string) (*entity.AdminUser, error) {
	var u entity.AdminUser
	row := r.db.QueryRow(ctx, `
		SELECT id, username, password, name, email, created_at, updated_at
		FROM admin_users WHERE username = $1
	`, username)
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *AdminUserRepository) GetByID(ctx context.Context, id int) (*entity.AdminUser, error) {
	var u entity.AdminUser
	row := r.db.QueryRow(ctx, `
		SELECT id, username, password, name, email, created_at, updated_at
		FROM admin_users WHERE id = $1
	`, id)
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

type ProfileRepository struct {
	db *pgxpool.Pool
}

func NewProfileRepository(db *pgxpool.Pool) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) Get(ctx context.Context) (*entity.Profile, error) {
	var p entity.Profile
	row := r.db.QueryRow(ctx, `
		SELECT id, name, title, bio, about, email, phone, location, image_url, created_at, updated_at
		FROM profiles WHERE id = $1
	`, 1)
	err := row.Scan(&p.ID, &p.Name, &p.Title, &p.Bio, &p.About, &p.Email, &p.Phone, &p.Location, &p.ImageURL, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProfileRepository) Update(ctx context.Context, p *entity.Profile) error {
	_, err := r.db.Exec(ctx, `
		UPDATE profiles SET name=$1, title=$2, bio=$3, about=$4, email=$5, phone=$6, location=$7, image_url=$8, updated_at=CURRENT_TIMESTAMP
		WHERE id=1
	`, p.Name, p.Title, p.Bio, p.About, p.Email, p.Phone, p.Location, p.ImageURL)
	return err
}

func (r *ProfileRepository) Create(ctx context.Context, p *entity.Profile) error {
	return r.db.QueryRow(ctx, `
		INSERT INTO profiles (name, title, bio, about, email, phone, location, image_url)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at
	`, p.Name, p.Title, p.Bio, p.About, p.Email, p.Phone, p.Location, p.ImageURL).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
}

func (r *ProfileRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM profiles WHERE id = $1`, id)
	return err
}

// Experience CRUD
func (r *ExperienceRepository) GetByID(ctx context.Context, id int) (*entity.Experience, error) {
	var e entity.Experience
	row := r.db.QueryRow(ctx, `
		SELECT id, profile_id, title, company, location, start_date, end_date, current, description, created_at, updated_at
		FROM experiences WHERE id = $1
	`, id)
	err := row.Scan(&e.ID, &e.ProfileID, &e.Title, &e.Company, &e.Location, &e.StartDate, &e.EndDate, &e.Current, &e.Description, &e.CreatedAt, &e.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

func (r *ExperienceRepository) Create(ctx context.Context, e *entity.Experience) error {
	return r.db.QueryRow(ctx, `
		INSERT INTO experiences (profile_id, title, company, location, start_date, end_date, current, description)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at
	`, e.ProfileID, e.Title, e.Company, e.Location, e.StartDate, e.EndDate, e.Current, e.Description).Scan(&e.ID, &e.CreatedAt, &e.UpdatedAt)
}

func (r *ExperienceRepository) Update(ctx context.Context, e *entity.Experience) error {
	_, err := r.db.Exec(ctx, `
		UPDATE experiences SET title=$1, company=$2, location=$3, start_date=$4, end_date=$5, current=$6, description=$7, updated_at=CURRENT_TIMESTAMP
		WHERE id=$8
	`, e.Title, e.Company, e.Location, e.StartDate, e.EndDate, e.Current, e.Description, e.ID)
	return err
}

func (r *ExperienceRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM experiences WHERE id = $1`, id)
	return err
}

// Skill CRUD
func (r *SkillRepository) GetByID(ctx context.Context, id int) (*entity.Skill, error) {
	var s entity.Skill
	row := r.db.QueryRow(ctx, `
		SELECT id, name, level, category, icon, created_at, updated_at
		FROM skills WHERE id = $1
	`, id)
	err := row.Scan(&s.ID, &s.Name, &s.Level, &s.Category, &s.Icon, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *SkillRepository) Create(ctx context.Context, s *entity.Skill) error {
	return r.db.QueryRow(ctx, `
		INSERT INTO skills (name, level, category, icon)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`, s.Name, s.Level, s.Category, s.Icon).Scan(&s.ID, &s.CreatedAt, &s.UpdatedAt)
}

func (r *SkillRepository) Update(ctx context.Context, s *entity.Skill) error {
	_, err := r.db.Exec(ctx, `
		UPDATE skills SET name=$1, level=$2, category=$3, icon=$4, updated_at=CURRENT_TIMESTAMP
		WHERE id=$5
	`, s.Name, s.Level, s.Category, s.Icon, s.ID)
	return err
}

func (r *SkillRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM skills WHERE id = $1`, id)
	return err
}

// Project CRUD
func (r *ProjectRepository) GetByID(ctx context.Context, id int) (*entity.Project, error) {
	var p entity.Project
	row := r.db.QueryRow(ctx, `
		SELECT id, title, description, tech_stack, image_url, demo_url, repo_url, featured, created_at, updated_at
		FROM projects WHERE id = $1
	`, id)
	err := row.Scan(&p.ID, &p.Title, &p.Description, &p.TechStack, &p.ImageURL, &p.DemoURL, &p.RepoURL, &p.Featured, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProjectRepository) Create(ctx context.Context, p *entity.Project) error {
	return r.db.QueryRow(ctx, `
		INSERT INTO projects (title, description, tech_stack, image_url, demo_url, repo_url, featured)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at
	`, p.Title, p.Description, p.TechStack, p.ImageURL, p.DemoURL, p.RepoURL, p.Featured).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
}

func (r *ProjectRepository) Update(ctx context.Context, p *entity.Project) error {
	_, err := r.db.Exec(ctx, `
		UPDATE projects SET title=$1, description=$2, tech_stack=$3, image_url=$4, demo_url=$5, repo_url=$6, featured=$7, updated_at=CURRENT_TIMESTAMP
		WHERE id=$8
	`, p.Title, p.Description, p.TechStack, p.ImageURL, p.DemoURL, p.RepoURL, p.Featured, p.ID)
	return err
}

func (r *ProjectRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM projects WHERE id = $1`, id)
	return err
}

// SocialLink CRUD
func (r *SocialLinkRepository) GetByID(ctx context.Context, id int) (*entity.SocialLink, error) {
	var s entity.SocialLink
	row := r.db.QueryRow(ctx, `
		SELECT id, platform, url, icon, created_at, updated_at
		FROM social_links WHERE id = $1
	`, id)
	err := row.Scan(&s.ID, &s.Platform, &s.URL, &s.Icon, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *SocialLinkRepository) Create(ctx context.Context, s *entity.SocialLink) error {
	return r.db.QueryRow(ctx, `
		INSERT INTO social_links (platform, url, icon)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`, s.Platform, s.URL, s.Icon).Scan(&s.ID, &s.CreatedAt, &s.UpdatedAt)
}

func (r *SocialLinkRepository) Update(ctx context.Context, s *entity.SocialLink) error {
	_, err := r.db.Exec(ctx, `
		UPDATE social_links SET platform=$1, url=$2, icon=$3, updated_at=CURRENT_TIMESTAMP
		WHERE id=$4
	`, s.Platform, s.URL, s.Icon, s.ID)
	return err
}

func (r *SocialLinkRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM social_links WHERE id = $1`, id)
	return err
}

// Tool CRUD
func (r *ToolRepository) GetByID(ctx context.Context, id int) (*entity.Tool, error) {
	var t entity.Tool
	row := r.db.QueryRow(ctx, `
		SELECT id, name, icon, url, created_at, updated_at
		FROM tools WHERE id = $1
	`, id)
	err := row.Scan(&t.ID, &t.Name, &t.Icon, &t.URL, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *ToolRepository) Create(ctx context.Context, t *entity.Tool) error {
	return r.db.QueryRow(ctx, `
		INSERT INTO tools (name, icon, url)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`, t.Name, t.Icon, t.URL).Scan(&t.ID, &t.CreatedAt, &t.UpdatedAt)
}

func (r *ToolRepository) Update(ctx context.Context, t *entity.Tool) error {
	_, err := r.db.Exec(ctx, `
		UPDATE tools SET name=$1, icon=$2, url=$3, updated_at=CURRENT_TIMESTAMP
		WHERE id=$4
	`, t.Name, t.Icon, t.URL, t.ID)
	return err
}

func (r *ToolRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM tools WHERE id = $1`, id)
	return err
}

// ContactMessage CRUD
func (r *ContactMessageRepository) GetByID(ctx context.Context, id int) (*entity.ContactMessage, error) {
	var m entity.ContactMessage
	row := r.db.QueryRow(ctx, `
		SELECT id, name, email, subject, message, read, created_at
		FROM contact_messages WHERE id = $1
	`, id)
	err := row.Scan(&m.ID, &m.Name, &m.Email, &m.Subject, &m.Message, &m.Read, &m.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *ContactMessageRepository) UpdateRead(ctx context.Context, id int, read bool) error {
	_, err := r.db.Exec(ctx, `UPDATE contact_messages SET read=$1 WHERE id=$2`, read, id)
	return err
}

func (r *ContactMessageRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM contact_messages WHERE id = $1`, id)
	return err
}

// Service CRUD
func (r *ServiceRepository) GetByID(ctx context.Context, id int) (*entity.Service, error) {
	var s entity.Service
	row := r.db.QueryRow(ctx, `
		SELECT id, title, description, icon, service_order, created_at, updated_at
		FROM services WHERE id = $1
	`, id)
	err := row.Scan(&s.ID, &s.Title, &s.Description, &s.Icon, &s.Order, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *ServiceRepository) Create(ctx context.Context, s *entity.Service) error {
	return r.db.QueryRow(ctx, `
		INSERT INTO services (title, description, icon, service_order)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`, s.Title, s.Description, s.Icon, s.Order).Scan(&s.ID, &s.CreatedAt, &s.UpdatedAt)
}

func (r *ServiceRepository) Update(ctx context.Context, s *entity.Service) error {
	_, err := r.db.Exec(ctx, `
		UPDATE services SET title=$1, description=$2, icon=$3, service_order=$4, updated_at=CURRENT_TIMESTAMP
		WHERE id=$5
	`, s.Title, s.Description, s.Icon, s.Order, s.ID)
	return err
}

func (r *ServiceRepository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM services WHERE id = $1`, id)
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

type ServiceRepository struct {
	db *pgxpool.Pool
}

func NewServiceRepository(db *pgxpool.Pool) *ServiceRepository {
	return &ServiceRepository{db: db}
}

func (r *ServiceRepository) GetAll(ctx context.Context) ([]entity.Service, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, title, description, icon, service_order, created_at, updated_at
		FROM services ORDER BY service_order ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []entity.Service
	for rows.Next() {
		var s entity.Service
		if err := rows.Scan(&s.ID, &s.Title, &s.Description, &s.Icon, &s.Order, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}
		services = append(services, s)
	}
	return services, nil
}