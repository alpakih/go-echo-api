package http

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"go-echo-api/infrastructure/database"
	"go-echo-api/infrastructure/validator"
	"go-echo-api/user/usecase"
	"go-echo-api/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func init() {
	database.RegisterTxDB("txdb")
}
func TestNewUserController(t *testing.T) {
	//prepare db test
	db, _ := database.PrepareTestDB("txdb")
	defer database.CleanTestDB(db)

	// setup expectations
	s := t.Run("success", func(t *testing.T) {
		// success scenario create object
		c := NewUserController(usecase.NewUserService(db), db)
		assert.NotNil(t, c.userRepository, "Null object created")
	})

	f := t.Run("error-failed", func(t *testing.T) {
		// failed scenario create object
		c := NewUserController(nil, db)
		assert.Nil(t, c.userRepository)
	})

	assert.Equal(t, true, s, "Success scenario failed run")
	assert.Equal(t, true, f, "Failed scenario failed run")
}

func TestUserController_FindAll(t *testing.T) {
	//prepare db test
	db, _ := database.PrepareTestDB("txdb")
	defer database.CleanTestDB(db)

	e := echo.New()
	limit := "5"
	offset := "0"
	req := httptest.NewRequest(echo.GET, "/api/v1/user?limit="+limit+"&offset="+offset, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	controller := NewUserController(usecase.NewUserService(db), db)

	// Assertions
	if assert.NoError(t, controller.FindAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUserController_FindById(t *testing.T) {
	//prepare db test
	db, _ := database.PrepareTestDB("txdb")
	defer database.CleanTestDB(db)

	// create an instance of our test object
	controller := NewUserController(usecase.NewUserService(db), db)
	e := echo.New()

	req := httptest.NewRequest(echo.GET, "/", nil)

	s := t.Run("success", func(t *testing.T) {
		// success scenario
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/user/:id")
		c.SetParamNames("id")
		c.SetParamValues("7dd77cc4-f786-4be0-b5a5-0c203b9c62c5")
		if assert.NoError(t, controller.FindById(c)) {
			fmt.Println("BODY ", rec.Body.String())
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	f := t.Run("error-failed", func(t *testing.T) {
		// failed scenario
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("api/v1/user/:id")
		c.SetParamNames("id")
		c.SetParamValues("not found")
		if assert.NoError(t, controller.FindById(c)) {
			fmt.Println("BODY ", rec.Body.String())
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})
	assert.Equal(t, true, s, "Success scenario failed run")
	assert.Equal(t, true, f, "Failed scenario failed run")
}

func TestUserController_Store(t *testing.T) {
	//prepare db test
	db, _ := database.PrepareTestDB("txdb")
	defer database.CleanTestDB(db)

	// create an instance of our test object
	controller := NewUserController(usecase.NewUserService(db), db)
	hashPassword, _ := utils.HashPassword("password")
	userJSON := `{"name":"Jon Snow","email":"jon@labstack.com","password":"` + hashPassword + `"}`
	userJSONFailed := `{"name":"Jon Snow","email":"","password":"` + hashPassword + `"}`
	userJSONDuplicateEmail := `{"name":"Jon Snow","email":"ipan@email.com","password":"` + hashPassword + `"}`

	s := t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = validator.NewValidator()
		req := httptest.NewRequest(echo.POST, "/api/v1/user", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		if assert.NoError(t, controller.Store(c)) {
			fmt.Println("BODY ", rec.Body.String())
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NotEqual(t, http.StatusInternalServerError, rec.Code)
			assert.NotEqual(t, http.StatusUnprocessableEntity, rec.Code)
			assert.NotEqual(t, http.StatusBadRequest, rec.Code)
		}
	})

	v := t.Run("error-validation", func(t *testing.T) {
		e := echo.New()
		e.Validator = validator.NewValidator()
		req := httptest.NewRequest(echo.POST, "/api/v1/user", strings.NewReader(userJSONFailed))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		if assert.NoError(t, controller.Store(c)) {
			fmt.Println("BODY ", rec.Body.String())
			assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		}
	})

	f := t.Run("error-bad-request", func(t *testing.T) {
		e := echo.New()
		e.Validator = validator.NewValidator()
		req := httptest.NewRequest(echo.POST, "/api/v1/user", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		if assert.NoError(t, controller.Store(c)) {
			fmt.Println("BODY ", rec.Body.String())
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	i := t.Run("error-failed", func(t *testing.T) {
		e := echo.New()
		e.Validator = validator.NewValidator()
		req := httptest.NewRequest(echo.POST, "/api/v1/user", strings.NewReader(userJSONDuplicateEmail))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		if assert.NoError(t, controller.Store(c)) {
			fmt.Println("BODY ", rec.Body.String())
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})
	assert.Equal(t, true, s, "Success scenario failed run")
	assert.Equal(t, true, f, "Failed scenario failed run")
	assert.Equal(t, true, v, "Validator scenario failed run")
	assert.Equal(t, true, i, "Failed save scenario failed run")
}

func TestUserController_Update(t *testing.T) {
	//prepare db test
	db, _ := database.PrepareTestDB("txdb")
	defer database.CleanTestDB(db)

	// create an instance of our test object
	controller := NewUserController(usecase.NewUserService(db), db)
	hashPassword, _ := utils.HashPassword("password")
	userJSON := `{"name":"Jon Snow","email":"jon@labstack.com","password":"` + hashPassword + `"}`
	userJSONFailed := `{"name":"Jon Snow","email":"","password":"` + hashPassword + `"}`
	userJSONDuplicateEmail := `{"name":"Jon Snow","email":"ipan@email.com","password":"` + hashPassword + `"}`

	s := t.Run("success", func(t *testing.T) {
		e := echo.New()
		e.Validator = validator.NewValidator()
		req := httptest.NewRequest(echo.PUT, "/api/v1/user/:id", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("7dd77cc4-f786-4be0-b5a5-0c203b9c62c5")
		if assert.NoError(t, controller.Update(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NotEqual(t, http.StatusInternalServerError, rec.Code)
			assert.NotEqual(t, http.StatusUnprocessableEntity, rec.Code)
			assert.NotEqual(t, http.StatusBadRequest, rec.Code)
		}
	})

	v := t.Run("error-validation", func(t *testing.T) {
		e := echo.New()
		e.Validator = validator.NewValidator()
		req := httptest.NewRequest(echo.PUT, "/api/v1/user/:id", strings.NewReader(userJSONFailed))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("7dd77cc4-f786-4be0-b5a5-0c203b9c62c5")
		if assert.NoError(t, controller.Update(c)) {
			assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		}
	})

	f := t.Run("error-bad-request", func(t *testing.T) {
		e := echo.New()
		e.Validator = validator.NewValidator()
		req := httptest.NewRequest(echo.POST, "/api/v1/user/:id", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("7dd77cc4-f786-4be0-b5a5-0c203b9c62c5")
		if assert.NoError(t, controller.Update(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	i := t.Run("error-failed", func(t *testing.T) {
		e := echo.New()
		e.Validator = validator.NewValidator()
		req := httptest.NewRequest(echo.POST, "/api/v1/user/:id", strings.NewReader(userJSONDuplicateEmail))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("7dd77cc4-f786-4be0-b5a5-0c203b9c62c5")
		if assert.NoError(t, controller.Update(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})
	assert.Equal(t, true, s, "Success scenario failed run")
	assert.Equal(t, true, f, "Failed scenario failed run")
	assert.Equal(t, true, v, "Validator scenario failed run")
	assert.Equal(t, true, i, "Failed update scenario failed run")
}

func TestUserController_Delete(t *testing.T) {
	//prepare db test
	db, _ := database.PrepareTestDB("txdb")
	defer database.CleanTestDB(db)

	// create an instance of our test object
	controller := NewUserController(usecase.NewUserService(db), db)

	s := t.Run("success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/api/v1/user/:id", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")
		c.SetParamValues("7dd77cc4-f786-4be0-b5a5-0c203b9c62c5")
		if assert.NoError(t, controller.Delete(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.NotEqual(t, http.StatusInternalServerError, rec.Code)
		}
	})

	f := t.Run("error-failed", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.DELETE, "/api/v1/user/:id", strings.NewReader(""))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("id")

		if assert.NoError(t, controller.Delete(c)) {
			assert.NotEqual(t, http.StatusInternalServerError, rec.Code)
		}
	})
	assert.Equal(t, true, s, "Success scenario failed run")
	assert.Equal(t, true, f, "Failed scenario failed run")

}