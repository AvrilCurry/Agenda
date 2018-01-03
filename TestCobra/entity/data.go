package entity

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"sort"
)

// User define the user struct
type User struct {
	Username  string
	Password  string // password cannot be seen
	Email     string
	Telephone string
	IsLogin   bool // whether user has logged in
}

// Meeting define the meeting struct
type Meeting struct {
	Title        string
	Participator []string
	StartTime    string
	EndTime      string
}

// UserRegister 0--Succeed 1--UserExisted
func UserRegister(username, password, email, telephone string) (int, error) {
	var user User
	var data []byte
	var err error
	var isUserExisted = false

	fin, err := os.OpenFile("./entity/userDatabase.data", os.O_RDWR, os.ModePerm)

	if err != nil {
		os.Exit(1)
	}

	buffreader := bufio.NewReader(fin)
	buffwriter := bufio.NewWriter(fin)

	for true {
		data, err = buffreader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}

		err = json.Unmarshal(data, &user)

		if username == user.Username {
			isUserExisted = true
			break
		}
	}

	if isUserExisted {
		return 1, err
	}

	user = User{username, password, email, telephone, false}
	result, err := json.Marshal(user)

	fin.Seek(0, os.SEEK_END)
	buffwriter.Write(result)
	buffwriter.WriteString("\n")
	buffwriter.Flush()
	fin.Close()
	return 0, err
}

// UserLogin 0--Succeed 1--UserNotExisted 2--PasswordWrong 3--UserlogInAlready
func UserLogin(username, password interface{}) (int, error) {
	var user User
	var data []byte
	var err error
	var isUserExisted = false
	var isUserLogin = false
	var isPasswordWrong = false

	fin, err := os.OpenFile("./entity/userDatabase.data", os.O_RDWR, os.ModePerm)
	fout, err := os.OpenFile("./entity/tempDatabase.data", os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		os.Exit(1)
	}

	buffreader := bufio.NewReader(fin)
	buffwriter := bufio.NewWriter(fout)

	for true {
		data, err = buffreader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}

		err = json.Unmarshal(data, &user)

		//fmt.Println(user, user.IsLogin)
		if username == user.Username && password == user.Password {
			isUserExisted = true
			isUserLogin = user.IsLogin

			if !isUserLogin {
				newUser := User{user.Username, user.Password, user.Email, user.Telephone, true}
				info, _ := json.Marshal(newUser)
				buffwriter.Write(info)
				buffwriter.WriteString("\n")
				buffwriter.Flush()
			} else {
				buffwriter.Write(data)
				buffwriter.Flush()
			}
		} else if user.Username == username && user.Password != password {
			isUserExisted = true
			isUserLogin = user.IsLogin
			isPasswordWrong = true

			buffwriter.Write(data)
			buffwriter.Flush()
		} else {
			buffwriter.Write(data)
			buffwriter.Flush()
		}
	}

	if !isUserExisted {
		err = os.Remove("./entity/tempDatabase.data")
		return 1, err
	}
	if isPasswordWrong {
		err = os.Remove("./entity/tempDatabase.data")
		return 2, err
	}
	if isUserLogin {
		err = os.Remove("./entity/tempDatabase.data")
		return 3, err
	}

	fin.Close()
	fout.Close()
	err = os.Remove("./entity/userDatabase.data")
	err = os.Rename("./entity/tempDatabase.data", "./entity/userDatabase.data")
	return 0, err
}

