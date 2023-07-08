package data

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	authn "github.com/tx7do/kratos-authn/engine"
	authnEngine "github.com/tx7do/kratos-authn/engine"

	v1 "kratos-uba/gen/api/go/user/service/v1"
)

type AppTokenRepo struct {
	data          *Data
	log           *log.Helper
	authenticator authnEngine.Authenticator
}

func NewApplicationTokenRepo(data *Data, authenticator authnEngine.Authenticator, logger log.Logger) *AppTokenRepo {
	l := log.NewHelper(log.With(logger, "module", "app-token/repo/agent-service"))
	return &AppTokenRepo{
		data:          data,
		log:           l,
		authenticator: authenticator,
	}
}

func (r *AppTokenRepo) createAccessJwtToken(_ string, appId uint32) string {
	principal := authn.AuthClaims{
		Subject: strconv.FormatUint(uint64(appId), 10),
		Scopes:  make(authn.ScopeSet),
	}

	signedToken, err := r.authenticator.CreateIdentity(principal)
	if err != nil {
		return ""
	}

	return signedToken
}

func (r *AppTokenRepo) GenerateToken(ctx context.Context, app *v1.Application) (string, error) {
	token := r.createAccessJwtToken(app.GetName(), app.GetId())
	if token == "" {
		return "", errors.New("create token failed")
	}

	err := r.setToken(ctx, app.GetId(), token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *AppTokenRepo) GetToken(ctx context.Context, appId uint32) string {
	return r.getToken(ctx, appId)
}

func (r *AppTokenRepo) RemoveToken(ctx context.Context, appId uint32) error {
	validToken := r.getToken(ctx, appId)
	if validToken == "" {
		return v1.ErrorTokenNotExist("令牌不存在")
	}

	return r.deleteToken(ctx, appId)
}

func (r *AppTokenRepo) RemoveApplicationToken(ctx context.Context, appId uint32) error {
	validToken := r.getToken(ctx, appId)
	if validToken == "" {
		return v1.ErrorTokenNotExist("令牌不存在")
	}

	return r.deleteToken(ctx, appId)
}

const appTokenKeyPrefix = "bi_app_token_"

func (r *AppTokenRepo) setToken(ctx context.Context, appId uint32, token string) error {
	key := fmt.Sprintf("%s%d", appTokenKeyPrefix, appId)
	return r.data.rdb.Set(ctx, key, token, 0).Err()
}

func (r *AppTokenRepo) getToken(ctx context.Context, appId uint32) string {
	key := fmt.Sprintf("%s%d", appTokenKeyPrefix, appId)
	result, err := r.data.rdb.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("get redis app token failed: %s", err.Error())
		}
		return ""
	}
	return result
}

func (r *AppTokenRepo) deleteToken(ctx context.Context, appId uint32) error {
	key := fmt.Sprintf("%s%d", appTokenKeyPrefix, appId)
	return r.data.rdb.Del(ctx, key).Err()
}
