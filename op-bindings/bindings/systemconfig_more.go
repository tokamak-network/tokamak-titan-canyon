// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const SystemConfigStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1003,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1005,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"overhead\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"scalar\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"batcherHash\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_bytes32\"},{\"astId\":1008,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"gasLimit\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_uint64\"},{\"astId\":1009,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_resourceConfig\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_struct(ResourceConfig)1011_storage\"},{\"astId\":1010,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"startBlock\",\"offset\":0,\"slot\":\"106\",\"type\":\"t_uint256\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_struct(ResourceConfig)1011_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ResourceMetering.ResourceConfig\",\"numberOfBytes\":\"32\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var SystemConfigStorageLayout = new(solc.StorageLayout)

var SystemConfigDeployedBin = "0x608060405234801561001057600080fd5b506004361061025c5760003560e01c8063935f029e11610145578063dac6e63a116100bd578063f45e65d81161008c578063f8c68de011610071578063f8c68de0146105bb578063fd32aa0f146105c3578063ffa1ad74146105cb57600080fd5b8063f45e65d81461059e578063f68016b7146105a757600080fd5b8063dac6e63a14610567578063e4a88b9f1461056f578063e81b2c6d14610582578063f2fde38b1461058b57600080fd5b8063bc49ce5f11610114578063c71973f6116100f9578063c71973f61461040d578063c9b26f6114610420578063cc731b021461043357600080fd5b8063bc49ce5f146103fd578063c4e8ddfa1461040557600080fd5b8063935f029e146103c75780639b7d7f0a146103da578063a7119869146103e2578063b40a817c146103ea57600080fd5b80634d0047ee116101d85780635d73369c116101a7578063697844c61161018c578063697844c614610399578063715018a6146103a15780638da5cb5b146103a957600080fd5b80635d73369c1461038957806361d157681461039157600080fd5b80634d0047ee146103095780634d9f1559146103115780634f16540b1461031957806354fd4d501461034057600080fd5b806318d139181161022f5780631fd19ee1116102145780631fd19ee1146102d757806348cd4cb1146102df5780634add321d146102e857600080fd5b806318d13918146102ba57806319f5cea8146102cf57600080fd5b806306c9265714610261578063078f29cf1461027c5780630a49cb03146102a95780630c18c162146102b1575b600080fd5b6102696105d3565b6040519081526020015b60405180910390f35b610284610601565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610273565b61028461063a565b61026960655481565b6102cd6102c836600461174f565b61066a565b005b61026961067e565b6102846106a9565b610269606a5481565b6102f06106d3565b60405167ffffffffffffffff9091168152602001610273565b6102846106f9565b610284610729565b6102697f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0881565b61037c6040518060400160405280600681526020017f312e31302e30000000000000000000000000000000000000000000000000000081525081565b60405161027391906117dc565b610269610759565b610269610784565b6102696107af565b6102cd6107da565b60335473ffffffffffffffffffffffffffffffffffffffff16610284565b6102cd6103d53660046117ef565b6107ee565b610284610804565b610284610834565b6102cd6103f8366004611829565b610864565b610269610875565b6102846108a0565b6102cd61041b36600461199a565b6108d0565b6102cd61042e3660046119b6565b6108e1565b6104f76040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c08101825260695463ffffffff8082168352640100000000820460ff9081166020850152650100000000008304169383019390935266010000000000008104831660608301526a0100000000000000000000810490921660808201526e0100000000000000000000000000009091046fffffffffffffffffffffffffffffffff1660a082015290565b6040516102739190600060c08201905063ffffffff80845116835260ff602085015116602084015260ff6040850151166040840152806060850151166060840152806080850151166080840152506fffffffffffffffffffffffffffffffff60a08401511660a083015292915050565b6102846108f2565b6102cd61057d3660046119cf565b610922565b61026960675481565b6102cd61059936600461174f565b610cef565b61026960665481565b6068546102f09067ffffffffffffffff1681565b610269610da3565b610269610dce565b610269600081565b6105fe60017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d611b48565b81565b600061063561063160017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad6377611b48565b5490565b905090565b600061063561063160017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad611b48565b610672610dfd565b61067b81610e7e565b50565b6105fe60017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a8611b48565b60006106357f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c085490565b6069546000906106359063ffffffff6a0100000000000000000000820481169116611b5f565b600061063561063160017fe1e3a95fb10ed56538cc130c2250de9823e7716d1142b8521655d7f7317b8ef1611b48565b600061063561063160017fe52a667f71ec761b9b381c7b76ca9b852adf7e8905da0e0ad49986a0a6871816611b48565b6105fe60017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce9580637611b48565b6105fe60017fe52a667f71ec761b9b381c7b76ca9b852adf7e8905da0e0ad49986a0a6871816611b48565b6105fe60017fe1e3a95fb10ed56538cc130c2250de9823e7716d1142b8521655d7f7317b8ef1611b48565b6107e2610dfd565b6107ec6000610f3b565b565b6107f6610dfd565b6108008282610fb2565b5050565b600061063561063160017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d611b48565b600061063561063160017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce9580637611b48565b61086c610dfd565b61067b81611043565b6105fe60017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc598611b48565b600061063561063160017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a8611b48565b6108d8610dfd565b61067b81611121565b6108e9610dfd565b61067b81611595565b600061063561063160017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc598611b48565b600054600390610100900460ff16158015610944575060005460ff8083169116105b6109d5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff831617610100179055610a0e6115bd565b610a178b610cef565b610a2088611595565b610a2a8a8a610fb2565b610a3387611043565b610a3c86610e7e565b610a6f610a6a60017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc598611b48565b849055565b610aa3610a9d60017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce9580637611b48565b83519055565b610ada610ad160017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a8611b48565b60208401519055565b610b11610b0860017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad6377611b48565b60408401519055565b610b48610b3f60017fe52a667f71ec761b9b381c7b76ca9b852adf7e8905da0e0ad49986a0a6871816611b48565b60608401519055565b610b7f610b7660017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad611b48565b60808401519055565b610bb6610bad60017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d611b48565b60a08401519055565b610bed610be460017fe1e3a95fb10ed56538cc130c2250de9823e7716d1142b8521655d7f7317b8ef1611b48565b60c08401519055565b610bf68461165c565b610bff85611121565b610c076106d3565b67ffffffffffffffff168767ffffffffffffffff161015610c84576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064016109cc565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050505050505050505050565b610cf7610dfd565b73ffffffffffffffffffffffffffffffffffffffff8116610d9a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016109cc565b61067b81610f3b565b6105fe60017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad6377611b48565b6105fe60017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad611b48565b9055565b60335473ffffffffffffffffffffffffffffffffffffffff1633146107ec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016109cc565b610ea77f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c08829055565b6040805173ffffffffffffffffffffffffffffffffffffffff8316602082015260009101604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052905060035b60007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be83604051610f2f91906117dc565b60405180910390a35050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b606582905560668190556040805160208101849052908101829052600090606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529050600160007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be8360405161103691906117dc565b60405180910390a3505050565b61104b6106d3565b67ffffffffffffffff168167ffffffffffffffff1610156110c8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064016109cc565b606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff83169081179091556040805160208082019390935281518082039093018352810190526002610efe565b8060a001516fffffffffffffffffffffffffffffffff16816060015163ffffffff1611156111d1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f53797374656d436f6e6669673a206d696e206261736520666565206d7573742060448201527f6265206c657373207468616e206d61782062617365000000000000000000000060648201526084016109cc565b6001816040015160ff1611611268576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a2064656e6f6d696e61746f72206d757374206260448201527f65206c6172676572207468616e2031000000000000000000000000000000000060648201526084016109cc565b6068546080820151825167ffffffffffffffff909216916112899190611b8b565b63ffffffff1611156112f7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064016109cc565b6000816020015160ff161161138e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a20656c6173746963697479206d756c7469706c60448201527f6965722063616e6e6f742062652030000000000000000000000000000000000060648201526084016109cc565b8051602082015163ffffffff82169160ff909116906113ae908290611baa565b6113b89190611bf4565b63ffffffff161461144b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f53797374656d436f6e6669673a20707265636973696f6e206c6f73732077697460448201527f6820746172676574207265736f75726365206c696d697400000000000000000060648201526084016109cc565b805160698054602084015160408501516060860151608087015160a09097015163ffffffff9687167fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000009095169490941764010000000060ff94851602177fffffffffffffffffffffffffffffffffffffffffffff0000000000ffffffffff166501000000000093909216929092027fffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffff1617660100000000000091851691909102177fffff0000000000000000000000000000000000000000ffffffffffffffffffff166a010000000000000000000093909416929092027fffff00000000000000000000000000000000ffffffffffffffffffffffffffff16929092176e0100000000000000000000000000006fffffffffffffffffffffffffffffffff90921691909102179055565b6067819055604080516020808201849052825180830390910181529082019091526000610efe565b600054610100900460ff16611654576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016109cc565b6107ec611686565b801580159061166b5750606a54155b1561167557606a55565b606a5460000361067b5743606a5550565b600054610100900460ff1661171d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016109cc565b6107ec33610f3b565b803573ffffffffffffffffffffffffffffffffffffffff8116811461174a57600080fd5b919050565b60006020828403121561176157600080fd5b61176a82611726565b9392505050565b6000815180845260005b818110156117975760208185018101518683018201520161177b565b818111156117a9576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061176a6020830184611771565b6000806040838503121561180257600080fd5b50508035926020909101359150565b803567ffffffffffffffff8116811461174a57600080fd5b60006020828403121561183b57600080fd5b61176a82611811565b60405160e0810167ffffffffffffffff8111828210171561188e577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405290565b803563ffffffff8116811461174a57600080fd5b803560ff8116811461174a57600080fd5b600060c082840312156118cb57600080fd5b60405160c0810181811067ffffffffffffffff82111715611915577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405290508061192483611894565b8152611932602084016118a8565b6020820152611943604084016118a8565b604082015261195460608401611894565b606082015261196560808401611894565b608082015260a08301356fffffffffffffffffffffffffffffffff8116811461198d57600080fd5b60a0919091015292915050565b600060c082840312156119ac57600080fd5b61176a83836118b9565b6000602082840312156119c857600080fd5b5035919050565b6000806000806000806000806000808a8c036102a08112156119f057600080fd5b6119f98c611726565b9a5060208c0135995060408c0135985060608c01359750611a1c60808d01611811565b9650611a2a60a08d01611726565b9550611a398d60c08e016118b9565b94506101808c01359350611a506101a08d01611726565b925060e07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe4082011215611a8257600080fd5b50611a8b611844565b611a986101c08d01611726565b8152611aa76101e08d01611726565b6020820152611ab96102008d01611726565b6040820152611acb6102208d01611726565b6060820152611add6102408d01611726565b6080820152611aef6102608d01611726565b60a0820152611b016102808d01611726565b60c0820152809150509295989b9194979a5092959850565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015611b5a57611b5a611b19565b500390565b600067ffffffffffffffff808316818516808303821115611b8257611b82611b19565b01949350505050565b600063ffffffff808316818516808303821115611b8257611b82611b19565b600063ffffffff80841680611be8577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600063ffffffff80831681851681830481118215151615611c1757611c17611b19565b0294935050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(SystemConfigStorageLayoutJSON), SystemConfigStorageLayout); err != nil {
		panic(err)
	}

	layouts["SystemConfig"] = SystemConfigStorageLayout
	deployedBytecodes["SystemConfig"] = SystemConfigDeployedBin
}
