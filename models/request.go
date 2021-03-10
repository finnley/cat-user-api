package models

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	ID          uint
	NickName    string
	AuthorityId uint // 角色
	jwt.StandardClaims
}
