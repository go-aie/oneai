package vectorstore

import (
	"context"
	"fmt"

	"github.com/RussellLuo/appx"
	"github.com/RussellLuo/kun/pkg/appx/httpapp"
	"github.com/RussellLuo/kun/pkg/werror"
	"github.com/RussellLuo/kun/pkg/werror/gcode"
	"github.com/RussellLuo/structool"
	"github.com/go-chi/chi"

	"github.com/go-aie/oneai/auth"
	"github.com/go-aie/oneai/vectorstore/api"
	"github.com/go-aie/oneai/vectorstore/controller/http"
	"github.com/go-aie/oneai/vectorstore/memory"
)

func init() {
	appx.MustRegister(
		httpapp.New("vector_store", new(VectorStore)).
			MountOn("oneai", "/vector").
			Require("auth").App,
	)
}

type VectorStore struct {
	stores map[string]api.VectorStore
	router chi.Router
}

func (v *VectorStore) Router() chi.Router {
	return v.router
}

func (v *VectorStore) Init(ctx appx.Context) error {
	codec := structool.New()

	var configs map[string]map[string]interface{}
	if err := codec.Decode(ctx.Config(), &configs); err != nil {
		return err
	}

	v.stores = make(map[string]api.VectorStore)
	for vendor := range configs {
		switch vendor {
		case "memory":
			fmt.Println("init memory vector store")
			store := memory.New()
			v.stores[vendor] = store
		}
	}

	a := ctx.MustLoad("auth").(*auth.Auth)
	v.router = http.NewHTTPRouter(v, a.Codecs)

	return nil
}

func (v *VectorStore) Upsert(ctx context.Context, vendor string, documents []*api.Document) error {
	store, ok := v.stores[vendor]
	if !ok {
		return werror.Wrapf(gcode.ErrInvalidArgument, "unsupported vendor: %s", vendor)
	}
	return store.Upsert(ctx, vendor, documents)
}

func (v *VectorStore) Query(ctx context.Context, vendor string, vector []float64, topK int, minScore float64) (similarities []*api.Similarity, err error) {
	store, ok := v.stores[vendor]
	if !ok {
		return nil, werror.Wrapf(gcode.ErrInvalidArgument, "unsupported vendor: %s", vendor)
	}
	return store.Query(ctx, vendor, vector, topK, minScore)
}

func (v *VectorStore) Delete(ctx context.Context, vendor string, sourceIDs ...string) error {
	store, ok := v.stores[vendor]
	if !ok {
		return werror.Wrapf(gcode.ErrInvalidArgument, "unsupported vendor: %s", vendor)
	}
	return store.Delete(ctx, vendor, sourceIDs...)
}
