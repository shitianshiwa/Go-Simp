package subscriber

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/robfig/cron.v2"

	"github.com/JustHumanz/Go-simp/config"
	"github.com/JustHumanz/Go-simp/database"
	"github.com/JustHumanz/Go-simp/engine"
	log "github.com/sirupsen/logrus"
)

var (
	BiliSession string
	Bot         *discordgo.Session
	Data        []database.GroupName
)

func Start(c *cron.Cron) {
	Bot = engine.BotSession
	BiliSession = config.BiliBiliSes
	Data = database.GetGroup()
	if BiliSession == "" {
		log.Error("BiliBili Session not found")
		os.Exit(1)
	}
	c.AddFunc("@every 1h0m0s", CheckYtSubsCount)
	c.AddFunc("@every 0h15m0s", CheckTwFollowCount)
	c.AddFunc("@every 0h20m0s", CheckBiliFollowCount)
	log.Info("Subs&Follow Checker module ready")
}

func SendNude(Embed *discordgo.MessageEmbed, Group database.GroupName) {
	for _, Channel := range Group.GetChannelByGroup() {
		msg, err := Bot.ChannelMessageSendEmbed(Channel, Embed)
		if err != nil {
			log.Error(msg, err)
		}
	}
}

type Subs struct {
	Success bool   `json:"success"`
	Service string `json:"service"`
	T       int64  `json:"t"`
	Data    struct {
		LvIdentifier string `json:"lv_identifier"`
		Subscribers  int    `json:"subscribers"`
		Videos       int    `json:"videos"`
		Views        int    `json:"views"`
	} `json:"data"`
}

type BiliBiliStat struct {
	LikeView LikeView
	Follow   BiliFollow
	Videos   int
}

type LikeView struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Archive struct {
			View int `json:"view"`
		} `json:"archive"`
		Article struct {
			View int `json:"view"`
		} `json:"article"`
		Likes int `json:"likes"`
	} `json:"data"`
}

type BiliFollow struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Mid       int `json:"mid"`
		Following int `json:"following"`
		Whisper   int `json:"whisper"`
		Black     int `json:"black"`
		Follower  int `json:"follower"`
	} `json:"data"`
}
