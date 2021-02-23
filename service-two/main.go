package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

type ResponseData struct {
	Message string `json:"message"`
}

func RequestToServiceOne() (ResponseData, error) {

	// Define error for reuse
	var err error

	url := fmt.Sprintf("%s", os.Getenv("SERVICE_ONE_URL"))
	// url := fmt.Sprintf("%s", "http://localhost:4000")
	var resData ResponseData

	// Setup the client
	client := &http.Client{}

	// Setup the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("GET failed with %s\n", err)
		return resData, err
	}

	// Make request to pricing service
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("GET failed with %s\n", err)
		return resData, err
	}

	// Read the response
	err = json.NewDecoder(res.Body).Decode(&resData)
	if err != nil {
		fmt.Printf("Failed to parse response JSON %s\n", err)
		return resData, err
	}
	return resData, nil
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {

		// Get the message from service one
		resData, err := RequestToServiceOne()
		if err != nil {
			fmt.Println("There was a problem")
			c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(&fiber.Map{
			"message":     "Hello from service 2.",
			"service_one": resData,
		})
	})

	app.Listen(":4002")
}
