package embedding

import (
	"context"

	"github.com/RussellLuo/appx"
	"github.com/RussellLuo/kun/pkg/appx/httpapp"
	"github.com/RussellLuo/kun/pkg/werror"
	"github.com/RussellLuo/kun/pkg/werror/gcode"
	"github.com/RussellLuo/structool"
	"github.com/go-chi/chi"

	"github.com/go-aie/oneai/auth"
	"github.com/go-aie/oneai/embedding/api"
	"github.com/go-aie/oneai/embedding/controller/http"
	"github.com/go-aie/oneai/embedding/rocketqa"
)

func init() {
	appx.MustRegister(
		httpapp.New("embedding", new(Embedding)).
			MountOn("oneai", "/embeddings").
			Require("auth").App,
	)
}

type Embedding struct {
	encoders map[string]api.Embedding
	router   chi.Router
}

func (e *Embedding) Router() chi.Router {
	return e.router
}

func (e *Embedding) Init(ctx appx.Context) error {
	codec := structool.New()

	var configs map[string]map[string]interface{}
	if err := codec.Decode(ctx.Config(), &configs); err != nil {
		return err
	}

	e.encoders = make(map[string]api.Embedding)
	for model, config := range configs {
		switch model {
		case "rocketqa":
			var cfg *rocketqa.Config
			if err := codec.Decode(config, &cfg); err != nil {
				return err
			}
			encoder, err := rocketqa.New(cfg)
			if err != nil {
				return err
			}
			e.encoders["rocketqa-query"] = encoder
			e.encoders["rocketqa-document"] = encoder
		}
	}

	a := ctx.MustLoad("auth").(*auth.Auth)
	e.router = http.NewHTTPRouter(e, a.Codecs)

	return nil
}

func (e *Embedding) Encode(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	encoder, ok := e.encoders[req.Model]
	if !ok {
		return nil, werror.Wrapf(gcode.ErrInvalidArgument, "unsupported model: %s", req.Model)
	}
	return encoder.Encode(ctx, req)
}
