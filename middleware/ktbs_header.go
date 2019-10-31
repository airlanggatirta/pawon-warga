package middleware

import (
	"net/http"

	"github.com/airlanggatirta/pawon-warga/common"
	"github.com/airlanggatirta/pawon-warga/handler"
	logger "github.com/sirupsen/logrus"
)

// KtbsHeader will check if mandatory Ktbs headers are exist or not.
func (m *Middleware) KtbsHeader(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.Header.Get("X-Ktbs-Request-ID") == "" {
		m.ktbsInvalidHeader(w, "X-Ktbs-Request-ID")
		return
	}

	if r.Header.Get("X-Ktbs-Api-Version") == "" {
		m.ktbsInvalidHeader(w, "X-Ktbs-Api-Version")
		return
	}

	if r.Header.Get("X-Ktbs-Client-Name") == "" {
		m.ktbsInvalidHeader(w, "X-Ktbs-Client-Name")
		return
	}

	// exclude client name from client-version & platform-name checking
	blacklistClientName := m.KtbsHeaderConfig.ExcludeMobileHeaderCheck
	exist, _ := common.InArray(r.Header.Get("X-Ktbs-Client-Name"), blacklistClientName)

	if !exist {
		if r.Header.Get("X-Ktbs-Client-Version") == "" {
			m.ktbsInvalidHeader(w, "X-Ktbs-Client-Version")
			return
		}

		if r.Header.Get("X-Ktbs-Platform-Name") == "" {
			m.ktbsInvalidHeader(w, "X-Ktbs-Platform-Name")
			return
		}
	}

	next(w, r)
}

func (m *Middleware) ktbsInvalidHeader(w http.ResponseWriter, headerName string) {
	logger.Warnf("Invalid header: header %s is required.", headerName)

	statusError := handler.StatusError{
		Code: http.StatusBadRequest,
		Err:  common.ErrInvalidHeader,
	}
	handler.WriteErrorResponse(w, statusError)
}
