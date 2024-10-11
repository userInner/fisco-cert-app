// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
)

// EvidenceFactoryABI is the input ABI used to generate the binding from.
const EvidenceFactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addSignatures\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getSigner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getEvidence\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"verify\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSigners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"evi\",\"type\":\"string\"}],\"name\":\"newEvidence\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSignersSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"evidenceSigners\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"newEvidenceEvent\",\"type\":\"event\"}]"

// EvidenceFactoryBin is the compiled bytecode used for deploying new contracts.
var EvidenceFactoryBin = "0x608060405234801561001057600080fd5b50604051611c9f380380611c9f8339810180604052810190808051820192919050505060008090505b81518110156100ca576000828281518110151561005257fe5b9060200190602002015190806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050806001019050610039565b5050611bc4806100db6000396000f3006080604052600436106200008b576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063203e492214620000905780633ffefe4e14620000ee57806360004bd1146200015e57806363a9c3d714620002b457806394cf795e1462000312578063a3aaada51462000383578063fa69efbd146200042f575b600080fd5b3480156200009d57600080fd5b50620000d4600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506200045d565b604051808215151515815260200191505060405180910390f35b348015620000fb57600080fd5b506200011c6004803603810190808035906020019092919050505062000508565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156200016b57600080fd5b50620001a2600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506200056a565b60405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b83811015620001ec578082015181840152602081019050620001cf565b50505050905090810190601f1680156200021a5780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019060200280838360005b83811015620002585780820151818401526020810190506200023b565b50505050905001848103825285818151815260200191508051906020019060200280838360005b838110156200029c5780820151818401526020810190506200027f565b50505050905001965050505050505060405180910390f35b348015620002c157600080fd5b50620002f8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506200071b565b604051808215151515815260200191505060405180910390f35b3480156200031f57600080fd5b506200032a620007c1565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156200036f57808201518184015260208101905062000352565b505050509050019250505060405180910390f35b3480156200039057600080fd5b50620003ed600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505062000851565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156200043c57600080fd5b506200044762000991565b6040518082815260200191505060405180910390f35b60008173ffffffffffffffffffffffffffffffffffffffff166369e8094b6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015620004c457600080fd5b505af1158015620004d9573d6000803e3d6000fd5b505050506040513d6020811015620004f057600080fd5b81019080805190602001909291905050509050919050565b6000806000805490509050808310156200055f576000838154811015156200052c57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16915062000564565b600091505b50919050565b60608060608373ffffffffffffffffffffffffffffffffffffffff1663596f21f86040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b158015620005d457600080fd5b505af1158015620005e9573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525060608110156200061457600080fd5b8101908080516401000000008111156200062d57600080fd5b828101905060208101848111156200064457600080fd5b81518560018202830111640100000000821117156200066257600080fd5b505092919060200180516401000000008111156200067f57600080fd5b828101905060208101848111156200069657600080fd5b8151856020820283011164010000000082111715620006b457600080fd5b50509291906020018051640100000000811115620006d157600080fd5b82810190506020810184811115620006e857600080fd5b81518560208202830111640100000000821117156200070657600080fd5b50509291905050509250925092509193909250565b600080600090505b600080549050811015620007b6576000818154811015156200074157fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415620007aa5760019150620007bb565b80600101905062000723565b600091505b50919050565b606060008054806020026020016040519081016040528092919081815260200182805480156200084757602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311620007fc575b5050505050905090565b6000808230620008606200099d565b80806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b83811015620008d1578082015181840152602081019050620008b4565b50505050905090810190601f168015620008ff5780820380516001836020036101000a031916815260200191505b509350505050604051809103906000f08015801562000922573d6000803e3d6000fd5b5090507f8b94c7f6b3fadc764673ea85b4bfef3e17ce928d13e51b818ddfa891ad0f1fcc81604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a180915050919050565b60008080549050905090565b6040516111ea80620009af83390190560060806040523480156200001157600080fd5b50604051620011ea380380620011ea833981018060405281019080805182019291906020018051906020019092919050505080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506200009e33620002d8640100000000026401000000009004565b15620001fc578160009080519060200190620000bc929190620003dc565b5060013390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550507f68a6f2dd22a3385ea4d4fd8b590e4890ff82d99853226c67df889e74dda33e67828260405180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b83811015620001ba5780820151818401526020810190506200019d565b50505050905090810190601f168015620001e85780820380516001836020036101000a031916815260200191505b50935050505060405180910390a1620002d0565b7f5fab2506db5bd68c3de713fb73321249ea57e4130d6fb55c532f61edc6727ab7828260405180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825284818151815260200191508051906020019080838360005b838110156200029357808201518184015260208101905062000276565b50505050905090810190601f168015620002c15780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15b50506200048b565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166363a9c3d7836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b1580156200039857600080fd5b505af1158015620003ad573d6000803e3d6000fd5b505050506040513d6020811015620003c457600080fd5b81019080805190602001909291905050509050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200041f57805160ff191683800117855562000450565b8280016001018555821562000450579182015b828111156200044f57825182559160200191906001019062000432565b5b5090506200045f919062000463565b5090565b6200048891905b80821115620004845760008160009055506001016200046a565b5090565b90565b610d4f806200049b6000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063596f21f81461007257806369e8094b1461019257806394cf795e146101c1578063963671be1461022d578063dc58ab1114610284575b600080fd5b34801561007e57600080fd5b506100876102df565b60405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b838110156100cf5780820151818401526020810190506100b4565b50505050905090810190601f1680156100fc5780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019060200280838360005b8381101561013857808201518184015260208101905061011d565b50505050905001848103825285818151815260200191508051906020019060200280838360005b8381101561017a57808201518184015260208101905061015f565b50505050905001965050505050505060405180910390f35b34801561019e57600080fd5b506101a7610646565b604051808215151515815260200191505060405180910390f35b3480156101cd57600080fd5b506101d66109c6565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156102195780820151818401526020810190506101fe565b505050509050019250505060405180910390f35b34801561023957600080fd5b50610242610bfc565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561029057600080fd5b506102c5600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c22565b604051808215151515815260200191505060405180910390f35b6060806060600060606000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fa69efbd6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561037057600080fd5b505af1158015610384573d6000803e3d6000fd5b505050506040513d602081101561039a57600080fd5b81019080805190602001909291905050509250826040519080825280602002602001820160405280156103dc5781602001602082028038833980820191505090505b509150600090505b8281101561050f57600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633ffefe4e826040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b15801561047d57600080fd5b505af1158015610491573d6000803e3d6000fd5b505050506040513d60208110156104a757600080fd5b810190808051906020019092919050505082828151811015156104c657fe5b9060200190602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505080806001019150506103e4565b6000826001828054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105a95780601f1061057e576101008083540402835291602001916105a9565b820191906000526020600020905b81548152906001019060200180831161058c57829003601f168201915b505050505092508080548060200260200160405190810160405280929190818152602001828054801561063157602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116105e7575b50505050509050955095509550505050909192565b600080600090505b6001805490508110156107985760018181548110151561066a57fe5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561078b577ffcf2bd202c015cb2348d6064af1fe7a3dc404134b63f8950d16e0f52399b06ae600060405180806020018281038252838181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156107745780601f1061074957610100808354040283529160200191610774565b820191906000526020600020905b81548152906001019060200180831161075757829003601f168201915b50509250505060405180910390a1600191506109c2565b808060010191505061064e565b6107a133610c22565b156108cf5760013390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550507f20816d1858bab038efc3fee7f47dbb0de0e5dc3425549996d36d148220973e02600060405180806020018281038252838181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156108b85780601f1061088d576101008083540402835291602001916108b8565b820191906000526020600020905b81548152906001019060200180831161089b57829003601f168201915b50509250505060405180910390a1600191506109c2565b7f4263d1d87ff73aebbce6fa4c04cd5b0a3cfed5c33e8a4fac06979f94842046f260003360405180806020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281038252848181546001816001161561010002031660029004815260200191508054600181600116156101000203166002900480156109ae5780601f10610983576101008083540402835291602001916109ae565b820191906000526020600020905b81548152906001019060200180831161099157829003601f168201915b5050935050505060405180910390a1600091505b5090565b6060600060606000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fa69efbd6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015610a5457600080fd5b505af1158015610a68573d6000803e3d6000fd5b505050506040513d6020811015610a7e57600080fd5b8101908080519060200190929190505050925082604051908082528060200260200182016040528015610ac05781602001602082028038833980820191505090505b509150600090505b82811015610bf357600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633ffefe4e826040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050602060405180830381600087803b158015610b6157600080fd5b505af1158015610b75573d6000803e3d6000fd5b505050506040513d6020811015610b8b57600080fd5b81019080805190602001909291905050508282815181101515610baa57fe5b9060200190602002019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508080600101915050610ac8565b81935050505090565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166363a9c3d7836040518263ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050602060405180830381600087803b158015610ce157600080fd5b505af1158015610cf5573d6000803e3d6000fd5b505050506040513d6020811015610d0b57600080fd5b810190808051906020019092919050505090509190505600a165627a7a723058206cd8bb3b1d442d6ac0bda3bca4fd684c9914be9ba89deff83c6dc6a65b6dc1990029a165627a7a72305820a104c4e4e276ed335729ba24b2591d573850634344d2530c523cf6f83a5455660029"

