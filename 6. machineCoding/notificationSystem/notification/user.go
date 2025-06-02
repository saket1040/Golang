package notification

// User holds name and preferred notification channels.
type User struct {
	Name     string
	Channels []Notification
}