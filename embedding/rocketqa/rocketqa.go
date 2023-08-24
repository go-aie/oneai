package rocketqa

import (
	"context"

	"github.com/go-aie/oneai/embedding/api"
	"github.com/go-aie/rocketqa"
	"github.com/go-aie/xslices"
)

type Encoder struct {
	*rocketqa.DualEncoder
}

type Config struct {
	ModelPath  string `structool:"model_path"`
	ParamsPath string `structool:"params_path"`
	VocabFile  string `structool:"vocab_file"`
}

func New(cfg *Config) (*Encoder, error) {
	de, err := rocketqa.NewDualEncoder(&rocketqa.DualEncoderConfig{
		ModelPath:         cfg.ModelPath,
		ParamsPath:        cfg.ParamsPath,
		VocabFile:         cfg.VocabFile,
		DoLowerCase:       true,
		QueryMaxSeqLength: 32,
		ParaMaxSeqLength:  384,
		ForCN:             true,
		//MaxConcurrency:    maxConcurrency,
	})
	if err != nil {
		return nil, err
	}

	return &Encoder{DualEncoder: de}, nil
}

func (e *Encoder) Encode(ctx context.Context, req *api.Request) (*api.Response, error) {
	resp := &api.Response{
		Model:  req.Model,
		Object: "list",
	}

	if req.Model == "rocketqa-query" {
		vectors := e.DualEncoder.EncodeQuery(req.Input)
		for i, vector := range vectors {
			resp.Data = append(resp.Data, &api.Data{
				Index:     i,
				Object:    "embedding",
				Embedding: vector.ToFloat64(),
			})
		}
		return resp, nil
	}

	// Model == "rocketqa-document"

	titles := xslices.Repeat([]string{""}, len(req.Input))
	vectors, err := e.DualEncoder.EncodePara(req.Input, titles)
	if err != nil {
		return nil, err
	}
	for i, vector := range vectors {
		resp.Data = append(resp.Data, &api.Data{
			Index:     i,
			Object:    "embedding",
			Embedding: vector.ToFloat64(),
		})
	}

	return resp, nil
}
