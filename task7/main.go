package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtSecret = []byte("a-string-secret-at-least-256-bits-long")

type CustomClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func AddJWTToContext(ctx context.Context, userID int) (context.Context, error) {
	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return ctx, err
	}

	newCtx := context.WithValue(ctx, "jwt", tokenString)

	return newCtx, nil
}

func ExtractUserForContext(ctx context.Context) (int, error) {
	// TODO извлекаем токен из context
	tokenStr, ok := ctx.Value("jwt").(string)

	if !ok || tokenStr == "" {
		return 0, errors.New("Токен не добавлен в контекст")
	}

	// TODO парсим токен
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return 0, errors.New("Токен не валидный")
	}

	// TODO получаем токен
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return 0, errors.New("Не успешное получение токена")
	}

	return claims.UserID, nil
}

// todo Пример использования
func main() {
	ctx := context.Background()

	ctxWithJWT, err := AddJWTToContext(ctx, 42)
	if err != nil {
		panic(err)
	}

	userID, err := ExtractUserForContext(ctxWithJWT)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Извлеченный userID: %d\n", userID)
}
