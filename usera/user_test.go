package usera_test

import (
	"testing"

	"example.com/hello/doer/mocks"
	"example.com/hello/usera"
)

func TestUse(t *testing.T) {
	mockDoer := new(mocks.Doer)
	mockDoer.On("DoSomething", 123, "Hello GoMock").Return(nil)
	testUser := usera.Usera{mockDoer}
	testUser.Use()
	mockDoer.AssertExpectations(t)

}
