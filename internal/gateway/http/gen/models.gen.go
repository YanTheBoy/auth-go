// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package gen

// BuildInfo defines model for BuildInfo.
type BuildInfo struct {
	// Arch Architecture of the machine used for the build
	Arch string `json:"arch"`

	// BuildDate Date of the build
	BuildDate string `json:"build_date"`

	// CommitHash Commit hash of the source code
	CommitHash string `json:"commit_hash"`

	// Compiler Compiler used for the build
	Compiler string `json:"compiler"`

	// GoVersion Go programming language version used for the build
	GoVersion string `json:"go_version"`

	// Os Operating system used for the build
	Os string `json:"os"`

	// Version Version of the build
	Version string `json:"version"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	// Error Description of the error
	Error string `json:"error"`
}

// LoginUserRequest defines model for LoginUserRequest.
type LoginUserRequest struct {
	// Password Password of the existing user
	Password string `json:"password"`

	// Username Username of the existing user
	Username string `json:"username"`
}

// LoginUserResponse defines model for LoginUserResponse.
type LoginUserResponse struct {
	// AccessToken JWT token for accessing protected routes
	AccessToken string `json:"accessToken"`

	// RefreshToken JWT token for refresh uses
	RefreshToken string `json:"refreshToken"`
}

// RegisterUserRequest defines model for RegisterUserRequest.
type RegisterUserRequest struct {
	Age *int `json:"age,omitempty"`

	// Password Password of the new user
	Password string `json:"password"`

	// Username Username of the new user
	Username string `json:"username"`
}

// RegisterUserResponse defines model for RegisterUserResponse.
type RegisterUserResponse struct {
	// Id Unique identifier for the registered user
	Id int `json:"id"`

	// Username Username of the newly registered user
	Username string `json:"username"`
}

// PostLoginJSONRequestBody defines body for PostLogin for application/json ContentType.
type PostLoginJSONRequestBody = LoginUserRequest

// PostRegisterJSONRequestBody defines body for PostRegister for application/json ContentType.
type PostRegisterJSONRequestBody = RegisterUserRequest
