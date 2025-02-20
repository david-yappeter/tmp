package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	jwtInternal "myapp/internal/jwt"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface {
	Login(ctx context.Context, request dto_request.AuthLoginRequest) (*model.Token, *dto_response.ErrorResponse, error)
	Register(ctx context.Context, request dto_request.AuthRegisterRequest) (*model.Token, *dto_response.ErrorResponse, error)
	Parse(ctx context.Context, token string) (*model.User, *dto_response.ErrorResponse, error)
}

type authUseCase struct {
	repositoryManager repository.RepositoryManager
	jwt               jwtInternal.Jwt
}

func NewAuthUseCase(
	repositoryManager repository.RepositoryManager,
	jwt jwtInternal.Jwt,
) AuthUseCase {
	return &authUseCase{
		repositoryManager: repositoryManager,
		jwt:               jwt,
	}
}

func (u *authUseCase) generateJWT(ctx context.Context, userId string) (*jwtInternal.Token, error) {
	currentTime := time.Now()
	accessToken, err := u.jwt.Generate(jwtInternal.Payload{
		UserId:    userId,
		CreatedAt: currentTime,
		ExpiredAt: currentTime.Add(time.Hour * 24), // 1 day
	})
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (u *authUseCase) validateEmailUnique(ctx context.Context, email string) (bool, error) {
	isExist, err := u.repositoryManager.UserRepository().IsExistByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	return !isExist, nil
}

func (u *authUseCase) Login(ctx context.Context, request dto_request.AuthLoginRequest) (*model.Token, *dto_response.ErrorResponse, error) {
	user, err := u.repositoryManager.UserRepository().GetByEmail(ctx, request.Email)
	if err != nil {
		return nil, nil, util.NewInternalServerError(err.Error())
	}
	if user == nil {
		return nil, dto_response.NewNotFoundErrorResponseP("User not found"), nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil, dto_response.NewBadRequestErrorResponseP("Incorrect password"), nil
		}
		return nil, nil, util.NewInternalServerError(err.Error())
	}

	jwtToken, err := u.generateJWT(ctx, user.Id)
	if err != nil {
		return nil, nil, util.NewInternalServerError(err.Error())
	}

	return &model.Token{
		AccessToken: jwtToken.AccessToken,
		ExpiredAt:   jwtToken.ExpiredAt,
		TokenType:   jwtToken.TokenType,
	}, nil, nil
}

func (u *authUseCase) Register(ctx context.Context, request dto_request.AuthRegisterRequest) (*model.Token, *dto_response.ErrorResponse, error) {
	isUnique, err := u.validateEmailUnique(ctx, request.Email)
	if err != nil {
		return nil, nil, util.NewInternalServerError(err.Error())
	}
	if !isUnique {
		return nil, dto_response.NewBadRequestErrorResponseP("Email already exist"), nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, util.NewInternalServerError(err.Error())
	}

	user := model.User{
		Id:       util.NewUuid(),
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	err = u.repositoryManager.UserRepository().Insert(ctx, &user)
	if err != nil {
		return nil, nil, util.NewInternalServerError("failed to create user")
	}

	jwtToken, err := u.generateJWT(ctx, user.Id)
	if err != nil {
		return nil, nil, util.NewInternalServerError(err.Error())
	}

	return &model.Token{
		AccessToken: jwtToken.AccessToken,
		ExpiredAt:   jwtToken.ExpiredAt,
		TokenType:   jwtToken.TokenType,
	}, nil, nil
}

func (u *authUseCase) Parse(ctx context.Context, token string) (*model.User, *dto_response.ErrorResponse, error) {
	payload, err := u.jwt.Parse(token)
	if err != nil {
		return nil, dto_response.NewBadRequestErrorResponseP("not authenticated"), nil
	}

	var (
		userId = payload.UserId
	)

	if payload.ExpiredAt.Before(time.Now()) {
		return nil, dto_response.NewBadRequestErrorResponseP("not authenticated"), nil
	}

	user, err := u.repositoryManager.UserRepository().Get(ctx, userId)
	if err != nil {
		return nil, nil, err
	}
	if user == nil {
		return nil, dto_response.NewBadRequestErrorResponseP("not authenticated"), nil
	}

	return user, nil, nil
}
