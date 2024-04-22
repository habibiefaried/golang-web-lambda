package ssmparam

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

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
