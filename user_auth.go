package inforusms

// UserAuth holds fields for User Authentication
type UserAuth struct {
	UserName string `xml:"User>Username"`
	Password string `xml:"User>Password"`
}
