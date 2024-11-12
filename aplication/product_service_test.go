package aplication_test

import (
	"testing"

	"github.com/Soter-Tec/go-hexagonal/aplication"
	mock_aplication "github.com/Soter-Tec/go-hexagonal/aplication/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestProductProductService_Get(t *testing.T) {
	cntrl := gomock.NewController(t)
	defer cntrl.Finish()

	product := mock_aplication.NewMockProductInterface(cntrl)
	persistence := mock_aplication.NewMockProductPersistenceInterface(cntrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil)

	service := aplication.ProductService{
		Persistence: persistence,
	}

	result, err := service.Get("1")
	require.Nil(t, err)
	require.Equal(t, product, result)

}

func TestProductService_Create(t *testing.T) {
	cntrl := gomock.NewController(t)
	defer cntrl.Finish()

	product := mock_aplication.NewMockProductInterface(cntrl)
	persistence := mock_aplication.NewMockProductPersistenceInterface(cntrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil)

	service := aplication.ProductService{
		Persistence: persistence,
	}

	result, err := service.Create("Product 1", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_EnableDisable(t *testing.T) {
	cntrl := gomock.NewController(t)
	defer cntrl.Finish()

	product := mock_aplication.NewMockProductInterface(cntrl)
	product.EXPECT().Enable().Return(nil)
	product.EXPECT().Disable().Return(nil)

	persistence := mock_aplication.NewMockProductPersistenceInterface(cntrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()

	service := aplication.ProductService{
		Persistence: persistence,
	}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

}
