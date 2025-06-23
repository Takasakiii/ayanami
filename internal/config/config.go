package config

const configFile = "config.json"

type Config struct {
	Server  Server  `json:"server"`
	Senders Senders `json:"senders"`
	File    File    `json:"file"`
}

type Server struct {
	BindHost  string `json:"bindHost"`
	BindPort  int    `json:"bindPort"`
	ServerUrl string `json:"url"`
}

type Senders struct {
	FileBin FileBin `json:"fileBin"`
}

type FileBin struct {
	BaseUrl string `json:"baseUrl"`
}

type File struct {
	Secret string `json:"secret"`
}

func newConfig() Config {
	return Config{
		Server: Server{
			BindHost:  "127.0.0.1",
			BindPort:  2000,
			ServerUrl: "http://127.0.0.1:2000",
		},
		Senders: Senders{
			FileBin: FileBin{
				BaseUrl: "https://filebin.net",
			},
		},
		File: File{
			Secret: "changeThisSecret",
		},
	}
}