// UserLogout 0--Succeed 1--UserNotExisted 2--UserNotlogIn
func UserLogout(username string) (int, error) {
	var user User
	var data []byte
	var isUserExisted = false
	var isUserLogin = false

	fin, err := os.OpenFile("./entity/userDatabase.data", os.O_RDWR, os.ModePerm)
	fout, err := os.OpenFile("./entity/tempDatabase.data", os.O_RDWR|os.O_CREATE, os.ModePerm)

	buffreader := bufio.NewReader(fin)
	buffWriter := bufio.NewWriter(fout)

	if err != nil {
		os.Exit(1)
	}

	for true {
		data, err = buffreader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}

		err = json.Unmarshal(data, &user)

		if user.Username == username {
			isUserExisted = true
			isUserLogin = user.IsLogin

			if isUserLogin {
				newuser := User{user.Username, user.Password, user.Email, user.Telephone, false}
				info, _ := json.Marshal(newuser)

				buffWriter.Write(info)
				buffWriter.WriteString("\n")
				buffWriter.Flush()
			}
		} else {
			buffWriter.Write(data)
			buffWriter.Flush()
		}
	}

	if !isUserExisted {
		err = os.Remove("./entity/tempDatabase.data")
		return 1, err
	}
	if !isUserLogin {
		err = os.Remove("./entity/tempDatabase.data")
		return 2, err
	}

	fin.Close()
	fout.Close()
	err = os.Remove("./entity/userDatabase.data")
	err = os.Rename("./entity/tempDatabase.data", "./entity/userDatabase.data")
	return 0, err
}

// UserSearch 0--Succeed 1--UserNotExisted 2--UserNotLogin
func UserSearch(username string, val interface{}) (int, interface{}, error) {
	var err error
	var fin *os.File
	var data []byte
	var user User
	var singleMap map[string]string
	var result []map[string]string
	var isUserExisted = false
	var isUserLogin = false
	var isAdd = false

	singleMap = make(map[string]string)
	searchInfo, ok := val.([]string)
	if ok {
		fin, err = os.OpenFile("./entity/userDatabase.data", os.O_RDWR, os.ModePerm)
		buffreader := bufio.NewReader(fin)

		if err != nil {
			os.Exit(1)
		}

		for true {
			data, err = buffreader.ReadBytes('\n')

			if err != nil {
				if err != io.EOF {
					os.Exit(2)
				}
				break
			}

			err = json.Unmarshal(data, &user)
			//fmt.Println(user)

			if user.Username == username {
				isUserExisted = true
				isUserLogin = user.IsLogin
			}

			if len(searchInfo) == 0 {
				isAdd = true
			} else {
				for _, item := range searchInfo {
					if item == user.Username || item == user.Email || item == user.Telephone {
						isAdd = true
					}
				}
			}

			if isAdd {
				singleMap["Username"] = user.Username
				singleMap["Email"] = user.Email
				singleMap["Telephone"] = user.Telephone
				result = append(result, singleMap)
				singleMap = make(map[string]string, 0)
				isAdd = false
			}

		}
	}

	if !isUserExisted {
		return 1, 0, err
	}
	if !isUserLogin {
		return 2, 0, err
	}
	return 0, result, err
}

// UserDelete 0--succeed 1--UserNotExisted 2--UserNotLogin
func UserDelete(username string) (int, error) {
	var user User
	var meeting Meeting
	var userdata, meetingdata []byte
	var isUserExisted = false
	var isUserLogin = false

	fin, err := os.OpenFile("./entity/userDatabase.data", os.O_RDWR, os.ModePerm)
	fout, err := os.OpenFile("./entity/meetingDatabase.data", os.O_RDWR, os.ModePerm)
	fs, err := os.OpenFile("./entity/tempUserDatabase.data", os.O_RDWR|os.O_CREATE, os.ModePerm)
	fs1, err := os.OpenFile("./entity/tempMeetingDatabase.data", os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		os.Exit(1)
	}

	buffUserReader := bufio.NewReader(fin)
	buffMeetingReader := bufio.NewReader(fout)
	buffUserWriter := bufio.NewWriter(fs)
	buffMeetingWriter := bufio.NewWriter(fs1)

	for true {
		userdata, err = buffUserReader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}

		err = json.Unmarshal(userdata, &user)

		if username == user.Username {
			isUserExisted = true
			isUserLogin = user.IsLogin

			if user.IsLogin {
				for true {
					meetingdata, err = buffMeetingReader.ReadBytes('\n')

					if err != nil {
						if err != io.EOF {

						}
						break
					}

					err = json.Unmarshal(meetingdata, &meeting)

					if meeting.Participator[0] == username {
						continue
					} else {
						var tempparticipator []string
						for _, item := range meeting.Participator {
							if item != username {
								tempparticipator = append(tempparticipator, item)
							}
						}
						if len(tempparticipator) != 0 {
							newmeeting := Meeting{meeting.Title, tempparticipator, meeting.StartTime, meeting.EndTime}
							info, _ := json.Marshal(newmeeting)
							buffMeetingWriter.Write(info)
							buffMeetingWriter.WriteString("\n")
							buffMeetingWriter.Flush()
						}
					}
				}
				continue
			} else {
				buffUserWriter.Write(userdata)
				buffUserWriter.Flush()
			}
		} else {
			buffUserWriter.Write(userdata)
			buffUserWriter.Flush()
		}
	}

	if !isUserExisted {
		err = os.Remove("./entity/tempUserDatabase.data")
		err = os.Remove("./entity/tempMeetingDatabase.data")
		return 1, err
	}
	if !isUserLogin {
		err = os.Remove("./entity/tempUserDatabase.data")
		err = os.Remove("./entity/tempMeetingDatabase.data")
		return 2, err
	}

	fin.Close()
	fout.Close()
	fs.Close()
	err = os.Remove("./entity/userDatabase.data")
	err = os.Rename("./entity/tempUserDatabase.data", "./entity/userDatabase.data")
	err = os.Remove("./entity/meetingDatabase.data")
	err = os.Rename("./entity/tempMeetingDatabase.data", "./entity/meetingDatabase.data")
	return 0, err
}

