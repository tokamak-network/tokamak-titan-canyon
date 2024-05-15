// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/tokamak-network/tokamak-thanos/op-bindings/solc"
)

const OptimismMintableERC721FactoryStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/universal/OptimismMintableERC721Factory.sol:OptimismMintableERC721Factory\",\"label\":\"isOptimismMintableERC721\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_address,t_bool)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_bool\"}}}"

var OptimismMintableERC721FactoryStorageLayout = new(solc.StorageLayout)

var OptimismMintableERC721FactoryDeployedBin = "0x60806040523480156200001157600080fd5b50600436106200006f5760003560e01c80637d1d0c5b11620000565780637d1d0c5b1462000100578063d97df6521462000137578063ee9a31a2146200017457600080fd5b806354fd4d5014620000745780635572acae14620000c9575b600080fd5b620000b16040518060400160405280600581526020017f312e342e3000000000000000000000000000000000000000000000000000000081525081565b604051620000c0919062000435565b60405180910390f35b620000ef620000da3660046200047b565b60006020819052908152604090205460ff1681565b6040519015158152602001620000c0565b620001287f000000000000000000000000000000000000000000000000000000000000000081565b604051908152602001620000c0565b6200014e620001483660046200057b565b6200019c565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001620000c0565b6200014e7f000000000000000000000000000000000000000000000000000000000000000081565b600073ffffffffffffffffffffffffffffffffffffffff84166200026d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526044602482018190527f4f7074696d69736d4d696e7461626c65455243373231466163746f72793a204c908201527f3120746f6b656e20616464726573732063616e6e6f742062652061646472657360648201527f7328302900000000000000000000000000000000000000000000000000000000608482015260a40160405180910390fd5b60008484846040516020016200028693929190620005f8565b6040516020818303038152906040528051906020012090506000817f00000000000000000000000000000000000000000000000000000000000000007f0000000000000000000000000000000000000000000000000000000000000000888888604051620002f490620003b9565b6200030495949392919062000647565b8190604051809103906000f590508015801562000325573d6000803e3d6000fd5b5073ffffffffffffffffffffffffffffffffffffffff8181166000818152602081815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590513381529394509189169290917fe72783bb8e0ca31286b85278da59684dd814df9762a52f0837f89edd1483b299910160405180910390a395945050505050565b6131bf80620006a983390190565b6000815180845260005b81811015620003ef57602081850181015186830182015201620003d1565b8181111562000402576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006200044a6020830184620003c7565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146200047657600080fd5b919050565b6000602082840312156200048e57600080fd5b6200044a8262000451565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112620004da57600080fd5b813567ffffffffffffffff80821115620004f857620004f862000499565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171562000541576200054162000499565b816040528381528660208588010111156200055b57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000606084860312156200059157600080fd5b6200059c8462000451565b9250602084013567ffffffffffffffff80821115620005ba57600080fd5b620005c887838801620004c8565b93506040860135915080821115620005df57600080fd5b50620005ee86828701620004c8565b9150509250925092565b73ffffffffffffffffffffffffffffffffffffffff84168152606060208201526000620006296060830185620003c7565b82810360408401526200063d8185620003c7565b9695505050505050565b600073ffffffffffffffffffffffffffffffffffffffff808816835286602084015280861660408401525060a060608301526200068860a0830185620003c7565b82810360808401526200069c8185620003c7565b9897505050505050505056fe60e06040523480156200001157600080fd5b50604051620031bf380380620031bf83398101604081905262000034916200062d565b8181600062000044838262000756565b50600162000053828262000756565b5050506001600160a01b038516620000d85760405162461bcd60e51b815260206004820152603360248201527f4f7074696d69736d4d696e7461626c654552433732313a20627269646765206360448201527f616e6e6f7420626520616464726573732830290000000000000000000000000060648201526084015b60405180910390fd5b83600003620001505760405162461bcd60e51b815260206004820152603660248201527f4f7074696d69736d4d696e7461626c654552433732313a2072656d6f7465206360448201527f6861696e2069642063616e6e6f74206265207a65726f000000000000000000006064820152608401620000cf565b6001600160a01b038316620001ce5760405162461bcd60e51b815260206004820152603960248201527f4f7074696d69736d4d696e7461626c654552433732313a2072656d6f7465207460448201527f6f6b656e2063616e6e6f742062652061646472657373283029000000000000006064820152608401620000cf565b60808490526001600160a01b0383811660a081905290861660c0526200020290601462000256602090811b62000eed17901c565b62000218856200041660201b620011301760201c565b6040516020016200022b92919062000822565b604051602081830303815290604052600a90816200024a919062000756565b50505050505062000993565b6060600062000267836002620008ac565b62000274906002620008ce565b6001600160401b038111156200028e576200028e62000553565b6040519080825280601f01601f191660200182016040528015620002b9576020820181803683370190505b509050600360fc1b81600081518110620002d757620002d7620008e9565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110620003095762000309620008e9565b60200101906001600160f81b031916908160001a90535060006200032f846002620008ac565b6200033c906001620008ce565b90505b6001811115620003be576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110620003745762000374620008e9565b1a60f81b8282815181106200038d576200038d620008e9565b60200101906001600160f81b031916908160001a90535060049490941c93620003b681620008ff565b90506200033f565b5083156200040f5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401620000cf565b9392505050565b6060816000036200043e5750506040805180820190915260018152600360fc1b602082015290565b8160005b81156200046e5780620004558162000919565b9150620004669050600a836200094b565b915062000442565b6000816001600160401b038111156200048b576200048b62000553565b6040519080825280601f01601f191660200182016040528015620004b6576020820181803683370190505b5090505b84156200052e57620004ce60018362000962565b9150620004dd600a866200097c565b620004ea906030620008ce565b60f81b818381518110620005025762000502620008e9565b60200101906001600160f81b031916908160001a90535062000526600a866200094b565b9450620004ba565b949350505050565b80516001600160a01b03811681146200054e57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b83811015620005865781810151838201526020016200056c565b8381111562000596576000848401525b50505050565b600082601f830112620005ae57600080fd5b81516001600160401b0380821115620005cb57620005cb62000553565b604051601f8301601f19908116603f01168101908282118183101715620005f657620005f662000553565b816040528381528660208588010111156200061057600080fd5b6200062384602083016020890162000569565b9695505050505050565b600080600080600060a086880312156200064657600080fd5b620006518662000536565b945060208601519350620006686040870162000536565b60608701519093506001600160401b03808211156200068657600080fd5b6200069489838a016200059c565b93506080880151915080821115620006ab57600080fd5b50620006ba888289016200059c565b9150509295509295909350565b600181811c90821680620006dc57607f821691505b602082108103620006fd57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200075157600081815260208120601f850160051c810160208610156200072c5750805b601f850160051c820191505b818110156200074d5782815560010162000738565b5050505b505050565b81516001600160401b0381111562000772576200077262000553565b6200078a81620007838454620006c7565b8462000703565b602080601f831160018114620007c25760008415620007a95750858301515b600019600386901b1c1916600185901b1785556200074d565b600085815260208120601f198616915b82811015620007f357888601518255948401946001909101908401620007d2565b5085821015620008125787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6832ba3432b932bab69d60b91b8152600083516200084881600985016020880162000569565b600160fe1b60099184019182015283516200086b81600a84016020880162000569565b712f746f6b656e5552493f75696e743235363d60701b600a9290910191820152601c01949350505050565b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615620008c957620008c962000896565b500290565b60008219821115620008e457620008e462000896565b500190565b634e487b7160e01b600052603260045260246000fd5b60008162000911576200091162000896565b506000190190565b6000600182016200092e576200092e62000896565b5060010190565b634e487b7160e01b600052601260045260246000fd5b6000826200095d576200095d62000935565b500490565b60008282101562000977576200097762000896565b500390565b6000826200098e576200098e62000935565b500690565b60805160a05160c0516127d9620009e6600039600081816103e20152818161047a01528181610b210152610c430152600081816101e001526103bc015260008181610329015261040801526127d96000f3fe608060405234801561001057600080fd5b50600436106101ae5760003560e01c80637d1d0c5b116100ee578063c87b56dd11610097578063e78cea9211610071578063e78cea92146103e0578063e951819614610406578063e985e9c51461042c578063ee9a31a21461047557600080fd5b8063c87b56dd1461039f578063d547cfb7146103b2578063d6c0b2c4146103ba57600080fd5b8063a1448194116100c8578063a144819414610366578063a22cb46514610379578063b88d4fde1461038c57600080fd5b80637d1d0c5b1461032457806395d89b411461034b5780639dc29fac1461035357600080fd5b806323b872dd1161015b5780634f6ccce7116101355780634f6ccce7146102af57806354fd4d50146102c25780636352211e146102fe57806370a082311461031157600080fd5b806323b872dd146102765780632f745c591461028957806342842e0e1461029c57600080fd5b8063081812fc1161018c578063081812fc1461023c578063095ea7b31461024f57806318160ddd1461026457600080fd5b806301ffc9a7146101b3578063033964be146101db57806306fdde0314610227575b600080fd5b6101c66101c1366004612226565b61049c565b60405190151581526020015b60405180910390f35b6102027f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101d2565b61022f6104fa565b6040516101d291906122b9565b61020261024a3660046122cc565b61058c565b61026261025d36600461230e565b6105c0565b005b6008545b6040519081526020016101d2565b610262610284366004612338565b610751565b61026861029736600461230e565b6107f2565b6102626102aa366004612338565b6108c1565b6102686102bd3660046122cc565b6108dc565b61022f6040518060400160405280600581526020017f312e332e3000000000000000000000000000000000000000000000000000000081525081565b61020261030c3660046122cc565b61099a565b61026861031f366004612374565b610a2c565b6102687f000000000000000000000000000000000000000000000000000000000000000081565b61022f610afa565b61026261036136600461230e565b610b09565b61026261037436600461230e565b610c2b565b61026261038736600461238f565b610d42565b61026261039a3660046123fa565b610d51565b61022f6103ad3660046122cc565b610df9565b61022f610e5f565b7f0000000000000000000000000000000000000000000000000000000000000000610202565b7f0000000000000000000000000000000000000000000000000000000000000000610202565b7f0000000000000000000000000000000000000000000000000000000000000000610268565b6101c661043a3660046124f4565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260056020908152604080832093909416825291909152205460ff1690565b6102027f000000000000000000000000000000000000000000000000000000000000000081565b60007f74259ebf000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000083168114806104f357506104f38361126d565b9392505050565b60606000805461050990612527565b80601f016020809104026020016040519081016040528092919081815260200182805461053590612527565b80156105825780601f1061055757610100808354040283529160200191610582565b820191906000526020600020905b81548152906001019060200180831161056557829003601f168201915b5050505050905090565b6000610597826112c3565b5060009081526004602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b60006105cb8261099a565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361068d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560448201527f720000000000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff821614806106b657506106b6813361043a565b610742576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206e6f7220617070726f76656420666f7220616c6c00006064820152608401610684565b61074c8383611351565b505050565b61075b33826113f1565b6107e7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560448201527f72206e6f7220617070726f7665640000000000000000000000000000000000006064820152608401610684565b61074c8383836114b0565b60006107fd83610a2c565b821061088b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201527f74206f6620626f756e64730000000000000000000000000000000000000000006064820152608401610684565b5073ffffffffffffffffffffffffffffffffffffffff919091166000908152600660209081526040808320938352929052205490565b61074c83838360405180602001604052806000815250610d51565b60006108e760085490565b8210610975576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60448201527f7574206f6620626f756e647300000000000000000000000000000000000000006064820152608401610684565b600882815481106109885761098861257a565b90600052602060002001549050919050565b60008181526002602052604081205473ffffffffffffffffffffffffffffffffffffffff1680610a26576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4552433732313a20696e76616c696420746f6b656e20494400000000000000006044820152606401610684565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff8216610ad1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f74206120766160448201527f6c6964206f776e657200000000000000000000000000000000000000000000006064820152608401610684565b5073ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b60606001805461050990612527565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610bce576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f4f7074696d69736d4d696e7461626c654552433732313a206f6e6c792062726960448201527f6467652063616e2063616c6c20746869732066756e6374696f6e0000000000006064820152608401610684565b610bd781611722565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca582604051610c1f91815260200190565b60405180910390a25050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610cf0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f4f7074696d69736d4d696e7461626c654552433732313a206f6e6c792062726960448201527f6467652063616e2063616c6c20746869732066756e6374696f6e0000000000006064820152608401610684565b610cfa82826117fb565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688582604051610c1f91815260200190565b610d4d338383611815565b5050565b610d5b33836113f1565b610de7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560448201527f72206e6f7220617070726f7665640000000000000000000000000000000000006064820152608401610684565b610df384848484611942565b50505050565b6060610e04826112c3565b6000610e0e6119e5565b90506000815111610e2e57604051806020016040528060008152506104f3565b80610e3884611130565b604051602001610e499291906125a9565b6040516020818303038152906040529392505050565b600a8054610e6c90612527565b80601f0160208091040260200160405190810160405280929190818152602001828054610e9890612527565b8015610ee55780601f10610eba57610100808354040283529160200191610ee5565b820191906000526020600020905b815481529060010190602001808311610ec857829003601f168201915b505050505081565b60606000610efc836002612607565b610f07906002612644565b67ffffffffffffffff811115610f1f57610f1f6123cb565b6040519080825280601f01601f191660200182016040528015610f49576020820181803683370190505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110610f8057610f8061257a565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110610fe357610fe361257a565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600061101f846002612607565b61102a906001612644565b90505b60018111156110c7577f303132333435363738396162636465660000000000000000000000000000000085600f166010811061106b5761106b61257a565b1a60f81b8282815181106110815761108161257a565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060049490941c936110c08161265c565b905061102d565b5083156104f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610684565b60608160000361117357505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b811561119d578061118781612691565b91506111969050600a836126f8565b9150611177565b60008167ffffffffffffffff8111156111b8576111b86123cb565b6040519080825280601f01601f1916602001820160405280156111e2576020820181803683370190505b5090505b8415611265576111f760018361270c565b9150611204600a86612723565b61120f906030612644565b60f81b8183815181106112245761122461257a565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535061125e600a866126f8565b94506111e6565b949350505050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f780e9d63000000000000000000000000000000000000000000000000000000001480610a265750610a26826119f4565b60008181526002602052604090205473ffffffffffffffffffffffffffffffffffffffff1661134e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4552433732313a20696e76616c696420746f6b656e20494400000000000000006044820152606401610684565b50565b600081815260046020526040902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff841690811790915581906113ab8261099a565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000806113fd8361099a565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16148061146b575073ffffffffffffffffffffffffffffffffffffffff80821660009081526005602090815260408083209388168352929052205460ff165b8061126557508373ffffffffffffffffffffffffffffffffffffffff166114918461058c565b73ffffffffffffffffffffffffffffffffffffffff1614949350505050565b8273ffffffffffffffffffffffffffffffffffffffff166114d08261099a565b73ffffffffffffffffffffffffffffffffffffffff1614611573576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201527f6f776e65720000000000000000000000000000000000000000000000000000006064820152608401610684565b73ffffffffffffffffffffffffffffffffffffffff8216611615576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610684565b611620838383611ad7565b61162b600082611351565b73ffffffffffffffffffffffffffffffffffffffff8316600090815260036020526040812080546001929061166190849061270c565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600090815260036020526040812080546001929061169c908490612644565b909155505060008181526002602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff86811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b600061172d8261099a565b905061173b81600084611ad7565b611746600083611351565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260036020526040812080546001929061177c90849061270c565b909155505060008281526002602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555183919073ffffffffffffffffffffffffffffffffffffffff8416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b610d4d828260405180602001604052806000815250611bdd565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036118aa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610684565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526005602090815260408083209487168084529482529182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b61194d8484846114b0565b61195984848484611c80565b610df3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e74657200000000000000000000000000006064820152608401610684565b6060600a805461050990612527565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f80ac58cd000000000000000000000000000000000000000000000000000000001480611a8757507fffffffff0000000000000000000000000000000000000000000000000000000082167f5b5e139f00000000000000000000000000000000000000000000000000000000145b80610a2657507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610a26565b73ffffffffffffffffffffffffffffffffffffffff8316611b3f57611b3a81600880546000838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b611b7c565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614611b7c57611b7c8382611e73565b73ffffffffffffffffffffffffffffffffffffffff8216611ba05761074c81611f2a565b8273ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161461074c5761074c8282611fd9565b611be7838361202a565b611bf46000848484611c80565b61074c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e74657200000000000000000000000000006064820152608401610684565b600073ffffffffffffffffffffffffffffffffffffffff84163b15611e68576040517f150b7a0200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85169063150b7a0290611cf7903390899088908890600401612737565b6020604051808303816000875af1925050508015611d50575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252611d4d91810190612780565b60015b611e1d573d808015611d7e576040519150601f19603f3d011682016040523d82523d6000602084013e611d83565b606091505b508051600003611e15576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e74657200000000000000000000000000006064820152608401610684565b805181602001fd5b7fffffffff00000000000000000000000000000000000000000000000000000000167f150b7a0200000000000000000000000000000000000000000000000000000000149050611265565b506001949350505050565b60006001611e8084610a2c565b611e8a919061270c565b600083815260076020526040902054909150808214611eea5773ffffffffffffffffffffffffffffffffffffffff841660009081526006602090815260408083208584528252808320548484528184208190558352600790915290208190555b50600091825260076020908152604080842084905573ffffffffffffffffffffffffffffffffffffffff9094168352600681528383209183525290812055565b600854600090611f3c9060019061270c565b60008381526009602052604081205460088054939450909284908110611f6457611f6461257a565b906000526020600020015490508060088381548110611f8557611f8561257a565b6000918252602080832090910192909255828152600990915260408082208490558582528120556008805480611fbd57611fbd61279d565b6001900381819060005260206000200160009055905550505050565b6000611fe483610a2c565b73ffffffffffffffffffffffffffffffffffffffff9093166000908152600660209081526040808320868452825280832085905593825260079052919091209190915550565b73ffffffffffffffffffffffffffffffffffffffff82166120a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610684565b60008181526002602052604090205473ffffffffffffffffffffffffffffffffffffffff1615612133576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610684565b61213f60008383611ad7565b73ffffffffffffffffffffffffffffffffffffffff82166000908152600360205260408120805460019290612175908490612644565b909155505060008181526002602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b7fffffffff000000000000000000000000000000000000000000000000000000008116811461134e57600080fd5b60006020828403121561223857600080fd5b81356104f3816121f8565b60005b8381101561225e578181015183820152602001612246565b83811115610df35750506000910152565b60008151808452612287816020860160208601612243565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006104f3602083018461226f565b6000602082840312156122de57600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461230957600080fd5b919050565b6000806040838503121561232157600080fd5b61232a836122e5565b946020939093013593505050565b60008060006060848603121561234d57600080fd5b612356846122e5565b9250612364602085016122e5565b9150604084013590509250925092565b60006020828403121561238657600080fd5b6104f3826122e5565b600080604083850312156123a257600080fd5b6123ab836122e5565b9150602083013580151581146123c057600080fd5b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806000806080858703121561241057600080fd5b612419856122e5565b9350612427602086016122e5565b925060408501359150606085013567ffffffffffffffff8082111561244b57600080fd5b818701915087601f83011261245f57600080fd5b813581811115612471576124716123cb565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156124b7576124b76123cb565b816040528281528a60208487010111156124d057600080fd5b82602086016020830137600060208483010152809550505050505092959194509250565b6000806040838503121561250757600080fd5b612510836122e5565b915061251e602084016122e5565b90509250929050565b600181811c9082168061253b57607f821691505b602082108103612574577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600083516125bb818460208801612243565b8351908301906125cf818360208801612243565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561263f5761263f6125d8565b500290565b60008219821115612657576126576125d8565b500190565b60008161266b5761266b6125d8565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036126c2576126c26125d8565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082612707576127076126c9565b500490565b60008282101561271e5761271e6125d8565b500390565b600082612732576127326126c9565b500690565b600073ffffffffffffffffffffffffffffffffffffffff808716835280861660208401525083604083015260806060830152612776608083018461226f565b9695505050505050565b60006020828403121561279257600080fd5b81516104f3816121f8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c634300080f000aa164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(OptimismMintableERC721FactoryStorageLayoutJSON), OptimismMintableERC721FactoryStorageLayout); err != nil {
		panic(err)
	}

	layouts["OptimismMintableERC721Factory"] = OptimismMintableERC721FactoryStorageLayout
	deployedBytecodes["OptimismMintableERC721Factory"] = OptimismMintableERC721FactoryDeployedBin
}
