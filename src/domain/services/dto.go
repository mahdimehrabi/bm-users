package services

type JWTDTO struct {
	AccessToken     string
	RefreshToken    string
	AccessTokenExp  uint64
	RefreshTokenExp uint64
}
