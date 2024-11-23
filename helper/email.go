package helper

import "gopkg.in/gomail.v2"

func SendPasswordResetEmail(to string, resetURL string) {
	e := gomail.NewMessage()
	e.SetHeader("From", "your-email@example.com")
	e.SetHeader("To", to)
	e.SetHeader("Subject", "Password Reset Request")
	e.SetBody("text/html", "Click <a href='"+resetURL+"'>here</a> to reset your password.")

	d := gomail.NewDialer("smtp.example.com", 587, "your-email@example.com", "your-password")
	if err := d.DialAndSend(e); err != nil {
		panic(err)
	}
}