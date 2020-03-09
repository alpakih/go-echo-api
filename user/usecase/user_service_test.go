package usecase

import (
	"github.com/stretchr/testify/assert"
	"go-echo-api/models"
	"go-echo-api/infrastructure/database"
	"go-echo-api/user"
	"go-echo-api/utils"
	"testing"
)

func init() {
	database.RegisterTxDB("txdb")
}

func TestUserServiceFindAll(t *testing.T) {
	//prepare db test
	db, _ := database.PrepareTestDB("txdb")
	defer database.CleanTestDB(db)

	// scenario find all success
	u := NewUserService(db)
	list, err := u.FindAll()
	assert.NotEmpty(t, list, "No Empty")
	assert.NoError(t, err, "Error")

}

func TestUserService_FindById(t *testing.T) {
	//prepare db test
	db, _ := database.PrepareTestDB("txdb")
	defer database.CleanTestDB(db)

	// create an instance of our test object
	mockUser := models.User{
		ID:    "7dd77cc4-f786-4be0-b5a5-0c203b9c62c5",
		Name:  "Uje",
		Email: "uje@email.com",
	}
	u := NewUserService(db)

	// setup expectations
	s := t.Run("success", func(t *testing.T) {
		// success scenario find by id
		data, err := u.FindById(mockUser.ID)
		assert.NoError(t, err)
		assert.NotNil(t, data)
	})

	f := t.Run("error-failed", func(t *testing.T) {
		// failed scenario find by id
		data, err := u.FindById("test")
		assert.Error(t, err)
		assert.Nil(t, data)
	})
	assert.Equal(t, true, s, "Success scenario failed run")
	assert.Equal(t, true, f, "Failed scenario failed run")
}

func TestUserService_Save(t *testing.T) {
	//prepare db test
	db, _ := database.PrepareTestDB("txdb")
	defer database.CleanTestDB(db)

	// create an instance of our test object
	hashPassword, _ := utils.HashPassword("password")
	mockUser := user.Dto{
		Name:     "Ahmad",
		Email:    "ahmad@email.com",
		Password: hashPassword,
	}
	mockUserFailed := user.Dto{
		Name:     "Ahmad",
		Email:    "ipan@email.com",
		Password: hashPassword,
	}
	u := NewUserService(db)

	// setup expectations
	s := t.Run("success", func(t *testing.T) {
		// success scenario save
		data, err := u.Save(mockUser)
		assert.NoError(t, err)
		assert.NotEmpty(t, data)
	})
	f := t.Run("error-failed", func(t *testing.T) {
		// failed scenario save (duplicate)
		_, err := u.Save(mockUserFailed)
		assert.Error(t, err)
	})
	assert.Equal(t, true, s, "Success scenario failed run")
	assert.Equal(t, true, f, "Failed scenario failed run")
}

func TestUserService_Update(t *testing.T) {
	//prepare db test
	db, _ := database.PrepareTestDB("txdb")
	defer database.CleanTestDB(db)

	// create an instance of our test object
	hashPassword, _ := utils.HashPassword("password")
	mockUser := user.Dto{
		Name:     "Ahmad",
		Email:    "ahmad@email.com",
		Password: hashPassword,
	}
	mockUserFailed := user.Dto{
		Name:     "Ahmad",
		Email:    "ipan@email.com",
		Password: hashPassword,
	}
	u := NewUserService(db)

	// setup expectations
	s := t.Run("success", func(t *testing.T) {
		// success scenario update
		data, err := u.Update("7dd77cc4-f786-4be0-b5a5-0c203b9c62c5", mockUser)
		assert.NoError(t, err)
		assert.NotEmpty(t, data)
	})
	f := t.Run("error-failed", func(t *testing.T) {
		// failed scenario update
		_, err := u.Update("7dd77cc4-f786-4be0-b5a5-0c203b9c62c5", mockUserFailed)
		assert.Error(t, err)
	})
	assert.Equal(t, true, s, "Success scenario failed run")
	assert.Equal(t, true, f, "Failed scenario failed run")
}

func TestUserService_Delete(t *testing.T) {
	//prepare database test
	db, _ := database.PrepareTestDB("txdb")
	defer database.CleanTestDB(db)
	//create instance of our test object
	u := NewUserService(db)

	// setup expectations
	s := t.Run("success", func(t *testing.T) {
		// success scenario delete
		success, err := u.Delete("7dd77cc4-f786-4be0-b5a5-0c203b9c62c5")
		assert.Equal(t, true, success)
		assert.Nil(t, err)
		assert.NoError(t, err)
	})
	f := t.Run("error-failed", func(t *testing.T) {
		// failed scenario delete
		success, err := u.Delete("7dd77cc4-f786-4be0-b5a5-0c203b9e")
		assert.NotNil(t, err)
		assert.Equal(t, false, success)
	})
	assert.Equal(t, true, s, "Success scenario failed run")
	assert.Equal(t, true, f, "Failed scenario failed run")
}
