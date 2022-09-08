package middleware

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware ...
/*func AuthMiddleware(app *firebase.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		idToken := c.Request.Header.Get("Authorization")
		_, err := auth.VerifyIDToken(c, app, idToken)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"error":  err.Error(),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}*/

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

//MaxSizeAllowed ... Avoid a large file from loading into memory
//If the file size is greater than 8MB dont allow it to even load into memory and waste our time.
func MaxSizeAllowed(n int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, n)
		buff, errRead := c.GetRawData()
		if errRead != nil {
			//c.JSON(http.StatusRequestEntityTooLarge,"too large")
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{
				"status":     http.StatusRequestEntityTooLarge,
				"upload_err": "too large: upload an image less than 8MB",
			})
			c.Abort()
			return
		}
		buf := bytes.NewBuffer(buff)
		c.Request.Body = ioutil.NopCloser(buf)
	}
}

// GinContextToContextMiddleware ...
func GinContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), userCtxKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// GinContextFromContext ...
func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(userCtxKey)
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}
