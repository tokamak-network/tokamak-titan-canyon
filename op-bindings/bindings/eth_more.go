// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/tokamak-network/tokamak-thanos/op-bindings/solc"
)

const ETHStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2/ETH.sol:ETH\",\"label\":\"_balances\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_address,t_uint256)\"},{\"astId\":1001,\"contract\":\"src/L2/ETH.sol:ETH\",\"label\":\"_allowances\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1002,\"contract\":\"src/L2/ETH.sol:ETH\",\"label\":\"_totalSupply\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_uint256\"},{\"astId\":1003,\"contract\":\"src/L2/ETH.sol:ETH\",\"label\":\"_name\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_string_storage\"},{\"astId\":1004,\"contract\":\"src/L2/ETH.sol:ETH\",\"label\":\"_symbol\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_string_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var ETHStorageLayout = new(solc.StorageLayout)

var ETHDeployedBin = "0x608060405234801561001057600080fd5b50600436106101775760003560e01c806370a08231116100d8578063ae1f6aaf1161008c578063dd62ed3e11610066578063dd62ed3e14610361578063e78cea9214610315578063ee9a31a2146103a757600080fd5b8063ae1f6aaf14610315578063c01e1bd61461033b578063d6c0b2c41461033b57600080fd5b80639dc29fac116100bd5780639dc29fac146102dc578063a457c2d7146102ef578063a9059cbb1461030257600080fd5b806370a082311461029e57806395d89b41146102d457600080fd5b806323b872dd1161012f5780633950935111610114578063395093511461026e57806340c10f191461028157806354fd4d501461029657600080fd5b806323b872dd1461022a578063313ce5671461023d57600080fd5b806306fdde031161016057806306fdde03146101f0578063095ea7b31461020557806318160ddd1461021857600080fd5b806301ffc9a71461017c578063033964be146101a4575b600080fd5b61018f61018a366004611329565b6103ce565b60405190151581526020015b60405180910390f35b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f86104bf565b60405161019b919061139e565b61018f610213366004611418565b610551565b6002545b60405190815260200161019b565b61018f610238366004611442565b610569565b60405160ff7f000000000000000000000000000000000000000000000000000000000000000016815260200161019b565b61018f61027c366004611418565b61058d565b61029461028f366004611418565b6105d9565b005b6101f8610701565b61021c6102ac36600461147e565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6101f86107a4565b6102946102ea366004611418565b6107b3565b61018f6102fd366004611418565b6108ca565b61018f610310366004611418565b61099b565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b61021c61036f366004611499565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007f1d1d8b63000000000000000000000000000000000000000000000000000000007fec4fc8e3000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000851683148061048757507fffffffff00000000000000000000000000000000000000000000000000000000858116908316145b806104b657507fffffffff00000000000000000000000000000000000000000000000000000000858116908216145b95945050505050565b6060600380546104ce906114cc565b80601f01602080910402602001604051908101604052809291908181526020018280546104fa906114cc565b80156105475780601f1061051c57610100808354040283529160200191610547565b820191906000526020600020905b81548152906001019060200180831161052a57829003601f168201915b5050505050905090565b60003361055f8185856109a9565b5060019392505050565b600033610577858285610b5d565b610582858585610c34565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061055f90829086906105d490879061154e565b6109a9565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146106a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f4f7074696d69736d4d696e7461626c6545524332303a206f6e6c79206272696460448201527f67652063616e206d696e7420616e64206275726e00000000000000000000000060648201526084015b60405180910390fd5b6106ad8282610ee7565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885826040516106f591815260200190565b60405180910390a25050565b606061072c7f0000000000000000000000000000000000000000000000000000000000000000611007565b6107557f0000000000000000000000000000000000000000000000000000000000000000611007565b61077e7f0000000000000000000000000000000000000000000000000000000000000000611007565b60405160200161079093929190611566565b604051602081830303815290604052905090565b6060600480546104ce906114cc565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610878576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f4f7074696d69736d4d696e7461626c6545524332303a206f6e6c79206272696460448201527f67652063616e206d696e7420616e64206275726e000000000000000000000000606482015260840161069a565b6108828282611144565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5826040516106f591815260200190565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091908381101561098e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f000000000000000000000000000000000000000000000000000000606482015260840161069a565b61058282868684036109a9565b60003361055f818585610c34565b73ffffffffffffffffffffffffffffffffffffffff8316610a4b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f7265737300000000000000000000000000000000000000000000000000000000606482015260840161069a565b73ffffffffffffffffffffffffffffffffffffffff8216610aee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f7373000000000000000000000000000000000000000000000000000000000000606482015260840161069a565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610c2e5781811015610c21576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000604482015260640161069a565b610c2e84848484036109a9565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610cd7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f6472657373000000000000000000000000000000000000000000000000000000606482015260840161069a565b73ffffffffffffffffffffffffffffffffffffffff8216610d7a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f6573730000000000000000000000000000000000000000000000000000000000606482015260840161069a565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610e30576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e63650000000000000000000000000000000000000000000000000000606482015260840161069a565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260208190526040808220858503905591851681529081208054849290610e7490849061154e565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610eda91815260200190565b60405180910390a3610c2e565b73ffffffffffffffffffffffffffffffffffffffff8216610f64576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640161069a565b8060026000828254610f76919061154e565b909155505073ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604081208054839290610fb090849061154e565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b60608160000361104a57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115611074578061105e816115dc565b915061106d9050600a83611643565b915061104e565b60008167ffffffffffffffff81111561108f5761108f611657565b6040519080825280601f01601f1916602001820160405280156110b9576020820181803683370190505b5090505b841561113c576110ce600183611686565b91506110db600a8661169d565b6110e690603061154e565b60f81b8183815181106110fb576110fb6116b1565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611135600a86611643565b94506110bd565b949350505050565b73ffffffffffffffffffffffffffffffffffffffff82166111e7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f7300000000000000000000000000000000000000000000000000000000000000606482015260840161069a565b73ffffffffffffffffffffffffffffffffffffffff82166000908152602081905260409020548181101561129d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f6365000000000000000000000000000000000000000000000000000000000000606482015260840161069a565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604081208383039055600280548492906112d9908490611686565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610b50565b60006020828403121561133b57600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461136b57600080fd5b9392505050565b60005b8381101561138d578181015183820152602001611375565b83811115610c2e5750506000910152565b60208152600082518060208401526113bd816040850160208701611372565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461141357600080fd5b919050565b6000806040838503121561142b57600080fd5b611434836113ef565b946020939093013593505050565b60008060006060848603121561145757600080fd5b611460846113ef565b925061146e602085016113ef565b9150604084013590509250925092565b60006020828403121561149057600080fd5b61136b826113ef565b600080604083850312156114ac57600080fd5b6114b5836113ef565b91506114c3602084016113ef565b90509250929050565b600181811c908216806114e057607f821691505b602082108103611519577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156115615761156161151f565b500190565b60008451611578818460208901611372565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516115b4816001850160208a01611372565b600192019182015283516115cf816002840160208801611372565b0160020195945050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361160d5761160d61151f565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261165257611652611614565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000828210156116985761169861151f565b500390565b6000826116ac576116ac611614565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(ETHStorageLayoutJSON), ETHStorageLayout); err != nil {
		panic(err)
	}

	layouts["ETH"] = ETHStorageLayout
	deployedBytecodes["ETH"] = ETHDeployedBin
}
