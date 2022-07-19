package main

import (
	"net/http"
	"time"
	"math/rand"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type Greet struct{
	Name string `json:"name"`
	Time string `json:"time"`
	SendQuote string `json:"sendquote"`
}

type album struct {
    ID     int 
    Quote  string  
}

var albums = []album{
    {ID: 1, Quote: "You define your own life. Don't let other people write your script."},
    {ID: 2, Quote: "You don't always need a plan. Sometimes you just need to breathe, trust, let go and see what happens."},
    {ID: 3, Quote: "Nothing is impossible. The word itself says 'I'm possible!'"},
	{ID: 4, Quote: "When you have a dream, you've got to grab it and never let go."},
	{ID: 5, Quote:  "Life is like riding a bicycle. To keep your balance, you must keep moving."},
}
func Greetins(c *gin.Context){
    
	var Wish string
	var sendGreet Greet
	t:=time.Now()

	if t.Hour() >= 21 && t.Hour() < 5 {
      Wish="Good Nigth!!!"
	}else if t.Hour() >= 5 && t.Hour() < 12{
		Wish="Good Morning!!!"
	}else if t.Hour() >= 12 && t.Hour() < 18{
		Wish="Good Afternoon!!!"
	}else if t.Hour() >= 18 && t.Hour() < 21{
		Wish="Good Evening!!!"
	}

	sendGreet.Name=Wish
	sendGreet.Time=t.Format(time.Kitchen)
	sendGreet.SendQuote=albums[rand.Intn(4)].Quote 
	
	c.JSON(http.StatusOK,sendGreet)

}
func main(){

	router:=gin.Default()
	router.Use(cors.Default())
	router.GET("/wish",Greetins)
	router.Run("localhost:8080")
}