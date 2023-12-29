package public_service_test

import (
	"Farashop/internal/contract"
	"Farashop/internal/dto"
	"Farashop/internal/entity"
	"Farashop/internal/service/public_service"
	storemock "Farashop/pkg/mocks/store"
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
)

func setup(t *testing.T) (contract.PublicInteractor, *storemock.MockPublicStore, func()) {
	mockCtr1 := gomock.NewController(t)
	mockPublicStore := storemock.NewMockPublicStore(mockCtr1)

	service := public_service.NewPublic(mockPublicStore)
	return service, mockPublicStore, func() {
		mockCtr1.Finish()
	}
}

func TestRegister(t *testing.T) {
	t.Run("Successful Registre", func(t *testing.T) {
		interactor, mockPublicStore, teardown := setup(t)
		defer teardown()

		req := dto.RegisterUserRequest{
			Username: faker.Username(),
			Email:    faker.Email(),
			Password: faker.Password(),
		}

		ctx := context.Background()
		passedUser := entity.User{
			ID:                0,
			Verification_code: 55458,
			Username:          req.Username,
			Email:             req.Email,
			Password:          req.Password,
		}

		mockPublicStore.EXPECT().Register(ctx, passedUser).Return(true, nil)
		resp, err := interactor.Register(ctx, req)

		if resp.Result != true && err != nil {
			t.Fatal("Error")
		}
	})
}
