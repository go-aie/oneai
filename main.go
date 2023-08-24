package main

import (
	_ "github.com/go-aie/oneai/embedding"
	_ "github.com/go-aie/oneai/llm"
	_ "github.com/go-aie/oneai/vectorstore"
)

func main() {
	Main()
}
