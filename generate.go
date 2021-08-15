package overviewsdk

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetItems() error {
	f, err := os.Open("items")
	if err != nil {
		return err
	}
	defer f.Close()

	c := http.DefaultClient

	scanner := bufio.NewScanner(f)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://esi.evetech.net/latest/universe/groups/%s/?datasource=tranquility&language=en", scanner.Text()), nil)
		if err != nil {
			return err
		}
		req.Header.Set("accept", "application/json")
		req.Header.Set("Accept-Language", "en")
		req.Header.Set("Cache-Control", "no-cache")

		resp, err := c.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return errors.New(resp.Status)
		}
		buf, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(buf))
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
