package server

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func generateUID() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

type EmailPayload struct {
	From     EmailAddress   `json:"from"`
	To       []EmailAddress `json:"to"`
	Subject  string         `json:"subject"`
	HTML     string         `json:"HTML"`
	Category string         `json:"category"`
}

type EmailAddress struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}

func sendConfirmationEmail(uid string) error {
	confirmationLink := fmt.Sprintf("http://localhost:3000/validate?uid=%s", uid)
	body := `<html>
				<head>
					<link rel="preconnect" href="https://fonts.googleapis.com">
					<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
					<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family=Dosis:wght@100;400;500;600;700&display=swap"/>
					<style>
						@import url('https://fonts.googleapis.com/css2?family=Cardo&family=Roboto&display=swap');
						@import url('https://fonts.googleapis.com/css2?family=Dosis:wght@100;400;500;600;700&display=swap');
					</style>
				</head>
				<body>
					<div style="width: 1000px; padding: 20px; border-radius: 20px; text-align:center; background-color: #ebc61c;">
						<h1 style="margin-bottom: 0;font-family:Dosis; font-style: normal;font-size:36pt;"> Welcome to Altitude Test Task! </h1>
						<p style="margin-left:50px; margin-right:50px;margin-top:20px; font-family:Dosis; font-style: normal;font-size:20pt;"> Please click the button below to verify your account </p>
						<button style="font-family:Roboto, sans-serif; border-radius:15px; margin-bottom:20px; letter-spacing:0.25em; font-size: 16pt; background-color:#212121; color: white; width: 300px; padding: 10px;">
							<a href="` + confirmationLink + `" style="text-decoration: none; color:white; display: inline-block; width:100%;"> VERIFY </a>
						</button>
					</div>
				</body>
			</html>`

	method := "POST"
	payload := EmailPayload{
		From: EmailAddress{
			Email: "hello@demomailtrap.com",
			Name:  "Mailtrap Test",
		},
		To: []EmailAddress{
			{
				Email: "lazzarmilanovic@gmail.com",
			},
		},
		Subject:  "Confirm your account",
		HTML:     body,
		Category: "Integration Test",
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error marshalling payload: %v", err)
	}

	requestBody := strings.NewReader(string(jsonPayload))

	client := &http.Client{}
	req, err := http.NewRequest(method, MAIL_URL, requestBody)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Authorization", "Bearer "+MAIL_TOKEN)
	req.Header.Add("Content-Type", "application/json")

	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}