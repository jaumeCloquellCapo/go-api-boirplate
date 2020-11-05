package model

//TokenDetails ...
type TokenDetails struct {
	Token     string
	UUID      string
	AtExpires int64
	RtExpires int64
}

//AccessDetails ...
type AccessDetails struct {
	AccessUUID string
	UserID     int64
}

//Token ...
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
