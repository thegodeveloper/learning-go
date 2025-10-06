package errors

import (
	"errors"
	"fmt"
)

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
	Err     error
}

func (se StatusErr) Error() string {
	return se.Message
}

func (se StatusErr) Unwrap() error {
	return se.Err
}

func UseLoginAndGetData() {
	_, err := LoginAndGetData("bill", "xyz", "mozart")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
			Err:     err,
		}
	}
	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", file),
			Err:     err,
		}
	}
	return data, nil
}

func login(id, password string) error {
	fmt.Println("id", id)
	fmt.Println("Password:", password)
	return errors.New("error authenticating the user")
}

func getData(file string) ([]byte, error) {
	fmt.Println("File:", file)
	return []byte("data"), nil
}
