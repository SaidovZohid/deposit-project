package api

import (
	"context"
	"fmt"

	v1 "github.com/SaidovZohid/deposit-project/api/v1"
	"github.com/SaidovZohid/deposit-project/config"
	"github.com/SaidovZohid/deposit-project/storage"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Cfg  *config.Config
	Strg storage.StorageI
}

type something func(ctx context.Context, arg int)

func Handle(arg ...something) {
	for _, fn := range arg {
		fn(context.Background(), 10)
	}
}

func Print(ctx context.Context, arg int) {
	fmt.Println(arg)
}

func Add(ctx context.Context, arg int) {
	fmt.Println(arg + 10)
}

func New(h *Handler) *gin.Engine {
	engine := gin.Default()

	handlerV1 := v1.New(&v1.HandleV1{
		Cfg:  h.Cfg,
		Strg: h.Strg,
	})

	engine.GET("/user/:id", handlerV1.GetUserById)

	Handle(Print, Add)

	return engine
}
