package posts

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

var rdb *redis.Client
var ctx context.Context

var redisHost = flag.String("redishost", "192.168.99.102:6379", "Redis address")
var redisPass = flag.String("redispass", "", "Redis password")
var redisDB = flag.Int("redisdb", 0, "Redis database")

func setup() {
	rdb = redis.NewClient(&redis.Options{
		Addr: *redisHost, Password: *redisPass, // no password set
		DB: *redisDB, // use default DB
	})

	ctx = context.TODO()

	rdb.FlushDB(ctx)
	rdb.HSet(ctx, "post1", map[string]interface{}{
		"name":    "foo",
		"message": "first!!!",
		"terms":   "true",
	})

	rdb.HSet(ctx, "post2", map[string]interface{}{
		"name":    "bar",
		"message": "hello w0rld",
		"terms":   "true",
	})
}

func shutdown() {
	rdb.FlushDB(ctx)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

func TestPostStorageGetPosts(t *testing.T) {
	ps := RedisStorage{Rdb: rdb}
	posts := ps.GetPosts()

	assert.Len(t, posts, 2)

	assert.Equal(t, posts[0].Name, "bar")
	assert.Equal(t, posts[0].Message, "hello w0rld")
	assert.Equal(t, posts[0].Terms, "true")

	assert.Equal(t, posts[1].Name, "foo")
	assert.Equal(t, posts[1].Message, "first!!!")
	assert.Equal(t, posts[1].Terms, "true")
}

func TestPostStorageNewPost(t *testing.T) {
	p := Post{Name: "Foo", Message: "Bar", Terms: ""}
	ps := RedisStorage{Rdb: rdb}

	ps.SavePost(p)
}
