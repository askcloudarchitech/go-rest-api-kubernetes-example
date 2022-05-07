package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YoutubeStats struct {
	Subscribers int    `json:"subscribers"`
	ChannelName string `json:"channelName"`
	Views       int    `json:"views"`
	SomeJunk    string `json:"somejunk"`
}

func getChannelStats(k string, channelID string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		ctx := context.Background()
		yts, err := youtube.NewService(ctx, option.WithAPIKey(k))
		if err != nil {
			fmt.Println("failed to create service")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		call := yts.Channels.List([]string{"snippet, contentDetails, statistics"})
		response, err := call.Id(channelID).Do()
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		yt := YoutubeStats{}
		if len(response.Items) > 0 {
			val := response.Items[0]
			yt = YoutubeStats{
				Subscribers: int(val.Statistics.SubscriberCount),
				ChannelName: val.Snippet.Title,
				Views:       int(val.Statistics.ViewCount),
				SomeJunk:    "howdy!",
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(yt); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	}
}
