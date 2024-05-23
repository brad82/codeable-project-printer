package codeable

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
)

type LoginRequest struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	TimezoneOffset int    `json:"timezone_offset"`
}

func (l *LoginRequest) calculateTimezoneOffset() {
	t := time.Now()
	_, offset := t.Zone()

	l.TimezoneOffset = -(offset / 60)
}

func (l *LoginRequest) IsValid() bool {
	return len(l.Email) > 0 && len(l.Password) > 0
}

func (c *ProjectClient) Login(loginRequest LoginRequest) error {
	loginRequest.calculateTimezoneOffset()

	body, err := json.Marshal(loginRequest)
	if err != nil {
		return err
	}
	req, _ := http.NewRequest("POST", CDBL_API_URL+"/users/login", bytes.NewBuffer(body))

	withJsonHeaders(req)

	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		return errors.New("unexpected status code received: " + strconv.Itoa(resp.StatusCode))
	}

	token := resp.Header.Get("auth-token")
	if len(token) == 0 {
		return errors.New("expected auth-token in login response, received null")
	}

	c.authToken = token

	return nil
}
