package data

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	authn "github.com/tx7do/kratos-authn/engine"
	authnEngine "github.com/tx7do/kratos-authn/engine"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"

	"kratos-bi/app/admin/service/internal/biz"
	v1 "kratos-bi/gen/api/go/user/service/v1"
)

var _ biz.UserTokenRepo = (*userTokenRepo)(nil)

type userTokenRepo struct {
	data          *Data
	log           *log.Helper
	authenticator authnEngine.Authenticator
}

func NewUserTokenRepo(data *Data, authenticator authnEngine.Authenticator, logger log.Logger) biz.UserTokenRepo {
	l := log.NewHelper(log.With(logger, "module", "user-token/repo/admin-service"))
	return &userTokenRepo{
		data:          data,
		log:           l,
		authenticator: authenticator,
	}
}

func (r *userTokenRepo) createAccessJwtToken(_ string, userId uint32) string {
	principal := authn.AuthClaims{
		Subject: strconv.FormatUint(uint64(userId), 10),
		Scopes:  make(authn.ScopeSet),
	}

	signedToken, err := r.authenticator.CreateIdentity(principal)
	if err != nil {
		return ""
	}

	return signedToken
}

func (r *userTokenRepo) GenerateToken(ctx context.Context, user *v1.User) (string, error) {
	token := r.createAccessJwtToken(user.GetUserName(), user.GetId())
	if token == "" {
		return "", errors.New("create token failed")
	}

	err := r.setToken(ctx, user.GetId(), token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (r *userTokenRepo) GetToken(ctx context.Context, userId uint32) string {
	return r.getToken(ctx, userId)
}

func (r *userTokenRepo) RemoveToken(ctx context.Context, userId uint32) error {
	validToken := r.getToken(ctx, userId)
	if validToken == "" {
		return v1.ErrorTokenNotExist("令牌不存在")
	}

	return r.deleteToken(ctx, userId)
}

func (r *userTokenRepo) RemoveUserToken(ctx context.Context, userId uint32) error {
	validToken := r.getToken(ctx, userId)
	if validToken == "" {
		return v1.ErrorTokenNotExist("令牌不存在")
	}

	return r.deleteToken(ctx, userId)
}

const userTokenKeyPrefix = "editor_admin_ut_"

func (r *userTokenRepo) setToken(ctx context.Context, userId uint32, token string) error {
	key := fmt.Sprintf("%s%d", userTokenKeyPrefix, userId)
	return r.data.rdb.Set(ctx, key, token, 0).Err()
}

func (r *userTokenRepo) getToken(ctx context.Context, userId uint32) string {
	key := fmt.Sprintf("%s%d", userTokenKeyPrefix, userId)
	result, err := r.data.rdb.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			r.log.Errorf("get redis user token failed: %s", err.Error())
		}
		return ""
	}
	return result
}

func (r *userTokenRepo) deleteToken(ctx context.Context, userId uint32) error {
	key := fmt.Sprintf("%s%d", userTokenKeyPrefix, userId)
	return r.data.rdb.Del(ctx, key).Err()
}
