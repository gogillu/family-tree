package dag

import (
	"fmt"
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
		{
			"Invalid Item",
			Node{
				"",
				"name-2",
				nil,
				nil,
			},
			fmt.Errorf("error : expected non empty value id: cannot be blank."),
		},
	}

	for _, tc := range tests {
		actualErr := tc.node.validate()
		if tc.expErr == nil || true {
			assert.Equal(t, actualErr, tc.expErr, tc.name)
		} else {
			assert.NotEqual(t, actualErr, tc.expErr, tc.name)
		}
	}
}
