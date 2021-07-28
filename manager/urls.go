package manager

const (
	USER_DATA_STATE   = "user-data-state"
	LOGIN_TOKEN_STATE = "login-token-state"

	SECRET_STORE               = "google-secret-state"
	GOOGLE_OAUTH_PUBLIC_SECRET = "google-oauth-public"
	PASETO_COMMON_KEY_SECRET   = ""
	SEED_SECRET                = "seed-value"
)

var GoogleOauthKey string
var SeedValue string
