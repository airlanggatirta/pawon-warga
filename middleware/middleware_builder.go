package middleware

import (
	"github.com/airlanggatirta/pawon-warga/common"
	"github.com/airlanggatirta/pawon-warga/config"
	"github.com/airlanggatirta/pawon-warga/service"
)

// Builder defines struct builder for middleware
type Builder struct {
	JWTSignKey       string               `validate:"required"`
	HmacConfig       config.HmacSignature `validate:"required"`
	KtbsHeaderConfig config.KtbsHeader    `validate:"required"`
	UserService      service.IUserService `validate:"required"`
}

// NewMiddlewareBuilder initializes middleware builder instance
func NewMiddlewareBuilder() *Builder {
	return &Builder{}
}

// SetJwtSignKey will set jwt sign key
func (b *Builder) SetJwtSignKey(key string) *Builder {
	b.JWTSignKey = key
	return b
}

// SetHmacConfig will set hmac configuration
func (b *Builder) SetHmacConfig(cfg config.HmacSignature) *Builder {
	b.HmacConfig = cfg
	return b
}

// SetKtbsHeaderConfig will set Ktbs header configuration
func (b *Builder) SetKtbsHeaderConfig(cfg config.KtbsHeader) *Builder {
	b.KtbsHeaderConfig = cfg
	return b
}

// SetUserService will set user service
func (b *Builder) SetUserService(svc service.IUserService) *Builder {
	b.UserService = svc
	return b
}

// Build will build middleware instance based on Builder
func (b *Builder) Build() (*Middleware, error) {
	err := common.Validate.Struct(b)
	if err != nil {
		return nil, err
	}

	return &Middleware{
		JWTSignKey:       b.JWTSignKey,
		HmacConfig:       b.HmacConfig,
		KtbsHeaderConfig: b.KtbsHeaderConfig,
		UserService:      b.UserService,
	}, nil
}
