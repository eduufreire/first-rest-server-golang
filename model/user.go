package model

import (
	"fmt"
	"time"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Birthday string `json:"birthday"`
}

func ToUserEntity(name string, birthday string) (*User, error) {
	if len(name) < 3 {
		return nil, fmt.Errorf("invalid name")
	}

	birthdayParsedTime, err := time.Parse(time.DateOnly, birthday)
	if err != nil {
		fmt.Println(err)
	}

	now := time.Now()
	if birthdayParsedTime.Compare(now) > 0 {
		return nil, fmt.Errorf("ruim ruim")
	}

	return &User{
		ID: 0,
		Name: name,
		Age: now.Year() - birthdayParsedTime.Year(),
		Birthday: birthday,
	}, nil
}
