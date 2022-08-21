package httpserver

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/xeoncross/project-templates/database/db"
	"github.com/xeoncross/project-templates/internal/mocks"
	"github.com/xeoncross/project-templates/internal/service"
)

var userJSON = `{"name":"Jon","email":"jon@example.com"}`

var userStruct = db.User{
	Name:  sql.NullString{String: "Jon", Valid: true},
	Email: "jon@example.com",
}

// TODO: this is broken

// https://gist.github.com/thiagozs/4276432d12c2e5b152ea15b3f8b0012e

func TestCreateUser(t *testing.T) {

	// var u db.User
	// err := json.Unmarshal([]byte(userJSON), &u)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Setup mock db
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := mocks.NewMockQuerier(ctrl)

	// Asserts that the first and only call to Bar() is passed 99.
	// Anything else will fail.
	m.
		EXPECT().
		InsertUser(gomock.Eq(req.Context()), gomock.Eq(userStruct)).
		Return(1, nil)

	h := &Handler{&service.User{DB: m}}

	// Assertions
	if assert.NoError(t, h.CreateUser(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}

// func TestGetUser(t *testing.T) {
// 	// Setup
// 	e := echo.New()
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.SetPath("/users/:email")
// 	c.SetParamNames("email")
// 	c.SetParamValues("jon@labstack.com")
// 	h := &handler{mockDB}

// 	// Assertions
// 	if assert.NoError(t, h.getUser(c)) {
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		assert.Equal(t, userJSON, rec.Body.String())
// 	}
// }
