# uvdecoder

[![Build Status](https://travis-ci.org/nowk/uvdecoder.svg?branch=master)](https://travis-ci.org/nowk/uvdecoder)
[![GoDoc](https://godoc.org/github.com/nowk/uvdecoder?status.svg)](http://godoc.org/github.com/nowk/uvdecoder)

`url.Values` DecodeHook for [mapstructure](https://github.com/mitchellh/mapstructure)

## Examples

    type User struct {
      Email string `mapstructure:"email"`
    }

    postData := url.Values{
      "email": {"email@company.com"},
    }

---

    var user User
    conf := uvdecoder.Config(&user)
    decoder, err := mapstructure.NewDecoder(conf)
    if err != nil {
      // handle err
    }

    err := decoder.Decode(postData)
    if err != nil {
      // handle err
    }

## License

MIT