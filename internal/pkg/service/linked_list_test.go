package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_LinkedListServiceBasicUsage(t *testing.T) {

	t.Run("Success", func(t *testing.T) {

		var (
			source  []int64
			service = NewLinkedListService()
		)

		for i := 1; i <= 100; i++ {
			source = append(source, int64(i))
		}

		for _, v := range source {
			service.InsertRight(v)
		}

		first, getFirstError := service.GetFirst()
		last, getLastError := service.GetLast()

		require.NoError(t, getFirstError)
		require.NoError(t, getLastError)

		require.Equal(t, service.GetLength(), len(source))
		require.Equal(t, first, source[0])
		require.Equal(t, last, source[len(source)-1])
	})
}
