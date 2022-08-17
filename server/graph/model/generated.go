// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AcceptInvitation struct {
	InvitationID string `json:"invitationId"`
}

type ActivateUser struct {
	ID   string `json:"id"`
	Code string `json:"code"`
}

type AuthMutation struct {
	Login    *Token `json:"login"`
	Register *Token `json:"register"`
}

type CommentMutation struct {
	Like *Comment `json:"like"`
}

type CommentPost struct {
	CommenterID        string  `json:"commenterId"`
	Text               string  `json:"text"`
	PostID             string  `json:"postId"`
	RepliedToCommentID *string `json:"repliedToCommentId"`
}

type ConnectionMutation struct {
	Create *ConnectInvitation `json:"create"`
	Accept *ConnectInvitation `json:"accept"`
	Reject *ConnectInvitation `json:"reject"`
}

type CreateEducation struct {
	UserID string `json:"userId"`
	School string `json:"school"`
}

type CreateExperience struct {
	UserID         string  `json:"userId"`
	Title          string  `json:"title"`
	EmploymentType string  `json:"employmentType"`
	IsActive       bool    `json:"isActive"`
	CompanyName    string  `json:"companyName"`
	LocationCity   string  `json:"locationCity"`
	LocationRegion string  `json:"locationRegion"`
	StartDateMonth string  `json:"startDateMonth"`
	StartDateYear  int     `json:"startDateYear"`
	EndDateMonth   *string `json:"endDateMonth"`
	EndDateYear    *int    `json:"endDateYear"`
	Industry       string  `json:"industry"`
	Headline       *string `json:"headline"`
}

type CreateInvitation struct {
	FromUserID string  `json:"fromUserId"`
	ToUserID   string  `json:"toUserId"`
	Note       *string `json:"note"`
}

type CreateJob struct {
	Title          string `json:"title"`
	CompanyName    string `json:"companyName"`
	Workplace      string `json:"workplace"`
	LocationCity   string `json:"locationCity"`
	LocationRegion string `json:"locationRegion"`
	EmploymentType string `json:"employmentType"`
	Description    string `json:"description"`
}

type CreateMessage struct {
	SenderID   string  `json:"senderId"`
	ReceiverID string  `json:"receiverId"`
	Text       string  `json:"text"`
	ImageURL   *string `json:"imageUrl"`
}

type CreateNotification struct {
	UserID string `json:"userId"`
	Text   string `json:"text"`
}

type CreatePost struct {
	Text     string  `json:"text"`
	PosterID string  `json:"posterId"`
	PhotoURL *string `json:"photoUrl"`
	VideoURL *string `json:"videoUrl"`
}

type DeleteEducation struct {
	EducationID string `json:"educationId"`
}

type DeleteExperience struct {
	ExperienceID string `json:"experienceId"`
}

type EducationMutation struct {
	Create *Education `json:"create"`
	Update *Education `json:"update"`
	Delete *Education `json:"delete"`
}

type ExperienceMutation struct {
	Create *Experience `json:"create"`
	Update *Experience `json:"update"`
	Delete *Experience `json:"delete"`
}

type FollowUser struct {
	UserID      string `json:"userId"`
	FollowingID string `json:"followingId"`
}

type ForgotPasswordCode struct {
	Code string `json:"code"`
}

type ForgotPasswordEmail struct {
	Email string `json:"email"`
}

type JobMutation struct {
	Create *Job `json:"create"`
}

type LikeComment struct {
	LikerID   string `json:"likerId"`
	CommentID string `json:"commentId"`
}

type LikePost struct {
	LikerID string `json:"likerId"`
	PostID  string `json:"postId"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MessageMutation struct {
	Create *Message `json:"create"`
}

type NotificationMutation struct {
	Create *Notification `json:"create"`
}

type PostFeed struct {
	UserID string `json:"userId"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

type PostMutation struct {
	Create  *Post `json:"create"`
	Like    *Post `json:"like"`
	Comment *Post `json:"comment"`
	Share   *Post `json:"share"`
}

type RegisterUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type RejectInvitation struct {
	InvitationID string `json:"invitationId"`
}

type ResetPassword struct {
	UserID          string `json:"userId"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type SharePost struct {
	SourceID      string `json:"sourceId"`
	PostID        string `json:"postId"`
	DestinationID string `json:"destinationId"`
}

type Token struct {
	Token string `json:"token"`
}

type UpdateEducation struct {
	EducationID    string   `json:"educationId"`
	School         *string  `json:"school"`
	Degree         *string  `json:"degree"`
	Field          *string  `json:"field"`
	StartDateMonth *string  `json:"startDateMonth"`
	StartDateYear  *int     `json:"startDateYear"`
	EndDateMonth   *string  `json:"endDateMonth"`
	EndDateYear    *int     `json:"endDateYear"`
	Grade          *float64 `json:"grade"`
	Activities     *string  `json:"activities"`
	Description    *string  `json:"description"`
}

type UpdateExperience struct {
	ExperienceID   string  `json:"experienceId"`
	Title          *string `json:"title"`
	EmploymentType *string `json:"employmentType"`
	CompanyName    *string `json:"companyName"`
	IsActive       *bool   `json:"isActive"`
	LocationCity   *string `json:"locationCity"`
	LocationRegion *string `json:"locationRegion"`
	StartDateMonth *string `json:"startDateMonth"`
	StartDateYear  *int    `json:"startDateYear"`
	EndDateMonth   *string `json:"endDateMonth"`
	EndDateYear    *int    `json:"endDateYear"`
	Industry       *string `json:"industry"`
	Headline       *string `json:"headline"`
}

type UpdateUser struct {
	UserID             string  `json:"userId"`
	Email              *string `json:"email"`
	FirstName          *string `json:"firstName"`
	LastName           *string `json:"lastName"`
	AdditionalName     *string `json:"additionalName"`
	ProfilePhotoURL    *string `json:"profilePhotoUrl"`
	BackgroundPhotoURL *string `json:"backgroundPhotoUrl"`
	Headline           *string `json:"headline"`
	Pronouns           *string `json:"pronouns"`
	ProfileLink        *string `json:"profileLink"`
	About              *string `json:"about"`
	LocationCity       *string `json:"locationCity"`
	LocationRegion     *string `json:"locationRegion"`
	IsActive           *bool   `json:"isActive"`
}

type UserMutation struct {
	View                      *User `json:"view"`
	Follow                    *User `json:"follow"`
	UnFollow                  *User `json:"unFollow"`
	Update                    *User `json:"update"`
	Activate                  *User `json:"activate"`
	VerifyForgotPasswordEmail *User `json:"verifyForgotPasswordEmail"`
	VerifyForgotPasswordCode  *User `json:"verifyForgotPasswordCode"`
	ResetPassword             *User `json:"resetPassword"`
}

type ViewUser struct {
	ViewerID     string `json:"viewerId"`
	ViewedUserID string `json:"viewedUserId"`
}
