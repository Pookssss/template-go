package controllers

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ListProjectEndpoint(c *fiber.Ctx) error {
	// Customer ID
	customerKey := c.Params("customerKey")
	fmt.Println("customerKey: ", customerKey)
	// Customer secret
	customerSecret := c.Params("customerSecret")
	fmt.Println("customerSecret: ", customerSecret)
	// Concatenate customer key and customer secret and use base64 to encode the concatenated string
	plainCredentials := customerKey + ":" + customerSecret
	base64Credentials := base64.StdEncoding.EncodeToString([]byte(plainCredentials))

	url := "https://api.agora.io/dev/v1/projects"
	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	// Add Authorization header
	req.Header.Add("Authorization", "Basic "+base64Credentials)
	req.Header.Add("Content-Type", "application/json")

	// Send HTTP request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))

	data := map[string]string{"agora_token": string(body)}

	return c.JSON(fiber.Map{"status": 200, "data": data, "message": "success"})
}

func RuleListByAppIDEndpoint(c *fiber.Ctx) error {

	// App ID
	appid := c.Params("appId")
	// Customer ID
	customerKey := c.Params("customerKey")
	fmt.Println("customerKey: ", customerKey)
	// Customer secret
	customerSecret := c.Params("customerSecret")
	fmt.Println("customerSecret: ", customerSecret)

	// Concatenate customer key and customer secret and use base64 to encode the concatenated string
	plainCredentials := customerKey + ":" + customerSecret
	base64Credentials := base64.StdEncoding.EncodeToString([]byte(plainCredentials))

	url := "https://api.agora.io/dev/v1/kicking-rule?appid=" + appid
	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	// Add Authorization header
	req.Header.Add("Authorization", "Basic "+base64Credentials)
	req.Header.Add("Content-Type", "application/json")

	// Send HTTP request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))

	data := map[string]string{"agora_token": string(body)}

	return c.JSON(fiber.Map{"status": 200, "data": data, "message": "success"})
}
