package entity

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username  string
	password  string // password cannot be seen
	Email     string
	Telephone string
	isLogin   bool // whether user has logged in
}

func EncodeJSON(username, password, email, telephone string) ([]byte, error) {
	user := User{username, password, email, telephone, false}
	result, err := json.Marshal(user)

	// Todo 存储result

	fmt.Println(result)

	return result, err
}

func DecodeJSON(data []byte) []string {
	var user User
	var result []string

	err := json.Unmarshal(data, &user)

	if err == nil {
		result = append(result, user.Username, user.Email, user.Telephone)
	}

	return result
}

/*func main() {
	data, err := EncodeJSON("Avril Lavigne", "kristen_Stewart", "avril_wade@163.com", "1468766556")

	if err == nil {
		str := DecodeJSON(data)
		fmt.Println(str)
		fmt.Println(len(str))
		fmt.Println(str[0], str[1], str[2])
	}

}*/
