package models

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	News_id        uuid.UUID
	Category_id    uuid.UUID
	Author_id      uuid.UUID
	News_title     string
	News_image     string
	News_content   string
	Comment_status bool
	Created_at     time.Time
	Updated_at     time.Time
}

type Category struct {
	Category_id          uuid.UUID
	Category_title       string
	Category_description string
	Created_at           time.Time
	Updated_at           time.Time
}

type Author struct {
	Author_id     uuid.UUID
	First_name    string
	Last_name     string
	Gender        string
	Birth_date    time.Time
	Email         string
	Password      string
	Profile_image string
	Last_login    time.Time
	Created_at    time.Time
	Updated_at    time.Time
}
