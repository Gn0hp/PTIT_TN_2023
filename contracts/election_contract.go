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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"prvKey\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_startDate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_numCandidate\",\"type\":\"uint256\"}],\"name\":\"NewElectionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"electionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"voterHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"name\":\"Voted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numCandidate\",\"type\":\"uint256\"}],\"name\":\"createNewElection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"electionToResult\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"elections\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numCandidates\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"electionId\",\"type\":\"uint256\"}],\"name\":\"getElectionResult\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structCommonUtil.Result[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"electionId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_voterId\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"value\",\"type\":\"uint256[]\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040525f60015534801562000014575f80fd5b5060405162001eb938038062001eb983398181016040528101906200003a91906200044c565b6200004a6200012060201b60201c565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000bd575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401620000b49190620004de565b60405180910390fd5b620000ce816200012760201b60201c565b506200010881604051602001620000e6919062000543565b60405160208183030381529060405280519060200120620001e860201b60201c565b6002908162000118919062000788565b505062000912565b5f33905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60605f602067ffffffffffffffff811115620002095762000208620002f0565b5b6040519080825280601f01601f1916602001820160405280156200023c5781602001600182028036833780820191505090505b5090505f5b6020811015620002bd578381602081106200026157620002606200086c565b5b1a60f81b8282815181106200027b576200027a6200086c565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690815f1a90535080620002b590620008c6565b905062000241565b5080915050919050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6200032882620002e0565b810181811067ffffffffffffffff821117156200034a5762000349620002f0565b5b80604052505050565b5f6200035e620002c7565b90506200036c82826200031d565b919050565b5f67ffffffffffffffff8211156200038e576200038d620002f0565b5b6200039982620002e0565b9050602081019050919050565b5f5b83811015620003c5578082015181840152602081019050620003a8565b5f8484015250505050565b5f620003e6620003e08462000371565b62000353565b905082815260208101848484011115620004055762000404620002dc565b5b62000412848285620003a6565b509392505050565b5f82601f830112620004315762000430620002d8565b5b815162000443848260208601620003d0565b91505092915050565b5f60208284031215620004645762000463620002d0565b5b5f82015167ffffffffffffffff811115620004845762000483620002d4565b5b62000492848285016200041a565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620004c6826200049b565b9050919050565b620004d881620004ba565b82525050565b5f602082019050620004f35f830184620004cd565b92915050565b5f81519050919050565b5f81905092915050565b5f6200051982620004f9565b62000525818562000503565b935062000537818560208601620003a6565b80840191505092915050565b5f6200055082846200050d565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680620005a057607f821691505b602082108103620005b657620005b56200055b565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026200061a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620005dd565b620006268683620005dd565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f620006706200066a62000664846200063e565b62000647565b6200063e565b9050919050565b5f819050919050565b6200068b8362000650565b620006a36200069a8262000677565b848454620005e9565b825550505050565b5f90565b620006b9620006ab565b620006c681848462000680565b505050565b5b81811015620006ed57620006e15f82620006af565b600181019050620006cc565b5050565b601f8211156200073c576200070681620005bc565b6200071184620005ce565b8101602085101562000721578190505b620007396200073085620005ce565b830182620006cb565b50505b505050565b5f82821c905092915050565b5f6200075e5f198460080262000741565b1980831691505092915050565b5f6200077883836200074d565b9150826002028217905092915050565b6200079382620004f9565b67ffffffffffffffff811115620007af57620007ae620002f0565b5b620007bb825462000588565b620007c8828285620006f1565b5f60209050601f831160018114620007fe575f8415620007e9578287015190505b620007f585826200076b565b86555062000864565b601f1984166200080e86620005bc565b5f5b82811015620008375784890151825560018201915060208501945060208101905062000810565b8683101562000857578489015162000853601f8916826200074d565b8355505b6001600288020188555050505b505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f620008d2826200063e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820362000907576200090662000899565b5b600182019050919050565b61159980620009205f395ff3fe608060405234801561000f575f80fd5b50600436106100a7575f3560e01c80638da5cb5b1161006f5780638da5cb5b14610152578063affed0e014610170578063b2e429aa1461018e578063bc1c56d5146101be578063f2fde38b146101da578063fb2cfa6d146101f6576100a7565b806323fc4b36146100ab57806350f650d6146100db5780635e6fef01146100f7578063715018a61461012a5780638a4e376914610134575b5f80fd5b6100c560048036038101906100c09190610ad6565b610226565b6040516100d29190610be5565b60405180910390f35b6100f560048036038101906100f09190610c05565b610374565b005b610111600480360381019061010c9190610ad6565b61040e565b6040516101219493929190610c64565b60405180910390f35b61013261043a565b005b61013c61044d565b6040516101499190610d31565b60405180910390f35b61015a6104ed565b6040516101679190610d90565b60405180910390f35b610178610514565b6040516101859190610da9565b60405180910390f35b6101a860048036038101906101a39190610dc2565b61051a565b6040516101b59190610da9565b60405180910390f35b6101d860048036038101906101d39190610ff0565b61053a565b005b6101f460048036038101906101ef91906110a2565b6107df565b005b610210600480360381019061020b9190611100565b610863565b60405161021d9190611158565b60405180910390f35b60605f60035f8481526020019081526020015f2090505f60055f8581526020019081526020015f206040518060800160405290815f8201548152602001600182015481526020016002820154815260200160038201548152505090505f816060015190505f8167ffffffffffffffff8111156102a5576102a4610e08565b5b6040519080825280602002602001820160405280156102de57816020015b6102cb610a7a565b8152602001906001900390816102c35790505b5090505f5b828167ffffffffffffffff1610156103675760405180604001604052808267ffffffffffffffff168152602001865f8467ffffffffffffffff1681526020019081526020015f2054815250828267ffffffffffffffff168151811061034b5761034a611171565b5b602002602001018190525080610360906111de565b90506102e3565b5080945050505050919050565b61037c61088d565b428284610389919061120d565b1015610393575f80fd5b60015f8154809291906103a590611240565b9190505550604051806080016040528060015481526020018481526020018381526020018281525060055f60015481526020019081526020015f205f820151815f0155602082015181600101556040820151816002015560608201518160030155905050505050565b6005602052805f5260405f205f91509050805f0154908060010154908060020154908060030154905084565b61044261088d565b61044b5f610914565b565b606061045761088d565b6104e860028054610467906112b4565b80601f0160208091040260200160405190810160405280929190818152602001828054610493906112b4565b80156104de5780601f106104b5576101008083540402835291602001916104de565b820191905f5260205f20905b8154815290600101906020018083116104c157829003601f168201915b50505050506109d5565b905090565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60015481565b6003602052815f5260405f20602052805f5260405f205f91509150505481565b61054261088d565b5f60055f8581526020019081526020015f206040518060800160405290815f820154815260200160018201548152602001600282015481526020016003820154815250509050806020015142101580156105af5750806040015181602001516105ab919061120d565b4211155b6105ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e59061133e565b60405180910390fd5b5f6106818460028054610600906112b4565b80601f016020809104026020016040519081016040528092919081815260200182805461062c906112b4565b80156106775780601f1061064e57610100808354040283529160200191610677565b820191905f5260205f20905b81548152906001019060200180831161065a57829003601f168201915b50505050506109df565b905060045f8681526020019081526020015f205f8281526020019081526020015f205f9054906101000a900460ff16156106f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e7906113cc565b60405180910390fd5b5f5b835181101561075b5760035f8781526020019081526020015f205f8583815181106107205761071f611171565b5b602002602001015181526020019081526020015f205f81548092919061074590611240565b91905055508061075490611240565b90506106f2565b50600160045f8781526020019081526020015f205f8381526020019081526020015f205f6101000a81548160ff021916908315150217905550826040516107a2919061149b565b604051809103902081867f4998b44a19dc1dcee334d2f10f32bdd1b5d6b0c0544a6f34c04f115fb3f6656660405160405180910390a45050505050565b6107e761088d565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610857575f6040517f1e4fbdf700000000000000000000000000000000000000000000000000000000815260040161084e9190610d90565b60405180910390fd5b61086081610914565b50565b6004602052815f5260405f20602052805f5260405f205f915091509054906101000a900460ff1681565b610895610a73565b73ffffffffffffffffffffffffffffffffffffffff166108b36104ed565b73ffffffffffffffffffffffffffffffffffffffff1614610912576108d6610a73565b6040517f118cdaa70000000000000000000000000000000000000000000000000000000081526004016109099190610d90565b60405180910390fd5b565b5f805f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050815f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6060819050919050565b5f81836040516024016109f291906114f3565b60405160208183030381529060405290604051610a0f919061154d565b60405180910390207bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505080519060200120905092915050565b5f33905090565b60405180604001604052805f81526020015f81525090565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b610ab581610aa3565b8114610abf575f80fd5b50565b5f81359050610ad081610aac565b92915050565b5f60208284031215610aeb57610aea610a9b565b5b5f610af884828501610ac2565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610b3381610aa3565b82525050565b604082015f820151610b4d5f850182610b2a565b506020820151610b606020850182610b2a565b50505050565b5f610b718383610b39565b60408301905092915050565b5f602082019050919050565b5f610b9382610b01565b610b9d8185610b0b565b9350610ba883610b1b565b805f5b83811015610bd8578151610bbf8882610b66565b9750610bca83610b7d565b925050600181019050610bab565b5085935050505092915050565b5f6020820190508181035f830152610bfd8184610b89565b905092915050565b5f805f60608486031215610c1c57610c1b610a9b565b5b5f610c2986828701610ac2565b9350506020610c3a86828701610ac2565b9250506040610c4b86828701610ac2565b9150509250925092565b610c5e81610aa3565b82525050565b5f608082019050610c775f830187610c55565b610c846020830186610c55565b610c916040830185610c55565b610c9e6060830184610c55565b95945050505050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610cde578082015181840152602081019050610cc3565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610d0382610ca7565b610d0d8185610cb1565b9350610d1d818560208601610cc1565b610d2681610ce9565b840191505092915050565b5f6020820190508181035f830152610d498184610cf9565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610d7a82610d51565b9050919050565b610d8a81610d70565b82525050565b5f602082019050610da35f830184610d81565b92915050565b5f602082019050610dbc5f830184610c55565b92915050565b5f8060408385031215610dd857610dd7610a9b565b5b5f610de585828601610ac2565b9250506020610df685828601610ac2565b9150509250929050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610e3e82610ce9565b810181811067ffffffffffffffff82111715610e5d57610e5c610e08565b5b80604052505050565b5f610e6f610a92565b9050610e7b8282610e35565b919050565b5f67ffffffffffffffff821115610e9a57610e99610e08565b5b610ea382610ce9565b9050602081019050919050565b828183375f83830152505050565b5f610ed0610ecb84610e80565b610e66565b905082815260208101848484011115610eec57610eeb610e04565b5b610ef7848285610eb0565b509392505050565b5f82601f830112610f1357610f12610e00565b5b8135610f23848260208601610ebe565b91505092915050565b5f67ffffffffffffffff821115610f4657610f45610e08565b5b602082029050602081019050919050565b5f80fd5b5f610f6d610f6884610f2c565b610e66565b90508083825260208201905060208402830185811115610f9057610f8f610f57565b5b835b81811015610fb95780610fa58882610ac2565b845260208401935050602081019050610f92565b5050509392505050565b5f82601f830112610fd757610fd6610e00565b5b8135610fe7848260208601610f5b565b91505092915050565b5f805f6060848603121561100757611006610a9b565b5b5f61101486828701610ac2565b935050602084013567ffffffffffffffff81111561103557611034610a9f565b5b61104186828701610eff565b925050604084013567ffffffffffffffff81111561106257611061610a9f565b5b61106e86828701610fc3565b9150509250925092565b61108181610d70565b811461108b575f80fd5b50565b5f8135905061109c81611078565b92915050565b5f602082840312156110b7576110b6610a9b565b5b5f6110c48482850161108e565b91505092915050565b5f819050919050565b6110df816110cd565b81146110e9575f80fd5b50565b5f813590506110fa816110d6565b92915050565b5f806040838503121561111657611115610a9b565b5b5f61112385828601610ac2565b9250506020611134858286016110ec565b9150509250929050565b5f8115159050919050565b6111528161113e565b82525050565b5f60208201905061116b5f830184611149565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f67ffffffffffffffff82169050919050565b5f6111e8826111cb565b915067ffffffffffffffff82036112025761120161119e565b5b600182019050919050565b5f61121782610aa3565b915061122283610aa3565b925082820190508082111561123a5761123961119e565b5b92915050565b5f61124a82610aa3565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361127c5761127b61119e565b5b600182019050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806112cb57607f821691505b6020821081036112de576112dd611287565b5b50919050565b5f82825260208201905092915050565b7f5b566f74655d3a20496e76616c696420766f74696e672074696d6500000000005f82015250565b5f611328601b836112e4565b9150611333826112f4565b602082019050919050565b5f6020820190508181035f8301526113558161131c565b9050919050565b7f596f75206861766520766f7465642e204561636820766f7465722063616e206f5f8201527f6e6c7920766f7465206f6e636521000000000000000000000000000000000000602082015250565b5f6113b6602e836112e4565b91506113c18261135c565b604082019050919050565b5f6020820190508181035f8301526113e3816113aa565b9050919050565b5f81519050919050565b5f81905092915050565b5f819050602082019050919050565b61141681610aa3565b82525050565b5f611427838361140d565b60208301905092915050565b5f602082019050919050565b5f611449826113ea565b61145381856113f4565b935061145e836113fe565b805f5b8381101561148e578151611475888261141c565b975061148083611433565b925050600181019050611461565b5085935050505092915050565b5f6114a6828461143f565b915081905092915050565b5f81519050919050565b5f6114c5826114b1565b6114cf81856112e4565b93506114df818560208601610cc1565b6114e881610ce9565b840191505092915050565b5f6020820190508181035f83015261150b81846114bb565b905092915050565b5f81905092915050565b5f611527826114b1565b6115318185611513565b9350611541818560208601610cc1565b80840191505092915050565b5f611558828461151d565b91508190509291505056fea2646970667358221220abdd59ecfb716b191fa644d546c20a04675c2598578c4d458d46df8425baf0d064736f6c63430008150033",
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

