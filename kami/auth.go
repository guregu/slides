package auth

type Auth struct {
	Key     string
	Session string
	UserID  string
}

func Check(r *http.Request) (auth Auth, ok bool) {
	// lots of complicated checks
}
