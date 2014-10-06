package uvdecoder

import "net/url"
import "testing"
import "github.com/mitchellh/mapstructure"
import "github.com/nowk/assert"

type User struct {
	Email   string   `mapstructure:"email"`
	Numbers []string `mapstructure:"numbers"`
	Age     int      `mapstructure:"age"`
	Admin   bool     `mapstructure:"admin"`
	Auth    []int    `mapstructure:"auth"`
}

var p = url.Values{
	"email":   {"foo@bar.com"},
	"numbers": {"555callme", "555pizza"},
	"age":     {"19"},
	"admin":   {"1"},
	"auth":    {"45", "17", "95"},
}

func check(t *testing.T, err error) {
	if err != nil {
		t.Errorf("error: %s", err)
	}
}

func TestFlattensUrlValueMapValuesForNonSlices(t *testing.T) {
	var d User
	decoder, err := mapstructure.NewDecoder(Config(&d))
	check(t, err)

	err = decoder.Decode(p)
	check(t, err)

	assert.Equal(t, d, User{
		Email:   "foo@bar.com",
		Numbers: []string{"555callme", "555pizza"},
		Age:     19,
		Admin:   true,
		Auth:    []int{45, 17, 95},
	})
}
