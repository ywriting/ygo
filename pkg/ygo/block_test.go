package ygo_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"riguz.com/ygo/pkg/ygo"
)

func TestItemFlags_Default(t *testing.T) {
	flags := ygo.NewItemFlags(0)

	assert.Equal(t, false, flags.IsKeep())
	assert.Equal(t, false, flags.IsCountable())
	assert.Equal(t, false, flags.IsDeleted())
	assert.Equal(t, false, flags.IsMarked())
}

func TestItemFlags_SetFlag(t *testing.T) {
	flags := ygo.NewItemFlags(0)
	flags.SetCountable()
	flags.SetDeleted()
	assert.Equal(t, true, flags.IsCountable())
	assert.Equal(t, true, flags.IsDeleted())

	assert.Equal(t, false, flags.IsKeep())
	flags.Set(ygo.ITEM_FLAG_KEEP)
	assert.Equal(t, true, flags.IsKeep())
}

func TestItemFlags_Clear(t *testing.T) {
	flags := ygo.NewItemFlags(0b1111_1111)
	assert.Equal(t, true, flags.IsCountable())
	flags.Clear(ygo.ITEM_FLAG_COUNTABLE)
	assert.Equal(t, false, flags.IsCountable())
}
