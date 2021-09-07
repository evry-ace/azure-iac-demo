package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/Masterminds/sprig"
	"github.com/evry-ace/azure-iac-demo/app/src/posts"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var storageBackend = flag.String("storagebackend", "mongo", "Storage backend to use, redis or mongo")

var redisHost = flag.String("redishost", "192.168.99.103:6379", "Redis address")
var redisPass = flag.String("redispass", "", "Redis password")
var redisDB = flag.Int("redisdb", 0, "Redis database")

var mongoURI = flag.String("mongouri", "mongodb://root:example@192.168.99.103:27017", "MongoDB Connection String")
var mongoDatabase = flag.String("mongodatabase", "mongo", "MongoDB Database Nname")
var mongoCollection = flag.String("mongocollection", "posts", "MongoDB Collection Name")

func createRedisStorage() *posts.RedisStorage {
	rdb := redis.NewClient(&redis.Options{
		Addr:     *redisHost,
		Password: *redisPass,
		DB:       *redisDB,
	})

	return &posts.RedisStorage{Rdb: rdb}
}

func createMongoStorage() *posts.MongoStorage {
	ctx := context.Background()
	mongoClientOpts := options.Client().ApplyURI(*mongoURI)
	mongoClient, err := mongo.Connect(ctx, mongoClientOpts)
	if err != nil {
		panic(err)
	}
	//defer mongoClient.Disconnect(ctx)

	mongoDatabase := mongoClient.Database(*mongoDatabase)
	mongoPostsCollection := mongoDatabase.Collection(*mongoCollection)

	return &posts.MongoStorage{Coll: mongoPostsCollection}
}

func main() {
	router := gin.Default()
	router.SetFuncMap(sprig.HtmlFuncMap())
	router.HTMLRender = loadTemplates("./templates")
	router.Static("/images", "./images")
	// router.LoadHTMLGlob("templates/views/*")

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})

	var ps posts.IPostStorage

	if *storageBackend == "redis" {
		ps = createRedisStorage()
	} else if *storageBackend == "mongo" {
		ps = createMongoStorage()
	} else {
		panic("Unsupported storage backend")
	}

	router.GET("/", posts.GetPosts(ps))
	router.GET("/post", posts.GetNewPost)
	router.POST("/post", posts.PostNewPost(ps))

	host := ""
	port := "8080"

	// This is to prevent the macOS firewall from complaining!
	if gin.IsDebugging() {
		host = "localhost"
	}

	router.Run(fmt.Sprintf("%s:%s", host, port))
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFilesFuncs(filepath.Base(include), sprig.HtmlFuncMap(), files...)
	}
	return r
}
