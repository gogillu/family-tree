package dag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name   string
		node   Node
		expErr error
	}{
		{
			"Valid Item",
			Node{
				"id-1",
				"name-1",
				nil,
				nil,
			},
			nil,
		},
	}

	for _, tc := range tests {
		actualErr := tc.node.validate()
		if tc.expErr == nil {
			assert.Equal(t, actualErr, tc.expErr, tc.name)
		}
	}
}
