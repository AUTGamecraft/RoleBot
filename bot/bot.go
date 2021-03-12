package bot

import (
	"github.com/AUTGamecraft/RoleBot/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session
var err error

func Start() {
	goBot, err = discordgo.New(config.BotPrefix + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(reactionAddHandler)
	goBot.AddHandler(reactionRemoveHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}

func reactionAddHandler(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	if m.ChannelID == config.RolesChannelID {
		for _, emojiRoleID := range config.RoleIDs {
			if m.Emoji.Name == emojiRoleID.Emoji{
				fmt.Println(emojiRoleID.Emoji)
				goBot.GuildMemberRoleAdd(m.GuildID, m.UserID, emojiRoleID.RoleID)
			}
		}
	}
}
func reactionRemoveHandler(s *discordgo.Session, m *discordgo.MessageReactionRemove){
	if m.ChannelID == config.RolesChannelID {
		for _, emojiRoleID := range config.RoleIDs {
			if m.Emoji.Name == emojiRoleID.Emoji{

				fmt.Println(emojiRoleID.Emoji)
				goBot.GuildMemberRoleRemove(m.GuildID, m.UserID, emojiRoleID.RoleID)
			}
		}		
	}
}