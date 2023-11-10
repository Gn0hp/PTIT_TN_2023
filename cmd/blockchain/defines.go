package main

import (
	election_contract "PTIT_TN/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"logur.dev/logur"
	"math/big"
)

const (
	TxTypeCreateElection    = "create_election"
	TxTypeNonce             = "nonce"
	TxTypeSignature         = "get_signature"
	TxTypeVote              = "vote"
	TxTypeTransferOwnership = "transfer_ownership"
	TxTypeRenounceOwnership = "renounce_ownership"
	TxTypeElectionMapping   = "elections_mapping"
	TxTypeElectionToResult  = "election_to_result"
	TxTypeGetElectionResult = "get_election_result"
	TxTypeIsVoted           = "is_voted"
	TxTypeOwner             = "owner"
)

type BcElection struct {
	Id            *big.Int
	StartDate     *big.Int
	Duration      *big.Int
	NumCandidates *big.Int
}

type ContractInstance struct {
	ChainID           *big.Int
	AccountBoundWrite *bind.TransactOpts
	AccountBoundRead  *bind.CallOpts
	Instance          *election_contract.ElectionContract
	Logger            logur.LoggerFacade
}
