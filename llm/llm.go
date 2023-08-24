package llm

import (
	"context"

	"github.com/RussellLuo/appx"
	"github.com/RussellLuo/kun/pkg/appx/httpapp"
	"github.com/RussellLuo/kun/pkg/werror"
	"github.com/RussellLuo/kun/pkg/werror/gcode"
	"github.com/RussellLuo/structool"
	"github.com/go-chi/chi"

	"github.com/go-aie/oneai/auth"
	"github.com/go-aie/oneai/llm/api"
	"github.com/go-aie/oneai/llm/chatglm"
	"github.com/go-aie/oneai/llm/controller/http"
)

func init() {
	appx.MustRegister(
		httpapp.New("llm", new(LLM)).
			MountOn("oneai", "/chat").
			Require("auth").App,
	)
}

type LLM struct {
	llms   map[string]api.LLM
	router chi.Router
}

func (l *LLM) Router() chi.Router {
	return l.router
}

func (l *LLM) Init(ctx appx.Context) error {
	codec := structool.New()

	var configs map[string]map[string]interface{}
	if err := codec.Decode(ctx.Config(), &configs); err != nil {
		return err
	}

	l.llms = make(map[string]api.LLM)
	for model, config := range configs {
		switch model {
		case "chatglm":
			var cfg *chatglm.Config
			if err := codec.Decode(config, &cfg); err != nil {
				return err
			}
			llm, err := chatglm.New(cfg)
			if err != nil {
				return err
			}
			l.llms[model] = llm
		}
	}

	a := ctx.MustLoad("auth").(*auth.Auth)
	l.router = http.NewHTTPRouter(l, a.Codecs)

	return nil
}

func (l *LLM) Chat(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	llm, ok := l.llms[req.Model]
	if !ok {
		return nil, werror.Wrapf(gcode.ErrInvalidArgument, "unsupported model: %s", req.Model)
	}
	return llm.Chat(ctx, req)
}
