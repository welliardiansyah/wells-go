package usecases

import (
	"context"
	"errors"
	"strings"
	"time"
	"wells-go/application/dtos"
	"wells-go/application/mappers"
	"wells-go/domain/entities"
	"wells-go/domain/repositories"
	"wells-go/infrastructure/config"
	"wells-go/infrastructure/redis"
	"wells-go/util/security"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo     repositories.UserRepository
	repoRole repositories.RoleRepository
	cfg      *config.Config
	security security.Maker
}

func NewUserUsecase(repo repositories.UserRepository, repoRole repositories.RoleRepository, cfg *config.Config, securityMaker security.Maker) *UserUsecase {
	return &UserUsecase{
		repo:     repo,
		repoRole: repoRole,
		cfg:      cfg,
		security: securityMaker,
	}
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
		return dtos.UserResponse{}, errors.New("password harus kombinasi huruf besar/kecil/angka/simbol")
	}

	roleID, err := uuid.Parse(req.Role)
	if err != nil {
		return dtos.UserResponse{}, errors.New("role ID tidak valid")
	}

	roleEntity, err := uc.repoRole.FindByID(roleID)
	if err != nil || roleEntity == nil {
		return dtos.UserResponse{}, errors.New("role tidak ditemukan")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	user := entities.UserEntity{
		ID:       uuid.New(),
		Fullname: req.Name,
		Email:    req.Email,
		//Phone:     req.Phone,
		Password:  string(hash),
		RoleId:    roleEntity.ID,
		Role:      *roleEntity,
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

	roles := []string{user.Role.Name}
	perms := []security.Permission{}
	for _, p := range user.Role.Permissions {
		perms = append(perms, security.Permission{
			Name:      p.Name,
			CanCreate: p.CanCreate,
			CanRead:   p.CanRead,
			CanUpdate: p.CanUpdate,
			CanDelete: p.CanDelete,
			CanExport: p.CanExport,
			CanImport: p.CanImport,
			CanView:   p.CanView,
		})
	}

	token, err := uc.security.CreateToken(
		user.ID.String(),
		user.Email,
		roles,
		perms,
		24*time.Hour,
	)
	if err != nil {
		return "", err
	}

	err = redis.Rdb.Set(context.Background(), "jwt:"+token, "active", 24*time.Hour).Err()
	if err != nil {
		return "", errors.New("failed to save token in redis")
	}

	return token, nil
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
			return dtos.UserResponse{}, errors.New("old password wajib diisi")
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
			return dtos.UserResponse{}, errors.New("old password salah")
		}
		if req.Password != req.ConfirmPassword {
			return dtos.UserResponse{}, errors.New("password dan konfirmasi password tidak sama")
		}
		if req.Password == strings.ToLower(req.Password) {
			return dtos.UserResponse{}, errors.New("password harus kombinasi huruf besar/kecil/angka/simbol")
		}
		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) == nil {
			return dtos.UserResponse{}, errors.New("password baru tidak boleh sama dengan lama")
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
	//if req.Phone != "" {
	//	user.Phone = req.Phone
	//}
	if req.Role != "" {
		roleID, err := uuid.Parse(req.Role)
		if err != nil {
			return dtos.UserResponse{}, errors.New("role ID tidak valid")
		}
		roleEntity, err := uc.repoRole.FindByID(roleID)
		if err != nil || roleEntity == nil {
			return dtos.UserResponse{}, errors.New("role tidak ditemukan")
		}
		user.RoleId = roleEntity.ID
		user.Role = *roleEntity
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
