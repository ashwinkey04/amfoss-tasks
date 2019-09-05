package main
import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/kurrik/oauth1a"
	"github.com/kurrik/twittergo"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)
func LoadCredentials() (client *twittergo.Client, err error) {
	credentials, err := ioutil.ReadFile("CREDENTIALS")
	if err != nil {
		return
	}
	lines := strings.Split(string(credentials), "\n")
	config := &oauth1a.ClientConfig{
		ConsumerKey:    lines[0],
		ConsumerSecret: lines[1],
	}
	user := oauth1a.NewAuthorizedConfig(lines[2], lines[3])
	client = twittergo.NewClient(config, user)
	return
}
type Args struct {
	ScreenName string
	OutputFile string
}
func parseArgs() *Args {
	a := &Args{}
	flag.StringVar(&a.ScreenName, "ashwinkey04", "twitterapi", "Screen name")
	flag.StringVar(&a.OutputFile, "out", "user_timeline.json", "Output file")
	flag.Parse()
	return a
}
func main() {
	var (
		err     error
		client  *twittergo.Client
		req     *http.Request
		resp    *twittergo.APIResponse
		args    *Args
		max_id  uint64
		out     *os.File
		query   url.Values
		results *twittergo.Timeline
		text    []byte
	)
	args = parseArgs()
	if client, err = LoadCredentials(); err != nil {
		fmt.Printf("Could not parse CREDENTIALS file: %v\n", err)
		os.Exit(1)
	}
	if out, err = os.Create(args.OutputFile); err != nil {
		fmt.Printf("Could not create output file: %v\n", args.OutputFile)
		os.Exit(1)
	}
	defer out.Close()
	const (
		count   int = 100
		urltmpl     = "/1.1/statuses/user_timeline.json?%v"
		minwait     = time.Duration(10) * time.Second
	)
	query = url.Values{}
	query.Set("count", fmt.Sprintf("%v", count))
	query.Set("ashwinkey04", args.ScreenName)
	total := 0
	for {
		if max_id != 0 {
			query.Set("max_id", fmt.Sprintf("%v", max_id))
		}
		endpoint := fmt.Sprintf(urltmpl, query.Encode())
		if req, err = http.NewRequest("GET", endpoint, nil); err != nil {
			fmt.Printf("Could not parse request: %v\n", err)
			os.Exit(1)
		}
		if resp, err = client.SendRequest(req); err != nil {
			fmt.Printf("Could not send request: %v\n", err)
			os.Exit(1)
		}
		results = &twittergo.Timeline{}
		if err = resp.Parse(results); err != nil {
			if rle, ok := err.(twittergo.RateLimitError); ok {
				dur := rle.Reset.Sub(time.Now()) + time.Second
				if dur < minwait {
					// Don't wait less than minwait.
					dur = minwait
				}
				msg := "Rate limited. Reset at %v. Waiting for %v\n"
				fmt.Printf(msg, rle.Reset, dur)
				time.Sleep(dur)
				continue // Retry request.
			} else {
				fmt.Printf("Problem parsing response: %v\n", err)
			}
		}
		batch := len(*results)
		if batch == 0 {
			fmt.Printf("No more results, end of timeline.\n")
			break
		}
		for _, tweet := range *results {
			if text, err = json.Marshal(tweet); err != nil {
				fmt.Printf("Could not encode Tweet: %v\n", err)
				os.Exit(1)
			}
			out.Write(text)
			out.Write([]byte("\n"))
			max_id = tweet.Id() - 1
			total += 1
		}
		fmt.Printf("Got %v Tweets", batch)
		if resp.HasRateLimit() {
			fmt.Printf(", %v calls available", resp.RateLimitRemaining())
		}
		fmt.Printf(".\n")
	}
	fmt.Printf("--------------------------------------------------------\n")
	fmt.Printf("Wrote %v Tweets to %v\n", total, args.OutputFile)
}