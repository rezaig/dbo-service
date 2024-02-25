package helper

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rezaig/dbo-service/internal/config"
	"github.com/rezaig/dbo-service/internal/model"
	log "github.com/sirupsen/logrus"
)

func GenerateJWTToken(claims jwt.MapClaims) (token string, err error) {
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.JWTSigningKey()))
	if err != nil {
		log.Errorf("error signing token: %v", err)
	}
	return
}

func DecodeJWTToken(tokenString string) (*model.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSigningKey()), nil
	})
	if err != nil {
		log.Errorf("error parsing token: %v", err)
		return nil, err
	}

	if !token.Valid {
		log.Errorf("invalid token: %v", err)
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*model.CustomClaims)
	if !ok {
		log.Warn("error getting claims")
		return nil, errors.New("error getting claims")
	}

	return claims, nil
}
