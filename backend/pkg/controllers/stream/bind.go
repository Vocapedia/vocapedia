package stream

type _createStream struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	Lang            string `json:"lang"`
	TargetLang      string `json:"target_lang"`
	ScheduledAt     string `json:"scheduled_at"`
	Duration        int    `json:"duration"` // in minutes
	MaxParticipants int    `json:"max_participants"`
}
