package middleware

import (
	"context"
	"first-blog-api/auth"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Получаем токен из заголовка
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// Убираем префикс "Bearer "
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		// Валидация токена
		claims, err := auth.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Сохраняем данные токена в контекст (если нужно)
		ctx := context.WithValue(r.Context(), auth.UserIDKey, claims.UserId)
		ctx = context.WithValue(ctx, auth.JwtKey, tokenString)
		r = r.WithContext(ctx)

		// Передаём управление следующему обработчику
		next.ServeHTTP(w, r)
	})
}
