# OneAI

[![Go Reference](https://pkg.go.dev/badge/go-aie/oneai/vulndb.svg)][1]

Unified REST APIs for LLM/Embedding models and Vector stores.

Currently supported models and vector stores:

- LLM
    - [x] ChatGLM
        + [ChatGLM-6B][2]
        + [ChatGLM2-6B][3]
    - [ ] LLaMA
- Embedding
    - [x] [RocketQA][4]
- Vector Store
    - [x] [Memory](vectorstore/memory)
    - [ ] [Faiss][5]
    - [ ] [Milvus][6]


## Installation

```bash
go install github.com/go-aie/oneai@latest
```

## Documentation

Check out the [documentation][1].


## License

[MIT](LICENSE)


[1]: https://pkg.go.dev/github.com/go-aie/oneai
[2]: https://github.com/THUDM/ChatGLM-6B
[3]: https://github.com/THUDM/ChatGLM2-6B
[4]: https://github.com/PaddlePaddle/RocketQA
[5]: https://github.com/facebookresearch/faiss
[6]: https://github.com/milvus-io/milvus
