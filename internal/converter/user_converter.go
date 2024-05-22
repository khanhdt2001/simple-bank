package converter

import (
	db "simple_bank/db/sqlc"
	"simple_bank/internal/dto"
	pb "simple_bank/proto/pb/proto"

	"simple_bank/internal/validation"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertUserDpToPb(user db.User) *pb.User {
	return &pb.User{
		Username:         user.Username,
		FullName:         user.FullName,
		Email:            user.Email,
		CreatedAt:        timestamppb.New(user.CreatedAt),
		PasswordChangeAt: timestamppb.New(user.PasswordChangedAt),
	}
}

func ConvertUserPbToDto(user *pb.CreateUserRequest) (*dto.CreateUserRequest, error) {

	req := &dto.CreateUserRequest{
		UserName: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
		Password: user.Password,
	}

	err := validation.BankValidatior.Struct(req)
	if err != nil {
		return nil, err

	}

	return req, nil
}
