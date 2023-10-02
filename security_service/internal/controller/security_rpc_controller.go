package security_controller

import (
	"context"
	"github.com/catness812/e-petitions-project/security_service/internal/pb"

	models "github.com/catness812/e-petitions-project/security_service/internal/model"
	"github.com/catness812/e-petitions-project/security_service/pkg/jwtoken"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ISecurityService interface {
	Login(user *models.UserCredentialsModel) (map[string]string, error)
	RefreshUserToken(token string, email string) (map[string]string, error)
}

type SecurityRpcServer struct {
	securitySvc ISecurityService
}

func NewSecurityRpcServer(securitySvc ISecurityService) *SecurityRpcServer {
	return &SecurityRpcServer{securitySvc: securitySvc}
}

func (s *SecurityRpcServer) Login(ctx context.Context, req *pb.UserCredentials) (*pb.Tokens, error) {
	userLogin := models.UserCredentialsModel{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	token, err := s.securitySvc.Login(&userLogin)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &pb.Tokens{
		AccessToken:  token["access_token"],
		RefreshToken: token["refresh_token"],
	}, nil
}

func (s *SecurityRpcServer) RefreshSession(ctx context.Context, req *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	refToken := req.Token
	userEmail, err := jwtoken.IsTokenValid(refToken)
	if err != nil {
		return &pb.RefreshResponse{
			Tokens:  nil,
			Message: err.Error(),
		}, nil
	}
	tokenMap, err := s.securitySvc.RefreshUserToken(refToken, userEmail)
	if err != nil {
		return &pb.RefreshResponse{
			Tokens:  nil,
			Message: err.Error(),
		}, nil
	}
	return &pb.RefreshResponse{Tokens: tokenMap}, nil
}

func (s *SecurityRpcServer) ValidateToken(ctx context.Context, req *pb.Token) (*pb.ValidateTokenResponse, error) {
	email, err := jwtoken.IsTokenValid(req.Token)
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	result := &pb.ValidateTokenResponse{Token: req.Token, Email: email}
	return result, nil
}
