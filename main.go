package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

func HelloWorldHandler(ctx *fasthttp.RequestCtx) {
	_, err := fmt.Fprintln(ctx, "Hello world!")
	if err != nil {
		log.Println(err.Error())
	}
}

func main() {
	err := fasthttp.ListenAndServe(":8001", HelloWorldHandler)
	if err != nil {
		log.Fatal(err.Error())
	}
}
