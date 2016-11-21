package main

import (
	"github.com/bwmarrin/Discordgo"
	"fmt"
	"encoding/binary"
	"io"
	"os"
	"strings"
	"time"
)

var airhorn = make([][]byte, 0)
var docomo docomoTTS

func main() {
	docomo = docomoTTS{apiKey: "504f505564754b747a314b69736f2f4466532e484233314e7849644e3746344e754343736b665a4a7a3442"}
	discord, err := discordgo.New("Bot " + "MjUwMjEyODAyODM5NDQ1NTA0.CxRkWg.MDUAqixO-02P6qrqSl_BIGYuOlA")
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}
	loadSound("airhorn.dca")
	discord.AddHandler(ready)
	discord.AddHandler(messageReserve)
	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening Discord session: ", err)
	}
	fmt.Println("Airhorn is now running.  Press CTRL-C to exit.")
	<-make(chan struct{})
}

func loadSound(soundFile string) error {
	file, err := os.Open(soundFile)
	if err != nil {
		fmt.Println("Error opening dca file :", err)
		return err
	}
	var opuslen int16
	for {
		err = binary.Read(file, binary.LittleEndian, &opuslen)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return nil
		}
		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)
		if err != nil {
			fmt.Println("Error reading from dca file :", err)
			return err
		}
		airhorn = append(airhorn, InBuf)
	}
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "!airhorn")
}

func messageReserve(session *discordgo.Session, message *discordgo.MessageCreate) {
	channel, err := session.State.Channel(message.ChannelID)
	guild, err := session.State.Guild(channel.GuildID)
	if strings.HasPrefix(message.Content, "!airhorn") || strings.HasPrefix(message.Content, "!AIRHORN") {
		for _, voiceChannel := range guild.VoiceStates {
			if voiceChannel.UserID == message.Author.ID {
				err = playSound(session, guild.ID, voiceChannel.ChannelID)
				if err != nil {
					fmt.Println("Error playing sound:", err)
				}
				return
			}
		}
	}
}

func playSound(s *discordgo.Session, guildID, channelID string) (err error) {
	voiceChannel, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		return err
	}
	time.Sleep(250 * time.Millisecond)
	voiceChannel.Speaking(true)
	for _, buff := range airhorn {
		voiceChannel.OpusSend <- buff
	}
	voiceChannel.Speaking(false)
	time.Sleep(250 * time.Millisecond)
	voiceChannel.Disconnect()
	return nil
}