package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostNewPost(ps IPostStorage) func(c *gin.Context) {
	return func(c *gin.Context) {
		//var p Post
		p := NewPost()
		c.Bind(&p)

		ps.SavePost(p)

		c.Redirect(http.StatusFound, "/")
	}
}

func GetPosts(ps IPostStorage) func(c *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "guestbook.html", gin.H{
			"title": "Guestbook",
			"posts": ps.GetPosts(),
		})
	}
}

func GetNewPost(c *gin.Context) {
	c.HTML(http.StatusOK, "post.html", gin.H{
		"title": "New Post",
	})
}
