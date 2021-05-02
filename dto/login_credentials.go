package dto

// LoginCredentials are used when a user is logging in. Fields which declare a form tag are meant to be used by Gin's
// binding system, while fields without are updated post-login and are meant to be used in the creation of a JWT.
type LoginCredentials struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
