package prompt

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInput(t *testing.T) {
	tests := []struct {
		v string
		e bool
	}{
		{"hello\n", false},
		{"hello", true},
	}
	for _, test := range tests {
		input = strings.NewReader(test.v)
		output, err := GetInput()
		if err != nil {
			assert.Equal(t, test.e, err == io.EOF)
			continue
		}
		assert.Equal(t, test.v, output)
	}
}