/*
CreateMeeting 0--Succeed 1--UserNotExited 2--UserNotLogIn 3--MeetingExisted 4--NotAllRegister
			5--TimeConflict
*/
func CreateMeeting(username string, title string, participators []string, startTime string, endTime string) (int, interface{}, error) {
	var user User
	var meeting Meeting
	var userdata, meetingdata []byte
	var err error
	var isUserExisted = false
	var isLogin = false
	var isMeetingExisted = false
	var isNotAllRegister = false
	var isTimeConflict = false
	var UnrregisterUser []string
	var ConflictUser []string
	var array []bool

	fin, err := os.OpenFile("./entity/userDatabase.data", os.O_RDWR, os.ModePerm)
	fout, err := os.OpenFile("./entity/meetingDatabase.data", os.O_RDWR, os.ModePerm)

	if err != nil {
		os.Exit(1)
	}

	buffUserReader := bufio.NewReader(fin)
	buffMeetingReader := bufio.NewReader(fout)
	buffMeetingWriter := bufio.NewWriter(fout)

	for i := 0; i < len(participators); i++ {
		array = append(array, false)
	}

	for true {
		userdata, err = buffUserReader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}

		err = json.Unmarshal(userdata, &user)
		if user.Username == username {
			isUserExisted = true
			isLogin = user.IsLogin
		}
		for index, item := range participators {
			if item == user.Username {
				array[index] = true
			}
		}
	}

	// User Not Existed
	if !isUserExisted {
		return 1, 0, err
	}
	// User Not Login
	if !isLogin {
		return 2, 0, err
	}

	for index, item := range array {
		if !item {
			UnrregisterUser = append(UnrregisterUser, participators[index])
			isNotAllRegister = true
		}
	}

	for true {
		meetingdata, err = buffMeetingReader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}

		err = json.Unmarshal(meetingdata, &meeting)

		if meeting.Title == title {
			isMeetingExisted = true
			break
		} else {
			for _, item := range participators {
				for _, item2 := range meeting.Participator {
					if item == item2 && !(meeting.StartTime >= endTime || meeting.EndTime <= startTime) {
						isTimeConflict = true
						ConflictUser = append(ConflictUser, item)
						break
					}
				}
			}
		}
	}

	if isMeetingExisted {
		return 3, 0, err
	}
	if isNotAllRegister {
		return 4, UnrregisterUser, err
	}
	if isTimeConflict {
		return 5, ConflictUser, err
	}

	newmeeting := Meeting{title, participators, startTime, endTime}
	info, err := json.Marshal(newmeeting)

	fout.Seek(0, os.SEEK_END)
	buffMeetingWriter.Write(info)
	buffMeetingWriter.WriteString("\n")
	buffMeetingWriter.Flush()
	fin.Close()
	fout.Close()
	return 0, 0, err
}

