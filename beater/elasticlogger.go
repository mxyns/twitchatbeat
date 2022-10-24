package beater

import (
	"errors"
	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/gempir/go-twitch-irc/v3"
	"github.com/mxyns/justlog/filelog"
	"time"
)

type ElasticLogger struct {
	b  *beat.Beat
	bt *twitchatbeat
}

func NewElasticLogger(b *beat.Beat, bt *twitchatbeat) ElasticLogger {
	return ElasticLogger{
		b:  b,
		bt: bt,
	}
}

func (l *ElasticLogger) LogPrivateMessageForUser(user twitch.User, message twitch.PrivateMessage) error {

	event := beat.Event{
		Timestamp: time.Now(),
		Fields: common.MapStr{
			"type":              l.b.Info.Name,
			"event_type":        []string{"PrivateMessage", "User"},
			"user.id":           user.ID,
			"user.name":         user.Name,
			"user.display_name": user.DisplayName,
			"user.color":        user.Color,
			"user.badges":       user.Badges,
			"msg.raw":           message.Raw,
			"msg.raw_json":      message.Raw,
			"msg.type":          message.Type,
			"msg.raw_type":      message.RawType,
			"msg.tags":          message.Tags,
			"msg.text":          message.Message,
			"msg.channel":       message.Channel,
			"msg.room_id":       message.RoomID,
			"msg.id":            message.ID,
			"msg.time":          message.Time,
			"msg.emotes":        message.Emotes,
			"msg.bits":          message.Bits,
			"msg.action":        message.Action,
			"msg.first_message": message.FirstMessage,
		},
	}
	l.bt.AddMessageEvent(l.b, event)
	logp.Info("Event queued")

	return nil
}

func (l *ElasticLogger) LogClearchatMessageForUser(userID string, message twitch.ClearChatMessage) error {

	event := beat.Event{
		Timestamp: time.Now(),
		Fields: common.MapStr{
			"type":                 l.b.Info.Name,
			"event_type":           []string{"ClearchatMessage", "User"},
			"user.id":              userID,
			"msg.raw":              message.Raw,
			"msg.raw_json":         message.Raw,
			"msg.type":             message.Type,
			"msg.raw_type":         message.RawType,
			"msg.tags":             message.Tags,
			"msg.text":             message.Message,
			"msg.channel":          message.Channel,
			"msg.room_id":          message.RoomID,
			"msg.time":             message.Time,
			"msg.ban_duration":     message.BanDuration,
			"msg.target_user_id":   message.TargetUserID,
			"msg.target_user_name": message.TargetUsername,
		},
	}
	l.bt.AddMessageEvent(l.b, event)
	logp.Info("Event queued")
	return nil
}

func (l *ElasticLogger) LogUserNoticeMessageForUser(userID string, message twitch.UserNoticeMessage) error {

	event := beat.Event{
		Timestamp: time.Now(),
		Fields: common.MapStr{
			"type":              l.b.Info.Name,
			"event_type":        []string{"UserNotice", "User"},
			"user.id":           message.User.ID,
			"user.name":         message.User.Name,
			"user.display_name": message.User.DisplayName,
			"user.color":        message.User.Color,
			"user.badges":       message.User.Badges,
			"msg.raw":           message.Raw,
			"msg.raw_json":      message.Raw,
			"msg.type":          message.Type,
			"msg.raw_type":      message.RawType,
			"msg.tags":          message.Tags,
			"msg.text":          message.Message,
			"msg.channel":       message.Channel,
			"msg.room_id":       message.RoomID,
			"msg.id":            message.ID,
			"msg.time":          message.Time,
			"msg.emotes":        message.Emotes,
			"msg.msg_id":        message.MsgID,
			"msg.params":        message.MsgParams,
			"msg.sys_msg":       message.SystemMsg,
		},
	}
	l.bt.AddMessageEvent(l.b, event)
	logp.Info("Event queued")
	return nil
}

