package ssmparam

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestGetParam(t *testing.T) {
	counterSSMParam := "/app/service/counter"
	s, err := getParameter(counterSSMParam)
	if err != nil {
		t.Error(err)
	}
	t.Log(*s)
	num, err := strconv.Atoi(*s)
	if err != nil {
		t.Error(err)
	}

	err = IncreaseCounter(counterSSMParam)
	if err != nil {
		t.Error(err)
	}

	s2, err := getParameter(counterSSMParam)
	if err != nil {
		t.Error(err)
	}
	t.Log(*s2)
	num2, err := strconv.Atoi(*s2)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, num2, (num + 1))
}
