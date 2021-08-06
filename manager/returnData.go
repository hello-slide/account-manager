package manager

type ReturnData struct {
	RefreshToken string `json:"refresh_token"`
	Session      string `json:"session_token"`
}
