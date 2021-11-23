package reader

type ScoresDB struct {
	OsuVersion    int32
	BeatmapScores []beatmapScore
}

type beatmapScore struct {
	BeatmapMD5 string
	Scores     []Score
}

type Score struct {
	Gamemode      uint8   `json:"gamemode"`
	OsuVersion    int32   `json:"osuVersion"`
	BeatmapMD5    string  `json:"beatmapMD5"`
	PlayerName    string  `json:"playerName"`
	MD5           string  `json:"scoreMD5"`
	Hit300        uint16  `json:"hit300"`
	Hit100        uint16  `json:"hit100"`
	Hit50         uint16  `json:"hit50"`
	HitGeki       uint16  `json:"hitGeki"`
	HitKatu       uint16  `json:"hitKatu"`
	HitMiss       uint16  `json:"hitMiss"`
	Score         int32   `json:"score"`
	MaxCombo      uint16  `json:"maxCombo"`
	IsPerfect     bool    `json:"isPerfect"`
	ModsBitSum    int32   `json:"modsBitSum"`
	LifeBar       string  `json:"lifeBar"`
	DateTime      int64   `json:"dateTime"`
	ReplayData    []uint8 `json:"replayData"`
	OnlineScoreID int64   `json:"onlineScoreID"`
}

