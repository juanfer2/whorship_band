package read_image_repository

type ImageRedResponse struct {
	Ext      string `json:"ext"`
	FileName string `json:"fileName"`
	Lang     string `json:"lang"`
	Text     string `json:"text"`
	UUID     string `json:"uuid"`
}
