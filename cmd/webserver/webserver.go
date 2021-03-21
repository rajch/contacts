package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rajch/contacts/pkg/contact"

	// "github.com/rajch/contacts/pkg/filerepo"
	"github.com/rajch/contacts/pkg/gormrepo"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "All is well")
	})

	r.GET("/contacts", func(c *gin.Context) {
		g, err := getrepo()
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		defer g.Close()

		allcontacts, err := g.List()
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		c.JSON(200, allcontacts)

	})

	r.GET("/contacts/:id", func(c *gin.Context) {
		idparamvalue := c.Param("id")
		idparam, err := strconv.Atoi(idparamvalue)
		if err != nil {
			c.AbortWithError(500, err)
		}

		idparamuint := uint(idparam)

		g, err := getrepo()
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		defer g.Close()

		contact, err := g.Get(idparamuint)
		if err != nil {
			c.JSON(404, nil)
			return
		}

		c.JSON(200, contact)

	})

	r.POST("/contacts", func(c *gin.Context) {
		var newrecord contact.Contact

		err := c.ShouldBindJSON(&newrecord)
		if err != nil {
			c.AbortWithError(500, err)
		}

		g, err := getrepo()
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		defer g.Close()

		newcontact, err := g.New(&newrecord)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		c.JSON(200, newcontact)

	})

	r.Run("localhost:8080")
	// r.RunTLS("localhost:8080", "ws.crt", "ws.key")
}

func getrepo() (contact.Repository, error) {
	//return filerepo.New("testdb.db.json")
	return gormrepo.New("testdb.db")
}
