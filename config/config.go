package config

var (
	API_KEY  string = "apikey"
	SEC_KEY  string = "secret key"
	BASE_URL string = "https://api.mexc.com/api/v3"
	PROXY    string = ""
)

func Init(apikey string, secret string) {
	API_KEY = apikey
	SEC_KEY = secret
}
