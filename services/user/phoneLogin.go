package user

import (
	errorss "KpChatGpt/e"
	"KpChatGpt/e/errors_const"
	"KpChatGpt/middleware/jwt"
)

// FindUserByPhone
func FindUserByPhone(phone string) (string, string, error) {

	accessToken, refreshToken, err := jwt.GenerateToken(phone, "123456")
	if err != nil {
		return "", "", errorss.HandleError(errors_const.ErrGenerateToken, "zn", err)
	}

	return accessToken, refreshToken, nil
}
