package auth

import (
	"context"
	"coolcar/shared/auth/token"
	"coolcar/shared/id"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	// 定义一个header，用于在请求头中加上和辨识
	// 一个请求是内部服务以某个account-id的身份来请求行程服务
	ImpersonateAccountHeader = "impersonate-account-id"
	authorizationHeader      = "authorization"
	bearerPrefix             = "Bearer "
)

func Interceptor(publicKeyFile string) (grpc.UnaryServerInterceptor, error) {
	fp, err := os.Open(publicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("cannot open public key file: %v", err)
	}
	b, err := io.ReadAll(fp)
	if err != nil {
		return nil, fmt.Errorf("cannot read public key file: %v", err)
	}
	pbKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("cannot parse public key: %v", err)
	}
	i := &interceptor{
		verifier: &token.JWTTokenVerifier{
			PublicKey: pbKey,
		},
	}
	return i.HandleRequest, nil
}

type tokenVerifier interface {
	Verify(token string) (string, error)
}

type interceptor struct {
	verifier tokenVerifier
}

func (i *interceptor) HandleRequest(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 如果是特殊的内部请求，从特殊的请求头中获取aid并直接返回
	aid := impersonationFromContext(ctx)
	if aid != "" {
		return handler(ContextWithAccountID(ctx, id.AccountID(aid)), req)
	}

	tkn, err := tokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}

	aid, err = i.verifier.Verify(tkn)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token not valid: %v", err)
	}

	return handler(ContextWithAccountID(ctx, id.AccountID(aid)), req)
}

func impersonationFromContext(c context.Context) string {
	m, ok := metadata.FromIncomingContext(c)
	if !ok {
		return ""
	}

	imp := m[ImpersonateAccountHeader]
	if len(imp) == 0 {
		return ""
	}

	return imp[0]
}

func tokenFromContext(c context.Context) (string, error) {
	m, ok := metadata.FromIncomingContext(c)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "")
	}

	tkn := ""
	for _, v := range m[authorizationHeader] {
		if strings.HasPrefix(v, bearerPrefix) {
			tkn = v[len(bearerPrefix):]
		}
	}
	if tkn == "" {
		return "", status.Error(codes.Unauthenticated, "")
	}

	return tkn, nil
}

type accountIDKey struct{}

func ContextWithAccountID(c context.Context, aid id.AccountID) context.Context {
	return context.WithValue(c, accountIDKey{}, aid)
}

func AccountIDFromContext(c context.Context) (id.AccountID, error) {
	v := c.Value(accountIDKey{})
	aid, ok := v.(id.AccountID)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "")
	}
	return aid, nil
}
