package auth

import (
	"fmt"

	"github.com/RussellLuo/appx"
	"github.com/RussellLuo/kun/pkg/httpcodec"
	"github.com/RussellLuo/structool"
)

func init() {
	appx.MustRegister(appx.New("auth", new(Auth)))
}

// Auth implements httpcodec.ParamCodec to encode and decode the `Authorization` header.
type Auth struct {
	Token string `structool:"token"`

	Codecs *httpcodec.DefaultCodecs
}

func (a *Auth) Init(ctx appx.Context) error {
	if err := structool.New().Decode(ctx.Config(), a); err != nil {
		return err
	}

	// Use Auth to encode and decode the argument named "__", if exists,
	// for all the operations.
	a.Codecs = httpcodec.NewDefaultCodecs(nil).
		PatchAll(func(c httpcodec.Codec) *httpcodec.Patcher {
			return httpcodec.NewPatcher(c).Param("__", a)
		})

	return nil
}

// Decode decodes the `Authorization` header.
func (a *Auth) Decode(in []string, out interface{}) error {
	// NOTE: never use out, which is nil here.

	// Skip authentication if an empty token is specified.
	if a.Token == "" {
		return nil
	}

	if len(in) == 0 || in[0] != a.bearerToken() {
		return fmt.Errorf("authentication failed")
	}
	return nil
}

// Encode encodes the `Authorization` header.
func (a *Auth) Encode(in interface{}) (out []string) {
	// NOTE: never use in, which is nil here.

	// Skip authentication if an empty token is specified.
	if a.Token == "" {
		return nil
	}

	return []string{a.bearerToken()}
}

func (a *Auth) bearerToken() string {
	return "Bearer " + a.Token
}
