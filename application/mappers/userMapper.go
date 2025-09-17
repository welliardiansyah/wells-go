package mappers

import (
	"wells-go/application/dtos"
	"wells-go/domain/entities"
)

func ToUserResponse(user *entities.UserEntity) dtos.UserResponse {
	return dtos.UserResponse{
		ID:    user.ID.String(),
		Name:  user.Fullname,
		Email: user.Email,
		Role:  user.Role.Name,
	}
}
