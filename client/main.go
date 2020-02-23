package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")

	addr     = flag.String("addr", ":8080", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

func trackHandler(ctx *fasthttp.RequestCtx) {

	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)

	err := track(ctx.PostBody())
	if err != nil {
		fmt.Sprint(err)
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		// this can be more descriptive, depends how much you want to expose to production
		ctx.Response.SetBody([]byte(fmt.Sprint("Пук серк")))
	} else {
		ctx.Response.SetStatusCode(fasthttp.StatusAccepted)
	}
}

func main() {
	flag.Parse()

	initProducer()

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/track":
			trackHandler(ctx)
		default:
			ctx.Error("Пук секр", fasthttp.StatusNotFound)
		}
	}

	h := requestHandler
	if *compress {
		h = fasthttp.CompressHandler(h)
	}

	if err := fasthttp.ListenAndServe(*addr, h); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
