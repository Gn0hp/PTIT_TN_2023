package pkg

type AppName string

const (
	APIAppName            = AppName("ElectionAPI")
	MySQLConnectorAppName = AppName("MySQLConnector")
)
const (
	TxTypeCreateElection    = "create_election"
	TxTypeVote              = "vote"
	TxTypeNonce             = "nonce"
	TxTypeSignature         = "get_signature"
	TxTypeTransferOwnership = "transfer_ownership"
	TxTypeRenounceOwnership = "renounce_ownership"
	TxTypeElectionMapping   = "elections_mapping"
	TxTypeElectionToResult  = "election_to_result"
	TxTypeGetElectionResult = "get_election_result"
	TxTypeIsVoted           = "is_voted"
	TxTypeOwner             = "owner"
)

func (name AppName) ToString() string {
	return string(name)
}
