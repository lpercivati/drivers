package tests

import (
	"root/services"
	"root/test/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldGetDriversOk(t *testing.T) {

	driverService := services.DriverService{
		Repository: new(mocks.DriverRepository),
	}

	drivers, err := driverService.GetDrivers(1, 1)

	if err != nil {
		assert.True(t, false)
	}

	assert.Equal(t, len(drivers), 1)
	assert.Equal(t, drivers[0].Id, 10)
	assert.Equal(t, drivers[0].Fullname, "Leandro")
}

func Test_ShouldGetDriversError(t *testing.T) {

	driverService := services.DriverService{
		Repository: new(mocks.DriverRepository),
	}

	drivers, err := driverService.GetDrivers(1, 2)

	if err == nil {
		assert.True(t, false)
	}

	assert.Equal(t, len(drivers), 0)
	assert.NotNil(t, err)
}
