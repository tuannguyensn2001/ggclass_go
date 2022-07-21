//go:generate mockgen --source=service.go --destination=repository.mock.go --package=auth
package auth

import (
	"context"
	"errors"
	"ggclass_go/src/app"
	"ggclass_go/src/models"
	"ggclass_go/src/packages/hash"
	"ggclass_go/src/packages/jwt"
	"ggclass_go/src/packages/validate"
	"ggclass_go/src/services/profile"
	"github.com/go-playground/validator/v10"
)

type IRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindById(ctx context.Context, id int) (*models.User, error)
}

type service struct {
	repository     IRepository
	secretKey      string
	profileService IProfileService
}

type IProfileService interface {
	GetByUserId(ctx context.Context, userId int) (*models.Profile, error)
	Create(ctx context.Context, input profile.CreateProfileInput) (*models.Profile, error)
}

func NewService(repository IRepository, secretKey string) *service {
	return &service{repository: repository, secretKey: secretKey}
}

func (s *service) SetProfileService(profileService IProfileService) {
	s.profileService = profileService
}

func (s *service) Register(ctx context.Context, input RegisterInput) (*models.User, error) {

	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		return nil, err
	}

	userWithEmail, err := s.repository.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	if userWithEmail != nil {
		return nil, app.ConflictHttpError("user existed", err)
	}

	password, err := hash.Hash(input.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Username: input.Username,
		Password: password,
		Email:    input.Email,
	}
	err = s.repository.Create(ctx, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (s *service) Login(ctx context.Context, input LoginInput) (*LoginOutput, error) {

	err := validate.Exec(input)

	if err != nil {
		return nil, err
	}

	user, err := s.repository.FindByEmail(ctx, input.Email)

	if err != nil {
		return nil, app.BadRequestHttpError("username or password not valid", err)
	}

	if user == nil {
		return nil, app.BadRequestHttpError("username or password not valid", errors.New("username or password not valid"))
	}

	check := hash.CompareWithContext(ctx, input.Password, user.Password)

	if !check {
		return nil, app.BadRequestHttpError("username or password not valid", err)
	}

	accessToken, err := jwt.GenerateToken(s.secretKey, user.Id, 100000)

	if err != nil {
		return nil, err
	}

	checkProfile, err := s.profileService.GetByUserId(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	if checkProfile == nil {
		checkProfile, err = s.profileService.Create(ctx, profile.CreateProfileInput{
			UserId: user.Id,
			Avatar: "https://cdn-icons-png.flaticon.com/512/1053/1053244.png?w=360",
		})
		if err != nil {
			return nil, err
		}
	}

	user.Profile = *checkProfile

	return &LoginOutput{
		AccessToken: accessToken,
		User:        user,
	}, nil
}

func (s *service) GetUserById(ctx context.Context, id int) (*models.User, error) {
	return s.repository.FindById(ctx, id)
}
