package ballotRoleElect

type StatBreByCandidateDbFind struct {
	VoterId     uint64 `json:"voter_id"`
	VoterName   string `json:"voter_name"`
	CccdId      string `json:"cccd_id"`
	BallotId    uint64 `json:"ballot_id"`
	CandidateId uint64 `json:"candidate_id"`
	VoteTime    string `json:"vote_time"`
}

type StatByBallotResponse struct {
	VoterName     string `json:"voter_name"`
	RoleName      string `json:"role_name"`
	CandidateName string `json:"candidate_name"`
	CandidateCccd string `json:"candidate_cccd"`
}
