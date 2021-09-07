package posts

import (
	"fmt"
	"strconv"
	"time"

	"github.com/xeonx/timeago"
)

// IPostStorage is a generic poster interface
type IPostStorage interface {
	GetPosts() []Post
	SavePost(p Post)
}

// Post is a representation of a Guestbook entry
type Post struct {
	ID      string
	Created time.Time
	Name    string `form:"name" json:"name"`
	Message string `form:"message" json:"message"`
	Terms   string `form:"terms" json:"terms"`
}

// NewPost is a constructor function which sets default values
func NewPost() Post {
	return Post{
		ID:      fmt.Sprintf("%s%d", "post", time.Now().Unix()),
		Created: time.Now(),
	}
}

// FromMap creates a new Post instance from a map of strings. This is typically
// what is returned from external stores like Redis and is used by the
// PostStorage struct.
func FromMap(m map[string]string) Post {
	i, err := strconv.ParseInt(m["created"], 10, 64)
	if err != nil {
		// @TODO
	}

	t := time.Unix(i, 0)

	return Post{
		ID:      m["id"],
		Created: t,
		Name:    m["name"],
		Message: m["message"],
		Terms:   m["terms"],
	}
}

// GetID returns the ID of the post
func (p Post) GetID() string {
	return p.ID
}

func (p Post) GetTimeSince() string {
	return timeago.English.Format(p.Created)
}

// ToMap converts the Post instance to a map of strings. This is usefull for
// persisting data to external stores lik Redis and is used by the PostStorage
// struct.
func (p Post) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":      p.ID,
		"created": p.Created.Unix(),
		"name":    p.Name,
		"message": p.Message,
		"terms":   p.Terms,
	}
}
