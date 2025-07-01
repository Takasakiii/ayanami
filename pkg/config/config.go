package config

const configFile = "config.json"

type Config struct {
	Database Database `json:"database"`
	Server   Server   `json:"server"`
	Senders  Senders  `json:"senders"`
	File     File     `json:"file"`
}

type Server struct {
	BindHost  string `json:"bindHost"`
	BindPort  int    `json:"bindPort"`
	ServerUrl string `json:"url"`
}

type Senders struct {
	FileBin FileBin `json:"fileBin"`
	S3      S3      `json:"s3"`
}

type FileBin struct {
	BaseUrl string `json:"baseUrl"`
}

type File struct {
	Secret string `json:"secret"`
}

type Database struct {
	ConnectionString string `json:"connectionString"`
}

type S3 struct {
	Region          string `json:"region"`
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey"`
	Bucket          string `json:"bucket"`
	BucketPublicUrl string `json:"bucketPublicUrl"`
}

func newConfig() *Config {
	return &Config{
		Database: Database{
			ConnectionString: "database.db",
		},
		Server: Server{
			BindHost:  "127.0.0.1",
			BindPort:  2000,
			ServerUrl: "http://127.0.0.1:2000",
		},
		Senders: Senders{
			FileBin: FileBin{
				BaseUrl: "https://filebin.net",
			},
			S3: S3{
				Region: "auto",
			},
		},
		File: File{
			Secret: "changeThisSecret",
		},
	}
}
