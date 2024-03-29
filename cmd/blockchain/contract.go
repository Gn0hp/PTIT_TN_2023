package main

import (
	election_contract "PTIT_TN/contracts"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (i *ContractInstance) TransferOwnership(newOwner string) (*types.Transaction, error) {
	res, err := i.Instance.TransferOwnership(i.AccountBoundWrite, common.HexToAddress(newOwner))
	if err != nil {
		i.Logger.Error(fmt.Sprintf("[Contract Instance] Transfer Ownership failed, detail: %v", err))
		return nil, err
	}
	i.Logger.Info(fmt.Sprintf("[Contract Instance] Transfer Ownership success, detail: %v", res))
	return res, nil
}
func (i *ContractInstance) RenounceOwnership() (*types.Transaction, error) {
	res, err := i.Instance.RenounceOwnership(i.AccountBoundWrite)
	if err != nil {
		i.Logger.Error(fmt.Sprintf("[Contract Instance] Renounce Ownership failed, detail: %v", err))
		return nil, err
	}
	i.Logger.Info(fmt.Sprintf("[Contract Instance] Renounce Ownership success, detail: %v", res))
	return res, nil
}
func (i *ContractInstance) GetNonce() (*big.Int, error) {
	res, err := i.Instance.Nonce(i.AccountBoundRead)
	if err != nil {
		i.Logger.Error(fmt.Sprintf("[Contract Instance] Get Nonce failed, detail: %v", err))
		return nil, err
	}
	i.Logger.Info(fmt.Sprintf("[Contract Instance] Get Nonce success, detail: %v", res))
	return res, nil
}
func (i *ContractInstance) GetSignatureString() ([]byte, error) {
	res, err := i.Instance.GetSignature(i.AccountBoundRead)
	if err != nil {
		i.Logger.Error(fmt.Sprintf("[Contract Instance] Get Signature failed, detail: %v", err))
		return nil, err
	}
	i.Logger.Info(fmt.Sprintf("[Contract Instance] Get Signature success, detail: %v", res))
	return res, nil
}
func (i *ContractInstance) CreateNewElection(startDate, duration, numCandidate uint64) (*types.Transaction, error) {
	res, err := i.Instance.CreateNewElection(i.AccountBoundWrite,
		big.NewInt(int64(startDate)),
		big.NewInt(int64(duration)),
		big.NewInt(int64(numCandidate)),
	)
	if err != nil {
		i.Logger.Error(fmt.Sprintf("[Contract Instance] Create New Election failed, detail: %v", err))
		return nil, err
	}
	i.Logger.Info(fmt.Sprintf("[Contract Instance] Create New Election success, detail: %v", res))
	return res, nil

}
func convertInt64ToBigIntArray(arr []int64) []*big.Int {
	var res []*big.Int
	for _, v := range arr {
		res = append(res, big.NewInt(v))
	}
	return res
}
func (i *ContractInstance) Vote(electionId int64, voterId string, choiceValue []int64) (*types.Transaction, error) {
	res, err := i.Instance.Vote(i.AccountBoundWrite,
		big.NewInt(electionId),
		voterId,
		convertInt64ToBigIntArray(choiceValue))
	if err != nil {
		i.Logger.Error(fmt.Sprintf("[Contract Instance] Vote failed, detail: %v", err))
		return nil, err
	}
	i.Logger.Info(fmt.Sprintf("[Contract Instance] Vote success, detail: %v", res))
	return res, nil
}
func (i *ContractInstance) ElectionsMapping(electionId int64) (BcElection, error) {
	res, err := i.Instance.Elections(i.AccountBoundRead, big.NewInt(electionId))
	if err != nil {
		i.Logger.Error(fmt.Sprintf("[Contract Instance] Get Owner failed, detail: %v", err))
		return BcElection{}, err
	}
	i.Logger.Info(fmt.Sprintf("[Contract Instance] Get ElectionsMapping success, detail: %v", res))
	return res, nil
}
func (i *ContractInstance) ElectionToResult(electionId int64, choiceValue int64) (*big.Int, error) {
	res, err := i.Instance.ElectionToResult(i.AccountBoundRead, big.NewInt(electionId), big.NewInt(choiceValue))
	if err != nil {
		i.Logger.Error(fmt.Sprintf("[Contract Instance] Get Owner failed, detail: %v", err))
		return nil, err
	}
	i.Logger.Info(fmt.Sprintf("[Contract Instance] Get ElectionToResult success, detail: %v", res))
	return res, nil
}
func (i *ContractInstance) GetElectionResult(electionId int64) ([]election_contract.CommonUtilResult, error) {
	res, err := i.Instance.GetElectionResult(i.AccountBoundRead, big.NewInt(electionId))
	if err != nil {
		i.Logger.Error(fmt.Sprintf("[Contract Instance] Get Owner failed, detail: %v", err))
		return nil, err
	}
	i.Logger.Info(fmt.Sprintf("[Contract Instance] Get GetElectionResult success, detail: %v", res))
	return res, nil
}
func (i *ContractInstance) Owner() common.Address {
	res, err := i.Instance.Owner(i.AccountBoundRead)
	if err != nil {
		i.Logger.Error(fmt.Sprintf("[Contract Instance] Get Owner failed, detail: %v", err))
		return common.Address{}
	}
	i.Logger.Info(fmt.Sprintf("[Contract Instance] Get Owner success, detail: %v", res))
	return res
}
