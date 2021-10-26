```go
import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New() //新建一个实例
	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})//行内函数   
	e.GET("/users/:id", getUser)//行外函数
	
	e.Logger.Fatal(e.Start(":1323"))
}
```

```go
// e.GET("/users/:id", getUser) 
func getUser(c echo.Context) error {
// User ID from path `users/:id`
id := c.Param("id")
return c.String(http.StatusOK, id)
}
//Browse to http://localhost:1323/users/Joe and you should see ‘Joe’ on the page.
```
```go
//e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}
//Browse to http://localhost:1323/show?team=x-men&member=wolverine
//and you should see ‘team:x-men, member:wolverine’ on the page.
```
```go

```