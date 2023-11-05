package ctxutils

type (
	appContextKey string
)

const (
	// AppContextKey is the key for the app context.
	AppContextKey appContextKey = "appContext"
	// UsernameContextKey is the key for the username context.
	UsernameContextKey appContextKey = "username"
)
