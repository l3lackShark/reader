package reader

type OsuDB struct {
	BuildVer          int32
	SongsFolderSize   int32
	IsAccountUnlocked bool
	DateTime          int64
	Nickname          string
	BmInfo            []beatmapInfo
}

type beatmapInfo struct {
	Artist                    string
	ArtistU                   string
	Title                     string
	TitleU                    string
	Creator                   string
	Difficulty                string
	AudioName                 string
	MD5                       string
	Filename                  string
	RankedStatus              int8
	NumHitCircles             int16
	NumSliders                int16
	NumSpinners               int16
	DateTime                  int64
	ApproachRate              float32
	CircleSize                float32
	HpDrain                   float32
	OverallDifficulty         float32
	SliderVelocity            float64 //double
	StarRatingOsu             []starRating
	StarRatingTaiko           []starRating
	StarRatingCtb             []starRating
	StarRatingMania           []starRating
	DrainTime                 int32
	TotalTime                 int32
	PreviewTime               int32
	TimingPoints              []timingPoint
	BeatmapID                 int32
	BeatmapSetID              int32
	ThreadID                  int32
	GradeOsu                  int8
	GradeTaiko                int8
	GradeCtb                  int8
	GradeMania                int8
	LocalOffset               int16
	StackLeniency             float32
	GameMode                  int8
	SongSource                string
	SongTags                  string
	OnlineOffset              int16
	FontTitle                 string //?
	IsUnplayed                bool
	LastPlayed                int64
	IsOsz2                    bool
	FolderFromSongs           string
	LastCheckedAgainstOsuRepo int64
	IsBmSoundIgnored          bool
	IsBmSkinIgnored           bool
	IsBmStoryBoardDisabled    bool
	IsBmVideoDisabled         bool
	IsVisualOverride          bool
	LastClosedEditor          int32
	ManiaScrollSpeed          uint8
}

type starRating struct {
	Fixed0x08  uint8
	BitMods    int32
	Fixed0x0D  uint8
	StarRating float64 //double
}

type timingPoint struct {
	MsPerBeat            float64 //double
	SongOffset           float64 //double
	InheritedTimingPoint bool
}
