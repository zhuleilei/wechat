package message

//Video 视频消息
type Video struct {
	CommonToken

	Video struct {
		MediaID     string `xml:"MediaId" json:"media_id"`
		Title       string `xml:"Title,omitempty" json:"title,omitempty"`
		Description string `xml:"Description,omitempty" json:"description,omitempty"`
	} `xml:"Video" json:"video"`
}

//NewVideo 回复图片消息
func NewVideo(mediaID, title, description string) *Video {
	video := new(Video)
	video.Video.MediaID = mediaID
	video.Video.Title = title
	video.Video.Description = description
	return video
}
