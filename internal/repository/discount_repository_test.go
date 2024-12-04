package repository

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDiscountRepository(t *testing.T) {
	repo := NewDiscountRepository()

	t.Run("CreateDiscountCode", func(t *testing.T) {
		err := repo.CreateDiscountCode("SAVE10", 10.0)
		assert.NoError(t, err)

		discountCode, err := repo.GetDiscountCode("SAVE10")
		assert.NoError(t, err)
		assert.NotNil(t, discountCode)
		assert.Equal(t, "SAVE10", discountCode.Code)
		assert.Equal(t, float64(10.0), discountCode.Percentage)
		assert.False(t, discountCode.IsUsed)
	})

	t.Run("GetDiscountCode_Existing", func(t *testing.T) {
		err := repo.CreateDiscountCode("SAVE20", 20.0)
		assert.NoError(t, err)

		discountCode, err := repo.GetDiscountCode("SAVE20")
		assert.NoError(t, err)
		assert.NotNil(t, discountCode)
		assert.Equal(t, "SAVE20", discountCode.Code)
		assert.Equal(t, float64(20.0), discountCode.Percentage)
	})

	t.Run("GetDiscountCode_NonExistent", func(t *testing.T) {
		discountCode, err := repo.GetDiscountCode("NONEXISTENT")
		assert.NoError(t, err)
		assert.Nil(t, discountCode)
	})

	t.Run("UseDiscountCode", func(t *testing.T) {
		err := repo.CreateDiscountCode("SAVE15", 15.0)
		assert.NoError(t, err)

		err = repo.UseDiscountCode("SAVE15", "user1")
		assert.NoError(t, err)

		discountCode, err := repo.GetDiscountCode("SAVE15")
		assert.NoError(t, err)
		assert.NotNil(t, discountCode)
		assert.True(t, discountCode.IsUsed)
		assert.Equal(t, "user1", discountCode.UsedByUserID)
	})

	t.Run("UseDiscountCode_AlreadyUsed", func(t *testing.T) {
		err := repo.CreateDiscountCode("SAVE25", 25.0)
		assert.NoError(t, err)

		err = repo.UseDiscountCode("SAVE25", "user1")
		assert.NoError(t, err)

		err = repo.UseDiscountCode("SAVE25", "user2")
		assert.Error(t, err)
	})
}