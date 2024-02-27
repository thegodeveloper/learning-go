package werrorold

import "fmt"

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func errorsAndValues(show bool) {
	if show {
		println("\n-- Errors and Values")

		data, err := LoginAndGetData("uid", "pwd", "file")
		if err != nil {
			fmt.Println("error opening a file")
			return
		}
		fmt.Println(data)
	}
}

func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status:  InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for user %s", uid),
		}
	}
	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status:  NotFound,
			Message: fmt.Sprintf("file %s not found", file),
		}
	}
	return data, nil
}

func login(uid, pwd string) error {
	fmt.Printf("uid: %s, pwd: %s", uid, pwd)
	return nil
}

func getData(file string) ([]byte, error) {
	fmt.Printf("file: %s", file)
	return nil, nil
}
