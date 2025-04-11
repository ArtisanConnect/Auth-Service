package postgres

import (
	pb "auth/genproto/auth"
	"auth/storage"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) storage.UserStorage {
	return &UserRepo{
		db: db,
	}
}

//id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
//username VARCHAR(255) UNIQUE NOT NULL,
//email VARCHAR(255) UNIQUE NOT NULL,
//password VARCHAR(255) NOT NULL,
//full_name VARCHAR(255) NOT NULL,
//user_type VARCHAR(50) NOT NULL,
//bio TEXT,
//created_at TIMESTAMP DEFAULT now(),
//updated_at TIMESTAMP DEFAULT now()

func (p *UserRepo) Register(in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	id := uuid.New().String()
	query := `INSERT INTO users (id, username, email, password, full_name,
                   user_type, created_at) VALUES ($1, $2, $3, $4, $5, $6, NOW())
                   RETURNING id`
	_, err := p.db.Exec(query, id, in.Username, in.Email, in.Password, in.FullName, in.UserType)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResponse{Id: id}, nil
}

func (p *UserRepo) Login(in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}

func (p *UserRepo) GetProfile(in *pb.ProfileRequest) (*pb.ProfileResponse, error) {
	return nil, nil
}

func (p *UserRepo) EditProfile(in *pb.EditProfileRequest) (*pb.EditProfileResponse, error) {
	return nil, nil
}

func (p *UserRepo) ChangeUserType(in *pb.ChangeUserTypeRequest) (*pb.ChangeUserTypeResponse, error) {
	return nil, nil
}

func (p *UserRepo) ListUsers(in *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	return nil, nil
}

func (p *UserRepo) DeleteUser(in *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return nil, nil
}

func (p *UserRepo) ResetPassword(in *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {
	return nil, nil
}

func (p *UserRepo) RefreshToken(in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
	return nil, nil
}

func (p *UserRepo) Logout(in *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return nil, nil
}
