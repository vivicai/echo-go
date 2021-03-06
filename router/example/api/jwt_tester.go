package api

import (
	"net/http"
	"strconv"
	"github.com/foxiswho/echo-go/module/log"
	"github.com/foxiswho/echo-go/module/cache"
	"github.com/foxiswho/echo-go/middleware/session"
	"github.com/foxiswho/echo-go/service/user_service/auth"
	userService "github.com/foxiswho/echo-go/service/example_service"
	"github.com/foxiswho/echo-go/router/base"
	"time"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func JwtTesterApiHandler(c *base.BaseContext) error {
	log.Debugf("JwtApiHandler")
	log.Debugf("JwtApiHandler")
	log.Debugf("JwtApiHandler")
	log.Debugf("JwtApiHandler")
	log.Debugf("JwtApiHandler")
	user := c.Get("_user").(*jwt.Token)
	fmt.Println("ClaimsClaimsClaimsClaims",user)
	fmt.Println("ClaimsClaimsClaimsClaims",user)
	fmt.Println("ClaimsClaimsClaimsClaims",user.Claims)
	//fmt.Println("name",name)
	idStr := c.QueryParam("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	u := &auth.User{}
	if err != nil {
		log.Debugf("Render Error: %v", err)
	} else {
		u = userService.GetUserById(id)
	}

	// 缓存测试
	value := -1
	if err == nil {
		cacheStore := cache.Default(c)
		if id == 1 {
			value = 0
			cacheStore.Set("userId", 1, 5*time.Minute)
		} else {
			if err := cacheStore.Get("userId", &value); err != nil {
				log.Debugf("cache userId get err:%v", err)
			}
		}
	}

	// Flash测试
	s := session.Default(c)
	fmt.Println("session.Default",s)
	s.AddFlash("0")
	s.AddFlash("1")
	s.AddFlash("10", "key1")
	s.AddFlash("20", "key2")
	s.AddFlash("21", "key2")
	//c.Response().Header().Del("Access-Control-Allow-Origin")
	//c.Response().Header().Add("Access-Control-Allow-Origin","*")
	request := c.Request()
	c.JSON(http.StatusOK, map[string]interface{}{
		"title":        "Api Index",
		"Admin":         u,
		"CacheValue":   value,
		"URL":          request.URL,
		"Scheme":       request.URL.Scheme,
		"Host":         request.Host,
		"UserAgent":    request.UserAgent(),
		"Method":       request.Method,
		"URI":          request.RequestURI,
		"RemoteAddr":   request.RemoteAddr,
		"Path":         request.URL.Path,
		"QueryString":  request.URL.RawQuery,
		"QueryParams":  request.URL.Query(),
		"HeaderKeys":   request.Header,
		"FlashDefault": s.Flashes(),
		"Flash1":       s.Flashes("key1"),
		"Flash2":       s.Flashes("key2"),
	})
	return  nil
}