/*
AddParticipator 0--succeed 1--UserNotExisted 2--UserNotLogin 3--MeetingNotExisted 4--MeetingNotCreatedByYou
	5--NotAllRegister 6--TimeConflict
*/
func AddParticipator(username string, title string, participators []string) (int, interface{}, error) {
	var user User
	var meeting Meeting
	var userdata, meetingdata []byte
	var err error
	var startTime, endTime string
	var isUserExisted = false
	var isUserLogin = false
	var isMeetingExisted = false
	var isCreatedByYou = false
	var isNotAllRegister = false
	var isTimeConflict = false
	var UnregisterUser []string
	var ConflictUser []string
	var array []bool

	fin, err := os.OpenFile("./entity/userDatabase.data", os.O_RDWR, os.ModePerm)
	fout, err := os.OpenFile("./entity/meetingDatabase.data", os.O_RDWR, os.ModePerm)

	if err != nil {
		os.Exit(1)
	}

	buffUserReader := bufio.NewReader(fin)
	buffMeetingReader := bufio.NewReader(fout)

	// 去除重复的参与者
	var newparticipators []string
	sort.Strings(participators)
	newparticipators = append(newparticipators, participators[0])

	for i := 0; i < len(participators)-1; i++ {
		if participators[i] != participators[i+1] {
			newparticipators = append(newparticipators, participators[i+1])
		}
	}

	// 记录将要加入的参与者是否已经注册
	for j := 0; j < len(newparticipators); j++ {
		array = append(array, false)
	}

	for true {
		userdata, err = buffUserReader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}

		err = json.Unmarshal(userdata, &user)

		if user.Username == username {
			isUserExisted = true
			isUserLogin = user.IsLogin
		} else {
			for index := range newparticipators {
				if newparticipators[index] == user.Username {
					array[index] = true
				}
			}
		}
	}

	if !isUserExisted {
		return 1, 0, err
	}
	if !isUserLogin {
		return 2, 0, err
	}

	for true {
		meetingdata, err = buffMeetingReader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}

		err = json.Unmarshal(meetingdata, &meeting)

		if meeting.Title == title {
			isMeetingExisted = true

			if meeting.Participator[0] == username {
				isCreatedByYou = true
				startTime = meeting.StartTime
				endTime = meeting.EndTime
				fout.Seek(0, os.SEEK_SET)
			}
			break
		}
	}

	if !isMeetingExisted {
		return 3, 0, err
	}
	if !isCreatedByYou {
		return 4, 0, err
	}

	for index, item := range array {
		if !item {
			isNotAllRegister = true
			UnregisterUser = append(UnregisterUser, newparticipators[index])
		}
	}

	if isNotAllRegister {
		return 5, UnregisterUser, err
	}

	for true {
		meetingdata, err = buffMeetingReader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}

		err = json.Unmarshal(meetingdata, &meeting)

		if meeting.Title != title {
			for _, item := range newparticipators {
				for _, item2 := range meeting.Participator {
					if item == item2 && !(meeting.StartTime >= endTime || meeting.EndTime <= startTime) {
						isTimeConflict = true
						ConflictUser = append(ConflictUser, item)
						break
					}
				}
			}
		}
		if isTimeConflict {
			break
		}
	}

	if isTimeConflict {
		fin.Close()
		fout.Close()
		return 6, ConflictUser, err
	}

	fout.Seek(0, os.SEEK_SET)
	fs, err := os.OpenFile("./entity/tempDatabase.data", os.O_RDWR|os.O_CREATE, os.ModePerm)
	buffWriter := bufio.NewWriter(fs)

	for true {
		meetingdata, err = buffMeetingReader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}

		err = json.Unmarshal(meetingdata, &meeting)

		if meeting.Title == title {
			for _, item := range newparticipators {
				meeting.Participator = append(meeting.Participator, item)
			}
			newmeeting := Meeting{meeting.Title, meeting.Participator, meeting.StartTime, meeting.EndTime}
			info, _ := json.Marshal(newmeeting)

			buffWriter.Write(info)
			buffWriter.WriteString("\n")
			buffWriter.Flush()
		} else {
			buffWriter.Write(meetingdata)
			buffWriter.Flush()
		}
	}

	fin.Close()
	fout.Close()
	fs.Close()
	err = os.Remove("./entity/meetingDatabase.data")
	err = os.Rename("./entity/tempDatabase.data", "./entity/meetingDatabase.data")
	return 0, 0, err
}

