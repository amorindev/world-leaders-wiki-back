package domain

import "time"

// Leader represents a political or public figure within the system.
// It contains personal information, biography, social media links, and metadata.
// ImgUrl is used only for responses to the client (not stored in the database).
// ImgPath is stored in the database and used by the backend to build the final public URL.
type Leader struct {
	ID         interface{} `json:"id" bson:"_id"`
	Slug       string      `json:"slug" bson:"slug"`
	FullName   string      `json:"full_name" bson:"full_name"`
	NickName   *string     `json:"nickname" bson:"nickname"`
	Phrase     *string     `json:"phrase" bson:"phrase"`
	Biography  string      `json:"biography" bson:"biography"`
	AvatarUrl  *string     `json:"avatar_url" bson:"-"`
	BannerUrl  *string     `json:"banner_url" bson:"-"`
	AvatarPath *string     `json:"-" bson:"avatar_path"`
	BannerPath *string     `json:"-" bson:"banner_path"`

	BirthDate   *time.Time `json:"birth_date" bson:"birth_date"`
	Nationality string     `json:"nationality" bson:"nationality"`
	Gender      string     `json:"gender,omitempty" bson:"gender"`
	Ideology    *string    `json:"ideology,omitempty" bson:"ideology"`

	Facebook  string  `json:"facebook" bson:"facebook"`
	Instagram string  `json:"instagram" bson:"instagram"`
	Twitter   *string `json:"twitter" bson:"twitter"`
	YouTube   *string `json:"youtube" bson:"youtube"`
	Linkedin  *string `json:"linkedin" bson:"linkedin"`
	Website   *string `json:"website" bson:"website"`

	CreatedAt *time.Time `json:"created_at" bson:"created_at"`
}
