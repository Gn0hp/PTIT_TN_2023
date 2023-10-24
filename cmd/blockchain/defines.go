package main

import (
	election_contract "PTIT_TN/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"logur.dev/logur"
	"math/big"
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
