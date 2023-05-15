package session

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Service) AuthenticateCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				return c.JSON(http.StatusUnauthorized, "unauthorized: no session cookie found")
			}
			// For any other type of error, return a bad request status
			return c.JSON(http.StatusBadRequest, "bad request")
		}

		_, err = s.Get(cookie.Value)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "unauthorized: session does not exist")
		}

		expired := s.IsSessionExpired(cookie.Value)
		if expired {
			s.Delete(cookie.Value)
			return c.JSON(http.StatusUnauthorized, "unauthorized: session expired")
		}

		return next(c)
	}
}
