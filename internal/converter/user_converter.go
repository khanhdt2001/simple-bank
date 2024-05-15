package converter

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	db "simple_bank/db/sqlc"
	pb "simple_bank/proto/pb/proto"
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
