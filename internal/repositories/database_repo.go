package repositories

import (
	"PTIT_TN/internal/repositories/admin"
	"PTIT_TN/internal/repositories/ballot"
	"PTIT_TN/internal/repositories/ballotRoleElect"
	"PTIT_TN/internal/repositories/candidate"
	"PTIT_TN/internal/repositories/election"
	"PTIT_TN/internal/repositories/electionCandidate"
	"PTIT_TN/internal/repositories/electionResult"
	"PTIT_TN/internal/repositories/post"
	"PTIT_TN/internal/repositories/roleElect"
	"PTIT_TN/internal/repositories/user"
	"PTIT_TN/internal/repositories/voter"
	"PTIT_TN/internal/services/db"
	"logur.dev/logur"
)

type DatabaseRepo interface {
	Admin() admin.Repo
	Ballot() ballot.Repo
	BallotRoleElect() ballotRoleElect.Repo
	Candidate() candidate.Repo
	Election() election.Repo
	ElectionCandidate() electionCandidate.Repo
	ElectionResult() electionResult.Repo
	Post() post.Repo
	RoleElect() roleElect.Repo
	User() user.Repo
	Voter() voter.Repo
}
type dbImpl struct {
	admin             admin.Repo
	ballot            ballot.Repo
	ballotRoleElect   ballotRoleElect.Repo
	candidate         candidate.Repo
	election          election.Repo
	electionCandidate electionCandidate.Repo
	electionResult    electionResult.Repo
	roleElect         roleElect.Repo
	post              post.Repo
	user              user.Repo
	voter             voter.Repo
}

func (d dbImpl) Ballot() ballot.Repo {
	return d.ballot
}

func (d dbImpl) BallotRoleElect() ballotRoleElect.Repo {
	return d.ballotRoleElect
}

func (d dbImpl) ElectionCandidate() electionCandidate.Repo {
	return d.electionCandidate
}

func (d dbImpl) ElectionResult() electionResult.Repo {
	return d.electionResult
}

func (d dbImpl) Post() post.Repo {
	return d.post
}

func (d dbImpl) RoleElect() roleElect.Repo {
	return d.roleElect
}
func (d dbImpl) User() user.Repo {
	return d.user
}
func (d dbImpl) Voter() voter.Repo {
	return d.voter
}
func (d dbImpl) Candidate() candidate.Repo {
	return d.candidate
}
func (d dbImpl) Admin() admin.Repo {
	return d.admin
}
func (d dbImpl) Election() election.Repo {
	return d.election
}

func NewDatabaseRepo(logger logur.LoggerFacade, db *db.GormDB) DatabaseRepo {
	return &dbImpl{
		user:      user.New(logger, db),
		voter:     voter.New(logger, db),
		candidate: candidate.New(logger, db),
		admin:     admin.New(logger, db),
	}
}