// DeployEvidenceFactory deploys a new contract, binding an instance of EvidenceFactory to it.
func DeployEvidenceFactory(auth *bind.TransactOpts, backend bind.ContractBackend, evidenceSigners []common.Address) (common.Address, *types.Transaction, *EvidenceFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(EvidenceFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EvidenceFactoryBin), backend, evidenceSigners)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EvidenceFactory{EvidenceFactoryCaller: EvidenceFactoryCaller{contract: contract}, EvidenceFactoryTransactor: EvidenceFactoryTransactor{contract: contract}, EvidenceFactoryFilterer: EvidenceFactoryFilterer{contract: contract}}, nil
}

func AsyncDeployEvidenceFactory(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, evidenceSigners []common.Address) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(EvidenceFactoryABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(EvidenceFactoryBin), backend, evidenceSigners)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// EvidenceFactory is an auto generated Go binding around a Solidity contract.
type EvidenceFactory struct {
	EvidenceFactoryCaller     // Read-only binding to the contract
	EvidenceFactoryTransactor // Write-only binding to the contract
	EvidenceFactoryFilterer   // Log filterer for contract events
}

// EvidenceFactoryCaller is an auto generated read-only Go binding around a Solidity contract.
type EvidenceFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvidenceFactoryTransactor is an auto generated write-only Go binding around a Solidity contract.
type EvidenceFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvidenceFactoryFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type EvidenceFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EvidenceFactorySession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type EvidenceFactorySession struct {
	Contract     *EvidenceFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EvidenceFactoryCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type EvidenceFactoryCallerSession struct {
	Contract *EvidenceFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// EvidenceFactoryTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type EvidenceFactoryTransactorSession struct {
	Contract     *EvidenceFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EvidenceFactoryRaw is an auto generated low-level Go binding around a Solidity contract.
type EvidenceFactoryRaw struct {
	Contract *EvidenceFactory // Generic contract binding to access the raw methods on
}

// EvidenceFactoryCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type EvidenceFactoryCallerRaw struct {
	Contract *EvidenceFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// EvidenceFactoryTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type EvidenceFactoryTransactorRaw struct {
	Contract *EvidenceFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEvidenceFactory creates a new instance of EvidenceFactory, bound to a specific deployed contract.
func NewEvidenceFactory(address common.Address, backend bind.ContractBackend) (*EvidenceFactory, error) {
	contract, err := bindEvidenceFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EvidenceFactory{EvidenceFactoryCaller: EvidenceFactoryCaller{contract: contract}, EvidenceFactoryTransactor: EvidenceFactoryTransactor{contract: contract}, EvidenceFactoryFilterer: EvidenceFactoryFilterer{contract: contract}}, nil
}

// NewEvidenceFactoryCaller creates a new read-only instance of EvidenceFactory, bound to a specific deployed contract.
func NewEvidenceFactoryCaller(address common.Address, caller bind.ContractCaller) (*EvidenceFactoryCaller, error) {
	contract, err := bindEvidenceFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EvidenceFactoryCaller{contract: contract}, nil
}

// NewEvidenceFactoryTransactor creates a new write-only instance of EvidenceFactory, bound to a specific deployed contract.
func NewEvidenceFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*EvidenceFactoryTransactor, error) {
	contract, err := bindEvidenceFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EvidenceFactoryTransactor{contract: contract}, nil
}

// NewEvidenceFactoryFilterer creates a new log filterer instance of EvidenceFactory, bound to a specific deployed contract.
func NewEvidenceFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*EvidenceFactoryFilterer, error) {
	contract, err := bindEvidenceFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EvidenceFactoryFilterer{contract: contract}, nil
}

