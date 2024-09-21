package optional

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValue_MarshalJSON(t *testing.T) {
	t.Run("marshal nil value", func(t *testing.T) {
		bytes, err := json.Marshal(Nil[int]())
		require.NoError(t, err)
		assert.EqualValues(t, "null", string(bytes))
	})

	t.Run("marshal builtin value", func(t *testing.T) {
		data, err := json.Marshal(Some[int](42))
		require.NoError(t, err)
		assert.EqualValues(t, "42", string(data))
	})

	t.Run("marshal structure", func(t *testing.T) {
		type S struct {
			Value int           `json:"value"`
			Opt   Value[int]    `json:"opt"`
			Nil   Value[string] `json:"nil"`
		}

		s := S{
			Value: 18,
			Opt:   Some(24),
		}
		data, err := json.Marshal(s)
		require.NoError(t, err)
		assert.EqualValues(t, `{"value":18,"opt":24,"nil":null}`, string(data))
	})
}

func TestValue_UnmarshalJSON(t *testing.T) {
	t.Run("unmarshal nil value", func(t *testing.T) {
		var value Value[int]
		err := json.Unmarshal([]byte("null"), &value)
		require.NoError(t, err)

		assert.EqualValues(t, value, Nil[int]())
	})

	t.Run("unmarshal builtin value", func(t *testing.T) {
		var value Value[int]
		err := json.Unmarshal([]byte("42"), &value)
		require.NoError(t, err)
		assert.EqualValues(t, 42, value.Get())
	})

	t.Run("unmarshal structure", func(t *testing.T) {
		type S struct {
			Value int           `json:"value"`
			Opt   Value[int]    `json:"opt"`
			Nil   Value[string] `json:"nil"`
		}

		var value Value[S]
		err := json.Unmarshal([]byte(`{"value":18,"opt":24,"nil":null}`), &value)
		require.NoError(t, err)
		assert.EqualValues(t, S{
			Value: 18,
			Opt:   Some(24),
		}, value.Get())
	})
}
