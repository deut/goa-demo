package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"goa-demo/gen/hamster"
	genhamsters "goa-demo/gen/hamster"
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
	mux := gin.Default()

	requestDecoder := func(ctx context.Context, c *gin.Context) (interface{}, error) {
		var req hamster.HamsterPayload
		if err := c.ShouldBindJSON(&req); err != nil {
			return nil, err
		}
		return &req, nil
	}

	responseEncoder := func(ctx context.Context, v interface{}) {
		c := ctx.Value("gin").(*gin.Context)
		c.JSON(http.StatusOK, v)
	}

	mux.POST("/hamsters", func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "gin", c)
		req, err := requestDecoder(ctx, c)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		res, err := endpoints.Create(ctx, req.(*genhamsters.HamsterPayload))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		responseEncoder(ctx, res)
	})

	mux.GET("/hamsters", func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), "gin", c)
		res, err := endpoints.List(ctx, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		responseEncoder(ctx, res)
	})

	mux.Run(":8080")
}
