package model

type CreateMenuRequest struct {
	// User    string `bson:"user" json:"user" `
	// Title   string `bson:"title" json:"title" validate:"required"`
	// Slug    string `bson:"slug" json:"slug" `
	// Summary string `bson:"summary" json:"summary" `
	// Type    string `bson:"type" json:"type" `
	// Created string `bson:"created" json:"created" `
	// Updated string `bson:"updated" json:"updated" `
	// Content string `bson:"content" json:"content" `

	User    string `bson:"user" json:"user"`
	Title   string `bson:"title" json:"title" validate:"required"`
	Slug    string `bson:"slug" json:"slug"`
	Summary string `bson:"summary" json:"summary"`
	Type    string `bson:"type" json:"type"`
	Created string `bson:"created" json:"created"`
	Updated string `bson:"updated" json:"updated"`
	Content string `bson:"content" json:"content"`
}
