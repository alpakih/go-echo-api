package usecase

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-txdb"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	"go-echo-api/entity"
	"go-echo-api/user"
	"go-echo-api/user/mocks"
	"go-echo-api/utils"
	"testing"
	"time"
)

func init() {
	RegisterTxDB("txdb")
}

// RegisterTxDB registers new db for single transaction tests
func RegisterTxDB(name string) {
	txdb.Register(name, "postgres", "postgres://postgres:postgres@localhost:5433/echo_api?sslmode=disable")
}

// PrepareTestDB prepare test DB according to txdb name
func PrepareTestDB(withName string) (*gorm.DB, error) {
	sqlDB, err := sql.Open(withName, fmt.Sprintf("connection_%d", time.Now().UnixNano()))
	db, err := gorm.Open("postgres", sqlDB)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(entity.User{})
	return db, err
}

// CleanTestDB drops all tables from test DB
func CleanTestDB(db *gorm.DB) {
	_ = db.Close()
}

func TestUserServiceFindAll(t *testing.T) {
	db, _ := PrepareTestDB("txdb")
	defer CleanTestDB(db)
	// create an instance of our test object
	mockUserRepository := new(mocks.Repository)
	var mockListUser []entity.User
	// setup expectations
	t.Run("success", func(t *testing.T) {
		// call the code we are testing
		u := NewUserService(db)
		list, err := u.FindAll()
		assert.NotEmpty(t, list, "Empty object")
		assert.NoError(t, err, "Error")
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		// call the code we are testing
		assert.Empty(t, mockListUser, "EMPTY")
		assert.Len(t, mockListUser, 0, "LEN ERROR")
		mockUserRepository.AssertExpectations(t)
	})

}

func TestUserService_FindById(t *testing.T) {
	db, _ := PrepareTestDB("txdb")
	defer CleanTestDB(db)
	mockUserRepository := new(mocks.Repository)
	mockUser := entity.User{
		ID:    "7dd77cc4-f786-4be0-b5a5-0c203b9c62c5",
		Name:  "Uje",
		Email: "uje@email.com",
	}
	// setup expectations
	t.Run("success", func(t *testing.T) {
		// call the code we are testing
		u := NewUserService(db)
		data, err := u.FindById(mockUser.ID)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		// call the code we are testing
		u := NewUserService(db)
		data, err := u.FindById("test")
		assert.Error(t, err)
		assert.Nil(t, data)
		mockUserRepository.AssertExpectations(t)
	})
}

func TestUserService_Save(t *testing.T) {
	db, _ := PrepareTestDB("txdb")
	defer CleanTestDB(db)
	mockUserRepository := new(mocks.Repository)
	hashPassword, _ := utils.HashPassword("password")
	mockUser := user.Dto{
		Name:     "Ahmad",
		Email:    "ahmad@email.com",
		Password: hashPassword,
	}
	// setup expectations
	t.Run("success", func(t *testing.T) {
		// call the code we are testing
		u := NewUserService(db)
		data, err := u.Save(mockUser)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		mockUserRepository.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		// call the code we are testing
		u := NewUserService(db)
		_, err := u.Save(mockUser)
		assert.Error(t, err)
		mockUserRepository.AssertExpectations(t)
	})
}

func TestUserService_Update(t *testing.T) {
	db, _ := PrepareTestDB("txdb")
	defer CleanTestDB(db)
	mockUserRepository := new(mocks.Repository)
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
	// setup expectations
	t.Run("success", func(t *testing.T) {
		// call the code we are testing
		u := NewUserService(db)
		data, err := u.Update("7dd77cc4-f786-4be0-b5a5-0c203b9c62c5", mockUser)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		mockUserRepository.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		// call the code we are testing
		u := NewUserService(db)
		_, err := u.Update("7dd77cc4-f786-4be0-b5a5-0c203b9c62c5", mockUserFailed)
		assert.Error(t, err)
		mockUserRepository.AssertExpectations(t)
	})
}

func TestUserService_Delete(t *testing.T) {
	db, _ := PrepareTestDB("txdb")
	defer CleanTestDB(db)
	mockUserRepository := new(mocks.Repository)
	// setup expectations
	t.Run("success", func(t *testing.T) {
		// call the code we are testing
		u := NewUserService(db)
		success, err := u.Delete("7dd77cc4-f786-4be0-b5a5-0c203b9c62c5")
		assert.Equal(t, true, success)
		assert.Nil(t, err)
		assert.NoError(t, err)
		mockUserRepository.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		// call the code we are testing
		u := NewUserService(db)
		success, err := u.Delete("7dd77cc4-f786-4be0-b5a5-0c203b9e")
		assert.NotNil(t, err)
		assert.Equal(t, false, success)
		mockUserRepository.AssertExpectations(t)
	})
}
