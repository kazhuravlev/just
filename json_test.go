package just_test

import (
	"github.com/kazhuravlev/just"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

type SomeType struct {
	ID int `json:"id"`
}

func TestJsonParseType(t *testing.T) {
	t.Run("valid_type", func(t *testing.T) {
		res, err := just.JsonParseType[SomeType]([]byte(`{"id":42}`))
		require.NoError(t, err)
		require.Equal(t, SomeType{ID: 42}, *res)
	})

	t.Run("invalid_type", func(t *testing.T) {
		res, err := just.JsonParseType[SomeType]([]byte(`42`))
		require.Error(t, err)
		require.Nil(t, res)
	})
}

func TestJsonParseTypeF(t *testing.T) {
	t.Run("valid_type", func(t *testing.T) {
		tmpFile, err := os.CreateTemp(os.TempDir(), "test-json-file")
		require.NoError(t, err)
		{
			_, err := tmpFile.Write([]byte(`{"id":42}`))
			require.NoError(t, err)

			tmpFile.Close()
		}

		res, err := just.JsonParseTypeF[SomeType](tmpFile.Name())
		require.NoError(t, err)
		require.Equal(t, SomeType{ID: 42}, *res)
	})

	t.Run("invalid_type", func(t *testing.T) {
		res, err := just.JsonParseTypeF[SomeType]("/path/not-exists/png")
		require.Error(t, err)
		require.Nil(t, res)
	})
}
