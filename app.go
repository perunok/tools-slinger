package main

import (
	"context"
	"fmt"
	"slinger/types"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Get(request types.Request) types.Response {
	fmt.Printf("header %s body %s and url %s", request.Headers, request.Body, request.URL)
	headers := []types.KeyValue{
		{Key: "success", Value: "true"},
	}
	response := types.Response{
		StatusCode: 200,
		Body:       `{"status":"success","message":"User saved successfully"}`,
		Headers:    headers,
	}
	return response
}
