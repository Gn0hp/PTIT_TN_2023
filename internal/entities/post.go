package entities

type Post struct {
	DefaultModel
	Title   string `json:"title"`
	Content string `json:"content"`
	Note    string `json:"note"`

	CandidateId uint64 `json:"candidate_id"`
}
type PostDto struct {
	Post
}

func (e *PostDto) ToEntity() (interface{}, error) {
	return &e.Post, nil
}
func (e *Post) IsValid() bool {
	return true
}
