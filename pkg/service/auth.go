package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	todo "github.com/jpegShawty/go_todo_app/pkg"
	"github.com/jpegShawty/go_todo_app/pkg/repository"
)

const (
	salt = "skdffiqwfiksfk"
	signingKey = "sodfkoqsdsqoadkaod"
	tokenTTL = 12 * time.Hour
)
	
// зависит не от конкретного AuthPostgre, а от Authorization-интерфейса, Почему это круто?
// Протестировать AuthService, заменив repo на фейковый мок, а не реальную БД.
// поменять реализацию без боли.
type AuthService struct {
	repo repository.Authorization
}

type TokenClaims struct{
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService{
	return &AuthService{repo: repo}
}	

func (s *AuthService) CreateUser(user todo.User) (int, error){
	user.Password = generatePasswordHash(user.Password)
// Следующий CreateUser уже идет из 
	return s.repo.CreateUser(user)
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil{
		return 0, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *TokenClaims")
	}

	return claims.UserId, nil

}

func (s *AuthService) GenerateToken(username, password string) (string, error){
	//get user from DB
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil{
		return "", err
	}

// стандартный метод для подписи и claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),	
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func generatePasswordHash(password string) string{
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}