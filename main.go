package main

import (
	"github.com/gin-gonic/gin"
	"main/src/emailservice"
)

func main() {
	r := gin.Default()
	r.GET("/send", func(c *gin.Context) {

		subject := "Hello Send Grid!"
		nameTo	:= "Sango"
		mailTo := "sangoshek@gmail.com"
		plainMailBody := "This is a testing email."
		mailBody := "This is a testing <strong>email</strong>."
		
		email_send := emailservice.Send(subject, nameTo, mailTo, plainMailBody, mailBody) 
		if email_send == true {
			c.JSON(200, gin.H{
				"message": "sent",
			})
		}else{
			c.JSON(400, gin.H{
				"message": "fail",
			})
		}
	})
	r.Run(":9000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}