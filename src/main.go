package main

import (
	"context"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var echoLambda *echoadapter.EchoLambda

func init() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Response{From: "mainpage", Message: time.Now().Format(time.UnixDate)})
	})

	e.POST("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Response{From: "posttest", Message: time.Now().Format(time.UnixDate)})
	})

	echoLambda = echoadapter.New(e)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}

type Response struct {
	From    string `json:"from"`
	Message string `json:"message"`
}
