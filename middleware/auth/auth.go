package auth

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	resp "github.com/iamvkosarev/go-shared-utils/api/response"
	"github.com/iamvkosarev/go-shared-utils/logger/sl"
	"log/slog"
	"net/http"
)

func Auth(log *slog.Logger, verifyURL string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				const op = "middleware.auth.Auth"

				log = log.With(
					slog.String("op", op),
					slog.String("request_id", middleware.GetReqID(r.Context())),
				)

				req, err := http.NewRequest(http.MethodGet, verifyURL, nil)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
				}

				authHeader := r.Header.Get("Authorization")
				if authHeader == "" {
					cookie, err := r.Cookie("jwt")
					if err != nil {
						log.Error("missing token", sl.Err(err))
						render.JSON(
							w, r, resp.Error(
								"missing token. header (\"Authorization\") or cookie ("+
									"\"jwt\") must contain token",
							),
						)
						return
					}
					req.AddCookie(cookie)
				} else {
					req.Header.Add("Authorization", authHeader)
				}

				responce, err := http.DefaultClient.Do(req)
				if err != nil {
					log.Error("failed to contact sso", sl.Err(err))
					render.JSON(w, r, resp.Error("failed to contact sso"))
					return
				}
				defer responce.Body.Close()

				if responce.StatusCode != http.StatusOK {
					log.Error("unauthorized", sl.Err(err))
					render.JSON(w, r, resp.Error("unauthorized"))
					return
				}

				var vResp struct {
					resp.Response
					UserID int64 `json:"user_id"`
				}

				if err := json.NewDecoder(responce.Body).Decode(&vResp); err != nil {
					log.Error("invalid response from sso", sl.Err(err))
					render.JSON(w, r, resp.Error("invalid response from sso"))
					return
				}

				if vResp.Response.Status == resp.StatusError {
					if vResp.Response.Error == resp.ErrorTokenExpired.Error() {
						log.Error("token expired", sl.ErrMsg(vResp.Response.Error))
						render.JSON(w, r, resp.Error("token expired"))
						return
					}
					log.Error("unauthorized", sl.ErrMsg(vResp.Response.Error))
					render.JSON(w, r, resp.Error("unauthorized"))
					return
				}

				ctx := context.WithValue(r.Context(), "user_id", vResp.UserID)
				next.ServeHTTP(w, r.WithContext(ctx))
			},
		)
	}
}
