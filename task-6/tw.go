package main
import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/dghubble/oauth1/twitter""
)
config := oauth1.NewConfig("L7bfVIrublBMvEn7lh3DIbfuf", "s6syCC3p59bBlP1memLDvjdPAqlKNHP8ZRTvm6hK0QQD0nncfa")
token := oauth1.NewToken("742371767486808065-kDqot1MMGyxzi3Gne4Y6RHH0n79ttsm", "k866oiH2aYf6NCdzopiuTLyhtkeuMoiZuigAcBXK87BQf")


httpClient := config.Client(oauth1.NoContext, token)

	
client := twitter.NewClient(httpClient)
