package middleware

import (
	"github.com/airlanggatirta/pawon-warga/common"
	"github.com/airlanggatirta/pawon-warga/config"
	"github.com/airlanggatirta/pawon-warga/service"
)

// Builder defines struct builder for middleware
type Builder struct {
	JWTSignKey       string                `validate:"required"`
	HmacConfig       config.HmacSignature  `validate:"required"`
	KtbsHeaderConfig config.KtbsHeader     `validate:"required"`
	TokenService     service.ITokenService `validate:"required"`
	UserService      service.IUserService  `validate:"required"`
	OTPService       service.IOTPService   `validate:"required"`
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

// SetTokenService will set token service
func (b *Builder) SetTokenService(svc service.ITokenService) *Builder {
	b.TokenService = svc
	return b
}

// SetUserService will set user service
func (b *Builder) SetUserService(svc service.IUserService) *Builder {
	b.UserService = svc
	return b
}

// SetOTPService will set OTP service
func (b *Builder) SetOTPService(svc service.IOTPService) *Builder {
	b.OTPService = svc
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
		TokenService:     b.TokenService,
		UserService:      b.UserService,
		OTPService:       b.OTPService,
	}, nil
}
