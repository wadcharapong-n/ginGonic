package main

import (
			"log"
			"github.com/gin-gonic/gin"
			"net/http"
			"fmt"
			"gopkg.in/mgo.v2"
			"gopkg.in/mgo.v2/bson"
		)

/*type Users struct {
	id      string `bson:"_id"`
	class 	string `bson: _class`
	username    string
	loginname string
	password  string
	email  string
	role  string
}

type ShortUrl struct {
	id      string `bson:"_id"`
	originalUrl    string
	shortUrl string
	clickCount  string
}*/

func main() {
	r := gin.Default()
	r.GET("/users", getAllusers) //get API
	r.GET("/ping", ping) //get API
	r.POST("/ping", pingPost)  //post API
	r.GET("/user/:name", func(c *gin.Context) {  //get API have parameter
		name := c.Param("name")
		fmt.Println(name)
		c.String(http.StatusOK, "Hello %s", name)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
			"message": "pong",
			"method": "get",
		})
}

func pingPost(c *gin.Context) {
	c.JSON(200, gin.H{
			"message": "pong",
			"method": "post",
		})
}

func getAllusers(c *gin.Context) {
	getUserDb()
	

	//users := getUserDb()
	//fmt.Println(users[0].username)

	//shortUrl := getShortUrlDb()
	//fmt.Println(shortUrl[0].shortUrl)
}

func getUserDb(){
	session, err := mgo.Dial("mongodb://localhost:27017/yourdb")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("yourdb").C("users")
	resp := []bson.M{}
	err = c.Find(nil).All(&resp)
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	fmt.Println(resp)

}

/*func getUserDb() []Users{
	session, err := mgo.Dial("mongodb://localhost:27017/yourdb")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("yourdb").C("users")
	users := []Users{}
	err = c.Find(nil).All(&users)
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	return users
}

/*func getShortUrlDb() []ShortUrl{
	session, err := mgo.Dial("mongodb://localhost:27017/yourdb")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("yourdb").C("shortUrl")
	shortUrl := []ShortUrl{}
	err = c.Find(nil).All(&shortUrl)
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	return shortUrl
}*/