/*
DeleteParticipator 0--succeed 1--UserNotExisted 2--UserNotLogin 3--MeetingNotExisted 4--MeetingNotCreatedByUser
				5--UserUnableToDelete(Cause It's the Organizer)
*/
func DeleteParticipator(username string, title string, participators []string) (int, error) {
	var user User
	var meeting Meeting
	var userdata, meetingdata []byte
	var err error
	var isUserExisted = false
	var isUserLogin = false
	var isMeetingExisted = false
	var isUnableToDelete = false

	fin, err := os.OpenFile("./entity/userDatabase.data", os.O_RDWR, os.ModePerm)
	fout, err := os.OpenFile("./entity/meetingDatabase.data", os.O_RDWR, os.ModePerm)
	fs, err := os.OpenFile("./entity/tempMeetingDatabase.data", os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		os.Exit(1)
	}

	buffUserReader := bufio.NewReader(fin)
	buffMeetingReader := bufio.NewReader(fout)
	buffWriter := bufio.NewWriter(fs)

	for true {
		userdata, err = buffUserReader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}

		err = json.Unmarshal(userdata, &user)

		if user.Username == username {
			isUserExisted = true
			isUserLogin = user.IsLogin
			break
		}
	}

	if !isUserExisted {
		err = os.Remove("./entity/tempMeetingDatabase.data")
		return 1, err
	}
	if !isUserLogin {
		err = os.Remove("./entity/tempMeetingDatabase.data")
		return 2, err
	}

	for true {
		meetingdata, err = buffMeetingReader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}

		err = json.Unmarshal(meetingdata, &meeting)

		if meeting.Title == title {
			isMeetingExisted = true

			if meeting.Participator[0] != username {
				os.Remove("./entity/tempMeetingDatabase.data")
				return 4, err
			}

			var tempparticipator []string
			var flag = false  // 判断删除的参与者是否在原列表中
			var isNIL = false // 判断参与者列表是否为空

			if len(meeting.Participator) == 1 {
				for _, item := range participators {
					if item == username {
						isNIL = true
						break
					}
				}
				if isUnableToDelete {
					break
				}
				if !isNIL {
					buffWriter.Write(meetingdata)
					buffWriter.Flush()
				}
			} else {
				for index, item := range meeting.Participator {
					flag = false
					for _, item2 := range participators {
						if item == item2 && index == 0 {
							isUnableToDelete = true
							break
						} else if item == item2 && index != 0 {
							flag = true
							break
						}
					}
					if isUnableToDelete {
						break
					}
					if !flag {
						tempparticipator = append(tempparticipator, item)
					}
				}

				if isUnableToDelete {
					break
				}
				//fmt.Println(tempparticipator)
				newmeeting := Meeting{meeting.Title, tempparticipator, meeting.StartTime, meeting.EndTime}
				info, _ := json.Marshal(newmeeting)

				buffWriter.Write(info)
				buffWriter.WriteString("\n")
				buffWriter.Flush()
			}
		} else {
			buffWriter.Write(meetingdata)
			buffWriter.Flush()
		}
		if isUnableToDelete {
			break
		}
	}

	if !isMeetingExisted {
		err = os.Remove("./entity/tempMeetingDatabase.data")
		return 3, err
	}
	if isUnableToDelete {
		err = os.Remove("./entity/tempMeetingDatabase.data")
		return 5, err
	}

	fin.Close()
	fout.Close()
	fs.Close()

	err = os.Remove("./entity/meetingDatabase.data")
	err = os.Rename("./entity/tempMeetingDatabase.data", "./entity/meetingDatabase.data")
	return 0, err
}

