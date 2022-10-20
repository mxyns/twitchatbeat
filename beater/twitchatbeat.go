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
	done   chan struct{}
	events []beat.Event
	config config.Config
	client beat.Client
}

// New creates an instance of twitchatbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &twitchatbeat{
		done:   make(chan struct{}),
		events: make([]beat.Event, c.QueueCapacity),
		config: c,
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
			bt.client.PublishAll(bt.events)
			bt.ClearEvents()
		}
	}
}

func RunBot(bt *twitchatbeat, b *beat.Beat) {

	cfg := config.ConvertConfiguration(&bt.config)

	elasticLogger := NewElasticLogger(b, bt)
	helixClient := helix.NewClient(cfg.ClientID, cfg.ClientSecret)
	go helixClient.StartRefreshTokenRoutine()

	newBot := bot.NewBot(cfg, &helixClient, &elasticLogger)
	newBot.Connect()
}

func (bt *twitchatbeat) ClearEvents() {
	bt.events = make([]beat.Event, bt.config.QueueCapacity)
}

func (bt *twitchatbeat) AddEvent(event beat.Event) {
	bt.events = append(bt.events, event)
}

// Stop stops twitchatbeat.
func (bt *twitchatbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
