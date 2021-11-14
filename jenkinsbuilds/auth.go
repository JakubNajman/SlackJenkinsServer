package jenkinsbuilds

type Authentication struct {
	User  string `json:"user"`
	Token string `json:"token"`
}

func (auth *Authentication) SetUser(user string) {
	auth.User = user
}

func (auth *Authentication) SetToken(token string) {
	auth.Token = token
}

func (auth Authentication) GetUser() string {
	return auth.User
}

func (auth Authentication) GetToken() string {
	return auth.Token
}
