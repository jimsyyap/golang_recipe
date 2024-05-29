package shodan

const BaseURL = "https://api.shodan.io"

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}
