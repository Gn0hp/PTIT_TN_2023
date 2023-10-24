// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package election_contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// CommonUtilResult is an auto generated low-level Go binding around an user-defined struct.
type CommonUtilResult struct {
	Id    *big.Int
	Value *big.Int
}

// ElectionContractMetaData contains all meta data concerning the ElectionContract contract.
var ElectionContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"prvKey\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_startDate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_numCandidate\",\"type\":\"uint256\"}],\"name\":\"NewElectionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"electionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"voterHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numCandidate\",\"type\":\"uint256\"}],\"name\":\"createNewElection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"electionToResult\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"elections\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numCandidates\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"electionId\",\"type\":\"uint256\"}],\"name\":\"getElectionResult\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structCommonUtil.Result[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"electionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_voterId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040525f60015534801562000014575f80fd5b5060405162001b5238038062001b5283398181016040528101906200003a91906200044c565b6200004a6200012060201b60201c565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000bd575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401620000b49190620004de565b60405180910390fd5b620000ce816200012760201b60201c565b506200010881604051602001620000e6919062000543565b60405160208183030381529060405280519060200120620001e860201b60201c565b6002908162000118919062000788565b505062000912565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60605f602067ffffffffffffffff811115620002095762000208620002f0565b5b6040519080825280601f01601f1916602001820160405280156200023c5781602001600182028036833780820191505090505b5090505f5b6020811015620002bd578381602081106200026157620002606200086c565b5b1a60f81b8282815181106200027b576200027a6200086c565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a90535080620002b590620008c6565b905062000241565b5080915050919050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6200032882620002e0565b810181811067ffffffffffffffff821117156200034a5762000349620002f0565b5b80604052505050565b5f6200035e620002c7565b90506200036c82826200031d565b919050565b5f67ffffffffffffffff8211156200038e576200038d620002f0565b5b6200039982620002e0565b9050602081019050919050565b5f5b83811015620003c5578082015181840152602081019050620003a8565b5f8484015250505050565b5f620003e6620003e08462000371565b62000353565b905082815260208101848484011115620004055762000404620002dc565b5b62000412848285620003a6565b509392505050565b5f82601f830112620004315762000430620002d8565b5b815162000443848260208601620003d0565b91505092915050565b5f60208284031215620004645762000463620002d0565b5b5f82015167ffffffffffffffff811115620004845762000483620002d4565b5b62000492848285016200041a565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620004c6826200049b565b9050919050565b620004d881620004ba565b82525050565b5f602082019050620004f35f830184620004cd565b92915050565b5f81519050919050565b5f81905092915050565b5f6200051982620004f9565b62000525818562000503565b935062000537818560208601620003a6565b80840191505092915050565b5f6200055082846200050d565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680620005a057607f821691505b602082108103620005b657620005b56200055b565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026200061a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620005dd565b620006268683620005dd565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620006706200066a62000664846200063e565b62000647565b6200063e565b9050919050565b5f819050919050565b6200068b8362000650565b620006a36200069a8262000677565b848454620005e9565b825550505050565b5f90565b620006b9620006ab565b620006c681848462000680565b505050565b5b81811015620006ed57620006e15f82620006af565b600181019050620006cc565b5050565b601f8211156200073c576200070681620005bc565b6200071184620005ce565b8101602085101562000721578190505b620007396200073085620005ce565b830182620006cb565b50505b505050565b5f82821c905092915050565b5f6200075e5f198460080262000741565b1980831691505092915050565b5f6200077883836200074d565b9150826002028217905092915050565b6200079382620004f9565b67ffffffffffffffff811115620007af57620007ae620002f0565b5b620007bb825462000588565b620007c8828285620006f1565b5f60209050601f831160018114620007fe575f8415620007e9578287015190505b620007f585826200076b565b86555062000864565b601f1984166200080e86620005bc565b5f5b82811015620008375784890151825560018201915060208501945060208101905062000810565b8683101562000857578489015162000853601f8916826200074d565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f620008d2826200063e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820362000907576200090662000899565b5b600182019050919050565b61123280620009205f395ff3fe608060405234801561000f575f80fd5b5060043610610091575f3560e01c80638da5cb5b116100645780638da5cb5b1461011e578063b2e429aa1461013c578063f0f454571461016c578063f2fde38b14610188578063fb2cfa6d146101a457610091565b806323fc4b361461009557806350f650d6146100c55780635e6fef01146100e1578063715018a614610114575b5f80fd5b6100af60048036038101906100aa9190610988565b6101d4565b6040516100bc9190610a97565b60405180910390f35b6100df60048036038101906100da9190610ab7565b610322565b005b6100fb60048036038101906100f69190610988565b6103bc565b60405161010b9493929190610b16565b60405180910390f35b61011c6103e8565b005b6101266103fb565b6040516101339190610b98565b60405180910390f35b61015660048036038101906101519190610bb1565b610422565b6040516101639190610bef565b60405180910390f35b61018660048036038101906101819190610d44565b610442565b005b6101a2600480360381019061019d9190610dda565b61069b565b005b6101be60048036038101906101b99190610e38565b61071f565b6040516101cb9190610e90565b60405180910390f35b60605f60035f8481526020019081526020015f2090505f60055f8581526020019081526020015f206040518060800160405290815f8201548152602001600182015481526020016002820154815260200160038201548152505090505f816060015190505f8167ffffffffffffffff81111561025357610252610c20565b5b60405190808252806020026020018201604052801561028c57816020015b61027961092c565b8152602001906001900390816102715790505b5090505f5b828167ffffffffffffffff1610156103155760405180604001604052808267ffffffffffffffff168152602001865f8467ffffffffffffffff1681526020019081526020015f2054815250828267ffffffffffffffff16815181106102f9576102f8610ea9565b5b60200260200101819052508061030e90610f16565b9050610291565b5080945050505050919050565b61032a610749565b4282846103379190610f45565b1015610341575f80fd5b60015f81548092919061035390610f78565b9190505550604051806080016040528060015481526020018481526020018381526020018281525060055f60015481526020019081526020015f205f820151815f0155602082015181600101556040820151816002015560608201518160030155905050505050565b6005602052805f5260405f205f91509050805f0154908060010154908060020154908060030154905084565b6103f0610749565b6103f95f6107d0565b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6003602052815f5260405f20602052805f5260405f205f91509150505481565b61044a610749565b5f60055f8581526020019081526020015f206040518060800160405290815f820154815260200160018201548152602001600282015481526020016003820154815250509050806020015142101580156104b75750806040015181602001516104b39190610f45565b4211155b6104f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ed90611019565b60405180910390fd5b5f610589846002805461050890611064565b80601f016020809104026020016040519081016040528092919081815260200182805461053490611064565b801561057f5780601f106105565761010080835404028352916020019161057f565b820191905f5260205f20905b81548152906001019060200180831161056257829003601f168201915b5050505050610891565b905060045f8681526020019081526020015f205f8281526020019081526020015f205f9054906101000a900460ff16156105f8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ef90611104565b60405180910390fd5b60035f8681526020019081526020015f205f8481526020019081526020015f205f81548092919061062890610f78565b9190505550600160045f8781526020019081526020015f205f8381526020019081526020015f205f6101000a81548160ff0219169083151502179055508281867f29efd4aee5961a3fe4cc194e4ea1a4e4572531e3a2426910970f85d0240f356c60405160405180910390a45050505050565b6106a3610749565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610713575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161070a9190610b98565b60405180910390fd5b61071c816107d0565b50565b6004602052815f5260405f20602052805f5260405f205f915091509054906101000a900460ff1681565b610751610925565b73ffffffffffffffffffffffffffffffffffffffff1661076f6103fb565b73ffffffffffffffffffffffffffffffffffffffff16146107ce57610792610925565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016107c59190610b98565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f81836040516024016108a4919061118c565b604051602081830303815290604052906040516108c191906111e6565b60405180910390207bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505080519060200120905092915050565b5f33905090565b60405180604001604052805f81526020015f81525090565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b61096781610955565b8114610971575f80fd5b50565b5f813590506109828161095e565b92915050565b5f6020828403121561099d5761099c61094d565b5b5f6109aa84828501610974565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6109e581610955565b82525050565b604082015f8201516109ff5f8501826109dc565b506020820151610a1260208501826109dc565b50505050565b5f610a2383836109eb565b60408301905092915050565b5f602082019050919050565b5f610a45826109b3565b610a4f81856109bd565b9350610a5a836109cd565b805f5b83811015610a8a578151610a718882610a18565b9750610a7c83610a2f565b925050600181019050610a5d565b5085935050505092915050565b5f6020820190508181035f830152610aaf8184610a3b565b905092915050565b5f805f60608486031215610ace57610acd61094d565b5b5f610adb86828701610974565b9350506020610aec86828701610974565b9250506040610afd86828701610974565b9150509250925092565b610b1081610955565b82525050565b5f608082019050610b295f830187610b07565b610b366020830186610b07565b610b436040830185610b07565b610b506060830184610b07565b95945050505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610b8282610b59565b9050919050565b610b9281610b78565b82525050565b5f602082019050610bab5f830184610b89565b92915050565b5f8060408385031215610bc757610bc661094d565b5b5f610bd485828601610974565b9250506020610be585828601610974565b9150509250929050565b5f602082019050610c025f830184610b07565b92915050565b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610c5682610c10565b810181811067ffffffffffffffff82111715610c7557610c74610c20565b5b80604052505050565b5f610c87610944565b9050610c938282610c4d565b919050565b5f67ffffffffffffffff821115610cb257610cb1610c20565b5b610cbb82610c10565b9050602081019050919050565b828183375f83830152505050565b5f610ce8610ce384610c98565b610c7e565b905082815260208101848484011115610d0457610d03610c0c565b5b610d0f848285610cc8565b509392505050565b5f82601f830112610d2b57610d2a610c08565b5b8135610d3b848260208601610cd6565b91505092915050565b5f805f60608486031215610d5b57610d5a61094d565b5b5f610d6886828701610974565b935050602084013567ffffffffffffffff811115610d8957610d88610951565b5b610d9586828701610d17565b9250506040610da686828701610974565b9150509250925092565b610db981610b78565b8114610dc3575f80fd5b50565b5f81359050610dd481610db0565b92915050565b5f60208284031215610def57610dee61094d565b5b5f610dfc84828501610dc6565b91505092915050565b5f819050919050565b610e1781610e05565b8114610e21575f80fd5b50565b5f81359050610e3281610e0e565b92915050565b5f8060408385031215610e4e57610e4d61094d565b5b5f610e5b85828601610974565b9250506020610e6c85828601610e24565b9150509250929050565b5f8115159050919050565b610e8a81610e76565b82525050565b5f602082019050610ea35f830184610e81565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f67ffffffffffffffff82169050919050565b5f610f2082610f03565b915067ffffffffffffffff8203610f3a57610f39610ed6565b5b600182019050919050565b5f610f4f82610955565b9150610f5a83610955565b9250828201905080821115610f7257610f71610ed6565b5b92915050565b5f610f8282610955565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610fb457610fb3610ed6565b5b600182019050919050565b5f82825260208201905092915050565b7f5b566f74655d3a20496e76616c696420766f74696e672074696d6500000000005f82015250565b5f611003601b83610fbf565b915061100e82610fcf565b602082019050919050565b5f6020820190508181035f83015261103081610ff7565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061107b57607f821691505b60208210810361108e5761108d611037565b5b50919050565b7f596f75206861766520766f7465642e204561636820766f7465722063616e206f5f8201527f6e6c7920766f7465206f6e636521000000000000000000000000000000000000602082015250565b5f6110ee602e83610fbf565b91506110f982611094565b604082019050919050565b5f6020820190508181035f83015261111b816110e2565b9050919050565b5f81519050919050565b5f5b8381101561114957808201518184015260208101905061112e565b5f8484015250505050565b5f61115e82611122565b6111688185610fbf565b935061117881856020860161112c565b61118181610c10565b840191505092915050565b5f6020820190508181035f8301526111a48184611154565b905092915050565b5f81905092915050565b5f6111c082611122565b6111ca81856111ac565b93506111da81856020860161112c565b80840191505092915050565b5f6111f182846111b6565b91508190509291505056fea2646970667358221220005f420823b72090ef44ae56db667e5baa721d5d2aad874d4ffa6cf17cf82ec264736f6c63430008150033",
}

// ElectionContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ElectionContractMetaData.ABI instead.
var ElectionContractABI = ElectionContractMetaData.ABI

// ElectionContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ElectionContractMetaData.Bin instead.
var ElectionContractBin = ElectionContractMetaData.Bin

// DeployElectionContract deploys a new Ethereum contract, binding an instance of ElectionContract to it.
func DeployElectionContract(auth *bind.TransactOpts, backend bind.ContractBackend, prvKey string) (common.Address, *types.Transaction, *ElectionContract, error) {
	parsed, err := ElectionContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ElectionContractBin), backend, prvKey)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ElectionContract{ElectionContractCaller: ElectionContractCaller{contract: contract}, ElectionContractTransactor: ElectionContractTransactor{contract: contract}, ElectionContractFilterer: ElectionContractFilterer{contract: contract}}, nil
}

// ElectionContract is an auto generated Go binding around an Ethereum contract.
type ElectionContract struct {
	ElectionContractCaller     // Read-only binding to the contract
	ElectionContractTransactor // Write-only binding to the contract
	ElectionContractFilterer   // Log filterer for contract events
}

// ElectionContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ElectionContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ElectionContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ElectionContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ElectionContractSession struct {
	Contract     *ElectionContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ElectionContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ElectionContractCallerSession struct {
	Contract *ElectionContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ElectionContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ElectionContractTransactorSession struct {
	Contract     *ElectionContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ElectionContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ElectionContractRaw struct {
	Contract *ElectionContract // Generic contract binding to access the raw methods on
}

// ElectionContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ElectionContractCallerRaw struct {
	Contract *ElectionContractCaller // Generic read-only contract binding to access the raw methods on
}

// ElectionContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ElectionContractTransactorRaw struct {
	Contract *ElectionContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewElectionContract creates a new instance of ElectionContract, bound to a specific deployed contract.
func NewElectionContract(address common.Address, backend bind.ContractBackend) (*ElectionContract, error) {
	contract, err := bindElectionContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ElectionContract{ElectionContractCaller: ElectionContractCaller{contract: contract}, ElectionContractTransactor: ElectionContractTransactor{contract: contract}, ElectionContractFilterer: ElectionContractFilterer{contract: contract}}, nil
}

// NewElectionContractCaller creates a new read-only instance of ElectionContract, bound to a specific deployed contract.
func NewElectionContractCaller(address common.Address, caller bind.ContractCaller) (*ElectionContractCaller, error) {
	contract, err := bindElectionContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ElectionContractCaller{contract: contract}, nil
}

// NewElectionContractTransactor creates a new write-only instance of ElectionContract, bound to a specific deployed contract.
func NewElectionContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ElectionContractTransactor, error) {
	contract, err := bindElectionContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ElectionContractTransactor{contract: contract}, nil
}