// GetSignature is a free data retrieval call binding the contract method 0x8a4e3769.
//
// Solidity: function getSignature() view returns(bytes)
func (_ElectionContract *ElectionContractCaller) GetSignature(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _ElectionContract.contract.Call(opts, &out, "getSignature")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSignature is a free data retrieval call binding the contract method 0x8a4e3769.
//
// Solidity: function getSignature() view returns(bytes)
func (_ElectionContract *ElectionContractSession) GetSignature() ([]byte, error) {
	return _ElectionContract.Contract.GetSignature(&_ElectionContract.CallOpts)
}

// GetSignature is a free data retrieval call binding the contract method 0x8a4e3769.
//
// Solidity: function getSignature() view returns(bytes)
func (_ElectionContract *ElectionContractCallerSession) GetSignature() ([]byte, error) {
	return _ElectionContract.Contract.GetSignature(&_ElectionContract.CallOpts)
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

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_ElectionContract *ElectionContractCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ElectionContract.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_ElectionContract *ElectionContractSession) Nonce() (*big.Int, error) {
	return _ElectionContract.Contract.Nonce(&_ElectionContract.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_ElectionContract *ElectionContractCallerSession) Nonce() (*big.Int, error) {
	return _ElectionContract.Contract.Nonce(&_ElectionContract.CallOpts)
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

// Vote is a paid mutator transaction binding the contract method 0xbc1c56d5.
//
// Solidity: function vote(uint256 electionId, string _voterId, uint256[] value) returns()
func (_ElectionContract *ElectionContractTransactor) Vote(opts *bind.TransactOpts, electionId *big.Int, _voterId string, value []*big.Int) (*types.Transaction, error) {
	return _ElectionContract.contract.Transact(opts, "vote", electionId, _voterId, value)
}

// Vote is a paid mutator transaction binding the contract method 0xbc1c56d5.
//
// Solidity: function vote(uint256 electionId, string _voterId, uint256[] value) returns()
func (_ElectionContract *ElectionContractSession) Vote(electionId *big.Int, _voterId string, value []*big.Int) (*types.Transaction, error) {
	return _ElectionContract.Contract.Vote(&_ElectionContract.TransactOpts, electionId, _voterId, value)
}

// Vote is a paid mutator transaction binding the contract method 0xbc1c56d5.
//
// Solidity: function vote(uint256 electionId, string _voterId, uint256[] value) returns()
func (_ElectionContract *ElectionContractTransactorSession) Vote(electionId *big.Int, _voterId string, value []*big.Int) (*types.Transaction, error) {
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
	Value      []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoted is a free log retrieval operation binding the contract event 0x4998b44a19dc1dcee334d2f10f32bdd1b5d6b0c0544a6f34c04f115fb3f66566.
//
// Solidity: event Voted(uint256 indexed electionId, bytes32 indexed voterHash, uint256[] indexed value)
func (_ElectionContract *ElectionContractFilterer) FilterVoted(opts *bind.FilterOpts, electionId []*big.Int, voterHash [][32]byte, value [][]*big.Int) (*ElectionContractVotedIterator, error) {

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

// WatchVoted is a free log subscription operation binding the contract event 0x4998b44a19dc1dcee334d2f10f32bdd1b5d6b0c0544a6f34c04f115fb3f66566.
//
// Solidity: event Voted(uint256 indexed electionId, bytes32 indexed voterHash, uint256[] indexed value)
func (_ElectionContract *ElectionContractFilterer) WatchVoted(opts *bind.WatchOpts, sink chan<- *ElectionContractVoted, electionId []*big.Int, voterHash [][32]byte, value [][]*big.Int) (event.Subscription, error) {

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

// ParseVoted is a log parse operation binding the contract event 0x4998b44a19dc1dcee334d2f10f32bdd1b5d6b0c0544a6f34c04f115fb3f66566.
//
// Solidity: event Voted(uint256 indexed electionId, bytes32 indexed voterHash, uint256[] indexed value)
func (_ElectionContract *ElectionContractFilterer) ParseVoted(log types.Log) (*ElectionContractVoted, error) {
	event := new(ElectionContractVoted)
	if err := _ElectionContract.contract.UnpackLog(event, "Voted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
