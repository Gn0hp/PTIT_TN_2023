package entities

type Post struct {
	DefaultModel
	Title string `json:"title"`
	Content string `json:"content"`
	
	CandidateId uint64 `json:"candidate_id"`
}
type PostDto struct {
	Post
}