package greeter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juanpablopizarro/tudai-rest/internal/config"
)

// HTTPGreeter ...
type HTTPGreeter interface {
	Run()
}

type httpService struct {
	r *gin.Engine
}

// NewHTTPGreeter returns a new instance of HTTPGreeter
func NewHTTPGreeter(g Greeter, conf *config.Config) HTTPGreeter {
	r := gin.Default()
	group := r.Group(conf.Greeter.Version)
	group.GET("/sayhello", sayHello(g))
	return httpService{r}
}

func sayHello(g Greeter) gin.HandlerFunc {
	return func(c *gin.Context) {
		s := c.Query("name")
		c.JSON(http.StatusOK, gin.H{
			"greeting": g.SayHello(s),
		})
	}
}

func (s httpService) Run() {
	s.r.Run()
}