// NewElectionContractFilterer creates a new log filterer instance of ElectionContract, bound to a specific deployed contract.
func NewElectionContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ElectionContractFilterer, error) {
	contract, err := bindElectionContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ElectionContractFilterer{contract: contract}, nil
}

// bindElectionContract binds a generic wrapper to an already deployed contract.
func bindElectionContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ElectionContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElectionContract *ElectionContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ElectionContract.Contract.ElectionContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElectionContract *ElectionContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionContract.Contract.ElectionContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElectionContract *ElectionContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ElectionContract.Contract.ElectionContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ElectionContract *ElectionContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ElectionContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ElectionContract *ElectionContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ElectionContract *ElectionContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ElectionContract.Contract.contract.Transact(opts, method, params...)
}

// ElectionToResult is a free data retrieval call binding the contract method 0xb2e429aa.
//
// Solidity: function electionToResult(uint256 , uint256 ) view returns(uint256)
func (_ElectionContract *ElectionContractCaller) ElectionToResult(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ElectionContract.contract.Call(opts, &out, "electionToResult", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ElectionToResult is a free data retrieval call binding the contract method 0xb2e429aa.
//
// Solidity: function electionToResult(uint256 , uint256 ) view returns(uint256)
func (_ElectionContract *ElectionContractSession) ElectionToResult(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _ElectionContract.Contract.ElectionToResult(&_ElectionContract.CallOpts, arg0, arg1)
}

// ElectionToResult is a free data retrieval call binding the contract method 0xb2e429aa.
//
// Solidity: function electionToResult(uint256 , uint256 ) view returns(uint256)
func (_ElectionContract *ElectionContractCallerSession) ElectionToResult(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _ElectionContract.Contract.ElectionToResult(&_ElectionContract.CallOpts, arg0, arg1)
}

// Elections is a free data retrieval call binding the contract method 0x5e6fef01.
//
// Solidity: function elections(uint256 ) view returns(uint256 id, uint256 startDate, uint256 duration, uint256 numCandidates)
func (_ElectionContract *ElectionContractCaller) Elections(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id            *big.Int
	StartDate     *big.Int
	Duration      *big.Int
	NumCandidates *big.Int
}, error) {
	var out []interface{}
	err := _ElectionContract.contract.Call(opts, &out, "elections", arg0)

	outstruct := new(struct {
		Id            *big.Int
		StartDate     *big.Int
		Duration      *big.Int
		NumCandidates *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartDate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.NumCandidates = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Elections is a free data retrieval call binding the contract method 0x5e6fef01.
//
// Solidity: function elections(uint256 ) view returns(uint256 id, uint256 startDate, uint256 duration, uint256 numCandidates)
func (_ElectionContract *ElectionContractSession) Elections(arg0 *big.Int) (struct {
	Id            *big.Int
	StartDate     *big.Int
	Duration      *big.Int
	NumCandidates *big.Int
}, error) {
	return _ElectionContract.Contract.Elections(&_ElectionContract.CallOpts, arg0)
}

// Elections is a free data retrieval call binding the contract method 0x5e6fef01.
//
// Solidity: function elections(uint256 ) view returns(uint256 id, uint256 startDate, uint256 duration, uint256 numCandidates)
func (_ElectionContract *ElectionContractCallerSession) Elections(arg0 *big.Int) (struct {
	Id            *big.Int
	StartDate     *big.Int
	Duration      *big.Int
	NumCandidates *big.Int
}, error) {
	return _ElectionContract.Contract.Elections(&_ElectionContract.CallOpts, arg0)
}

// GetElectionResult is a free data retrieval call binding the contract method 0x23fc4b36.
//
// Solidity: function getElectionResult(uint256 electionId) view returns((uint256,uint256)[])
func (_ElectionContract *ElectionContractCaller) GetElectionResult(opts *bind.CallOpts, electionId *big.Int) ([]CommonUtilResult, error) {
	var out []interface{}
	err := _ElectionContract.contract.Call(opts, &out, "getElectionResult", electionId)

	if err != nil {
		return *new([]CommonUtilResult), err
	}

	out0 := *abi.ConvertType(out[0], new([]CommonUtilResult)).(*[]CommonUtilResult)

	return out0, err

}

// GetElectionResult is a free data retrieval call binding the contract method 0x23fc4b36.
//
// Solidity: function getElectionResult(uint256 electionId) view returns((uint256,uint256)[])
func (_ElectionContract *ElectionContractSession) GetElectionResult(electionId *big.Int) ([]CommonUtilResult, error) {
	return _ElectionContract.Contract.GetElectionResult(&_ElectionContract.CallOpts, electionId)
}

// GetElectionResult is a free data retrieval call binding the contract method 0x23fc4b36.
//
// Solidity: function getElectionResult(uint256 electionId) view returns((uint256,uint256)[])
func (_ElectionContract *ElectionContractCallerSession) GetElectionResult(electionId *big.Int) ([]CommonUtilResult, error) {
	return _ElectionContract.Contract.GetElectionResult(&_ElectionContract.CallOpts, electionId)
}

// IsVoted is a free data retrieval call binding the contract method 0xfb2cfa6d.
//
// Solidity: function isVoted(uint256 , bytes32 ) view returns(bool)
func (_ElectionContract *ElectionContractCaller) IsVoted(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _ElectionContract.contract.Call(opts, &out, "isVoted", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoted is a free data retrieval call binding the contract method 0xfb2cfa6d.
//
// Solidity: function isVoted(uint256 , bytes32 ) view returns(bool)
func (_ElectionContract *ElectionContractSession) IsVoted(arg0 *big.Int, arg1 [32]byte) (bool, error) {
	return _ElectionContract.Contract.IsVoted(&_ElectionContract.CallOpts, arg0, arg1)
}

// IsVoted is a free data retrieval call binding the contract method 0xfb2cfa6d.
//
// Solidity: function isVoted(uint256 , bytes32 ) view returns(bool)
func (_ElectionContract *ElectionContractCallerSession) IsVoted(arg0 *big.Int, arg1 [32]byte) (bool, error) {
	return _ElectionContract.Contract.IsVoted(&_ElectionContract.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ElectionContract *ElectionContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ElectionContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ElectionContract *ElectionContractSession) Owner() (common.Address, error) {
	return _ElectionContract.Contract.Owner(&_ElectionContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ElectionContract *ElectionContractCallerSession) Owner() (common.Address, error) {
	return _ElectionContract.Contract.Owner(&_ElectionContract.CallOpts)
}

// CreateNewElection is a paid mutator transaction binding the contract method 0x50f650d6.
//
// Solidity: function createNewElection(uint256 _startDate, uint256 _duration, uint256 _numCandidate) returns()
func (_ElectionContract *ElectionContractTransactor) CreateNewElection(opts *bind.TransactOpts, _startDate *big.Int, _duration *big.Int, _numCandidate *big.Int) (*types.Transaction, error) {
	return _ElectionContract.contract.Transact(opts, "createNewElection", _startDate, _duration, _numCandidate)
}

// CreateNewElection is a paid mutator transaction binding the contract method 0x50f650d6.
//
// Solidity: function createNewElection(uint256 _startDate, uint256 _duration, uint256 _numCandidate) returns()
func (_ElectionContract *ElectionContractSession) CreateNewElection(_startDate *big.Int, _duration *big.Int, _numCandidate *big.Int) (*types.Transaction, error) {
	return _ElectionContract.Contract.CreateNewElection(&_ElectionContract.TransactOpts, _startDate, _duration, _numCandidate)
}

// CreateNewElection is a paid mutator transaction binding the contract method 0x50f650d6.
//
// Solidity: function createNewElection(uint256 _startDate, uint256 _duration, uint256 _numCandidate) returns()
func (_ElectionContract *ElectionContractTransactorSession) CreateNewElection(_startDate *big.Int, _duration *big.Int, _numCandidate *big.Int) (*types.Transaction, error) {
	return _ElectionContract.Contract.CreateNewElection(&_ElectionContract.TransactOpts, _startDate, _duration, _numCandidate)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ElectionContract *ElectionContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ElectionContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ElectionContract *ElectionContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _ElectionContract.Contract.RenounceOwnership(&_ElectionContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ElectionContract *ElectionContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ElectionContract.Contract.RenounceOwnership(&_ElectionContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ElectionContract *ElectionContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ElectionContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ElectionContract *ElectionContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ElectionContract.Contract.TransferOwnership(&_ElectionContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ElectionContract *ElectionContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ElectionContract.Contract.TransferOwnership(&_ElectionContract.TransactOpts, newOwner)
}

// Vote is a paid mutator transaction binding the contract method 0xf0f45457.
//
// Solidity: function vote(uint256 electionId, string _voterId, uint256 value) returns()
func (_ElectionContract *ElectionContractTransactor) Vote(opts *bind.TransactOpts, electionId *big.Int, _voterId string, value *big.Int) (*types.Transaction, error) {
	return _ElectionContract.contract.Transact(opts, "vote", electionId, _voterId, value)
}

// Vote is a paid mutator transaction binding the contract method 0xf0f45457.
//
// Solidity: function vote(uint256 electionId, string _voterId, uint256 value) returns()
func (_ElectionContract *ElectionContractSession) Vote(electionId *big.Int, _voterId string, value *big.Int) (*types.Transaction, error) {
	return _ElectionContract.Contract.Vote(&_ElectionContract.TransactOpts, electionId, _voterId, value)
}

// Vote is a paid mutator transaction binding the contract method 0xf0f45457.
//
// Solidity: function vote(uint256 electionId, string _voterId, uint256 value) returns()
func (_ElectionContract *ElectionContractTransactorSession) Vote(electionId *big.Int, _voterId string, value *big.Int) (*types.Transaction, error) {
	return _ElectionContract.Contract.Vote(&_ElectionContract.TransactOpts, electionId, _voterId, value)
}

// ElectionContractNewElectionCreatedIterator is returned from FilterNewElectionCreated and is used to iterate over the raw logs and unpacked data for NewElectionCreated events raised by the ElectionContract contract.
type ElectionContractNewElectionCreatedIterator struct {
	Event *ElectionContractNewElectionCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ElectionContractNewElectionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionContractNewElectionCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ElectionContractNewElectionCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ElectionContractNewElectionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionContractNewElectionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionContractNewElectionCreated represents a NewElectionCreated event raised by the ElectionContract contract.
type ElectionContractNewElectionCreated struct {
	StartDate    *big.Int
	Duration     *big.Int
	NumCandidate *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewElectionCreated is a free log retrieval operation binding the contract event 0xc9473b53b8fd2efec13f0dba18c583c918a79ace1cb12fc31f2ce474d0788e29.
//
// Solidity: event NewElectionCreated(uint256 _startDate, uint256 _duration, uint256 _numCandidate)
func (_ElectionContract *ElectionContractFilterer) FilterNewElectionCreated(opts *bind.FilterOpts) (*ElectionContractNewElectionCreatedIterator, error) {

	logs, sub, err := _ElectionContract.contract.FilterLogs(opts, "NewElectionCreated")
	if err != nil {
		return nil, err
	}
	return &ElectionContractNewElectionCreatedIterator{contract: _ElectionContract.contract, event: "NewElectionCreated", logs: logs, sub: sub}, nil
}

// WatchNewElectionCreated is a free log subscription operation binding the contract event 0xc9473b53b8fd2efec13f0dba18c583c918a79ace1cb12fc31f2ce474d0788e29.
//
// Solidity: event NewElectionCreated(uint256 _startDate, uint256 _duration, uint256 _numCandidate)
func (_ElectionContract *ElectionContractFilterer) WatchNewElectionCreated(opts *bind.WatchOpts, sink chan<- *ElectionContractNewElectionCreated) (event.Subscription, error) {

	logs, sub, err := _ElectionContract.contract.WatchLogs(opts, "NewElectionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionContractNewElectionCreated)
				if err := _ElectionContract.contract.UnpackLog(event, "NewElectionCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewElectionCreated is a log parse operation binding the contract event 0xc9473b53b8fd2efec13f0dba18c583c918a79ace1cb12fc31f2ce474d0788e29.
//
// Solidity: event NewElectionCreated(uint256 _startDate, uint256 _duration, uint256 _numCandidate)
func (_ElectionContract *ElectionContractFilterer) ParseNewElectionCreated(log types.Log) (*ElectionContractNewElectionCreated, error) {
	event := new(ElectionContractNewElectionCreated)
	if err := _ElectionContract.contract.UnpackLog(event, "NewElectionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElectionContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ElectionContract contract.
type ElectionContractOwnershipTransferredIterator struct {
	Event *ElectionContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ElectionContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ElectionContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ElectionContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionContractOwnershipTransferred represents a OwnershipTransferred event raised by the ElectionContract contract.
type ElectionContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ElectionContract *ElectionContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ElectionContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ElectionContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ElectionContractOwnershipTransferredIterator{contract: _ElectionContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ElectionContract *ElectionContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ElectionContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ElectionContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionContractOwnershipTransferred)
				if err := _ElectionContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ElectionContract *ElectionContractFilterer) ParseOwnershipTransferred(log types.Log) (*ElectionContractOwnershipTransferred, error) {
	event := new(ElectionContractOwnershipTransferred)
	if err := _ElectionContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElectionContractVotedIterator is returned from FilterVoted and is used to iterate over the raw logs and unpacked data for Voted events raised by the ElectionContract contract.
type ElectionContractVotedIterator struct {
	Event *ElectionContractVoted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ElectionContractVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionContractVoted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ElectionContractVoted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ElectionContractVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionContractVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionContractVoted represents a Voted event raised by the ElectionContract contract.
type ElectionContractVoted struct {
	ElectionId *big.Int
	VoterHash  [32]byte
	Value      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x29efd4aee5961a3fe4cc194e4ea1a4e4572531e3a2426910970f85d0240f356c.
//
// Solidity: event Voted(uint256 indexed electionId, bytes32 indexed voterHash, uint256 indexed value)
func (_ElectionContract *ElectionContractFilterer) FilterVoted(opts *bind.FilterOpts, electionId []*big.Int, voterHash [][32]byte, value []*big.Int) (*ElectionContractVotedIterator, error) {

	var electionIdRule []interface{}
	for _, electionIdItem := range electionId {
		electionIdRule = append(electionIdRule, electionIdItem)
	}
	var voterHashRule []interface{}
	for _, voterHashItem := range voterHash {
		voterHashRule = append(voterHashRule, voterHashItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _ElectionContract.contract.FilterLogs(opts, "Voted", electionIdRule, voterHashRule, valueRule)
	if err != nil {
		return nil, err
	}
	return &ElectionContractVotedIterator{contract: _ElectionContract.contract, event: "Voted", logs: logs, sub: sub}, nil
}

// WatchVoted is a free log subscription operation binding the contract event 0x29efd4aee5961a3fe4cc194e4ea1a4e4572531e3a2426910970f85d0240f356c.
//
// Solidity: event Voted(uint256 indexed electionId, bytes32 indexed voterHash, uint256 indexed value)
func (_ElectionContract *ElectionContractFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *ElectionContractVoted, electionId []*big.Int, voterHash [][32]byte, value []*big.Int) (event.Subscription, error) {

	var electionIdRule []interface{}
	for _, electionIdItem := range electionId {
		electionIdRule = append(electionIdRule, electionIdItem)
	}
	var voterHashRule []interface{}
	for _, voterHashItem := range voterHash {
		voterHashRule = append(voterHashRule, voterHashItem)
	}
	var valueRule []interface{}
	for _, valueItem := range value {
		valueRule = append(valueRule, valueItem)
	}

	logs, sub, err := _ElectionContract.contract.WatchLogs(opts, "Voted", electionIdRule, voterHashRule, valueRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionContractVoted)
				if err := _ElectionContract.contract.UnpackLog(event, "Voted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoted is a log parse operation binding the contract event 0x29efd4aee5961a3fe4cc194e4ea1a4e4572531e3a2426910970f85d0240f356c.
//
// Solidity: event Voted(uint256 indexed electionId, bytes32 indexed voterHash, uint256 indexed value)
func (_ElectionContract *ElectionContractFilterer) ParseVoted(log types.Log) (*ElectionContractVoted, error) {
	event := new(ElectionContractVoted)
	if err := _ElectionContract.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
