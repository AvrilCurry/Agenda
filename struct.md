type User struct {
    Username    string
    password    string
    Email       string
    Phone       string
    isLogIn     bool
}

type Meeting struct {
    Title            string
    Paiticipator     []string
    StartTime        string
    EndTime          string
    organizer        string
}
