package httpserver

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/xeoncross/project-templates/database/db"
	"github.com/xeoncross/project-templates/internal/service"
)

type Handler struct {
	S service.Users
}

func (h *Handler) CreateUser(c echo.Context) error {

	type userInput struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	// Validate input
	u := userInput{}
	if err := c.Bind(&u); err != nil {
		return err
	}

	// TODO: we need an app entity instead of trying to use the db.User
	// directly in the handler. This hack of using a userInput would be replaced
	row := db.User{Name: sql.NullString{String: u.Name, Valid: true}, Email: u.Email}
	id, err := h.S.InsertUser(c.Request().Context(), row)
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
