package main

import (
	"monitor/slices"

	"git.fuyu.moe/Fuyu/router"
)

func middlewareCORS(allowOrigin []string) router.Middleware {
	return func(next router.Handle) router.Handle {
		return func(c *router.Context) error {
			if len(allowOrigin) == 0 {
				next(c)
			}

			origin, ok := slices.First(allowOrigin, func(origin string) bool {
				return origin == c.Request.Host
			})
			if !ok {
				origin = allowOrigin[0]
			}

			c.Response.Header().Set(`Access-Control-Allow-Origin`, origin)
			c.Response.Header().Set(`Access-Control-Allow-Methods`, `DELETE, PUT, PATCH, POST, GET, OPTIONS`)
			c.Response.Header().Set(`Vary`, `Origin`)

			return next(c)
		}
	}
}
