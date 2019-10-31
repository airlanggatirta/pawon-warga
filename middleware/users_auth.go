package middleware

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/airlanggatirta/pawon-warga/common"
	"github.com/airlanggatirta/pawon-warga/model"
)

func (m *Middleware) UsersAuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	ctx := r.Context()
	params := mux.Vars(r)
	stringUserID := params["id"]

	userID, err := strconv.ParseUint(stringUserID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	var userCtxID uint64
	userInfo, ok := ctx.Value("UserInfo").(*model.NewKitabisaClaims)
	if !ok {
		userInfo, _ := ctx.Value("UserInfo").(model.User)
		userCtxID = userInfo.ID
	} else {
		userCtxID = userInfo.UserID
	}

	if userID != userCtxID {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(common.ErrAccessDenied.Error()))
		return
	} else {
		next(w, r)
	}

}
