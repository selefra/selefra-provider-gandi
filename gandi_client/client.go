package gandi_client

type Client struct {
	Config *Config
}

func NewClients(config Config) ([]*Client, error) {
	return []*Client{&Client{Config: &config}}, nil
}
