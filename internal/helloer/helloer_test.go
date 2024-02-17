package helloer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelloer_SayHello(t *testing.T) {
	t.Parallel()

	type testCase struct {
		to   string
		want string
	}

	testCases := []testCase{
		{
			to:   "me",
			want: "say hello: me",
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run("case", func(t *testing.T) {
			t.Parallel()

			h := New(tt.to)
			got := h.SayHello()
			require.Equal(t, tt.want, got)
		})
	}
}
