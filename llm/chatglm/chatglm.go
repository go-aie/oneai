package chatglm

import (
	"context"
	"time"

	"github.com/RussellLuo/kun/pkg/werror"
	"github.com/RussellLuo/kun/pkg/werror/gcode"
	"github.com/go-aie/chatglm"

	"github.com/go-aie/oneai/llm/api"
)

type Config struct {
	ModelPath string `structool:"model_path"`
}

type LLM struct {
	*chatglm.ChatGLM
}

// New creates a ChatGLM model l. Remember to release the model by calling
// l.Delete() when not in use.
func New(cfg *Config) (*LLM, error) {
	return &LLM{
		ChatGLM: chatglm.New(cfg.ModelPath),
	}, nil
}

func (l *LLM) Chat(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	cnt := len(req.Messages)
	if cnt%2 != 1 {
		return nil, werror.Wrapf(gcode.ErrInvalidArgument, "length of messages expected to be odd")
	}

	var history []*chatglm.Turn
	for i := 0; i < cnt-1; i++ {
		msg := req.Messages[i]
		if i%2 == 0 {
			if msg.Role == "assistant" {
				return nil, werror.Wrapf(gcode.ErrInvalidArgument, "invalid messages")
			}
			history = append(history, &chatglm.Turn{Question: msg.Content})
		}
		if i%2 == 1 {
			if msg.Role == "user" {
				return nil, werror.Wrapf(gcode.ErrInvalidArgument, "invalid messages")
			}
			history[len(history)-1].Answer = msg.Content
		}
	}

	lastMsg := req.Messages[cnt-1]
	if lastMsg.Role != "user" {
		return nil, werror.Wrapf(gcode.ErrInvalidArgument, "invalid messages")
	}
	query := lastMsg.Content

	output := l.ChatGLM.Generate(chatglm.BuildPrompt(query, history))
	return &api.Response{
		ID:      "",
		Object:  "chat.completion",
		Created: int(time.Now().Unix()),
		Choices: []*api.Choice{
			{
				Index: 0,
				Message: &api.Message{
					Role:    "assistant",
					Content: output,
				},
				FinishReason: "stop",
			},
		},
	}, nil
}
