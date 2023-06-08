package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/patrickkoss/grpc-gateway-example/domain"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/api/common"
	proto "github.com/patrickkoss/grpc-gateway-example/internal/adapter/api/domain/proto"
	"github.com/patrickkoss/grpc-gateway-example/internal/adapter/logging"
	"github.com/patrickkoss/grpc-gateway-example/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type server struct {
	userService service.UserService
	logger      logging.Logger
}

func (s server) CreateUser(ctx context.Context, body *CreateUserBody) (*proto.Message, error) {
	domainUser := domain.User{
		Id:      uuid.NewString(),
		Name:    body.Name,
		Email:   body.Email,
		Surname: body.Surname,
	}

	err := s.userService.Create(&domainUser)
	if err != nil {
		return nil, common.GetStatusErrorFromDomainError(err)
	}

	_ = grpc.SetHeader(ctx, metadata.Pairs("x-http-code", "201"))

	return &proto.Message{Message: "User created"}, nil
}

func (s server) GetUsers(ctx context.Context, params *UserParams) (*ListUsersResponse, error) {
	domainUsers, err := s.userService.List()
	if err != nil {
		return nil, common.GetStatusErrorFromDomainError(err)
	}

	users := make([]*User, len(domainUsers))
	for i, domainUser := range domainUsers {
		users[i] = &User{
			Id:      domainUser.Id,
			Name:    domainUser.Name,
			Email:   domainUser.Email,
			Surname: domainUser.Surname,
		}
	}

	return &ListUsersResponse{Users: users}, nil
}

func (s server) GetUser(ctx context.Context, id *UserId) (*GetUserResponse, error) {
	domainUser, err := s.userService.Get(id.Id)
	if err != nil {
		return nil, common.GetStatusErrorFromDomainError(err)
	}

	user := User{
		Id:      domainUser.Id,
		Name:    domainUser.Name,
		Email:   domainUser.Email,
		Surname: domainUser.Surname,
	}

	return &GetUserResponse{User: &user}, nil
}

func (s server) DeleteUser(ctx context.Context, id *UserId) (*proto.Message, error) {
	err := s.userService.Delete(id.Id)
	if err != nil {
		return nil, common.GetStatusErrorFromDomainError(err)
	}

	return &proto.Message{Message: "User deleted"}, nil
}

func (s server) mustEmbedUnimplementedUserServiceServer() {}

func NewServer(userService service.UserService, logger logging.Logger) UserServiceServer {
	return server{
		userService: userService,
		logger:      logger,
	}
}
