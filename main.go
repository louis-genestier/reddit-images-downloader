package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Posts struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash  string `json:"modhash"`
		Dist     int    `json:"dist"`
		Children []struct {
			Kind string `json:"kind"`
			Data struct {
				ApprovedAtUtc              interface{}   `json:"approved_at_utc"`
				Subreddit                  string        `json:"subreddit"`
				Selftext                   string        `json:"selftext"`
				AuthorFullname             string        `json:"author_fullname"`
				Saved                      bool          `json:"saved"`
				ModReasonTitle             interface{}   `json:"mod_reason_title"`
				Gilded                     int           `json:"gilded"`
				Clicked                    bool          `json:"clicked"`
				Title                      string        `json:"title"`
				LinkFlairRichtext          []interface{} `json:"link_flair_richtext"`
				SubredditNamePrefixed      string        `json:"subreddit_name_prefixed"`
				Hidden                     bool          `json:"hidden"`
				Pwls                       int           `json:"pwls"`
				LinkFlairCSSClass          interface{}   `json:"link_flair_css_class"`
				Downs                      int           `json:"downs"`
				ThumbnailHeight            int           `json:"thumbnail_height"`
				TopAwardedType             interface{}   `json:"top_awarded_type"`
				HideScore                  bool          `json:"hide_score"`
				Name                       string        `json:"name"`
				Quarantine                 bool          `json:"quarantine"`
				LinkFlairTextColor         string        `json:"link_flair_text_color"`
				UpvoteRatio                float64       `json:"upvote_ratio"`
				AuthorFlairBackgroundColor interface{}   `json:"author_flair_background_color"`
				SubredditType              string        `json:"subreddit_type"`
				Ups                        int           `json:"ups"`
				TotalAwardsReceived        int           `json:"total_awards_received"`
				MediaEmbed                 struct {
				} `json:"media_embed"`
				ThumbnailWidth        int           `json:"thumbnail_width"`
				AuthorFlairTemplateID interface{}   `json:"author_flair_template_id"`
				IsOriginalContent     bool          `json:"is_original_content"`
				UserReports           []interface{} `json:"user_reports"`
				SecureMedia           interface{}   `json:"secure_media"`
				IsRedditMediaDomain   bool          `json:"is_reddit_media_domain"`
				IsMeta                bool          `json:"is_meta"`
				Category              interface{}   `json:"category"`
				SecureMediaEmbed      struct {
				} `json:"secure_media_embed"`
				LinkFlairText       interface{}   `json:"link_flair_text"`
				CanModPost          bool          `json:"can_mod_post"`
				Score               int           `json:"score"`
				ApprovedBy          interface{}   `json:"approved_by"`
				AuthorPremium       bool          `json:"author_premium"`
				Thumbnail           string        `json:"thumbnail"`
				AuthorFlairCSSClass interface{}   `json:"author_flair_css_class"`
				AuthorFlairRichtext []interface{} `json:"author_flair_richtext"`
				Gildings            struct {
					Gid2 int `json:"gid_2"`
				} `json:"gildings"`
				PostHint            string      `json:"post_hint"`
				ContentCategories   interface{} `json:"content_categories"`
				IsSelf              bool        `json:"is_self"`
				ModNote             interface{} `json:"mod_note"`
				Created             float64     `json:"created"`
				LinkFlairType       string      `json:"link_flair_type"`
				Wls                 int         `json:"wls"`
				RemovedByCategory   interface{} `json:"removed_by_category"`
				BannedBy            interface{} `json:"banned_by"`
				AuthorFlairType     string      `json:"author_flair_type"`
				Domain              string      `json:"domain"`
				AllowLiveComments   bool        `json:"allow_live_comments"`
				SelftextHTML        interface{} `json:"selftext_html"`
				Likes               interface{} `json:"likes"`
				SuggestedSort       interface{} `json:"suggested_sort"`
				BannedAtUtc         interface{} `json:"banned_at_utc"`
				URLOverriddenByDest string      `json:"url_overridden_by_dest"`
				ViewCount           interface{} `json:"view_count"`
				Archived            bool        `json:"archived"`
				NoFollow            bool        `json:"no_follow"`
				IsCrosspostable     bool        `json:"is_crosspostable"`
				Pinned              bool        `json:"pinned"`
				Over18              bool        `json:"over_18"`
				Preview             struct {
					Images []struct {
						Source struct {
							URL    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"source"`
						Resolutions []struct {
							URL    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"resolutions"`
						Variants struct {
						} `json:"variants"`
						ID string `json:"id"`
					} `json:"images"`
					Enabled bool `json:"enabled"`
				} `json:"preview"`
				AllAwardings []struct {
					GiverCoinReward          interface{} `json:"giver_coin_reward"`
					SubredditID              interface{} `json:"subreddit_id"`
					IsNew                    bool        `json:"is_new"`
					DaysOfDripExtension      int         `json:"days_of_drip_extension"`
					CoinPrice                int         `json:"coin_price"`
					ID                       string      `json:"id"`
					PennyDonate              interface{} `json:"penny_donate"`
					AwardSubType             string      `json:"award_sub_type"`
					CoinReward               int         `json:"coin_reward"`
					IconURL                  string      `json:"icon_url"`
					DaysOfPremium            int         `json:"days_of_premium"`
					TiersByRequiredAwardings interface{} `json:"tiers_by_required_awardings"`
					ResizedIcons             []struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"resized_icons"`
					IconWidth                        int         `json:"icon_width"`
					StaticIconWidth                  int         `json:"static_icon_width"`
					StartDate                        interface{} `json:"start_date"`
					IsEnabled                        bool        `json:"is_enabled"`
					AwardingsRequiredToGrantBenefits interface{} `json:"awardings_required_to_grant_benefits"`
					Description                      string      `json:"description"`
					EndDate                          interface{} `json:"end_date"`
					SubredditCoinReward              int         `json:"subreddit_coin_reward"`
					Count                            int         `json:"count"`
					StaticIconHeight                 int         `json:"static_icon_height"`
					Name                             string      `json:"name"`
					ResizedStaticIcons               []struct {
						URL    string `json:"url"`
						Width  int    `json:"width"`
						Height int    `json:"height"`
					} `json:"resized_static_icons"`
					IconFormat    interface{} `json:"icon_format"`
					IconHeight    int         `json:"icon_height"`
					PennyPrice    interface{} `json:"penny_price"`
					AwardType     string      `json:"award_type"`
					StaticIconURL string      `json:"static_icon_url"`
				} `json:"all_awardings"`
				Awarders                 []interface{} `json:"awarders"`
				MediaOnly                bool          `json:"media_only"`
				CanGild                  bool          `json:"can_gild"`
				Spoiler                  bool          `json:"spoiler"`
				Locked                   bool          `json:"locked"`
				AuthorFlairText          interface{}   `json:"author_flair_text"`
				TreatmentTags            []interface{} `json:"treatment_tags"`
				Visited                  bool          `json:"visited"`
				RemovedBy                interface{}   `json:"removed_by"`
				NumReports               interface{}   `json:"num_reports"`
				Distinguished            interface{}   `json:"distinguished"`
				SubredditID              string        `json:"subreddit_id"`
				ModReasonBy              interface{}   `json:"mod_reason_by"`
				RemovalReason            interface{}   `json:"removal_reason"`
				LinkFlairBackgroundColor string        `json:"link_flair_background_color"`
				ID                       string        `json:"id"`
				IsRobotIndexable         bool          `json:"is_robot_indexable"`
				ReportReasons            interface{}   `json:"report_reasons"`
				Author                   string        `json:"author"`
				DiscussionType           interface{}   `json:"discussion_type"`
				NumComments              int           `json:"num_comments"`
				SendReplies              bool          `json:"send_replies"`
				WhitelistStatus          string        `json:"whitelist_status"`
				ContestMode              bool          `json:"contest_mode"`
				ModReports               []interface{} `json:"mod_reports"`
				AuthorPatreonFlair       bool          `json:"author_patreon_flair"`
				AuthorFlairTextColor     interface{}   `json:"author_flair_text_color"`
				Permalink                string        `json:"permalink"`
				ParentWhitelistStatus    string        `json:"parent_whitelist_status"`
				Stickied                 bool          `json:"stickied"`
				URL                      string        `json:"url"`
				SubredditSubscribers     int           `json:"subreddit_subscribers"`
				CreatedUtc               float64       `json:"created_utc"`
				NumCrossposts            int           `json:"num_crossposts"`
				Media                    interface{}   `json:"media"`
				IsVideo                  bool          `json:"is_video"`
			} `json:"data"`
		} `json:"children"`
		After  string      `json:"after"`
		Before interface{} `json:"before"`
	} `json:"data"`
}

// sub set the subreddit
var sub string
// limit set the max number of posts (100 maximum)
var limit int
var wg sync.WaitGroup

func init() {
	flag.StringVar(&sub, "s", "all", "name of the subreddit to get images from")
	flag.IntVar(&limit, "l", 100, "number of posts to get")
	flag.Parse()
}

func main() {
	var p Posts
	url := fmt.Sprintf("https://www.reddit.com/r/%s/top/.json?t=all&limit=%d", sub, limit)

	body := getResponse(url)
	defer body.Close()

	err := json.NewDecoder(body).Decode(&p)

	if err != nil {
		log.Fatalln("can't convert json to struct: ", err.Error())
	}

	for _, post := range p.Data.Children {
		if len(post.Data.Preview.Images) >= 1 {
			wg.Add(1)
			imageUrl := strings.ReplaceAll(post.Data.Preview.Images[0].Source.URL, "amp;", "")
			go downloadImage(imageUrl, post.Data.Subreddit, post.Data.Title)
		}
	}

	wg.Wait()
}

// downloadImage is used to download an image from a give imageUrl
func downloadImage(imageUrl string, subreddit string, title string) {

	body := getResponse(imageUrl)

	filename := fmt.Sprintf("%s_%d.jpg", subreddit, rand.Int())
	path := filepath.Join(".", "images", subreddit)

	_ = os.MkdirAll(path, os.ModePerm)

	file, err := os.Create(fmt.Sprintf("%s/%s", path, filename))

	if err != nil {
		log.Println("can't create the image: ", err.Error())
		return
	}
	_, err = io.Copy(file, body)

	if err != nil {
		log.Println("can't write into the file: ", err.Error())
		return
	}

	_ = file.Close()
	fmt.Printf("Image from post %s on /r/%s downloaded.\n", title, subreddit)
	body.Close()

	wg.Done()
}

// getResponse returns the body of a http.Request and handle the errors checks
func getResponse(url string) io.ReadCloser {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalln("error while creating the request: ", err.Error())
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Android 4.4; Tablet; rv:41.0) Gecko/41.0 Firefox/41.0")
	resp, err := client.Do(req)

	if err != nil {
		log.Panicln("error while requesting reddit api: ", err.Error())
	}

	if resp.StatusCode != 200 {
		log.Panicf("response status code is not 200 but %d", resp.StatusCode)
	}

	return resp.Body
}