func (l *ElasticLogger) LogPrivateMessageForChannel(message twitch.PrivateMessage) error {

	event := beat.Event{
		Timestamp: time.Now(),
		Fields: common.MapStr{
			"type":              l.b.Info.Name,
			"event_type":        []string{"PrivateMessage", "Channel"},
			"user.id":           message.User.ID,
			"user.name":         message.User.Name,
			"user.display_name": message.User.DisplayName,
			"user.color":        message.User.Color,
			"user.badges":       message.User.Badges,
			"msg.raw":           message.Raw,
			"msg.raw_json":      message.Raw,
			"msg.type":          message.Type,
			"msg.raw_type":      message.RawType,
			"msg.tags":          message.Tags,
			"msg.text":          message.Message,
			"msg.channel":       message.Channel,
			"msg.room_id":       message.RoomID,
			"msg.id":            message.ID,
			"msg.time":          message.Time,
			"msg.emotes":        message.Emotes,
			"msg.bits":          message.Bits,
			"msg.action":        message.Action,
			"msg.first_message": message.FirstMessage,
		},
	}
	l.bt.AddMessageEvent(l.b, event)
	logp.Info("Event queued")

	return nil
}

func (l *ElasticLogger) LogClearchatMessageForChannel(message twitch.ClearChatMessage) error {

	event := beat.Event{
		Timestamp: time.Now(),
		Fields: common.MapStr{
			"type":                 l.b.Info.Name,
			"event_type":           []string{"ClearchatMessage", "Channel"},
			"user.id":              message.TargetUserID,
			"msg.raw":              message.Raw,
			"msg.raw_json":         message.Raw,
			"msg.type":             message.Type,
			"msg.raw_type":         message.RawType,
			"msg.tags":             message.Tags,
			"msg.text":             message.Message,
			"msg.channel":          message.Channel,
			"msg.room_id":          message.RoomID,
			"msg.time":             message.Time,
			"msg.ban_duration":     message.BanDuration,
			"msg.target_user_id":   message.TargetUserID,
			"msg.target_user_name": message.TargetUsername,
		},
	}
	l.bt.AddMessageEvent(l.b, event)
	logp.Info("Event queued")
	return nil
}

func (l *ElasticLogger) LogUserNoticeMessageForChannel(message twitch.UserNoticeMessage) error {

	event := beat.Event{
		Timestamp: time.Now(),
		Fields: common.MapStr{
			"type":              l.b.Info.Name,
			"event_type":        []string{"UserNotice", "Channel"},
			"user.id":           message.User.ID,
			"user.name":         message.User.Name,
			"user.display_name": message.User.DisplayName,
			"user.color":        message.User.Color,
			"user.badges":       message.User.Badges,
			"msg.raw":           message.Raw,
			"msg.raw_json":      message.Raw,
			"msg.type":          message.Type,
			"msg.raw_type":      message.RawType,
			"msg.tags":          message.Tags,
			"msg.text":          message.Message,
			"msg.channel":       message.Channel,
			"msg.room_id":       message.RoomID,
			"msg.id":            message.ID,
			"msg.time":          message.Time,
			"msg.emotes":        message.Emotes,
			"msg.msg_id":        message.MsgID,
			"msg.params":        message.MsgParams,
			"msg.sys_msg":       message.SystemMsg,
		},
	}
	l.bt.AddMessageEvent(l.b, event)
	logp.Info("Event queued")
	return nil
}

func (l *ElasticLogger) GetLastLogYearAndMonthForUser(channelID, userID string) (int, int, error) {
	return 0, 0, errors.New("NotImplementedError")
}

func (l *ElasticLogger) GetAvailableLogsForUser(channelID, userID string) ([]filelog.UserLogFile, error) {
	return nil, errors.New("NotImplementedError")
}

// ReadLogForUser fetch logs
func (l *ElasticLogger) ReadLogForUser(channelID, userID string, year string, month string) ([]string, error) {
	return nil, errors.New("NotImplementedError")
}

func (l *ElasticLogger) ReadRandomMessageForUser(channelID, userID string) (string, error) {
	return "NotImplementedError", errors.New("NotImplementedError")
}

func (l *ElasticLogger) ReadLogForChannel(channelID string, year int, month int, day int) ([]string, error) {
	return nil, errors.New("NotImplementedError")
}

func (l *ElasticLogger) ReadRandomMessageForChannel(channelID string) (string, error) {
	return "NotImplementedError", errors.New("NotImplementedError")
}
