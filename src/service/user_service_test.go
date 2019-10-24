package service_test

import (
	"database/sql"
	"errors"
	. "github.com/sashiyama/syncnow_server/model"
	. "github.com/sashiyama/syncnow_server/repository"
	"github.com/sashiyama/syncnow_server/service"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type UserRepositoryStub struct{ UserRepository }
type UserRepositoryErrorStub struct{ UserRepository }
type UserCredentialRepositoryStub struct{}
type AccessTokenRepositoryStub struct{ AccessTokenRepository }
type RefreshTokenRepositoryStub struct{ RefreshTokenRepository }
type TransactionRepositoryStub struct{}

func (urs *UserRepositoryStub) Create(p UserCreateParam) (string, error) {
	return "test-user-id", nil
}

func (ures *UserRepositoryErrorStub) Create(p UserCreateParam) (string, error) {
	return "", errors.New("User creation failed")
}

func (urs *UserRepositoryStub) FindByUser(u User) (string, error) {
	return "test-user-id", nil
}

func (urs *UserRepositoryErrorStub) FindByUser(u User) (string, error) {
	return "", errors.New("Not Exist")
}

func (ucr *UserCredentialRepositoryStub) Create(p UserCredentialCreateParam) error {
	return nil
}

func (ucr *UserCredentialRepositoryStub) ExistsByEmail(email string) (bool, error) {
	if email == "test@example.com" {
		return true, nil
	} else {
		return false, errors.New("Not Exist")
	}
}

func (atr *AccessTokenRepositoryStub) Create(p AccessTokenParam) (AccessToken, error) {
	return AccessToken{Id: "test-id", Token: "test-access-token", ExpiresAt: time.Now().Add(24 * time.Hour)}, nil
}

func (atr *AccessTokenRepositoryStub) Update(p AccessTokenParam) (AccessToken, error) {
	return AccessToken{Id: "test-id", Token: "test-access-token", ExpiresAt: time.Now().Add(24 * time.Hour)}, nil
}

func (atr *AccessTokenRepositoryStub) FindByToken(token string) (AccessToken, error) {
	return AccessToken{Id: "test-id", Token: "test-access-token", ExpiresAt: time.Now().Add(24 * time.Hour)}, nil
}

func (rtr *RefreshTokenRepositoryStub) Create(p RefreshTokenParam) (RefreshToken, error) {
	return RefreshToken{Id: "test-id", Token: "test-refresh-token", ExpiresAt: time.Now().Add(72 * time.Hour)}, nil
}

func (rtr *RefreshTokenRepositoryStub) Update(p RefreshTokenParam) (RefreshToken, error) {
	return RefreshToken{Id: "test-id", Token: "test-refresh-token", ExpiresAt: time.Now().Add(72 * time.Hour)}, nil
}

func (tr *TransactionRepositoryStub) Transaction(txFunc func(*sql.Tx) (interface{}, error)) (interface{}, error) {
	r, err := txFunc(&sql.Tx{})
	return r, err
}

func TestUserServicIsRegistered(t *testing.T) {
	ucr := &UserCredentialRepositoryStub{}
	us := service.UserService{UserCredentialRepository: ucr}

	t.Run("When it's registered", func(t *testing.T) {
		exists, err := us.IsRegistered("test@example.com")
		assert.Equal(t, exists, true)
		assert.Nil(t, err)
	})

	t.Run("When it isn't registered", func(t *testing.T) {
		exists, err := us.IsRegistered("test+1@example.com")
		assert.Equal(t, exists, false)
		assert.NotNil(t, err)
	})
}

func TestUserServiceSignUp(t *testing.T) {
	t.Run("When signup is successful", func(t *testing.T) {
		ur := &UserRepositoryStub{}
		ucr := &UserCredentialRepositoryStub{}
		atr := &AccessTokenRepositoryStub{}
		rtr := &RefreshTokenRepositoryStub{}
		tr := &TransactionRepositoryStub{}

		us := service.UserService{
			UserRepository:           ur,
			UserCredentialRepository: ucr,
			AccessTokenRepository:    atr,
			RefreshTokenRepository:   rtr,
			TransactionRepository:    tr,
		}

		authToken, err := us.SignUp(&User{Email: "test@example.com", Password: "P@ssword"}, time.Now())
		assert.Equal(t, authToken, AuthToken{AccessToken: "test-access-token", RefreshToken: "test-refresh-token"})
		assert.Nil(t, err)
	})

	t.Run("When signup fails", func(t *testing.T) {
		ur := &UserRepositoryErrorStub{}
		ucr := &UserCredentialRepositoryStub{}
		atr := &AccessTokenRepositoryStub{}
		rtr := &RefreshTokenRepositoryStub{}
		tr := &TransactionRepositoryStub{}

		us := service.UserService{
			UserRepository:           ur,
			UserCredentialRepository: ucr,
			AccessTokenRepository:    atr,
			RefreshTokenRepository:   rtr,
			TransactionRepository:    tr,
		}

		authToken, err := us.SignUp(&User{Email: "test@example.com", Password: "P@ssword"}, time.Now())
		assert.Empty(t, authToken)
		assert.NotNil(t, err)
	})
}

func TestUserServiceSignIn(t *testing.T) {
	t.Run("When signin is successful", func(t *testing.T) {
		ur := &UserRepositoryStub{}
		ucr := &UserCredentialRepositoryStub{}
		atr := &AccessTokenRepositoryStub{}
		rtr := &RefreshTokenRepositoryStub{}
		tr := &TransactionRepositoryStub{}

		us := service.UserService{
			UserRepository:           ur,
			UserCredentialRepository: ucr,
			AccessTokenRepository:    atr,
			RefreshTokenRepository:   rtr,
			TransactionRepository:    tr,
		}

		authToken, err := us.SignIn(&User{Email: "test@example.com", Password: "P@ssword"}, time.Now())
		assert.Equal(t, authToken, AuthToken{AccessToken: "test-access-token", RefreshToken: "test-refresh-token"})
		assert.Nil(t, err)
	})

	t.Run("When signin fails", func(t *testing.T) {
		ur := &UserRepositoryErrorStub{}
		ucr := &UserCredentialRepositoryStub{}
		atr := &AccessTokenRepositoryStub{}
		rtr := &RefreshTokenRepositoryStub{}
		tr := &TransactionRepositoryStub{}

		us := service.UserService{
			UserRepository:           ur,
			UserCredentialRepository: ucr,
			AccessTokenRepository:    atr,
			RefreshTokenRepository:   rtr,
			TransactionRepository:    tr,
		}

		authToken, err := us.SignUp(&User{Email: "test@example.com", Password: "P@ssword"}, time.Now())
		assert.Empty(t, authToken)
		assert.NotNil(t, err)
	})
}

func TestUserServiceIsAuthorized(t *testing.T) {
	t.Run("When authorized", func(t *testing.T) {
		ur := &UserRepositoryStub{}
		ucr := &UserCredentialRepositoryStub{}
		atr := &AccessTokenRepositoryStub{}
		rtr := &RefreshTokenRepositoryStub{}
		tr := &TransactionRepositoryStub{}

		us := service.UserService{
			UserRepository:           ur,
			UserCredentialRepository: ucr,
			AccessTokenRepository:    atr,
			RefreshTokenRepository:   rtr,
			TransactionRepository:    tr,
		}

		isAuthorized, err := us.IsAuthorized("test-access-token", time.Now())
		assert.Equal(t, isAuthorized, true)
		assert.Nil(t, err)
	})

	t.Run("When not authorized", func(t *testing.T) {
		ur := &UserRepositoryStub{}
		ucr := &UserCredentialRepositoryStub{}
		atr := &AccessTokenRepositoryStub{}
		rtr := &RefreshTokenRepositoryStub{}
		tr := &TransactionRepositoryStub{}

		us := service.UserService{
			UserRepository:           ur,
			UserCredentialRepository: ucr,
			AccessTokenRepository:    atr,
			RefreshTokenRepository:   rtr,
			TransactionRepository:    tr,
		}

		isAuthorized, err := us.IsAuthorized("test-access-token", time.Now().Add(72*time.Hour))
		assert.Equal(t, isAuthorized, false)
		assert.NotNil(t, err)
	})
}
