package starter_test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	starter "github.com/williaminfante/go_test_starter"
)

func TestSayHello(t *testing.T) {
	greeting := starter.SayHello("William")
	assert.Equal(t, greeting, "Hello William. Welcome!")
}

func TestPickAnInteger(t *testing.T) {
	t.Run("Check for Non-Negative Numbers", func(t *testing.T) {
		assert.Equal(t, starter.PickAnInteger(45), "45 is an odd number")
		assert.Equal(t, starter.PickAnInteger(42), "42 is an even number")
		assert.Equal(t, starter.PickAnInteger(0), "0 is an even number")
	})
	t.Run("Check for Negative Numbers", func(t *testing.T) {
		assert.Equal(t, starter.PickAnInteger(-45), "-45 is an odd number")
		assert.Equal(t, starter.PickAnInteger(-42), "-42 is an even number")
	})
}
func TestCheckhealth(t *testing.T) {
	t.Run("Check health status", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://mysite.com/example", nil)
		writer := httptest.NewRecorder()
		starter.Checkhealth(writer, req)
		response := writer.Result()
		body, err := io.ReadAll(response.Body)

		assert.Equal(t, response.StatusCode, 200)
		assert.Equal(t,
			response.Header.Get("Content-Type"),
			"text/plain; charset=utf-8")

		assert.Equal(t, string(body), "health check passed")
		assert.Equal(t, err, nil)
	})
}
