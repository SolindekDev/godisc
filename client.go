package godisc

type Client struct {
	Token   string
	Intents int
}

func NewClient(token string, intents int) Client {
	client := Client{}
	client.Token = token
	if intents == 0 {
		client.Intents = 7796
	} else {
		client.Intents = intents
	}

	return client
}
