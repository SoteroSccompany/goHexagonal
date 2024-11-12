package cli_test

import (
	"fmt"
	"testing"

	"github.com/Soter-Tec/go-hexagonal/adapters/cli"
	mock_aplication "github.com/Soter-Tec/go-hexagonal/aplication/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	cntrl := gomock.NewController(t)
	defer cntrl.Finish()

	productName := "Product Test"
	productPrice := 25.5
	productID := "abc"
	productStatus := "enabled"

	productMock := mock_aplication.NewMockProductInterface(cntrl)
	productMock.EXPECT().GetId().Return(productID).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	service := mock_aplication.NewMockProductServiceInterface(cntrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productID).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s", productID,
		productName, productPrice, productStatus)
	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been enabled", productName)
	result, err = cli.Run(service, "enable", productID, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been disabled", productName)
	result, err = cli.Run(service, "disable", productID, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID %s with the name %s has the price %f and status %s", productID,
		productName, productPrice, productStatus)
	result, err = cli.Run(service, "", productID, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

}
