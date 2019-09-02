package main
import ("github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"flag"
	"fmt"
    	"os")

var (
	newFile *os.File
	err     error
)

func main(){
	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)

	consumerKey := flags.String("consumer-key", "L7bfVIrublBMvEn7lh3DIbfuf", "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", "s6syCC3p59bBlP1memLDvjdPAqlKNHP8ZRTvm6hK0QQD0nncfa", "Twitter Consumer Secret")
	accessToken := flags.String("access-token", "742371767486808065-kDqot1MMGyxzi3Gne4Y6RHH0n79ttsm", "Twitter Access Token")
	accessSecret := flags.String("access-secret", "k866oiH2aYf6NCdzopiuTLyhtkeuMoiZuigAcBXK87BQf", "Twitter Access Secret")
	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnv(flags, "TWITTER")
	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)	
	client := twitter.NewClient(httpClient)
	followers, resp, err := client.Followers.List(&twitter.FollowerListParams{})
	fmt.Println("Ash")
}
