package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type User struct {
	Id			string
	AddressId	string
}

const VerifyMessage = "verified"

func main() {

	s := NewServer()

	s.HandleFunc("GET", "/", func(c *Context) {
		c.RenderTemplate("/src/gothic/web_000/public/index.html", map[string]interface{}{"time": time.Now()})
	})

	s.HandleFunc("GET", "/about", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "about")
	})

	s.HandleFunc("GET", "/users/:id", func(c *Context) {
		u := User{Id: c.Params["id"].(string)}
		c.RenderXml(u)
	})

	s.HandleFunc("GET", "/users/:user_id/addresses/:address_id", func(c *Context) {
		u := User{c.Params["user_id"].(string), c.Params["address_id"].(string)}
		c.RenderJson(u)
	})

	s.HandleFunc("POST", "/users", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, c.Params)
	})

	s.HandleFunc("POST", "/users/:user_id/addresses", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, c.Params)
	})

	s.HandleFunc("GET", "/login", func(c *Context) {
		c.RenderTemplate("src/gothic/web_000/public/login.html", map[string]interface{}{"message": "로그인이 필요합니다."})
	})

	s.HandleFunc("POST", "/login", func(c *Context) {
		if CheckLogin(c.Params["username"].(string), c.Params["password"].(string)) {
			http.SetCookie(c.ResponseWriter, &http.Cookie{
				Name:		"X_AUTH",
				Value:		Sign(VerifyMessage),
				Path:		"/",
			})
			c.Redirect("/")
		}
		c.RenderTemplate("src/gothic/web_000/public/login.html", map[string]interface{}{"message": "id 또는 password가 일치하지 않습니다."})

	})

	s.Use(AuthHandler)

	s.Run(":8000")

}

func CheckLogin(username, password string) bool {
	const (
		USERNAME = "gothic"
		PASSWORD = "dl159147!@"
	)

	return username == USERNAME && password == PASSWORD
}

func Sign(message string) string {
	secretkey := []byte("golang-book-secret-key2")
	if len(secretkey) == 0 {
		return ""
	}
	mac := hmac.New(sha1.New, secretkey)
	io.WriteString(mac, message)
	return hex.EncodeToString(mac.Sum(nil))
}

func AuthHandler(next HandlerFunc) HandlerFunc {
	ignore := []string{"/login", "src/gothic/web_000/public/index.html"}
	return func(c *Context) {
		for _, s := range ignore {
			if strings.HasPrefix(c.Request.URL.Path, s) {
				next(c)
				return
			}

			if v, err := c.Request.Cookie("X_AUTH"); err == http.ErrNoCookie {
				c.Redirect("/login")
				return
			} else if err != nil {
				c.RenderErr(http.StatusInternalServerError, err)
				return
			} else if Verify(VerifyMessage, v.Value) {
				next(c)
				return
			}

			c.Redirect("/login")

		}
	}
}

func Verify(message, sig string) bool {
	return hmac.Equal([]byte(sig), []byte(Sign(message)))
}