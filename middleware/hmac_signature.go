package middleware

import (
	"fmt"
	"net/http"

	"github.com/airlanggatirta/pawon-warga/common"
	"github.com/airlanggatirta/pawon-warga/handler"
	hmacSignature "github.com/kitabisa/perkakas/v2/signature"
	logger "github.com/sirupsen/logrus"
)

// HmacSignature defines middleware for Ktbs HMAC signature
func (m *Middleware) HmacSignature(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	signature := r.Header.Get("X-Ktbs-Signature")
	clientName := r.Header.Get("X-Ktbs-Client-Name")
	timestamp := r.Header.Get("X-Ktbs-Time")

	if signature == "" {
		m.ktbsInvalidHeader(w, "X-Ktbs-Signature")
		return
	}

	if clientName == "" {
		m.ktbsInvalidHeader(w, "X-Ktbs-Client-Name")
		return
	}

	if timestamp == "" {
		m.ktbsInvalidHeader(w, "X-Ktbs-Time")
		return
	}

	data := fmt.Sprintf("%s.%s", clientName, timestamp)

	if !hmacSignature.IsMatchHmac(data, signature, m.HmacConfig.SecretKey) {
		logger.Warn("Invalid signature header")

		statusError := handler.StatusError{
			Code: http.StatusUnauthorized,
			Err:  common.ErrInvalidSignature,
		}
		handler.WriteErrorResponse(w, statusError)
		return
	}

	next(w, r)
}
