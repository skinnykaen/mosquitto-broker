package app

import (
	"fmt"
	"net/http"
	u "github.com/skinnykaen/go.git/utils"
	"strings"
	models "github.com/skinnykaen/go.git/models"
	jwt "github.com/dgrijalva/jwt-go"
	"context"
)

var jwtKey = []byte("my_secret_key")

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/", "/registration" , "/login"}
		requestPath := r.URL.Path

		//проверяем, не требует ли запрос аутентификации, обслуживаем запрос, если он не нужен
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		response := make(map[string] interface{})
		tokenHeader := r.Header.Get("Authorization") //Получение токена

		if tokenHeader == "" { //Токен отсутствует, возвращаем  403 http-код Unauthorized
			response = u.Message(false, "Missing auth token")
			fmt.Println("Missing auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}
		fmt.Println("here jwt auth")
		splitted := strings.Split(tokenHeader, " ") //`Bearer {token-body}`,  соответствует ли полученный токен этому требованию
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			fmt.Println("Invalid/Malformed auth token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		tokenPart := splitted[1] //Получаем вторую часть токена
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil { //Неправильный токен, как правило, возвращает 403 http-код
			response = u.Message(false, "Malformed authentication token")
			fmt.Println("Malformed authentication token")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}

		if !token.Valid { //токен недействителен, возможно, не подписан на этом сервере
			response = u.Message(false, "Token is not valid.")
			fmt.Println("Token is not valid.")
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			u.Respond(w, response)
			return
		}
		//Всё прошло хорошо, продолжаем выполнение запроса
		//fmt.Sprintf("User %", tk.Username)
		fmt.Println("Succes!")
		ctx := context.WithValue(r.Context(), "user", tk.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //передать управление следующему обработчику
	});
}
