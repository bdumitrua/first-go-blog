package auth

import (
	"errors"
	"first-blog-api/utils"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type key string

const UserIDKey key = "userId"
const JwtKey key = "jwt"

var jwtKey = []byte("your_secret_key") // Секретный ключ для подписи токенов

// Claims структура для данных внутри JWT
type Claims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

// Генерация JWT токена
func GenerateJWT(userId int) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour) // Время жизни токена

	claims := &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // Время истечения
		},
	}

	// Создаём токен с алгоритмом подписи HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен секретным ключом
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// Парсим и проверяем токен
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Проверяем алгоритм подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Проверяем валидность токена
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func GetUserId(req *utils.Request) (int, error) {
	userID, ok := req.GetRequest().Context().Value(UserIDKey).(int)
	if !ok {
		http.Error(req.Writer(), "Unauthorized: user ID not found", http.StatusUnauthorized)
		return 0, errors.New("unauthorized")
	}

	return userID, nil
}

func GetJwtToken(req *utils.Request) (string, error) {
	userID, ok := req.GetRequest().Context().Value(jwtKey).(string)
	if !ok {
		http.Error(req.Writer(), "Unauthorized: user ID not found", http.StatusUnauthorized)
		return "", errors.New("unauthorized")
	}

	return userID, nil
}
