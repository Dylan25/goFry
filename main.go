package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	fry "github.com/goFry/imagefryer"
)

var router *gin.Engine

func main() {
	//set router to default gin
	router = gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./gofry/build", true)))

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		api.POST("/", FryImage)
	}

	//run server
	router.Run(":3000")

}

func FryImage(c *gin.Context) {
	fmt.Println("frying")
	//fmt.Printf("request is %v", c.JSON)
	timesFry, _ := strconv.Atoi(c.Request.FormValue("timesFry"))
	fmt.Printf("timesFry is %v", timesFry)
	//fmt.Printf("got %v", c.Request)
	file, header, err := c.Request.FormFile("image")
	//filetofry, err := header.Open()
	if err != nil {
		fmt.Printf("error opening file.FileHeader in FryImage\n")
	}
	filename := header.Filename
	fmt.Println(filename)
	frydimage, imageType := fry.Fry(&file, timesFry)
	//image_byte_buffer_to_send := make([]byte, header.Size)
	image_byte_buffer_to_send := new(bytes.Buffer)

	if imageType == "png" {
		png.Encode(image_byte_buffer_to_send, *frydimage)
	} else if imageType == "jpeg" {
		jpeg.Encode(image_byte_buffer_to_send, *frydimage, nil)
	} else {
		fmt.Println("ERROR: unrecognized file format")
	}

	// out, err := os.Create("TMP" + filename)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer out.Close()

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"fryedImage": image_byte_buffer_to_send.Bytes(),
		"message":    "ping",
	})

	// c.Header("Content-Type", "application/json")
	// c.JSON(http.StatusOK, gin.H{
	// 	"fryedImage":
	// 	"message": "ping",
	// })
}
