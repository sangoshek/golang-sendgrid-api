package emailservice

import (
	"github.com/joho/godotenv"
	"fmt"
	"log"
	"os"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func Send(subject string, nameTo string, mailTo string, plainMailBody string, mailBody string) bool {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	api_key := os.Getenv("SENDGRID_API_KEY")
	sender := os.Getenv("SENDGRID_SENDER_NAME")
	sender_email := os.Getenv("SENDGRID_SENDER_EMAIL")	
	from := mail.NewEmail(sender, sender_email)
	to := mail.NewEmail(nameTo, mailTo)
	plainTextContent := plainMailBody
	htmlContent := mailBody
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(api_key)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return false
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
		return true
	}
}