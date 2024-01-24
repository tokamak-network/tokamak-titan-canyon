// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1002,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":21,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1003,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_1_0_1600\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1004,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_51_0_20\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1005,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_52_0_1568\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1006,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_101_0_1\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":1007,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_102_0_1568\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1008,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_151_0_32\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_152_0_1568\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1010,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_201_0_32\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1011,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_202_0_32\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1012,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1013,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":1014,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_uint240\"},{\"astId\":1015,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"failedMessages\",\"offset\":0,\"slot\":\"206\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1016,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"207\",\"type\":\"t_array(t_uint256)42_storage\"},{\"astId\":1017,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"PORTAL\",\"offset\":0,\"slot\":\"249\",\"type\":\"t_contract(OptimismPortal)1019\"},{\"astId\":1018,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"l1TONAddress\",\"offset\":0,\"slot\":\"250\",\"type\":\"t_address\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)42_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[42]\",\"numberOfBytes\":\"1344\",\"base\":\"t_uint256\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_contract(OptimismPortal)1019\":{\"encoding\":\"inplace\",\"label\":\"contract OptimismPortal\",\"numberOfBytes\":\"20\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L1CrossDomainMessengerDeployedBin = "0x6080604052600436106101755760003560e01c80636425666b116100cb578063a4e7f8bd1161007f578063cb8398bf11610059578063cb8398bf14610450578063d764ad0b14610470578063ecc704281461048357600080fd5b8063a4e7f8bd146103c0578063b1b1b20914610400578063b28ade251461043057600080fd5b806383a74074116100b057806383a74074146103755780638cbeeef2146102b35780639fce812c1461038c57600080fd5b80636425666b146103355780636e296e451461036057600080fd5b80633dbb202b1161012d5780634c1d6a69116101075780634c1d6a69146102b357806354fd4d50146102c95780635644cfdf1461031f57600080fd5b80633dbb202b146102565780633f827a5a1461026b578063485cc9551461029357600080fd5b80630ff754ea1161015e5780630ff754ea146101c257806312fd7963146102145780632828d7e81461024157600080fd5b8063028f85f71461017a5780630c568498146101ad575b600080fd5b34801561018657600080fd5b5061018f601081565b60405167ffffffffffffffff90911681526020015b60405180910390f35b3480156101b957600080fd5b5061018f603f81565b3480156101ce57600080fd5b5060f9546101ef9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a4565b34801561022057600080fd5b5060fa546101ef9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561024d57600080fd5b5061018f604081565b610269610264366004611f2d565b6104e8565b005b34801561027757600080fd5b50610280600181565b60405161ffff90911681526020016101a4565b34801561029f57600080fd5b506102696102ae366004611f94565b61074c565b3480156102bf57600080fd5b5061018f619c4081565b3480156102d557600080fd5b506103126040518060400160405280600581526020017f312e372e3100000000000000000000000000000000000000000000000000000081525081565b6040516101a49190612043565b34801561032b57600080fd5b5061018f61138881565b34801561034157600080fd5b5060f95473ffffffffffffffffffffffffffffffffffffffff166101ef565b34801561036c57600080fd5b506101ef610a6e565b34801561038157600080fd5b5061018f62030d4081565b34801561039857600080fd5b506101ef7f000000000000000000000000000000000000000000000000000000000000000081565b3480156103cc57600080fd5b506103f06103db366004612056565b60ce6020526000908152604090205460ff1681565b60405190151581526020016101a4565b34801561040c57600080fd5b506103f061041b366004612056565b60cb6020526000908152604090205460ff1681565b34801561043c57600080fd5b5061018f61044b36600461206f565b610b55565b34801561045c57600080fd5b5061026961046b3660046120c3565b610bc5565b61026961047e366004612134565b610e2a565b34801561048f57600080fd5b506104da60cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b6040519081526020016101a4565b6106217f0000000000000000000000000000000000000000000000000000000000000000610517858585610b55565b347fd764ad0b0000000000000000000000000000000000000000000000000000000061058360cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c60405160240161059f9796959493929190612203565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526116b6565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a3385856106a660cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b866040516106b8959493929190612262565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b6000546003907501000000000000000000000000000000000000000000900460ff1615801561079a575060005460ff8083167401000000000000000000000000000000000000000090920416105b61082b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff60ff84167401000000000000000000000000000000000000000002167fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff90911617750100000000000000000000000000000000000000000017905560f9805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560fa80549285169290911691909117905561090861174d565b60f95473ffffffffffffffffffffffffffffffffffffffff1615801590610946575060fa5473ffffffffffffffffffffffffffffffffffffffff1615155b15610a0b5760fa5460f9546040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff602482015291169063095ea7b3906044016020604051808303816000875af11580156109e5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a0991906122b0565b505b600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b60cc5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff215301610b38576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f742073657400000000000000000000006064820152608401610822565b5060cc5473ffffffffffffffffffffffffffffffffffffffff1690565b6000611388619c4080603f610b71604063ffffffff8816612301565b610b7b9190612331565b610b86601088612301565b610b939062030d4061237f565b610b9d919061237f565b610ba7919061237f565b610bb1919061237f565b610bbb919061237f565b90505b9392505050565b610cfe7f0000000000000000000000000000000000000000000000000000000000000000610bf4858585610b55565b867fd764ad0b00000000000000000000000000000000000000000000000000000000610c6060cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338b8b898c8c604051602401610c7c9796959493929190612203565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611826565b8473ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a338585610d8360cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b86604051610d95959493929190612262565b60405180910390a260405184815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff000000000000000000000000000000000000000000000000000000000000909116179055505050565b60f087901c60028110610ee5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604d60248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2030206f722031206d657373616765732061726520737570706f7274656460648201527f20617420746869732074696d6500000000000000000000000000000000000000608482015260a401610822565b8061ffff16600003610fda576000610f36878986868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508f92506118aa915050565b600081815260cb602052604090205490915060ff1615610fd8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a206c65676163792077697460448201527f6864726177616c20616c72656164792072656c617965640000000000000000006064820152608401610822565b505b6000611020898989898989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506118c992505050565b905061102a6118ec565b156110625785341461103e5761103e6123ab565b600081815260ce602052604090205460ff161561105d5761105d6123ab565b6111b4565b3415611116576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a401610822565b600081815260ce602052604090205460ff166111b4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c61796564000000000000000000000000000000006064820152608401610822565b6111bd876119e2565b15611270576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a401610822565b600081815260cb602052604090205460ff161561130f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c61796564000000000000000000006064820152608401610822565b61133085611321611388619c4061237f565b67ffffffffffffffff16611a28565b1580611356575060cc5473ffffffffffffffffffffffffffffffffffffffff1661dead14155b1561146f57600081815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555182917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff3201611468576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d657373616765000000000000000000000000000000000000006064820152608401610822565b50506116a8565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a16179055600061150088619c405a6114c391906123da565b8988888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611a4692505050565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead1790559050801561159757600082815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a26116a4565b600082815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff32016116a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d657373616765000000000000000000000000000000000000006064820152608401610822565b5050505b50505050505050565b905090565b60f9546040517fe9e05c4200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063e9e05c42906117159087908690889060009088906004016123f1565b600060405180830381600087803b15801561172f57600080fd5b505af1158015611743573d6000803e3d6000fd5b5050505050505050565b6000547501000000000000000000000000000000000000000000900460ff166117f8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610822565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055565b60fa5461184b9073ffffffffffffffffffffffffffffffffffffffff16333085611a60565b60f9546040517f0608ee3a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690630608ee3a906117159087908690889060009088906004016123f1565b60006118b885858585611afb565b805190602001209050949350505050565b60006118d9878787878787611b94565b8051906020012090509695505050505050565b60f95460009073ffffffffffffffffffffffffffffffffffffffff16331480156116b1575060f954604080517f9bf62d82000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000008116931691639bf62d829160048083019260209291908290030181865afa1580156119a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119c6919061243e565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b600073ffffffffffffffffffffffffffffffffffffffff8216301480611a22575060f95473ffffffffffffffffffffffffffffffffffffffff8381169116145b92915050565b600080603f83619c4001026040850201603f5a021015949350505050565b600080600080845160208601878a8af19695505050505050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000179052611af5908590611c33565b50505050565b606084848484604051602401611b14949392919061245b565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcbd4ece9000000000000000000000000000000000000000000000000000000001790529050949350505050565b6060868686868686604051602401611bb1969594939291906124a5565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b6000611c95826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16611d449092919063ffffffff16565b805190915015611d3f5780806020019051810190611cb391906122b0565b611d3f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610822565b505050565b6060610bbb84846000858573ffffffffffffffffffffffffffffffffffffffff85163b611dcd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610822565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051611df691906124fc565b60006040518083038185875af1925050503d8060008114611e33576040519150601f19603f3d011682016040523d82523d6000602084013e611e38565b606091505b5091509150611e48828286611e53565b979650505050505050565b60608315611e62575081610bbe565b825115611e725782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108229190612043565b73ffffffffffffffffffffffffffffffffffffffff81168114611ec857600080fd5b50565b60008083601f840112611edd57600080fd5b50813567ffffffffffffffff811115611ef557600080fd5b602083019150836020828501011115611f0d57600080fd5b9250929050565b803563ffffffff81168114611f2857600080fd5b919050565b60008060008060608587031215611f4357600080fd5b8435611f4e81611ea6565b9350602085013567ffffffffffffffff811115611f6a57600080fd5b611f7687828801611ecb565b9094509250611f89905060408601611f14565b905092959194509250565b60008060408385031215611fa757600080fd5b8235611fb281611ea6565b91506020830135611fc281611ea6565b809150509250929050565b60005b83811015611fe8578181015183820152602001611fd0565b83811115611af55750506000910152565b60008151808452612011816020860160208601611fcd565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610bbe6020830184611ff9565b60006020828403121561206857600080fd5b5035919050565b60008060006040848603121561208457600080fd5b833567ffffffffffffffff81111561209b57600080fd5b6120a786828701611ecb565b90945092506120ba905060208501611f14565b90509250925092565b6000806000806000608086880312156120db57600080fd5b85356120e681611ea6565b945060208601359350604086013567ffffffffffffffff81111561210957600080fd5b61211588828901611ecb565b9094509250612128905060608701611f14565b90509295509295909350565b600080600080600080600060c0888a03121561214f57600080fd5b87359650602088013561216181611ea6565b9550604088013561217181611ea6565b9450606088013593506080880135925060a088013567ffffffffffffffff81111561219b57600080fd5b6121a78a828b01611ecb565b989b979a50959850939692959293505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a083015261225560c0830184866121ba565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff861681526080602082015260006122926080830186886121ba565b905083604083015263ffffffff831660608301529695505050505050565b6000602082840312156122c257600080fd5b81518015158114610bbe57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff80831681851681830481118215151615612328576123286122d2565b02949350505050565b600067ffffffffffffffff80841680612373577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600067ffffffffffffffff8083168185168083038211156123a2576123a26122d2565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b6000828210156123ec576123ec6122d2565b500390565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015267ffffffffffffffff84166040820152821515606082015260a060808201526000611e4860a0830184611ff9565b60006020828403121561245057600080fd5b8151610bbe81611ea6565b600073ffffffffffffffffffffffffffffffffffffffff8087168352808616602084015250608060408301526124946080830185611ff9565b905082606083015295945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a08301526124f060c0830184611ff9565b98975050505050505050565b6000825161250e818460208701611fcd565b919091019291505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1CrossDomainMessengerStorageLayoutJSON), L1CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1CrossDomainMessenger"] = L1CrossDomainMessengerStorageLayout
	deployedBytecodes["L1CrossDomainMessenger"] = L1CrossDomainMessengerDeployedBin
}
