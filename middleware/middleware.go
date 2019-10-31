package middleware

import (
	"github.com/airlanggatirta/pawon-warga/config"
	"github.com/airlanggatirta/pawon-warga/service"
)

// Middleware defines middleware struct data
type Middleware struct {
	JWTSignKey       string
	HmacConfig       config.HmacSignature
	KtbsHeaderConfig config.KtbsHeader
	TokenService     service.ITokenService
	UserService      service.IUserService
	OTPService       service.IOTPService
}
