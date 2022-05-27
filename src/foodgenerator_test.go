package webserver

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsMatchingCases(t *testing.T) {
	haystack := []string{"Foo", "foo", "bar"}
	needle := "Foo"
	fmt.Println(haystack, needle)
	assert.True(t, contains(haystack, needle))
}

func TestContainsDifferentCases(t *testing.T) {
	haystack := []string{"foo", "bar"}
	needle := "Bar"
	fmt.Println(haystack, needle)
	assert.True(t, contains(haystack, needle))
}
