package discordgo

type voiceStateUpdatePayload struct {
	GuildID   string `json:"guild_id"`
	ChannelID string `json:"channel_id"`
	SelfMute  bool   `json:"self_mute"`
	SelfDeaf  bool   `json:"self_deaf"`
}

func (s *Session) GatewayVoiceStateUpdate(guildID, channelID string, selfMute, selfDeaf bool) error {
	data := voiceStateUpdatePayload{
		GuildID:   guildID,
		ChannelID: channelID,
		SelfMute:  selfMute,
		SelfDeaf:  selfDeaf,
	}
	return s.wsConn.WriteJSON(map[string]any{
		"op": 4,
		"d":  data,
	})
}
