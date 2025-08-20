package steam

import (
	"fmt"
	"regexp"
	"strconv"
)

func ConvertSteamID3ToSteamID64(steamID3 string) (uint64, error) {
	re := regexp.MustCompile(`^\[U:1:(\d+)\]$`)
	matches := re.FindStringSubmatch(steamID3)

	if len(matches) != 2 {
		return 0, fmt.Errorf("invalid SteamID3 format")
	}

	accountID, err := strconv.ParseUint(matches[1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse SteamID3 number: %w", err)
	}

	// SteamID64 = Base + Account ID
	const steamID64Base uint64 = 76561197960265728

	steamID64 := steamID64Base + accountID

	return steamID64, nil
}
