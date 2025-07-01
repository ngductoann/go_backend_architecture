package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	var (
		input  = 1
		output = 2
	)

	actual := AddOne(input)
	if actual != output {
		t.Errorf("AddOne(%d) = %d; want %d", input, actual, output)
	}

	actual = AddOne2(input)
	if actual != output {
		t.Errorf("AddOne(%d) = %d; want %d", input, actual, output)
	}

	assert.Equal(t, AddOne(2), 3, "AddOne(2) should return 3")
	assert.NotEqual(t, AddOne(3), 5, "AddOne(3) should not return 5")
	assert.Greater(t, AddOne(4), 4, "AddOne(4) should return a value greater than 4")
}

func TestRequire(t *testing.T) {
	require.Equal(t, 2, 3)
	fmt.Println("Not executing")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 2, 3)
	fmt.Println("executing")
}
