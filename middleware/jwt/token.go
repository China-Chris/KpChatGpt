package jwt

import (
	"KpChatGpt/handle/response"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strings"
	"time"
)

var accessTokenSecret = []byte("accessTokenSecret")
var refreshTokenSecret = []byte("refreshTokenSecret")

const (
	Minute30 = 30 * time.Minute
	OneWeek  = 7 * 24 * time.Hour
)

type RefreshAtAndRt struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type User struct {
	Phone    string `json:"phone"`
	PassWord string `json:"passWord"`
}

type Claims struct {
	UserName string
	PassWord string
	jwt.StandardClaims
}

// GenerateToken 生成令牌
//accessTokenString(AT)，refreshTokenString(RT).
//当AT失效且RT有效时才会重新生产AT,RT.当AT和RT都失效时才会让用户重新登录.当AT未失效不会使用到RT.
func GenerateToken(userName, passWord string) (string, string, error) {
	shortExpireTime := time.Now().Add(Minute30)
	longExpireTime := time.Now().Add(OneWeek)
	accessTokenClaims := &Claims{
		UserName: userName,
		PassWord: passWord,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: shortExpireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",
			Subject:   "user token",
		},
	}
	refreshTokenClaims := &Claims{
		UserName: userName,
		PassWord: passWord,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: longExpireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",
			Subject:   "user token",
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(accessTokenSecret)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString(refreshTokenSecret)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	return accessTokenString, refreshTokenString, nil
}

// ParseToken 解析Token
func ParseToken(accessTokenString, refreshTokenString string) (*Claims, bool, error) {
	fmt.Println("In ParseToken")
	accessToken, err := jwt.ParseWithClaims(accessTokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return accessTokenSecret, nil
	})
	if claims, ok := accessToken.Claims.(*Claims); ok && accessToken.Valid {
		return claims, false, nil
	}
	fmt.Println("RefreshToken")
	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return refreshTokenSecret, nil
	})
	if err != nil {
		return nil, false, err
	}
	if claims, ok := refreshToken.Claims.(*Claims); ok && refreshToken.Valid {
		return claims, true, nil
	}
	return nil, false, errors.New("invalid token")
}

func authHandler(ctx *gin.Context) {
	fmt.Println("In authHandler")
	var user User
	err := ctx.ShouldBind(&user)
	if err != nil {
		response.JsonMessage(ctx, 1001, "请求头中Auth为空")
		return
	}
	fmt.Println("user = ", user)
	if !(user.Phone == "ar" && user.PassWord == "123456") {
		response.JsonMessage(ctx, 1002, "请求头中Auth为空")
		fmt.Println("User not exist or password error")
		return
	}
	accessTokenString, refreshTokenString, _ := GenerateToken(user.Phone, user.PassWord)
	response.JsonSuccess(ctx, RefreshAtAndRt{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	})
}

// AuthMiddleware 用鉴权到中间件
func AuthMiddleware(ctx *gin.Context) {
	// 默认双Token放在请求头Authorization的Bearer中，并以空格隔开
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		response.JsonMessage(ctx, 1002, "请求头中Auth为空")
		return
	}
	fmt.Println("authHeader = ", authHeader)
	parts := strings.Split(authHeader, " ")
	fmt.Println("len = ", len(parts))
	fmt.Println("parts[0] = ", parts[0])
	if !(len(parts) == 3 && parts[0] == "Bearer") {
		response.JsonMessage(ctx, 1003, "请求头中Auth格式有误")
		return
	}
	parseToken, isUpd, err := ParseToken(parts[1], parts[2])
	if err != nil {
		response.JsonMessage(ctx, 1004, "无效的Token")
		return
	}
	// AT失效，刷新双AT,RT
	if isUpd {
		parts[1], parts[2], err = GenerateToken(parseToken.UserName, parseToken.PassWord)
		if err != nil {
			response.JsonMessage(ctx, 1005, "获取令牌失败")
			return
		}
		response.JsonSuccess(ctx, RefreshAtAndRt{
			AccessToken:  parts[1],
			RefreshToken: parts[2],
		})
		ctx.Set("userName", parseToken.UserName)
		ctx.Next()
	}
}
