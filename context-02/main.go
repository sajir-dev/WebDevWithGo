package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = context.WithValue(ctx, "UserID", 777)
	ctx = context.WithValue(ctx, "fname", "Bond")

	results := dBAccess(ctx)

	fmt.Fprintln(res, results)
}

func dBAccess(ctx context.Context) int {
	uid := ctx.Value("UserID").(int)
	return uid
}

func bar(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	log.Println(ctx)
	fmt.Fprintln(res, ctx)
}
