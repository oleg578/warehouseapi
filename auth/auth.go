package auth

import (
	"net/http"
	"strings"

	"github.com/gomodule/redigo/redis"
	"whapi/config"
)

// Auth authorization test
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.ToUpper(r.Method) == "OPTIONS" {
			next.ServeHTTP(w, r)
		}
		user, pass, ok := r.BasicAuth()
		if !ok {
			NotAuthHandler(w, r)
			return
		}
		//test user valid
		if IsUserValid(user, pass) == false {
			NotAuthHandler(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// NotAuthHandler Not Auth
func NotAuthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Basic")
	w.WriteHeader(401)
	return
}

// IsUserValid Is User Valid ?
func IsUserValid(u, p string) bool {
	c, err := redis.Dial("tcp", config.RedisDSN)
	if err != nil {
		return false
	}
	defer c.Close()
	pass, err := redis.String(c.Do("GET", u))
	if err != nil {
		return false
	}
	return pass == p
}
