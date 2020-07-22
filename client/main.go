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
		fmt.Println(err)
		ctx.Response.SetStatusCode(fasthttp.StatusBadRequest)
		// this can be more descriptive
		ctx.Response.SetBody([]byte(fmt.Sprint("Пук серк")))
	} else {
		ctx.Response.SetStatusCode(fasthttp.StatusAccepted)
	}
}

func main() {
	flag.Parse()

	err := initProducer()
	if err != nil {
		log.Fatal(err)
	}

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
