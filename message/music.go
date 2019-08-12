package message

//Music 音乐消息
type Music struct {
	CommonToken

	Music struct {
		Title        string `xml:"Title" json:"title"`
		Description  string `xml:"Description" json:"description"`
		MusicURL     string `xml:"MusicUrl" json:"music_url"`
		HQMusicURL   string `xml:"HQMusicUrl" json:"hq_music_url"`
		ThumbMediaID string `xml:"ThumbMediaId" json:"thumb_media_id"`
	} `xml:"Music" json:"music"`
}

//NewMusic  回复音乐消息
func NewMusic(title, description, musicURL, hQMusicURL, thumbMediaID string) *Music {
	music := new(Music)
	music.Music.Title = title
	music.Music.Description = description
	music.Music.MusicURL = musicURL
	music.Music.ThumbMediaID = thumbMediaID
	return music
}
