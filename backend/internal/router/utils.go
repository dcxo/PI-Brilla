package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
)

// loginUserBody for route postLogin
type loginUserBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// activateUserBody for route postActivateUser
type activateUserBody struct {
	Token string `json:"token"`
}

func writeError(rw http.ResponseWriter, err string, status int) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(map[string]string{"error": err})
}

const emailBody = `
Confirme su correo electronico pulsando en el <a href="http://localhost/activate?token=%v">este enlace</a>
`

var emailFrom = os.Getenv("EMAIL_FROM")
var emailPass = os.Getenv("EMAIL_PASS")
var emailHost = os.Getenv("EMAIL_HOST")

func sendMail(to string, token string) (err error) {

	auth := smtp.PlainAuth("", emailFrom, emailPass, emailHost)

	toA := []string{to}

	headers := make(map[string]string)
	headers["From"] = emailFrom
	headers["To"] = to
	headers["Content-Type"] = "text/html"

	msg := ""
	for h, v := range headers {
		msg += fmt.Sprintf("%v: %v\t\n", h, v)
	}
	msg += fmt.Sprintf(emailBody, token)

	err = smtp.SendMail(emailHost+":25", auth, emailFrom, toA, []byte(msg))
	fmt.Println(err)
	return

}
