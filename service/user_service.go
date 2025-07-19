package service

import (
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/constant"
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/dto"
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/entity"
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/middleware"
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/repository"
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/utils"
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type (
	UserService interface {
		Register(ctx context.Context, req dto.RegisterUserRequest) (dto.RegisterUserResponse, error)
		Verify(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error)
	}

	userService struct {
		userRepo repository.UserRepository
	}
)

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) Register(ctx context.Context, req dto.RegisterUserRequest) (dto.RegisterUserResponse, error) {
	_, flag, _ := s.userRepo.CheckEmail(ctx, nil, req.Email)
	if flag {
		return dto.RegisterUserResponse{}, dto.ErrEmailAlreadyExists
	}

	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return dto.RegisterUserResponse{}, dto.ErrHashPass
	}

	user := entity.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hash,
		Role:     constant.ROLE_USER,
	}

	userReg, err := s.userRepo.Register(ctx, nil, user)
	if err != nil {
		return dto.RegisterUserResponse{}, dto.ErrCreateUser
	}

	return dto.RegisterUserResponse{
		Username: userReg.Username,
		Email:    userReg.Email,
	}, nil
}

func (s *userService) Verify(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	// Step 1: Check if email exists in the database
	user, exists, err := s.userRepo.CheckEmail(ctx, nil, req.Email)
	if err != nil {
		return dto.UserLoginResponse{}, dto.ErrInternalServer
	}
	if !exists {
		return dto.UserLoginResponse{}, dto.ErrInvalidCredentials
	}

	// Step 2; Validate the password
	if err := utils.VerifyPassword(user.Password, req.Password); err != nil {
		return dto.UserLoginResponse{}, dto.ErrInvalidCredentials
	}

	// Step 3: Generate a JWT token
	privateKeyBytes, err := middleware.DecodePrivateKeyBase64()
	if err != nil {
		return dto.UserLoginResponse{}, dto.ErrFailedDecodePrivateKey
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return dto.UserLoginResponse{}, dto.ErrInvalidPrivateKeyFormat
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user": user.ID.String(),
		"role": user.Role,
		"exp":  time.Now().Add(15 * time.Minute).Unix(),
		"iat":  time.Now().Unix(),
	})

	// Step 4: Sign the token with the secret key
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return dto.UserLoginResponse{}, dto.ErrFailedCreateToken
	}

	// Step 5: Return the response
	return dto.UserLoginResponse{
		Token: tokenString,
		Role:  user.Role,
	}, nil
}
