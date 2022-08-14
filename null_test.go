package just_test

import (
	"database/sql"
	"github.com/kazhuravlev/just"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNullConstructor(t *testing.T) {
	t.Run("null_null", func(t *testing.T) {
		v1 := just.NullNull[int]()
		assert.Equal(t, just.NullVal[int]{Val: 0, Valid: false}, v1)

		v2 := just.NullNull[string]()
		assert.Equal(t, just.NullVal[string]{Val: "", Valid: false}, v2)

		type user struct {
			Name  string
			Email string
		}

		v3 := just.NullNull[user]()
		assert.Equal(t, just.NullVal[user]{Val: user{Name: "", Email: ""}, Valid: false}, v3)
	})

	t.Run("null", func(t *testing.T) {
		v1 := just.Null[int](10)
		assert.Equal(t, just.NullVal[int]{Val: 10, Valid: true}, v1)

		v2 := just.Null[string]("Hello")
		assert.Equal(t, just.NullVal[string]{Val: "Hello", Valid: true}, v2)

		type user struct {
			Name  string
			Email string
		}

		v3 := just.Null[user](user{Name: "Joe", Email: "joe@example.com"})
		assert.Equal(t, just.NullVal[user]{Val: user{Name: "Joe", Email: "joe@example.com"}, Valid: true}, v3)
	})

	t.Run("set_default", func(t *testing.T) {
		v := just.NullNull[int]()
		assert.True(t, v.SetDefault(10))
		assert.Equal(t, 10, v.Val)

		v2 := just.Null[int](10)
		assert.False(t, v2.SetDefault(42))
		assert.Equal(t, 10, v2.Val)
	})

	t.Run("value_ok", func(t *testing.T) {
		v1, ok := just.NullNull[int]().ValueOk()
		assert.Equal(t, 0, v1)
		assert.False(t, ok)

		v2, ok := just.Null[int](10).ValueOk()
		assert.Equal(t, 10, v2)
		assert.True(t, ok)
	})

	t.Run("value", func(t *testing.T) {
		v1, err := just.NullNull[int]().Value()
		assert.Equal(t, nil, v1)
		assert.NoError(t, err)

		v2, err := just.NullNull[sql.NullBool]().Value()
		assert.Equal(t, nil, v2)
		assert.NoError(t, err)

		v3, err := just.Null[int](10).Value()
		assert.Equal(t, 10, v3)
		assert.NoError(t, err)

		v4, err := just.Null[sql.NullBool](sql.NullBool{Bool: true, Valid: true}).Value()
		assert.Equal(t, true, v4)
		assert.NoError(t, err)
	})

	t.Run("scan", func(t *testing.T) {
		v1 := just.NullNull[int]()
		assert.NoError(t, v1.Scan(10))
		assert.Equal(t, 10, v1.Val)
		assert.True(t, v1.Valid)

		v2 := just.NullNull[sql.NullString]()
		assert.NoError(t, v2.Scan("hi"))
		assert.Equal(t, sql.NullString{String: "hi", Valid: true}, v2.Val)
		assert.True(t, v2.Valid)

		v3 := just.NullNull[int]()
		assert.NoError(t, v3.Scan(nil))
		assert.Equal(t, 0, v3.Val)
		assert.True(t, v3.Valid)

		v4 := just.NullNull[int]()
		assert.Error(t, v4.Scan("this is not integer"))
		assert.Equal(t, 0, v4.Val)
		assert.False(t, v4.Valid)

		v5 := just.NullNull[sql.NullInt64]()
		assert.Error(t, v5.Scan("this is not int64"))
		assert.Equal(t, sql.NullInt64{Int64: 0, Valid: false}, v5.Val)
		assert.False(t, v5.Valid)
	})
}
