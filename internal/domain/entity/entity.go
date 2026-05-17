package entity

import (
	"database/sql"
	"time"
)

type Profile struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Title     string         `json:"title"`
	Bio       string         `json:"bio"`
	Email     string         `json:"email"`
	Phone     sql.NullString `json:"-"`
	Location  sql.NullString `json:"-"`
	ImageURL  sql.NullString `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

// Getters for JSON serialization
func (p *Profile) GetPhone() string {
	if p.Phone.Valid {
		return p.Phone.String
	}
	return ""
}

func (p *Profile) GetLocation() string {
	if p.Location.Valid {
		return p.Location.String
	}
	return ""
}

func (p *Profile) GetImageURL() string {
	if p.ImageURL.Valid {
		return p.ImageURL.String
	}
	return ""
}

type Experience struct {
	ID          int        `json:"id"`
	ProfileID   int        `json:"profile_id"`
	Title       string     `json:"title"`
	Company     string     `json:"company"`
	Location    string     `json:"location"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Current     bool       `json:"current"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type Skill struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Level     int       `json:"level"`
	Category  string    `json:"category"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Project struct {
	ID          int
	Title       string
	Description string
	TechStack   sql.NullString `json:"-"`
	ImageURL    sql.NullString `json:"-"`
	DemoURL     sql.NullString `json:"-"`
	RepoURL     sql.NullString `json:"-"`
	Featured    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Getters for JSON serialization
func (p *Project) GetTechStack() string {
	if p.TechStack.Valid {
		return p.TechStack.String
	}
	return ""
}

func (p *Project) GetImageURL() string {
	if p.ImageURL.Valid {
		return p.ImageURL.String
	}
	return ""
}

func (p *Project) GetDemoURL() string {
	if p.DemoURL.Valid {
		return p.DemoURL.String
	}
	return ""
}

func (p *Project) GetRepoURL() string {
	if p.RepoURL.Valid {
		return p.RepoURL.String
	}
	return ""
}

type SocialLink struct {
	ID        int       `json:"id"`
	Platform  string    `json:"platform"`
	URL       string    `json:"url"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Tool struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ContactMessage struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Subject   string    `json:"subject"`
	Message   string    `json:"message"`
	Read      bool      `json:"read"`
	CreatedAt time.Time `json:"created_at"`
}