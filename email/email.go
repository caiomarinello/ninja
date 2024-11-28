package email

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)



func SendSuccessEmail(emailTo string) {

	godotenv.Load()

	// Read the HTML template (which now includes the inlined CSS)
	htmlTemplate, err := os.ReadFile("./email/email_template.html")
	if err != nil {
		log.Fatalf("Failed to read HTML template: %v", err)
	}
	
	m := gomail.NewMessage()
	m.SetHeader("From", "caiomarinello1@gmail.com")
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", "Your Purchase Was Successful!!")
	m.SetBody("text/html", string(htmlTemplate))
	
	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL_LOGIN"), os.Getenv("EMAIL_PASSWORD"))
	
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}