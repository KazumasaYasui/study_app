package concerns

type BookReq struct {
	Genre       string `json:"genre"`
	Publisher   string `json:"publisher"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	ImageUrl    string `json:"image_url"`
	PublishDate string `json:"publish_date"`
}

type BookMyInfoReq struct {
	Id           string  `json:"id,omitempty"`
	Status       string  `json:"status"`
	ProgressRate float64 `json:"progress_rate"`
	Urgency      int     `json:"urgency"`
	Priority     int     `json:"priority"`
}

type SignupReq struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
