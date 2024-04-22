package ssmparam

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("RULEGROUPNAME", "test")
	os.Setenv("COUNTERSSMPATH", "/app/service/counter")

	// Run all the tests
	code := m.Run()

	// Exit with the return code determined by the tests
	os.Exit(code)
}

func TestGetParam(t *testing.T) {
	counterSSMParam := os.Getenv("COUNTERSSMPATH")
	num, err := GetCounter(counterSSMParam)
	if err != nil {
		t.Error(err)
	}

	err = IncreaseCounter(counterSSMParam)
	if err != nil {
		t.Error(err)
	}

	num2, err := GetCounter(counterSSMParam)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, num2, (num + 1))
}
