package httpserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xeoncross/project-templates/database/db"
	"github.com/xeoncross/project-templates/internal/service"
)

type Handler struct {
	S service.Users
}

func (h *Handler) CreateUser(c echo.Context) error {

	// Validate input
	u := db.User{}
	if err := c.Bind(&u); err != nil {
		return err
	}

	id, err := h.S.InsertUser(c.Request().Context(), u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, id)
}

func (h *Handler) GetUser(c echo.Context) error {
	email := c.Param("email")
	user, err := h.S.GetUserByEmail(c.Request().Context(), email)
	if err != nil {
		return err
	}
	// if user == nil {
	// 	return echo.NewHTTPError(http.StatusNotFound, "user not found")
	// }
	return c.JSON(http.StatusOK, user)
}
