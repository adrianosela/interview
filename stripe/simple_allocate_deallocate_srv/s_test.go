package simple_allocate_deallocate_srv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllocateManySameName(t *testing.T) {
	s := NewSystem()
	assert.Equal(t, 0, s.Allocate("seal"))
	assert.Equal(t, 1, s.Allocate("seal"))
	assert.Equal(t, 2, s.Allocate("seal"))
}

func TestAllocateManyDifferentName(t *testing.T) {
	s := NewSystem()
	assert.Equal(t, 0, s.Allocate("seal"))
	assert.Equal(t, 1, s.Allocate("seal"))
	assert.Equal(t, 0, s.Allocate("goat"))
	assert.Equal(t, 1, s.Allocate("goat"))
}

func TestAllocateDeallocateAllocate(t *testing.T) {
	s := NewSystem()
	assert.Equal(t, 0, s.Allocate("seal"))
	assert.Equal(t, 1, s.Allocate("seal"))
	s.Deallocate("seal0")
	s.Deallocate("seal1")
	assert.Equal(t, 0, s.Allocate("seal"))
	assert.Equal(t, 1, s.Allocate("seal"))
}

func TestAllocateDeallocateSomeAllocate(t *testing.T) {
	s := NewSystem()
	assert.Equal(t, 0, s.Allocate("seal"))
	assert.Equal(t, 1, s.Allocate("seal"))
	assert.Equal(t, 2, s.Allocate("seal"))
	assert.Equal(t, 3, s.Allocate("seal"))
	assert.Equal(t, 4, s.Allocate("seal"))
	s.Deallocate("seal3")
	assert.Equal(t, 3, s.Allocate("seal"))
	s.Deallocate("seal2")
	assert.Equal(t, 2, s.Allocate("seal"))
}
