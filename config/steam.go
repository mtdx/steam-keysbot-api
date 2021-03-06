package config

import "os"

// SteamAPIKey ...
func SteamAPIKey() string {
	key := os.Getenv("KEYC_STEAM_API_KEY")
	if key == "" {
		key = "285C61E0D684F696A53B67F3B81B07E9"
	}
	return key
}

// SteamBotsAPIKey ...
func SteamBotsAPIKey() string {
	key := os.Getenv("KEYC_STEAM_BOTS_API_KEY")
	if key == "" {
		key = "JK7GgUgxt5wTtEYLKG6VdwyUQS4kLVRpLtzy" // Insecure Default Bots API key
	}
	return key
}
