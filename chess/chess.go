package chess

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ChessPiece struct {
	PieceName string `json:"name"`
	Value     int    `json:"value"`
}

type CustomContext struct {
	echo.Context //This is equvivalent CustomContext extends echo.Context"
}

func (c *CustomContext) Foo() {
	println("foo")
}

func (c *CustomContext) Bar() {
	println("bar")
}

var pieces []*ChessPiece = []*ChessPiece{}

func Start() {
	e := echo.New()

	e.Use(func(hFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return hFunc(cc)
		}
	})
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	e.GET("/", homePage, middleWareFunc)
	e.GET("/pieces", getPieces)
	e.GET("/pieces/:name", getPiece)
	e.POST("/pieces", addPiece)
	e.PUT("/pieces/:name", updatePiece)
	e.DELETE("/pieces/:name", deletePiece)
	e.Logger.Print("Listening To Port 8080")
	e.Logger.Fatal(e.Start(":8080"))
}

func homePage(c echo.Context) error {
	c.Response().Before(func() {
		fmt.Println("Before Func")
	})
	c.Response().After(func() {
		fmt.Println("After Func")
	})
	fmt.Println("Inside Home page")
	return c.String(http.StatusOK, "Hello to Chess world")
}
func getPiece(c echo.Context) error {
	_, foundPiece := getPieceByName(c.Param("name"))
	if foundPiece == nil {
		return c.String(http.StatusNotAcceptable, "Piece not found")
	}
	return c.JSON(http.StatusAccepted, *foundPiece)
}

func getPieces(c echo.Context) error {
	return c.JSON(http.StatusOK, pieces)
}

func updatePiece(c echo.Context) error {
	type Rating struct {
		Value int `json:"value"`
	}
	var newValue Rating
	if err := c.Bind(&newValue); err != nil {
		return err
	}
	_, matchingPiece := getPieceByName(c.Param("name"))
	if matchingPiece == nil {
		return c.String(http.StatusNotAcceptable, "Piece not found")
	}
	matchingPiece.Value = newValue.Value
	return c.JSON(http.StatusAccepted, *matchingPiece)
}

func middleWareFunc(hFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Inside middleware")
		cc := &CustomContext{c}
		return hFunc(cc)
	}
}

func deletePiece(c echo.Context) error {
	foundIndex, foundPiece := getPieceByName(c.Param("name"))
	if foundPiece == nil {
		return c.String(http.StatusNotAcceptable, "Piece not found")
	}
	pieces = append(pieces[:foundIndex], pieces[foundIndex+1:]...)
	return c.JSON(http.StatusOK, pieces)
}

func getPieceByName(name string) (int, *ChessPiece) {
	for index, piece := range pieces {
		if piece.PieceName == name {
			return index, piece
		}
	}
	return -1, nil
}

func addPiece(c echo.Context) error {
	var newPiece ChessPiece
	if err := c.Bind(&newPiece); err != nil {
		return err
	}
	for _, piece := range pieces {
		if piece.PieceName == newPiece.PieceName {
			return c.String(http.StatusNotAcceptable, "Piece already present")
		}
	}
	pieces = append(pieces, &newPiece)
	return c.JSON(http.StatusOK, newPiece)
}
