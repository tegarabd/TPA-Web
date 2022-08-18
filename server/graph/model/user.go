package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID                 string               `json:"id"`
	Email              string               `json:"email"`
	FirstName          string               `json:"firstName"`
	LastName           string               `json:"lastName"`
	Password           string               `json:"password"`
	AdditionalName     *string              `json:"additionalName"`
	ProfilePhotoURL    *string              `json:"profilePhotoUrl"`
	BackgroundPhotoURL *string              `json:"backgroundPhotoUrl"`
	Headline           *string              `json:"headline"`
	Pronouns           *string              `json:"pronouns"`
	ProfileLink        string               `json:"profileLink"`
	About              *string              `json:"about"`
	Location           *Location            `json:"location" gorm:"embedded"`
	ProfileViews       int                  `json:"profileViews"`
	IsActive           bool                 `json:"isActive"`
	ActivationCode ActivationCode
	ResetPasswordCode ResetPasswordCode
	Experiences        []*Experience        `json:"experiences"`
	Educations         []*Education         `json:"educations"`
	Connections        []*User              `json:"connections" gorm:"many2many:user_connections"`
	Following          []*User              `json:"following" gorm:"many2many:user_follow"`
	Posts              []*Post              `json:"posts" gorm:"foreignKey:PosterID"`
	Invitations        []*ConnectInvitation `json:"invitations" gorm:"foreignKey:FromID;foreignKey:ToID"`
	Notifications      []*Notification      `json:"notifications" gorm:"foreignKey:FromID"`
	Messages           []*Message           `json:"messages" gorm:"foreignKey:SenderID;foreignKey:ReceiverID"`
}