// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/tokamak-network/tokamak-thanos/op-bindings/solc"
)

const L1CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1002,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":21,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1003,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_1_0_1600\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1004,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_51_0_20\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1005,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_52_0_1568\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1006,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_101_0_1\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":1007,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_102_0_1568\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1008,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_151_0_32\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":1009,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_152_0_1568\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1010,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_201_0_32\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1011,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_202_0_32\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1012,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1013,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":1014,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_uint240\"},{\"astId\":1015,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"failedMessages\",\"offset\":0,\"slot\":\"206\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1016,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"otherMessenger\",\"offset\":0,\"slot\":\"207\",\"type\":\"t_contract(CrossDomainMessenger)1021\"},{\"astId\":1017,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"208\",\"type\":\"t_array(t_uint256)43_storage\"},{\"astId\":1018,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"superchainConfig\",\"offset\":0,\"slot\":\"251\",\"type\":\"t_contract(SuperchainConfig)1023\"},{\"astId\":1019,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"portal\",\"offset\":0,\"slot\":\"252\",\"type\":\"t_contract(OptimismPortal)1022\"},{\"astId\":1020,\"contract\":\"src/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"systemConfig\",\"offset\":0,\"slot\":\"253\",\"type\":\"t_contract(SystemConfig)1024\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)43_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[43]\",\"numberOfBytes\":\"1376\",\"base\":\"t_uint256\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_contract(CrossDomainMessenger)1021\":{\"encoding\":\"inplace\",\"label\":\"contract CrossDomainMessenger\",\"numberOfBytes\":\"20\"},\"t_contract(OptimismPortal)1022\":{\"encoding\":\"inplace\",\"label\":\"contract OptimismPortal\",\"numberOfBytes\":\"20\"},\"t_contract(SuperchainConfig)1023\":{\"encoding\":\"inplace\",\"label\":\"contract SuperchainConfig\",\"numberOfBytes\":\"20\"},\"t_contract(SystemConfig)1024\":{\"encoding\":\"inplace\",\"label\":\"contract SystemConfig\",\"numberOfBytes\":\"20\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L1CrossDomainMessengerDeployedBin = "0x6080604052600436106101c25760003560e01c80635c975abb116100f7578063a4e7f8bd11610095578063d764ad0b11610064578063d764ad0b14610559578063db505d801461056c578063e0e593c514610599578063ecc70428146105b957600080fd5b8063a4e7f8bd146104b9578063b1b1b209146104e9578063b28ade2514610519578063c0c53b8b1461053957600080fd5b806383a74074116100d157806383a74074146104455780638cbeeef21461035757806392a162cf1461045c5780639fce812c1461048e57600080fd5b80635c975abb146103ee5780636425666b146104035780636e296e451461043057600080fd5b80633dbb202b116101645780634c1d6a691161013e5780634c1d6a69146103575780634d0047ee1461036d57806354fd4d50146103825780635644cfdf146103d857600080fd5b80633dbb202b146102fa5780633f827a5a1461030f5780634273ca161461033757600080fd5b80630ff754ea116101a05780630ff754ea1461023f5780632828d7e81461028b57806333d7e2bd146102a057806335e80ab3146102cd57600080fd5b806301ffc9a7146101c7578063028f85f7146101fc5780630c5684981461022a575b600080fd5b3480156101d357600080fd5b506101e76101e2366004612592565b61061e565b60405190151581526020015b60405180910390f35b34801561020857600080fd5b50610211601081565b60405167ffffffffffffffff90911681526020016101f3565b34801561023657600080fd5b50610211603f81565b34801561024b57600080fd5b5060fc5473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101f3565b34801561029757600080fd5b50610211604081565b3480156102ac57600080fd5b5060fd546102669073ffffffffffffffffffffffffffffffffffffffff1681565b3480156102d957600080fd5b5060fb546102669073ffffffffffffffffffffffffffffffffffffffff1681565b61030d61030836600461265b565b6106b7565b005b34801561031b57600080fd5b50610324600181565b60405161ffff90911681526020016101f3565b34801561034357600080fd5b506101e76103523660046126c2565b610914565b34801561036357600080fd5b50610211619c4081565b34801561037957600080fd5b50610266610ac8565b34801561038e57600080fd5b506103cb6040518060400160405280600581526020017f322e332e3000000000000000000000000000000000000000000000000000000081525081565b6040516101f391906127ab565b3480156103e457600080fd5b5061021161138881565b3480156103fa57600080fd5b506101e7610b61565b34801561040f57600080fd5b5060fc546102669073ffffffffffffffffffffffffffffffffffffffff1681565b34801561043c57600080fd5b50610266610bf5565b34801561045157600080fd5b5061021162030d4081565b34801561046857600080fd5b5061047c6104773660046127be565b610cdc565b6040516101f396959493929190612849565b34801561049a57600080fd5b5060cf5473ffffffffffffffffffffffffffffffffffffffff16610266565b3480156104c557600080fd5b506101e76104d43660046128a1565b60ce6020526000908152604090205460ff1681565b3480156104f557600080fd5b506101e76105043660046128a1565b60cb6020526000908152604090205460ff1681565b34801561052557600080fd5b506102116105343660046128ba565b610dc5565b34801561054557600080fd5b5061030d61055436600461290e565b610e35565b61030d610567366004612959565b611072565b34801561057857600080fd5b5060cf546102669073ffffffffffffffffffffffffffffffffffffffff1681565b3480156105a557600080fd5b5061030d6105b43660046129df565b611a95565b3480156105c557600080fd5b5061061060cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b6040519081526020016101f3565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f4273ca160000000000000000000000000000000000000000000000000000000014806106b157507fffffffff0000000000000000000000000000000000000000000000000000000082167f01ffc9a700000000000000000000000000000000000000000000000000000000145b92915050565b60cf546107e99073ffffffffffffffffffffffffffffffffffffffff166106df858585610dc5565b347fd764ad0b0000000000000000000000000000000000000000000000000000000061074b60cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c6040516024016107679796959493929190612a50565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611aaa565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a33858561086e60cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b86604051610880959493929190612aaf565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b600061091e610ac8565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146109dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f6f6e6c7920616363657074206e617469766520746f6b656e20617070726f766560448201527f2063616c6c6261636b000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6000806000803660006109f08989610cdc565b9550955095509550955095508573ffffffffffffffffffffffffffffffffffffffff168c73ffffffffffffffffffffffffffffffffffffffff16148015610a365750838a145b8015610a425750600084115b610aa8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f696e76616c6964206f6e417070726f766520646174610000000000000000000060448201526064016109d4565b610ab6868686868686611bab565b5060019b9a5050505050505050505050565b60fd54604080517f4d0047ee000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff1691634d0047ee9160048083019260209291908290030181865afa158015610b38573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b5c9190612afd565b905090565b60fb54604080517f5c975abb000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff1691635c975abb9160048083019260209291908290030181865afa158015610bd1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b5c9190612b1a565b60cc5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff215301610cbf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f7420736574000000000000000000000060648201526084016109d4565b5060cc5473ffffffffffffffffffffffffffffffffffffffff1690565b60008080803681604c871015610d74576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603160248201527f496e76616c6964206f6e417070726f7665206461746120666f72204c3143726f60448201527f7373446f6d61696e4d657373656e67657200000000000000000000000000000060648201526084016109d4565b5050508435606090811c96601487013590911c95602881013595604882013560e01c9550604c90910193507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffb4019150565b6000611388619c4080603f610de1604063ffffffff8816612b6b565b610deb9190612b9b565b610df6601088612b6b565b610e039062030d40612be9565b610e0d9190612be9565b610e179190612be9565b610e219190612be9565b610e2b9190612be9565b90505b9392505050565b6000546003907501000000000000000000000000000000000000000000900460ff16158015610e83575060005460ff8083167401000000000000000000000000000000000000000090920416105b610f0f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016109d4565b600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff60ff84167401000000000000000000000000000000000000000002167fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff90911617750100000000000000000000000000000000000000000017905560fb805473ffffffffffffffffffffffffffffffffffffffff8087167fffffffffffffffffffffffff00000000000000000000000000000000000000009283161790925560fc805486841690831617905560fd80549285169290911691909117905561100e734200000000000000000000000000000000000007611e77565b600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050565b3415611100576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f00000000000000000000000000000000000000000000000060648201526084016109d4565b60f087901c600281106111bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604d60248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2030206f722031206d657373616765732061726520737570706f7274656460648201527f20617420746869732074696d6500000000000000000000000000000000000000608482015260a4016109d4565b8061ffff166000036112b057600061120c878986868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508f9250611fb3915050565b600081815260cb602052604090205490915060ff16156112ae576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a206c65676163792077697460448201527f6864726177616c20616c72656164792072656c6179656400000000000000000060648201526084016109d4565b505b60006112f6898989898989898080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611fd292505050565b90506000611302610ac8565b905061130c611ff5565b1561136357600082815260ce602052604090205460ff161561133057611330612c15565b861561135e5760fc5461135e9073ffffffffffffffffffffffffffffffffffffffff8381169116308a6120d1565b611401565b600082815260ce602052604090205460ff16611401576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c617965640000000000000000000000000000000060648201526084016109d4565b61140a8861216c565b15801561144357508073ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff1614155b6114f5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605960248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f657373206f72206e6174697665546f6b656e4164647265737300000000000000608482015260a4016109d4565b600082815260cb602052604090205460ff1615611594576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c617965640000000000000000000060648201526084016109d4565b6115b5866115a6611388619c40612be9565b67ffffffffffffffff166121af565b15806115db575060cc5473ffffffffffffffffffffffffffffffffffffffff1661dead14155b156116f557600082815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff32016116ed576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d6573736167650000000000000000000000000000000000000060648201526084016109d4565b505050611a8c565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8b16179055600187156117d9576040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8a81166004830152602482018a905283169063095ea7b3906044016020604051808303816000875af11580156117b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117d69190612b1a565b90505b600061182b8a619c405a6117ed9190612c44565b60008a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506121cd92505050565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead17905590508815801590611866575080155b15611907576040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8b811660048301526000602483015284169063095ea7b3906044016020604051808303816000875af11580156118e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119049190612b1a565b91505b8080156119115750815b1561197957600084815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555185917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a2611a86565b600084815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555185917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff3201611a86576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f43726f7373446f6d61696e4d657373656e6765723a206661696c656420746f2060448201527f72656c6179206d6573736167650000000000000000000000000000000000000060648201526084016109d4565b50505050505b50505050505050565b611aa3338686848787611bab565b5050505050565b3415611b12576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f44656e79206465706f736974696e67204554480000000000000000000000000060448201526064016109d4565b60fc546040517fb9e5595800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063b9e5595890611b739087908690819089906000908990600401612c5b565b600060405180830381600087803b158015611b8d57600080fd5b505af1158015611ba1573d6000803e3d6000fd5b5050505050505050565b8315611c7f576000611bbb610ac8565b9050611bdf73ffffffffffffffffffffffffffffffffffffffff82168830886120d1565b60fc546040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152602481018790529082169063095ea7b3906044016020604051808303816000875af1158015611c58573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c7c9190612b1a565b50505b60cf54611d2f9073ffffffffffffffffffffffffffffffffffffffff16611ca7848487610dc5565b867fd764ad0b00000000000000000000000000000000000000000000000000000000611d1360cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b8b8b8b8b8b8b6040516024016107679796959493929190612a50565b8473ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a878484611db460cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b88604051611dc6959493929190612aaf565b60405180910390a28573ffffffffffffffffffffffffffffffffffffffff167f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d54685604051611e1691815260200190565b60405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff00000000000000000000000000000000000000000000000000000000000090911617905550505050565b6000547501000000000000000000000000000000000000000000900460ff16611f22576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016109d4565b60cc5473ffffffffffffffffffffffffffffffffffffffff16611f6c5760cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead1790555b60cf80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6000611fc1858585856121e7565b805190602001209050949350505050565b6000611fe2878787878787612280565b8051906020012090509695505050505050565b60fc5460009073ffffffffffffffffffffffffffffffffffffffff1633148015610b5c575060cf5460fc54604080517f9bf62d82000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff9384169390921691639bf62d82916004808201926020929091908290030181865afa158015612091573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120b59190612afd565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261216690859061231f565b50505050565b600073ffffffffffffffffffffffffffffffffffffffff82163014806106b157505060fc5473ffffffffffffffffffffffffffffffffffffffff90811691161490565b600080603f83619c4001026040850201603f5a021015949350505050565b600080600080845160208601878a8af19695505050505050565b6060848484846040516024016122009493929190612cae565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fcbd4ece9000000000000000000000000000000000000000000000000000000001790529050949350505050565b606086868686868660405160240161229d96959493929190612cf8565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b6000612381826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166124309092919063ffffffff16565b80519091501561242b578080602001905181019061239f9190612b1a565b61242b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016109d4565b505050565b6060610e2b84846000858573ffffffffffffffffffffffffffffffffffffffff85163b6124b9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016109d4565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516124e29190612d43565b60006040518083038185875af1925050503d806000811461251f576040519150601f19603f3d011682016040523d82523d6000602084013e612524565b606091505b509150915061253482828661253f565b979650505050505050565b6060831561254e575081610e2e565b82511561255e5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d491906127ab565b6000602082840312156125a457600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610e2e57600080fd5b73ffffffffffffffffffffffffffffffffffffffff811681146125f657600080fd5b50565b60008083601f84011261260b57600080fd5b50813567ffffffffffffffff81111561262357600080fd5b60208301915083602082850101111561263b57600080fd5b9250929050565b803563ffffffff8116811461265657600080fd5b919050565b6000806000806060858703121561267157600080fd5b843561267c816125d4565b9350602085013567ffffffffffffffff81111561269857600080fd5b6126a4878288016125f9565b90945092506126b7905060408601612642565b905092959194509250565b6000806000806000608086880312156126da57600080fd5b85356126e5816125d4565b945060208601356126f5816125d4565b935060408601359250606086013567ffffffffffffffff81111561271857600080fd5b612724888289016125f9565b969995985093965092949392505050565b60005b83811015612750578181015183820152602001612738565b838111156121665750506000910152565b60008151808452612779816020860160208601612735565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610e2e6020830184612761565b600080602083850312156127d157600080fd5b823567ffffffffffffffff8111156127e857600080fd5b6127f4858286016125f9565b90969095509350505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b600073ffffffffffffffffffffffffffffffffffffffff808916835280881660208401525085604083015263ffffffff8516606083015260a0608083015261289560a083018486612800565b98975050505050505050565b6000602082840312156128b357600080fd5b5035919050565b6000806000604084860312156128cf57600080fd5b833567ffffffffffffffff8111156128e657600080fd5b6128f2868287016125f9565b9094509250612905905060208501612642565b90509250925092565b60008060006060848603121561292357600080fd5b833561292e816125d4565b9250602084013561293e816125d4565b9150604084013561294e816125d4565b809150509250925092565b600080600080600080600060c0888a03121561297457600080fd5b873596506020880135612986816125d4565b95506040880135612996816125d4565b9450606088013593506080880135925060a088013567ffffffffffffffff8111156129c057600080fd5b6129cc8a828b016125f9565b989b979a50959850939692959293505050565b6000806000806000608086880312156129f757600080fd5b8535612a02816125d4565b945060208601359350604086013567ffffffffffffffff811115612a2557600080fd5b612a31888289016125f9565b9094509250612a44905060608701612642565b90509295509295909350565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a0830152612aa260c083018486612800565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff86168152608060208201526000612adf608083018688612800565b905083604083015263ffffffff831660608301529695505050505050565b600060208284031215612b0f57600080fd5b8151610e2e816125d4565b600060208284031215612b2c57600080fd5b81518015158114610e2e57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff80831681851681830481118215151615612b9257612b92612b3c565b02949350505050565b600067ffffffffffffffff80841680612bdd577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600067ffffffffffffffff808316818516808303821115612c0c57612c0c612b3c565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600082821015612c5657612c56612b3c565b500390565b73ffffffffffffffffffffffffffffffffffffffff8716815285602082015284604082015267ffffffffffffffff84166060820152821515608082015260c060a0820152600061289560c0830184612761565b600073ffffffffffffffffffffffffffffffffffffffff808716835280861660208401525060806040830152612ce76080830185612761565b905082606083015295945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a083015261289560c0830184612761565b60008251612d55818460208701612735565b919091019291505056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(L1CrossDomainMessengerStorageLayoutJSON), L1CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1CrossDomainMessenger"] = L1CrossDomainMessengerStorageLayout
	deployedBytecodes["L1CrossDomainMessenger"] = L1CrossDomainMessengerDeployedBin
	immutableReferences["L1CrossDomainMessenger"] = false
}
