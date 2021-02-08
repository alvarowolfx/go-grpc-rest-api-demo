package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"com.aviebrantz.tvtime/internal/application/users"
)

type AuthService interface {
	Login(ctx context.Context, username, password string) (string, error)
	Verify(ctx context.Context, token string) error
	GetUserID(ctx context.Context, token string) (string, error)
}

type ServiceDeps struct {
	UserService users.Service
	JWTSecret   string
}

type authService struct {
	ServiceDeps
}

func NewService(deps ServiceDeps) AuthService {
	return &authService{deps}
}

func (s *authService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.UserService.CheckUser(ctx, username, password)
	if err != nil {
		return "", fmt.Errorf("user credential check failed: %v", err)
	}
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return "", err
	}
	return t, nil
}

func (s *authService) keyFunc(t *jwt.Token) (interface{}, error) {
	if t.Method.Alg() != jwt.SigningMethodHS256.Name {
		return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
	}
	return []byte(s.JWTSecret), nil
}

func (s *authService) Verify(ctx context.Context, token string) error {
	_, err := s.parseToken(ctx, token)
	if err != nil {
		return fmt.Errorf("failed to parse jwt :%v", err)
	}
	return err
}

func (s *authService) GetUserID(ctx context.Context, token string) (string, error) {
	jwtToken, err := s.parseToken(ctx, token)
	if err != nil {
		return "", err
	}
	claims := jwtToken.Claims.(jwt.MapClaims)
	return claims["sub"].(string), nil
}

func (s *authService) parseToken(ctx context.Context, token string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(token, s.keyFunc)
	if err != nil {
		return nil, fmt.Errorf("failed to parse jwt :%v", err)
	}
	return jwtToken, nil
}
