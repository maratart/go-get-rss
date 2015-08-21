package main

import
(
	"fmt"
	"os"

	rss "github.com/jteeuwen/go-pkg-rss"
	"github.com/jteeuwen/go-pkg-xmlx"
)

func main() {
	if len(os.Args) > 1 {
		myFeed := os.Args[1]
		PollFeed(myFeed, 5, nil)
	} else {

		fmt.Println("Usage:\n" + os.Args[0] + " rss_uri")
	}

}

func PollFeed(uri string, timeout int, cr xmlx.CharsetFunc) {
	feed := rss.New(timeout, true, chanHandler, itemHandler)
	feed.Fetch(uri, cr)
}

func chanHandler(feed *rss.Feed, newchannels []*rss.Channel) {
	fmt.Printf("%d new channel(s) in %s\n", len(newchannels), feed.Url)
}

func itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	fmt.Printf("%d new item(s) in %s\n", len(newitems), feed.Url)
	for _, i := range newitems {
		if i.Title != "" {
			fmt.Println(i.Title)
		}
		if i.Description != "" {
			fmt.Println(i.Description)
		}
		fmt.Println("\n")
	}
}

