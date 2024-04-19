package entities

type Timeline struct {
	Tweets []TweetInTimeline
}

type TweetInTimeline struct {
	ID     string
    UserId string
	Username string
    Text   string
    Ts     int64
}