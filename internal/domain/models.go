package domain

import (
	"time"
)

type User struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	IsActive  bool      `json:"IsActive" db:"is_active"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type Team struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type PullRequest struct {
	Id       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	AuthorID int    `json:"authorId" db:"author_id"`
	Status   string `json:"status" db:"status"`

	Reviewer1ID *int      `json:"reviewer1ID" db:"reviewr1_id"`
	Reviewer2ID *int      `json:"reviewer2ID" db:"reviewr2_id"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
}

type UserTeam struct {
	UserID int `db:"user_id"`
	TeamID int `db:"team_id"`
}
