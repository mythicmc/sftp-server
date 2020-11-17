package main

// AuthenticationRequest ... An authentication request to the SFTP server.
type AuthenticationRequest struct {
	User          string `json:"username"`
	Pass          string `json:"password"`
	IP            string `json:"ip"`
	SessionID     []byte `json:"session_id"`
	ClientVersion []byte `json:"client_version"`
}

// AuthenticationResponse ... An authentication response from the SFTP server.
type AuthenticationResponse struct {
	Server      string   `json:"server"`
	Token       string   `json:"token"`
	Permissions []string `json:"permissions"`
}

// InvalidCredentialsError ... An error emitted when credentials are invalid.
type InvalidCredentialsError struct {
}

func (ice InvalidCredentialsError) Error() string {
	return "the credentials provided were invalid"
}

// IsInvalidCredentialsError ... Checks if an error is a InvalidCredentialsError.
func IsInvalidCredentialsError(err error) bool {
	_, ok := err.(*InvalidCredentialsError)

	return ok
}