// MeetingSearch  1--UserNotExisted 2--UserNotLogin
func MeetingSearch(username, startTime, endTime string) (interface{}, error) {
	var user User
	var meeting Meeting
	var userdata, meetingdata []byte
	var result []Meeting
	var err error
	var isUserExisted = false
	var isUserLogin = false
	var isInMeetingParticipator = false

	fin, err := os.OpenFile("./entity/userDatabase.data", os.O_RDWR, os.ModePerm)
	fout, err := os.OpenFile("./entity/meetingDatabase.data", os.O_RDWR, os.ModePerm)

	if err != nil {
		os.Exit(1)
	}

	buffUserReader := bufio.NewReader(fin)
	buffMeetingReader := bufio.NewReader(fout)

	for true {
		userdata, err = buffUserReader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}

		err = json.Unmarshal(userdata, &user)

		if user.Username == username {
			isUserExisted = true
			isUserLogin = user.IsLogin
			break
		}
	}

	if !isUserExisted {
		return 1, err
	}
	if !isUserLogin {
		return 2, err
	}

	for true {
		meetingdata, err = buffMeetingReader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}

		err = json.Unmarshal(meetingdata, &meeting)

		//fmt.Println(meeting.Participator)
		for _, participator := range meeting.Participator {
			if participator == username {
				isInMeetingParticipator = true
				break
			}
		}

		if isInMeetingParticipator && startTime <= meeting.StartTime && meeting.EndTime <= endTime {
			result = append(result, meeting)
			isInMeetingParticipator = false
		}
	}

	fin.Close()
	fout.Close()
	return result, err
}

// CancelMeeting 0--Succeed 1--UserNotExisted 2--UserNotLogin 3--MeetingNotExisted 4--MeetingNotCreatedByYou
func CancelMeeting(username, title string) (int, error) {
	var user User
	var meeting Meeting
	var userdata, meetingdata []byte
	var err error
	var isUserExisted = false
	var isUserLogin = false
	var isMeetingExisted = false

	fin, err := os.OpenFile("./entity/userDatabase.data", os.O_RDWR, os.ModePerm)
	fout, err := os.OpenFile("./entity/meetingDatabase.data", os.O_RDWR, os.ModePerm)
	fs, err := os.OpenFile("./entity/tempMeetingDatabase.data", os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		os.Exit(1)
	}

	buffUserReader := bufio.NewReader(fin)
	buffMeetingReader := bufio.NewReader(fout)
	buffWriter := bufio.NewWriter(fs)

	for true {
		userdata, err = buffUserReader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}

		err = json.Unmarshal(userdata, &user)

		if user.Username == username {
			isUserExisted = true
			isUserLogin = user.IsLogin
			break
		}
	}

	if !isUserExisted {
		err = os.Remove("./entity/tempMeetingDatabase.data")
		return 1, err
	}
	if !isUserLogin {
		err = os.Remove("./entity/tempMeetingDatabase.data")
		return 2, err
	}

	for true {
		meetingdata, err = buffMeetingReader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}

		err = json.Unmarshal(meetingdata, &meeting)

		if meeting.Title == title {
			isMeetingExisted = true
			if meeting.Participator[0] == username {
				continue
			} else {
				err = os.Remove("./entity/tempMeetingDatabase.data")
				return 4, err
			}
		} else {
			buffWriter.Write(meetingdata)
			buffWriter.Flush()
		}
	}

	if !isMeetingExisted {
		err = os.Remove("./entity/tempMeetingDatabase.data")
		return 3, err
	}

	fin.Close()
	fout.Close()
	fs.Close()
	err = os.Remove("./entity/meetingDatabase.data")
	err = os.Rename("./entity/tempMeetingDatabase.data", "./entity/meetingDatabase.data")
	return 0, err
}

