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
	Gamemode      uint8
	OsuVersion    int32
	BeatmapMD5    string
	PlayerName    string
	MD5           string
	Hit300        uint16
	Hit100        uint16
	Hit50         uint16
	HitGeki       uint16
	HitKatu       uint16
	HitMiss       uint16
	Score         int32
	MaxCombo      uint16
	IsPerfect     bool
	ModsBitSum    int32
	LifeBar       string
	DateTime      int64
	ReplayData    []uint8
	OnlineScoreID int64
}
