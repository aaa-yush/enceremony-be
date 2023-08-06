package authorizer

import (
	"context"
	"enceremony-be/internal/auth/models"
	"enceremony-be/internal/common/logger"
	"enceremony-be/internal/config"
	models2 "enceremony-be/internal/database/mysql/models"
	"enceremony-be/internal/user/repo"
	userModel "enceremony-be/pkg/user"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/markbates/goth"
)

type Service interface {
	CreateJWTToken(ctx context.Context, claims *models.JwtPayload) (string, error)
	VerifyAndCreateUser(ctx context.Context, user *goth.User) (*userModel.ProfileResponse, error)
}

type impl struct {
	logg      *logger.Logger
	conf      *config.Config
	secretKey string
	userRepo  repo.UserRepo
}

func (i *impl) VerifyAndCreateUser(ctx context.Context, user *goth.User) (*userModel.ProfileResponse, error) {

	createUserReq := &models2.User{
		SourceId: user.UserID,
		Name:     user.Name,
		Email:    user.Email,
	}

	userDetails, err := i.userRepo.GetOrCreateUser(ctx, createUserReq)
	if err != nil {
		return nil, err
	}

	jwtClaims := &models.JwtPayload{
		Uid:          string(userDetails.Id),
		FirstName:    userDetails.Name,
		Email:        userDetails.Email,
		UserId:       string(userDetails.Id),
		GoogleUserId: user.UserID,
		//DOB:          "",
	}

	token, err := i.CreateJWTToken(ctx, jwtClaims)
	if err != nil {
		return nil, err
	}

	return &userModel.ProfileResponse{
		Id:           jwtClaims.Uid,
		Name:         userDetails.Name,
		Email:        userDetails.Email,
		SessionToken: token,
	}, nil
}

func NewAuthorizerService(logg *logger.Logger,
	userRepo repo.UserRepo,
	conf *config.Config) Service {
	return &impl{
		logg:      logg,
		conf:      conf,
		secretKey: conf.Auth.SecretKey,
		userRepo:  userRepo,
	}
}

func (i *impl) CreateJWTToken(ctx context.Context, claims *models.JwtPayload) (string, error) {

	secretBytes, err := json.Marshal(i.secretKey)

	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secretBytes)
	if err != nil {
		return "", fmt.Errorf("secret.JWTTokenString: sign JWT: %v", err)
	}

	return tokenString, nil
}
