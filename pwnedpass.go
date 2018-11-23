package pwnedpass

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// ClientV2 is used for making requests to the PwnedPass V2 API.
type ClientV2 struct {
	HTTPClient *http.Client
	BaseURL    string
}

// BaseURLV2 is the base URL for the V2 Pwned Passwords API.
const BaseURLV2 = "https://api.pwnedpasswords.com"

// DefaultClient is the default client used for making requests to the PwnedPass
// API. It uses the V2 client with the https://api.pwnedpasswords.com base URL.
// It has a sane timeout of 5 seconds.
var DefaultClient = &ClientV2{
	HTTPClient: &http.Client{
		Timeout: time.Second * 5,
	},
}

const (
	errUnableToWrite     = "unable to write password to sha1 hash.Hash"
	errShortHash         = "hex encoded hash is too short"
	errMakingHTTPRequest = "got error making http request"
	errParsingCount      = "error parsing count"
)

// Count returns the numbers of passwords associated with a given password.
// Count will use http.DefaultClient if the Client HTTPClient is nil and it will
// use BaseURLV2 if the Client BaseURL is empty. This means ClientV2 has a
// usable zero value.
func (c *ClientV2) Count(password string) (int, error) {
	h := sha1.New()
	if _, err := h.Write([]byte(password)); err != nil {
		return 0, errors.Wrap(err, errUnableToWrite)
	}
	sha := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))

	if len(sha) < 5 {
		return 0, errors.New(errShortHash)
	}

	prefix, rest := sha[:5], sha[5:]

	url := fmt.Sprintf("%s/range/%s", c.baseURL(), prefix)

	resp, err := c.httpClient().Get(url)
	if err != nil {
		return 0, errors.Wrap(err, errMakingHTTPRequest)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")
		if len(parts) < 2 {
			continue
		}

		if strings.ToUpper(parts[0]) == rest {
			n, err := strconv.Atoi(parts[1])
			return n, errors.Wrap(err, errParsingCount)
		}
	}

	return 0, nil
}

func (c *ClientV2) httpClient() *http.Client {
	if c.HTTPClient != nil {
		return c.HTTPClient
	}

	return http.DefaultClient
}

func (c *ClientV2) baseURL() string {
	if c.BaseURL != "" {
		return c.BaseURL
	}

	return BaseURLV2
}
