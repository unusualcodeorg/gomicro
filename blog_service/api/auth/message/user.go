package message

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoleCode string

const (
	RoleCodeLearner RoleCode = "LEARNER"
	RoleCodeAdmin   RoleCode = "ADMIN"
	RoleCodeAuthor  RoleCode = "AUTHOR"
	RoleCodeEditor  RoleCode = "EDITOR"
)

type User struct {
	ID            primitive.ObjectID `json:"_id"`
	Name          string             `json:"name"`
	Email         string             `json:"email"`
	ProfilePicURL *string            `json:"profilePicUrl,omitempty"`
}
