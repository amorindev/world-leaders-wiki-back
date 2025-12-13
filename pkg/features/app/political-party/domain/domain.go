package domain

import "time"

type PoliticalParty struct {
	ID          interface{}   `json:"id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	Acronym     *string       `json:"acronym" bson:"acronym"`
	Country     string        `json:"country" bson:"country"`
	LogoUrl     *string       `json:"logo_url" bson:"-"`
	LogoPath    *string       `json:"-" bson:"logo_path"`
	Description string        `json:"description" bson:"description"`
	Website     *string       `json:"website" bson:"website"`
	Founders    []interface{} `json:"founders" bson:"founders"`
	FoundedAt   *time.Time    `json:"founded_at" bson:"founded_at"`
	CreatedAt   *time.Time    `json:"created_at" bson:"created_at"`
}
