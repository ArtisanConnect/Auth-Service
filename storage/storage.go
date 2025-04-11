package storage

import (
	pb "auth/genproto/auth"
)

type UserStorage interface {
	Register(in *pb.RegisterRequest) (*pb.RegisterResponse, error)
	Login(in *pb.LoginRequest) (*pb.LoginResponse, error)
	GetProfile(in *pb.ProfileRequest) (*pb.ProfileResponse, error)
	EditProfile(in *pb.EditProfileRequest) (*pb.EditProfileResponse, error)
	ChangeUserType(in *pb.ChangeUserTypeRequest) (*pb.ChangeUserTypeResponse, error)
	ListUsers(in *pb.ListUsersRequest) (*pb.ListUsersResponse, error)
	DeleteUser(in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error)
	ResetPassword(in *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error)
	RefreshToken(in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error)
	Logout(in *pb.LogoutRequest) (*pb.LogoutResponse, error)
}
