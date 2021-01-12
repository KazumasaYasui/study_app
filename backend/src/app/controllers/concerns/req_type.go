package concerns

type BookReq struct {
	Genre       string `json:"genre"`
	Publisher   string `json:"publisher"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	ImageUrl    string `json:"image_url"`
	PublishDate string `json:"publish_date"`
}
