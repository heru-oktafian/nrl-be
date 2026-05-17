package entity

import "time"

type Profile struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Bio       string    `json:"bio"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Location  string    `json:"location"`
	ImageURL  string    `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TechStack   string    `json:"tech_stack"`
	ImageURL    string    `json:"image_url"`
	DemoURL     string    `json:"demo_url"`
	RepoURL     string    `json:"repo_url"`
	Featured    bool      `json:"featured"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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