// bindEvidenceFactory binds a generic wrapper to an already deployed contract.
func bindEvidenceFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EvidenceFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EvidenceFactory *EvidenceFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EvidenceFactory.Contract.EvidenceFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EvidenceFactory *EvidenceFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _EvidenceFactory.Contract.EvidenceFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EvidenceFactory *EvidenceFactoryRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _EvidenceFactory.Contract.EvidenceFactoryTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EvidenceFactory *EvidenceFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EvidenceFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EvidenceFactory *EvidenceFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _EvidenceFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EvidenceFactory *EvidenceFactoryTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _EvidenceFactory.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// GetEvidence is a free data retrieval call binding the contract method 0x60004bd1.
//
// Solidity: function getEvidence(address addr) constant returns(string, address[], address[])
func (_EvidenceFactory *EvidenceFactoryCaller) GetEvidence(opts *bind.CallOpts, addr common.Address) (string, []common.Address, []common.Address, error) {
	var (
		ret0 = new(string)
		ret1 = new([]common.Address)
		ret2 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _EvidenceFactory.contract.Call(opts, out, "getEvidence", addr)
	return *ret0, *ret1, *ret2, err
}

// GetEvidence is a free data retrieval call binding the contract method 0x60004bd1.
//
// Solidity: function getEvidence(address addr) constant returns(string, address[], address[])
func (_EvidenceFactory *EvidenceFactorySession) GetEvidence(addr common.Address) (string, []common.Address, []common.Address, error) {
	return _EvidenceFactory.Contract.GetEvidence(&_EvidenceFactory.CallOpts, addr)
}

// GetEvidence is a free data retrieval call binding the contract method 0x60004bd1.
//
// Solidity: function getEvidence(address addr) constant returns(string, address[], address[])
func (_EvidenceFactory *EvidenceFactoryCallerSession) GetEvidence(addr common.Address) (string, []common.Address, []common.Address, error) {
	return _EvidenceFactory.Contract.GetEvidence(&_EvidenceFactory.CallOpts, addr)
}

// GetSigner is a free data retrieval call binding the contract method 0x3ffefe4e.
//
// Solidity: function getSigner(uint256 index) constant returns(address)
func (_EvidenceFactory *EvidenceFactoryCaller) GetSigner(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EvidenceFactory.contract.Call(opts, out, "getSigner", index)
	return *ret0, err
}

// GetSigner is a free data retrieval call binding the contract method 0x3ffefe4e.
//
// Solidity: function getSigner(uint256 index) constant returns(address)
func (_EvidenceFactory *EvidenceFactorySession) GetSigner(index *big.Int) (common.Address, error) {
	return _EvidenceFactory.Contract.GetSigner(&_EvidenceFactory.CallOpts, index)
}

// GetSigner is a free data retrieval call binding the contract method 0x3ffefe4e.
//
// Solidity: function getSigner(uint256 index) constant returns(address)
func (_EvidenceFactory *EvidenceFactoryCallerSession) GetSigner(index *big.Int) (common.Address, error) {
	return _EvidenceFactory.Contract.GetSigner(&_EvidenceFactory.CallOpts, index)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() constant returns(address[])
func (_EvidenceFactory *EvidenceFactoryCaller) GetSigners(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _EvidenceFactory.contract.Call(opts, out, "getSigners")
	return *ret0, err
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() constant returns(address[])
func (_EvidenceFactory *EvidenceFactorySession) GetSigners() ([]common.Address, error) {
	return _EvidenceFactory.Contract.GetSigners(&_EvidenceFactory.CallOpts)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() constant returns(address[])
func (_EvidenceFactory *EvidenceFactoryCallerSession) GetSigners() ([]common.Address, error) {
	return _EvidenceFactory.Contract.GetSigners(&_EvidenceFactory.CallOpts)
}

// GetSignersSize is a free data retrieval call binding the contract method 0xfa69efbd.
//
// Solidity: function getSignersSize() constant returns(uint256)
func (_EvidenceFactory *EvidenceFactoryCaller) GetSignersSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EvidenceFactory.contract.Call(opts, out, "getSignersSize")
	return *ret0, err
}

// GetSignersSize is a free data retrieval call binding the contract method 0xfa69efbd.
//
// Solidity: function getSignersSize() constant returns(uint256)
func (_EvidenceFactory *EvidenceFactorySession) GetSignersSize() (*big.Int, error) {
	return _EvidenceFactory.Contract.GetSignersSize(&_EvidenceFactory.CallOpts)
}

// GetSignersSize is a free data retrieval call binding the contract method 0xfa69efbd.
//
// Solidity: function getSignersSize() constant returns(uint256)
func (_EvidenceFactory *EvidenceFactoryCallerSession) GetSignersSize() (*big.Int, error) {
	return _EvidenceFactory.Contract.GetSignersSize(&_EvidenceFactory.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0x63a9c3d7.
//
// Solidity: function verify(address addr) constant returns(bool)
func (_EvidenceFactory *EvidenceFactoryCaller) Verify(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EvidenceFactory.contract.Call(opts, out, "verify", addr)
	return *ret0, err
}

// Verify is a free data retrieval call binding the contract method 0x63a9c3d7.
//
// Solidity: function verify(address addr) constant returns(bool)
func (_EvidenceFactory *EvidenceFactorySession) Verify(addr common.Address) (bool, error) {
	return _EvidenceFactory.Contract.Verify(&_EvidenceFactory.CallOpts, addr)
}

// Verify is a free data retrieval call binding the contract method 0x63a9c3d7.
//
// Solidity: function verify(address addr) constant returns(bool)
func (_EvidenceFactory *EvidenceFactoryCallerSession) Verify(addr common.Address) (bool, error) {
	return _EvidenceFactory.Contract.Verify(&_EvidenceFactory.CallOpts, addr)
}

// AddSignatures is a paid mutator transaction binding the contract method 0x203e4922.
//
// Solidity: function addSignatures(address addr) returns(bool)
func (_EvidenceFactory *EvidenceFactoryTransactor) AddSignatures(opts *bind.TransactOpts, addr common.Address) (bool, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	transaction, receipt, err := _EvidenceFactory.contract.TransactWithResult(opts, out, "addSignatures", addr)
	return *ret0, transaction, receipt, err
}

func (_EvidenceFactory *EvidenceFactoryTransactor) AsyncAddSignatures(handler func(*types.Receipt, error), opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _EvidenceFactory.contract.AsyncTransact(opts, handler, "addSignatures", addr)
}

// AddSignatures is a paid mutator transaction binding the contract method 0x203e4922.
//
// Solidity: function addSignatures(address addr) returns(bool)
func (_EvidenceFactory *EvidenceFactorySession) AddSignatures(addr common.Address) (bool, *types.Transaction, *types.Receipt, error) {
	return _EvidenceFactory.Contract.AddSignatures(&_EvidenceFactory.TransactOpts, addr)
}

func (_EvidenceFactory *EvidenceFactorySession) AsyncAddSignatures(handler func(*types.Receipt, error), addr common.Address) (*types.Transaction, error) {
	return _EvidenceFactory.Contract.AsyncAddSignatures(handler, &_EvidenceFactory.TransactOpts, addr)
}

// AddSignatures is a paid mutator transaction binding the contract method 0x203e4922.
//
// Solidity: function addSignatures(address addr) returns(bool)
func (_EvidenceFactory *EvidenceFactoryTransactorSession) AddSignatures(addr common.Address) (bool, *types.Transaction, *types.Receipt, error) {
	return _EvidenceFactory.Contract.AddSignatures(&_EvidenceFactory.TransactOpts, addr)
}

func (_EvidenceFactory *EvidenceFactoryTransactorSession) AsyncAddSignatures(handler func(*types.Receipt, error), addr common.Address) (*types.Transaction, error) {
	return _EvidenceFactory.Contract.AsyncAddSignatures(handler, &_EvidenceFactory.TransactOpts, addr)
}

// NewEvidence is a paid mutator transaction binding the contract method 0xa3aaada5.
//
// Solidity: function newEvidence(string evi) returns(address)
func (_EvidenceFactory *EvidenceFactoryTransactor) NewEvidence(opts *bind.TransactOpts, evi string) (common.Address, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	transaction, receipt, err := _EvidenceFactory.contract.TransactWithResult(opts, out, "newEvidence", evi)
	return *ret0, transaction, receipt, err
}

func (_EvidenceFactory *EvidenceFactoryTransactor) AsyncNewEvidence(handler func(*types.Receipt, error), opts *bind.TransactOpts, evi string) (*types.Transaction, error) {
	return _EvidenceFactory.contract.AsyncTransact(opts, handler, "newEvidence", evi)
}

// NewEvidence is a paid mutator transaction binding the contract method 0xa3aaada5.
//
// Solidity: function newEvidence(string evi) returns(address)
func (_EvidenceFactory *EvidenceFactorySession) NewEvidence(evi string) (common.Address, *types.Transaction, *types.Receipt, error) {
	return _EvidenceFactory.Contract.NewEvidence(&_EvidenceFactory.TransactOpts, evi)
}

func (_EvidenceFactory *EvidenceFactorySession) AsyncNewEvidence(handler func(*types.Receipt, error), evi string) (*types.Transaction, error) {
	return _EvidenceFactory.Contract.AsyncNewEvidence(handler, &_EvidenceFactory.TransactOpts, evi)
}

// NewEvidence is a paid mutator transaction binding the contract method 0xa3aaada5.
//
// Solidity: function newEvidence(string evi) returns(address)
func (_EvidenceFactory *EvidenceFactoryTransactorSession) NewEvidence(evi string) (common.Address, *types.Transaction, *types.Receipt, error) {
	return _EvidenceFactory.Contract.NewEvidence(&_EvidenceFactory.TransactOpts, evi)
}

func (_EvidenceFactory *EvidenceFactoryTransactorSession) AsyncNewEvidence(handler func(*types.Receipt, error), evi string) (*types.Transaction, error) {
	return _EvidenceFactory.Contract.AsyncNewEvidence(handler, &_EvidenceFactory.TransactOpts, evi)
}

// EvidenceFactoryNewEvidenceEvent represents a NewEvidenceEvent event raised by the EvidenceFactory contract.
type EvidenceFactoryNewEvidenceEvent struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// WatchNewEvidenceEvent is a free log subscription operation binding the contract event 0x8b94c7f6b3fadc764673ea85b4bfef3e17ce928d13e51b818ddfa891ad0f1fcc.
//
// Solidity: event newEvidenceEvent(address addr)
func (_EvidenceFactory *EvidenceFactoryFilterer) WatchNewEvidenceEvent(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _EvidenceFactory.contract.WatchLogs(fromBlock, handler, "newEvidenceEvent")
}

func (_EvidenceFactory *EvidenceFactoryFilterer) WatchAllNewEvidenceEvent(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _EvidenceFactory.contract.WatchLogs(fromBlock, handler, "newEvidenceEvent")
}

// ParseNewEvidenceEvent is a log parse operation binding the contract event 0x8b94c7f6b3fadc764673ea85b4bfef3e17ce928d13e51b818ddfa891ad0f1fcc.
//
// Solidity: event newEvidenceEvent(address addr)
func (_EvidenceFactory *EvidenceFactoryFilterer) ParseNewEvidenceEvent(log types.Log) (*EvidenceFactoryNewEvidenceEvent, error) {
	event := new(EvidenceFactoryNewEvidenceEvent)
	if err := _EvidenceFactory.contract.UnpackLog(event, "newEvidenceEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchNewEvidenceEvent is a free log subscription operation binding the contract event 0x8b94c7f6b3fadc764673ea85b4bfef3e17ce928d13e51b818ddfa891ad0f1fcc.
//
// Solidity: event newEvidenceEvent(address addr)
func (_EvidenceFactory *EvidenceFactorySession) WatchNewEvidenceEvent(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _EvidenceFactory.Contract.WatchNewEvidenceEvent(fromBlock, handler)
}

func (_EvidenceFactory *EvidenceFactorySession) WatchAllNewEvidenceEvent(fromBlock *uint64, handler func(int, []types.Log)) (string, error) {
	return _EvidenceFactory.Contract.WatchAllNewEvidenceEvent(fromBlock, handler)
}

// ParseNewEvidenceEvent is a log parse operation binding the contract event 0x8b94c7f6b3fadc764673ea85b4bfef3e17ce928d13e51b818ddfa891ad0f1fcc.
//
// Solidity: event newEvidenceEvent(address addr)
func (_EvidenceFactory *EvidenceFactorySession) ParseNewEvidenceEvent(log types.Log) (*EvidenceFactoryNewEvidenceEvent, error) {
	return _EvidenceFactory.Contract.ParseNewEvidenceEvent(log)
}