/*
QuitMeeting 0--Succeed 1--UserNotExisted 2--UserNotLogin 3--MeetingNotExisted 4--UserNotInTheParticipators
	5--UserUnableToQuit(Cause It's the Organizer)
*/
func QuitMeeting(username, title string) (int, error) {
	var user User
	var meeting Meeting
	var userdata, meetingdata []byte
	var err error
	var isUserExisted = false
	var isUserLogin = false
	var isMeetingExisted = false
	var isIntheParticipators = false
	var isUnableToQuit = false

	fin, err := os.OpenFile("./entity/userDatabase.data", os.O_RDWR, os.ModePerm)
	fout, err := os.OpenFile("./entity/meetingDatabase.data", os.O_RDWR, os.ModePerm)
	fs, err := os.OpenFile("./entity/tempMeetingDatabase.data", os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		os.Exit(1)
	}

	buffUserReader := bufio.NewReader(fin)
	buffMeetingReader := bufio.NewReader(fout)
	buffWriter := bufio.NewWriter(fs)

	for true {
		userdata, err = buffUserReader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}

		err = json.Unmarshal(userdata, &user)

		if user.Username == username {
			isUserExisted = true
			isUserLogin = user.IsLogin
			break
		}
	}

	if !isUserExisted {
		err = os.Remove("./entity/tempMeetingDatabase.data")
		return 1, err
	}
	if !isUserLogin {
		err = os.Remove("./entity/tempMeetingDatabase.data")
		return 2, err
	}

	for true {
		meetingdata, err = buffMeetingReader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}

		err = json.Unmarshal(meetingdata, &meeting)

		if meeting.Title == title {
			isMeetingExisted = true

			if len(meeting.Participator) == 1 {
				if meeting.Participator[0] == username {
					isIntheParticipators = true
					continue
				} else {
					buffWriter.Write(meetingdata)
					buffWriter.Flush()
				}
			} else {
				var tempparticipator []string
				for index, item := range meeting.Participator {
					if item == username && index == 0 {
						isIntheParticipators = true
						isUnableToQuit = true
						buffWriter.Write(meetingdata)
						buffWriter.Flush()
						break
					} else if item == username && index != 0 {
						isIntheParticipators = true
					} else {
						tempparticipator = append(tempparticipator, item)
					}
				}
				if isUnableToQuit {
					break
				} else {
					newmeeting := Meeting{meeting.Title, tempparticipator, meeting.StartTime, meeting.EndTime}
					info, _ := json.Marshal(newmeeting)

					buffWriter.Write(info)
					buffWriter.WriteString("\n")
					buffWriter.Flush()
				}
			}
		} else {
			buffWriter.Write(meetingdata)
			buffWriter.Flush()
		}
		if isUnableToQuit {
			break
		}
	}

	if !isMeetingExisted {
		err = os.Remove("./entity/tempMeetingDatabase.data")
		return 3, err
	}
	if !isIntheParticipators {
		err = os.Remove("./entity/tempMeetingDatabase.data")
		return 4, err
	}
	if isUnableToQuit {
		err = os.Remove("./entity/tempMeetingDatabase.data")
		return 5, err
	}

	fin.Close()
	fout.Close()
	fs.Close()
	err = os.Remove("./entity/meetDatabase.data")
	err = os.Rename("./entity/tempMeetingDatabase.data", "./entity/meetingDatabase.data")

	return 0, err
}

// ClearMeeting 0--Succeed 1--UserNotExisted 2--UserNotLogin
func ClearMeeting(username string) (int, error) {
	var user User
	var meeting Meeting
	var userdata, meetingdata []byte
	var err error
	var isUserExisted = false
	var isUserLogin = false

	fin, err := os.OpenFile("./entity/userDatabase.data", os.O_RDWR, os.ModePerm)
	fout, err := os.OpenFile("./entity/meetingDatabase.data", os.O_RDWR, os.ModePerm)
	fs, err := os.OpenFile("./entity/tempMeetingDatabase.data", os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		os.Exit(1)
	}

	buffUserReader := bufio.NewReader(fin)
	buffMeetingReader := bufio.NewReader(fout)
	buffWriter := bufio.NewWriter(fs)

	for true {
		userdata, err = buffUserReader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}

		err = json.Unmarshal(userdata, &user)

		if user.Username == username {
			isUserExisted = true
			isUserLogin = user.IsLogin
			break
		}
	}

	if !isUserExisted {
		err = os.Remove("./entity/tempMeetingDatabase.data")
		return 1, err
	}
	if !isUserLogin {
		err = os.Remove("./entity/tempMeetingDatabase.data")
		return 2, err
	}

	for true {
		meetingdata, err = buffMeetingReader.ReadBytes('\n')

		if err != nil {
			if err != io.EOF {
				os.Exit(2)
			}
			break
		}

		err = json.Unmarshal(meetingdata, &meeting)

		//fmt.Println(meeting)
		if meeting.Participator[0] == username {
			continue
		} else {
			buffWriter.Write(meetingdata)
			buffWriter.Flush()
		}
	}

	fin.Close()
	fout.Close()
	fs.Close()
	err = os.Remove("./entity/meetingDatabase.data")
	err = os.Rename("./entity/tempMeetingDatabase.data", "./entity/meetingDatabase.data")
	return 0, err
}

