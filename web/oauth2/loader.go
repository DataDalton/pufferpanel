/*
 Copyright 2020 Padduck, LLC
  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at
  	http://www.apache.org/licenses/LICENSE-2.0
  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

package oauth2

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pufferpanel/pufferpanel/v3/oauth2"
	"github.com/spf13/cast"
	"net/http"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	rg.Use(setHeaders, recovery)
	registerTokens(rg)
	registerInfo(rg)
}

func setHeaders(c *gin.Context) {
	c.Header("Cache-Control", "no-store")
	c.Header("Pragma", "no-cache")
}

func recovery(c *gin.Context) {
	//override the recovery route, as we need to change the type returned
	defer func() {
		if err := recover(); err != nil {
			var msg string
			if e, ok := err.(error); ok {
				msg = e.Error()
			} else if e, ok := cast.ToStringE(msg); ok == nil {
				msg = e
			} else {
				msg = fmt.Sprintf("%v", err)
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, &oauth2.ErrorResponse{
				Error:            "internal_error",
				ErrorDescription: msg,
			})
		}
	}()
	c.Next()
}
