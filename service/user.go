package service

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
	pb "github.com/venomuz/project2/genproto"
	l "github.com/venomuz/project2/pkg/logger"
	"github.com/venomuz/project2/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
}

//NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	id1 := uuid.NewV4()
	id2 := uuid.NewV4()
	req.Id = id1.String()
	req.Address.Id = id2.String()
	user, err := s.storage.User().Create(req)
	if err != nil {
		s.logger.Error("Error while inserting user info", l.Error(err))
		return nil, status.Error(codes.Internal, "Error insert user")
	}
	return user, err
}
func (s *UserService) GetByID(ctx context.Context, req *pb.GetIdFromUser) (*pb.User, error) {
	user, err := s.storage.User().GetByID(req.Id)
	if err != nil {
		fmt.Println(err)
		s.logger.Error("Error while getting user info", l.Error(err))
		return nil, status.Error(codes.Internal, "Error insert user")
	}

	return user, err
}
func (s *UserService) DeleteByID(ctx context.Context, req *pb.GetIdFromUser) (*pb.GetIdFromUser, error) {
	user, err := s.storage.User().DeleteByID(req.Id)
	if err != nil {
		s.logger.Error("Error while getting user info", l.Error(err))
		return nil, status.Error(codes.Internal, "Error insert user")
	}
	return user, err
}
