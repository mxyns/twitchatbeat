package beater

import (
	"fmt"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/mxyns/justlog/bot"
	"github.com/mxyns/justlog/helix"
	"time"

	"github.com/mxyns/twitchatbeat/config"
)

// twitchatbeat configuration.
type twitchatbeat struct {
	done            chan struct{}
	events          []beat.Event
	config          config.Config
	client          beat.Client
	helixClient     helix.Client
	helixClientInit bool
}

// New creates an instance of twitchatbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &twitchatbeat{
		done:            make(chan struct{}),
		events:          make([]beat.Event, 0, c.QueueCapacity),
		config:          c,
		helixClientInit: false,
	}
	return bt, nil
}

// Run starts twitchatbeat.
func (bt *twitchatbeat) Run(b *beat.Beat) error {
	logp.Info("twitchatbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	go RunBot(bt, b)

	ticker := time.NewTicker(bt.config.Period)

	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
			bt.PushStreamStatuses(b)
			bt.client.PublishAll(bt.events)
			logp.Info("Pushed %v events", len(bt.events))
			bt.ClearEvents()
		}
	}
}

func RunBot(bt *twitchatbeat, b *beat.Beat) {

	cfg := config.ConvertConfiguration(&bt.config)

	elasticLogger := NewElasticLogger(b, bt)
	bt.helixClient = helix.NewClient(cfg.ClientID, cfg.ClientSecret)
	go bt.helixClient.StartRefreshTokenRoutine()
	bt.helixClientInit = true

	newBot := bot.NewBot(cfg, &bt.helixClient, &elasticLogger)
	newBot.Connect()
}

func (bt *twitchatbeat) ClearEvents() {
	bt.events = make([]beat.Event, 0, bt.config.QueueCapacity)
}

func (bt *twitchatbeat) AddEvent(event beat.Event) {
	bt.events = append(bt.events, event)
}

// Stop stops twitchatbeat.
func (bt *twitchatbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

func (bt *twitchatbeat) GetStreamStatuses() map[string]helix.ChannelData {
	if !bt.helixClientInit {
		return nil
	}

	statuses, err := bt.helixClient.GetChannelInformationByChannelIds(bt.config.Channels)
	if err != nil {
		logp.Error(err)
		return nil
	}

	return statuses
}

func (bt *twitchatbeat) PushStreamStatuses(b *beat.Beat) {

	statuses := bt.GetStreamStatuses()
	logp.Info("Statuses: %v", statuses)
	for _, status := range statuses {
		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"type":                               b.Info.Name,
				"event_type":                         []string{"ChannelStatus", "Channel"},
				"channel.status.channel_id":          status.BroadcasterID,
				"channel.status.channel_name":        status.BroadcasterName,
				"channel.status.channel_language":    status.BroadcasterLanguage,
				"channel.status.stream.title":        status.Title,
				"channel.status.stream.game_id":      status.GameID,
				"channel.status.stream.game_name":    status.GameName,
				"channel.status.stream.status_valid": status.StreamStatus,
				"channel.status.stream.is_live":      status.IsLive,
				"channel.status.stream.started_at":   status.StartedAt.Time,
				"channel.status.stream.delay":        status.Delay,
			},
		}

		bt.events = append(bt.events, event)
		logp.Info("Status Report Event: %v", event)
	}
}
