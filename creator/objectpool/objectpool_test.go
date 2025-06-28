package objectpool

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewObjectPool(t *testing.T) {
	assert := assert.New(t)

	pool := NewObjectPool(3)

	object1 := pool.AcquireObject()
	assert.Equal(0, object1.ID)

	object2 := pool.AcquireObject()
	assert.Equal(1, object2.ID)

	pool.ReleaseObject(object1)

	object3 := pool.AcquireObject()
	assert.Equal(2, object3.ID)

	object4 := pool.AcquireObject()
	assert.Equal(0, object4.ID)
}
