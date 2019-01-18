package models

/*
This struct is used to read the response
from the Google OAuth service
 */
type GoogleAuthResponse struct {
	Id string `json:"id"`
	Email string `json:"email"`
	VerifiedEmail bool `json:"verified_email"`
	Name string `json:"name"`
	GivenName string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Link string `json:"link"`
	Avatar string `json:"picture"`
	Gender string `json:"gender"`
	Locale string `json:"locale"`
}
