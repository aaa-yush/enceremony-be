package user

type ProfileResponse struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	SessionToken string `json:"session_token,omitempty"`
}
