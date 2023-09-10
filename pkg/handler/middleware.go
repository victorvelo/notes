package handler

import (
	"context"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) AuthenticationMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth_header := r.Header.Get(authorizationHeader)
		if !strings.HasPrefix(auth_header, "Bearer") {
			ResponseError(w, http.StatusInternalServerError, "Authorization header is missing")
			return
		}

		tokenString := strings.TrimPrefix(auth_header, "Bearer ")

		userId, err := h.service.Authorization.ParseToken(tokenString)
		if err != nil {
			ResponseError(w, http.StatusUnauthorized, err.Error())
			return
		}
		r = r.WithContext(SetJWTClaimsContext(r.Context(), userId))
		next.ServeHTTP(w, r)
	})
}

type claimskey int

var claimsKey claimskey

func SetJWTClaimsContext(ctx context.Context, userId int) context.Context {
	return context.WithValue(ctx, claimsKey, userId)
}

func JWTClaimsFromContext(ctx context.Context) (int, bool) {
	userId, ok := ctx.Value(claimsKey).(int)
	return userId, ok
}

// func (h *Handler) protectedHandler(w http.ResponseWriter, r *http.Request) {
// 	//claims, ok := JWTClaimsFromContext(r.Context())

// }

// func (h *Handler) JwtVerify(w http.ResponseWriter, r *http.Request) {
// 	var header = r.Header.Get(authorizationHeader)
// 	if header == "" {
// 		ResponseError(w, http.StatusInternalServerError, "Authorization header is missing")
// 		return
// 	}
// 	parts := strings.Split(header, " ")

// 	if len(parts) != 2 {
// 		ResponseError(w, http.StatusForbidden, "Malformed token")
// 		return
// 	}

// 	userId, err := h.service.Authorization.ParseToken(parts[1])
// 	if err != nil {
// 		ResponseError(w, http.StatusUnauthorized, err.Error())
// 		return
// 	}

// 	ctx := context.WithValue(r.Context(), "user", userId)
// 	r.WithContext(ctx)
// }
