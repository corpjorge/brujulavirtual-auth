package models

type Auth struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func (a *Auth) IsValid() bool {
	return a.User != "" && a.Password != ""
}