/*func main() {
//UserRegister
/*res, _ := UserRegister("Taylor Swift", "Country_Music_Queen", "Taylor_Swift@gmail.com", "124157998")
fmt.Println(res)
res, _ = UserRegister("Adele Queen", "British_Queen", "Adele_Queen@qq.com", "1906214878")
fmt.Println(res)*/

// Userlogin
/*res, _ := UserLogin("Kristen Stewart", "Iassc_Newtoen")
fmt.Println(res)
res, _ = UserLogin("Avril Lavigne", "kristen_Stewart")
fmt.Println(res)
res, _ = UserLogin("Avril Lavigne", "Jamie Lanster")
fmt.Println(res)
res, _ = UserLogin("Avril Lavigne", "Jamie Lanster")
fmt.Println(res)*/

//UserSearch
/*res, user, _ := UserSearch("Avril Lavigne", []string{"124157998"})
fmt.Println(res)
fmt.Println(user)*/

//MeetingSearch
/*res, _ := UserLogin("Steve Kerr", "Haly_Porter")
fmt.Println(res)*/
/*value, _ := MeetingSearch("Michelle Dockery", "2017-11-04 11:30", "2017-11-04 23:00")
switch result := value.(type) {
case int:
	fmt.Println(result)
case []Meeting:
	for _, meeting := range result {
		fmt.Println(meeting)
	}
default:
	fmt.Println("none")
}*/

// UserLogin
/*res, _ := UserLogin("Avril Lavigne", "Jamie Lanster")
fmt.Println(res)*/
/*res, _ = UserLogin("Steve Kerr", "Haly_Porter")
fmt.Println(res)*/

//UserLogout
/*res, _ := UserLogout("Hally POrter")
fmt.Println(res)
res, _ = UserLogout("Klay Thompson")
fmt.Println(res)
res, _ = UserLogout("Avril Lavigne")
fmt.Println(res)*/

// UserDelete
/*res, _ := UserDelete("Hason Jerry")
fmt.Println(res)*/

// CreateMeeting
/*res, value, _ := CreateMeeting("Michelle Dockery", "DownTown_Abbey", []string{"Michelle Dockery", "Taylor Swift", "Adele Queen"}, "2017-11-04 19:30", "2017-11-04 20:20")
fmt.Println(res)
fmt.Println(value)*/
/*res, value, _ := CreateMeeting("Klay Thompson", "King of God", []string{"Klay Thompson", "Taylor Swift", "Avril Lavigne"}, "2017-11-04 17:40", "2017-11-04 18:40")
fmt.Println(res)
fmt.Println(value)*/
/*res, value, _ = CreateMeeting("Avril Lavigne", "Let it Go", []string{"Avril Lavigne", "Steve Kerr", "Kristen Stewart"}, "2017-11-04 8:10", "2017-11-04 9:20")
fmt.Println(res)
fmt.Println(value)*/

// AddParticipators
/*res, value, _ := AddParticipator("Michelle Dockery", "DownTown_Abbey", []string{"Zack Landole", "Stephen Curry"})
fmt.Println(res)
fmt.Println(value)*/

// DeleteParticipators
/*res, _ := DeleteParticipator("Kristen Stewart", "Just for Fun", []string{"Kristen Stewart"})
fmt.Println(res)*/

// CancelMeeting
/*res, _ := CancelMeeting("Avril Lavigne", "Complicate")
fmt.Println(res)*/

// QuitMeeting
/*res, _ := UserLogin("Klay Tompason", "Sliver_Adam")
fmt.Println(res)*/
/*res, _ := QuitMeeting("Avril Lavigne", "King of God")
fmt.Println(res)*/

// ClearMeeting
/*res, _ := UserLogin("Taylor Swift", "Taylor_Country_Music")
fmt.Println(res)
res, _ := ClearMeeting("Klay Thompson")
fmt.Println(res)*/
/*}
 */
