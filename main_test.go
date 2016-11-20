package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNameToUnderscore(t *testing.T) {
	assert.Equal(t, "query", NameToUnderscore("Query"))
	assert.Equal(t, "read_buffer_from_file_descriptor_read", NameToUnderscore("ReadBufferFromFileDescriptorRead"))
}
