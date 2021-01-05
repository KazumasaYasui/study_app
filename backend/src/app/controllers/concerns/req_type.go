package concerns

type TmpBookReq struct {
	Genre     string `json:"genre"`
	Publisher string `json:"publisher"`
	Author    string `json:"author"`
}
