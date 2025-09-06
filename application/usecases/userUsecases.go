package usecases

import (
	"errors"
	"strings"
	"time"
	"wells-go/application/dtos"
	"wells-go/application/mappers"
	"wells-go/domain/entities"
	"wells-go/domain/repositories"
	"wells-go/infrastructure/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo repositories.UserRepository
	cfg  *config.Config
}

func NewUserUsecase(repo repositories.UserRepository, cfg *config.Config) *UserUsecase {
	return &UserUsecase{repo: repo, cfg: cfg}
}

func (uc *UserUsecase) Register(req dtos.RegisterUserRequest) (dtos.UserResponse, error) {
	existing, _ := uc.repo.FindByEmail(req.Email)
	if existing != nil && existing.ID != uuid.Nil {
		return dtos.UserResponse{}, errors.New("email already registered")
	}

	if req.Password != req.ConfirmPassword {
		return dtos.UserResponse{}, errors.New("password dan konfirmasi password tidak sama")
	}

	if req.Password == strings.ToLower(req.Password) {
		return dtos.UserResponse{}, errors.New("password tidak boleh semua huruf kecil, gunakan kombinasi huruf besar/kecil/angka/simbol")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := entities.UserEntity{
		ID:        uuid.New(),
		Fullname:  req.Name,
		Email:     req.Email,
		Password:  string(hash),
		Role:      req.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.repo.Create(&user); err != nil {
		return dtos.UserResponse{}, err
	}

	return mappers.ToUserResponse(&user), nil
}

func (uc *UserUsecase) Login(req dtos.LoginRequest) (string, error) {
	user, err := uc.repo.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		return "", errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"sub":   user.ID.String(),
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(uc.cfg.JWTSecret))
}

func (uc *UserUsecase) GetUsers() ([]dtos.UserResponse, error) {
	users, err := uc.repo.List()
	if err != nil {
		return nil, err
	}

	var res []dtos.UserResponse
	for _, u := range users {
		res = append(res, mappers.ToUserResponse(&u))
	}
	return res, nil
}

func (uc *UserUsecase) GetUserByID(id string) (dtos.UserResponse, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return dtos.UserResponse{}, errors.New("invalid id format")
	}

	user, err := uc.repo.FindByID(uid)
	if err != nil {
		return dtos.UserResponse{}, err
	}

	return mappers.ToUserResponse(user), nil
}

func (uc *UserUsecase) UpdateUser(id string, req dtos.RegisterUserRequest) (dtos.UserResponse, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return dtos.UserResponse{}, errors.New("invalid id format")
	}

	user, err := uc.repo.FindByID(uid)
	if err != nil {
		return dtos.UserResponse{}, err
	}

	if req.Password != "" {
		if req.OldPassword == "" {
			return dtos.UserResponse{}, errors.New("old password wajib diisi untuk update password")
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
			return dtos.UserResponse{}, errors.New("old password salah")
		}

		if req.Password != req.ConfirmPassword {
			return dtos.UserResponse{}, errors.New("password dan konfirmasi password tidak sama")
		}

		if req.Password == strings.ToLower(req.Password) {
			return dtos.UserResponse{}, errors.New("password tidak boleh semua huruf kecil, gunakan kombinasi huruf besar/kecil/angka/simbol")
		}

		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) == nil {
			return dtos.UserResponse{}, errors.New("password baru tidak boleh sama dengan password lama")
		}

		hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		user.Password = string(hash)
	}

	if req.Name != "" {
		user.Fullname = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Role != "" {
		user.Role = req.Role
	}
	user.UpdatedAt = time.Now()

	if err := uc.repo.Update(user); err != nil {
		return dtos.UserResponse{}, err
	}

	return mappers.ToUserResponse(user), nil
}

func (uc *UserUsecase) DeleteUser(id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid id format")
	}
	return uc.repo.Delete(uid)
}
