package scraper

import (
	"context"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/swayamduhan/rssagg-go/internal/db"
	"github.com/swayamduhan/rssagg-go/internal/models"
	"github.com/swayamduhan/rssagg-go/internal/utils"
)

func InitScraper(concurrency int, interval time.Duration) {

	log.Printf("Collecting feeds every %v on %d goroutines...\n", interval, concurrency)
	ctx := context.Background()

	ticker := time.NewTicker(interval)

	for ; ; <-ticker.C {

		feeds, err := utils.Queries.GetFeedsToFetch(ctx, int32(concurrency))
		if err != nil {
			log.Println("unable to get feeds to scrape!")
		}
		log.Printf("Found %v feeds to fetch!\n", len(feeds))

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(ctx, wg, feed)
		}
		wg.Wait()
	}

}

func scrapeFeed(ctx context.Context,wg *sync.WaitGroup, feed db.Feed){
	defer wg.Done()

	// mark feed fetched
	_, err := utils.Queries.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		log.Printf("error marking feed %v as fetched\n", feed.ID)
		return
	}

	// get data into RSS struct
	rssData, err := fetchFeed(feed.Url)
	if err != nil {
		return
	}

	// create post
	
	for _, post := range rssData.Channel.Items {
		_, err := utils.Queries.CreatePost(ctx, db.CreatePostParams{
			FeedID: feed.ID,
			Title: post.Title,
			Link: post.Link,
			Description: pgtype.Text{String: post.Description, Valid: true},
			PublishedAt: pgtype.Timestamp{Time: post.PublishDate.Time, Valid: true},
		})

		if err != nil {
			log.Printf("couldn't create post : %v\n", err)
			continue
		}
	}

	log.Printf("Scraped feed %v, %v posts found", feed.Name, len(rssData.Channel.Items))
}


func fetchFeed(feedUrl string) (*models.RSS, error) {
	httpClient := http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := httpClient.Get(feedUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	
	var RSS models.RSS
	err = xml.Unmarshal(data, &RSS)
	if err != nil {
		return nil, err
	}

	return &RSS, nil
	
}
