package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	goahttp "goa.design/goa/v3/http"

	genhamsters "goa-demo/gen/hamster"
	genhttp "goa-demo/gen/http/hamster/server"
)

type HamsterService struct{}

func (m *HamsterService) List(_ context.Context) (res []*genhamsters.Hamster, err error) {
	hs := []*genhamsters.Hamster{
		{
			ID:     uuid.New().String(),
			Name:   "Fluffy",
			Colors: []string{"white", "black"},
		},
		{
			ID:     uuid.New().String(),
			Name:   "Fuzzy",
			Colors: []string{"brown", "white"},
		},
		{
			ID:     uuid.New().String(),
			Name:   "Squeaky",
			Colors: []string{"gray", "white"},
		},
	}
	return hs, nil
}

func (m *HamsterService) Create(_ context.Context, p *genhamsters.HamsterPayload) (res *genhamsters.Hamster, err error) {
	newConcert := &genhamsters.Hamster{
		ID:     uuid.New().String(),
		Name:   *p.Name,
		Colors: p.Colors,
	}

	return newConcert, nil
}

// Get a single concert by ID.
func (m *HamsterService) Show(_ context.Context, p *genhamsters.ShowPayload) (res *genhamsters.Hamster, err error) {
	return nil, genhamsters.MakeNotFound(fmt.Errorf("concert not found: %s", p.HamsterID))
}

// main instantiates the service and starts the HTTP server.
func main() {
	svc := &HamsterService{}
	endpoints := genhamsters.NewEndpoints(svc)
	mux := goahttp.NewMuxer()
	requestDecoder := goahttp.RequestDecoder
	responseEncoder := goahttp.ResponseEncoder
	handler := genhttp.New(endpoints, mux, requestDecoder, responseEncoder, nil, nil)

	genhttp.Mount(mux, handler)

	port := "3000"
	server := &http.Server{Addr: ":" + port, Handler: mux}

	for _, mount := range handler.Mounts {
		log.Printf("%q mounted on %s %s", mount.Method, mount.Verb, mount.Pattern)
	}

	log.Printf("Starting selling hamsters :%s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
