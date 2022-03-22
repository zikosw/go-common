package echo

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

func DumpBody(c echo.Context) ([]byte, error) {
	// Request
	reqBody := []byte{}
	if c.Request().Body != nil { // Read
		var err error
		reqBody, err = ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return nil, fmt.Errorf("read request - %w", err)
		}
	}
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(reqBody)) // Reset
	return reqBody, nil
}

func SetRequestBody(c echo.Context, body []byte) {
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
}

// Bind will dump/copy request body, bind and then set body back to the request
// so you can bind the request multiple times
func Bind(c echo.Context, i interface{}) error {
	body, err := DumpBody(c)
	if err != nil {
		return err
	}
	bindErr := c.Bind(i)

	SetRequestBody(c, body)
	return bindErr
}
