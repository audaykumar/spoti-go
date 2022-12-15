package config

import (
	"log"
	"os"

	"golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
)

var DebugMode bool

func init() {
	// log.Println("Inside config init")
	// DebugMode = os.Getenv("DEBUG") == "true"
	DebugMode = true
}

// Config represents what we pull in from the 'xero:' key in config.yml.
type Config struct {
	ClientID     string   `yaml:"client_id"`
	ClientSecret string   `yaml:"client_secret"`
	RedirectURI  string   `yaml:"redirect_uri"`
	Scopes       []string `yaml:"scopes"`

	OAuth2Config *oauth2.Config
}

// New - returns an instance of Config from the given filePath - If empty, will default to 'config.yml'.
func New() *Config {

	filePath := "config.yaml"

	// Check that the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Println("Config file does not exist:", filePath)
		log.Fatalln(err)
	}
	configFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("An error occurred while trying to read the config file:", filePath)
		log.Fatalln(err)
	}

	var config Config

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Println("Unable to parse contents of YAML config file:", filePath)
		log.Fatalln(err)
	}
	if DebugMode {
		log.Println("Loaded config:")
		config.Print()
	}

	config.OAuth2Config = &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Scopes:       config.Scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.spotify.com/authorize",
			TokenURL: "https://accounts.spotify.com/api/token",
		},
	}

	// Validate that the config has the required values
	validateConfig(&config)

	// Attempt to parse the YAML file for the details we need.
	return &config
}

func validateConfig(c *Config) {
	var fieldsMissing []string
	if c.ClientID == "" {
		fieldsMissing = append(fieldsMissing, "client_id")
	}
	if c.ClientSecret == "" {
		fieldsMissing = append(fieldsMissing, "client_secret")
	}
	if len(fieldsMissing) > 0 {
		log.Println("The follow fields appear to be missing from your config file:")
		for _, configKey := range fieldsMissing {
			log.Println("-", configKey)
		}
		log.Fatalln("Please ensure all required config values are present.")
	}
}

// Print outputs the loaded client struct.
func (c *Config) Print() {
	log.Println("Client ID:    ", c.ClientID)
	log.Println("Client Secret:", c.ClientSecret)
	log.Println("Redirect URI:", c.RedirectURI)
	log.Println("Scopes ", c.Scopes)
}
