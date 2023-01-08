package main

// package handler

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-playground/validator"
	"github.com/kevincobain2000/go-gibberish/gibberish"
	"github.com/knuppe/vader"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := Echo()
	e.Logger.Fatal(e.Start("localhost:3000"))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	e := Echo()
	e.ServeHTTP(w, r)
}

func Echo() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = HTTPErrorHandler

	h := NewGibberishHandler()
	e.GET("/", h.GibberishHandler)
	return e
}

type GibberishHandler struct {
	vader *vader.Vader
}

func NewGibberishHandler() GibberishHandler {
	v, err := vader.NewVader("./lexicons/en.zip")
	if err != nil {
		panic(err)
	}
	return GibberishHandler{
		vader: v,
	}
}

// GibberishRequest for the /
type GibberishRequest struct {
	Query string `json:"q" query:"q" form:"q" validate:"required" message:"No value for the query param. q is required"`
}

// GibberishResponse
// returns success if server is healthy
type GibberishResponse struct {
	Gibberish *gibberish.Gibberish `json:"gibberish"`
	Sentiment struct {
		Compound float64 `json:"compound"`
		Positive float64 `json:"positive"`
		Negative float64 `json:"negative"`
		Neutral  float64 `json:"neutral"`
	} `json:"sentiment"`
}

func (h *GibberishHandler) GibberishHandler(c echo.Context) error {
	request := &GibberishRequest{}
	err := c.Bind(request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// validate request
	msgs, err := ValidateRequest(request)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, msgs)
	}

	sentimentScores := h.vader.PolarityScores(request.Query)
	response := &GibberishResponse{
		Gibberish: gibberish.NewGibberish().Detect(request.Query),
	}

	response.Sentiment.Compound = sentimentScores.Compound
	response.Sentiment.Positive = sentimentScores.Positive
	response.Sentiment.Negative = sentimentScores.Negative
	response.Sentiment.Neutral = sentimentScores.Neutral

	return c.JSON(http.StatusOK, response)
}

func ValidateRequest[T any](request T) (map[string]string, error) {
	errs := validator.New().Struct(request)
	msgs := make(map[string]string)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			field, _ := reflect.TypeOf(request).Elem().FieldByName(err.Field())
			queryTag := getStructTag(field, "query")
			message := getStructTag(field, "message")
			msgs[queryTag] = message
		}
		return msgs, errs
	}
	return nil, nil
}

// getStructTag returns the value of the tag with the given name
func getStructTag(f reflect.StructField, tagName string) string {
	return string(f.Tag.Get(tagName))
}

// HTTPErrorResponse is the response for HTTP errors
type HTTPErrorResponse struct {
	Error interface{} `json:"error"`
}

// HTTPErrorHandler handles HTTP errors for entire application
func HTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var message interface{}
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message
	} else {
		message = err.Error()
	}

	if err = c.JSON(code, &HTTPErrorResponse{Error: message}); err != nil {
		fmt.Print("error: ", err)
	}
}
