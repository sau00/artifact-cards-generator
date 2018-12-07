package alphabets

type Alphabet struct {
	Filename string `json:"filename"`
	Height int `json:"height"`
	Letters map[string]struct {
		StartPos int    `json:"start_pos"`
		EndPos   int    `json:"end_pos"`
	} `json:"letter"`
}
