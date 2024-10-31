package types

type Game struct {
	Bin               string                `json:"bin"`
	IsDaytime         bool                  `json:"is_daytime,omitempty"`
	ExplanationSeen   bool                  `json:"explanation_seen,omitempty"`
	FirstDayCompleted bool                  `json:"first_day_completed,omitempty"`
	Level             string                `json:"level"`
	CurrentStep       string                `json:"current_step"`
	GroupId           string                `json:"group_id"`
	Status            string                `json:"status"`
	ServerName        string                `json:"server_name,omitempty"`
	ServerAddress     string                `json:"server_ip,omitempty"`
	ServerPort        int                   `json:"server_port,omitempty"`
	StartTime         string                `json:"start_time,omitempty"`
	EndTime           string                `json:"end_time,omitempty"`
	Creator           *Player               `json:"creator,omitempty"`
	Cycles            int                   `json:"cycles,omitempty"`
	Invited           []string              `json:"invited,omitempty"`
	AllVotesSubmitted bool                  `json:"all_votes_submitted,omitempty"`
	Steps             map[string]*Step      `json:"steps,omitempty"`
	Gamers            map[string]*Gamer     `json:"gamers,omitempty"`
	Characters        map[string]*Character `json:"characters,omitempty"`
}
