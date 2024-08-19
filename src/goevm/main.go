package main

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	goevm "github.com/payalkanyan/goevm/pkg"
)

func main() {
	block := goevm.NewBlock(common.HexToAddress("0x"), 2, 1, 0, 1, time.Now())
	evm := goevm.NewEVM(common.HexToAddress("0x"), 500_000, 2e5, 1, 8000000, []byte{}, []byte{}, block)

	// Switch the bytecodes around to use them.
	ADDCODE := []byte{0x60, 0x42, 0x60, 0xFF, 0x01}
	PUSHCODE := []byte{0x60, 0x42}
	evm.Code = PUSHCODE
	evm.Code = ADDCODE

	REVERTCODE := []byte{0x60, 0x1f, 0x60, 0x01, 0x01, 0x60, 0x00, 0x60, 0x00, 0xFD, 0x60, 0x20}
	evm.Code = REVERTCODE

	STORECODE := []byte{0x60, 0x20, 0x5f, 0x55}
	STORE2CODE := []byte{0x60, 0x20, 0x5f, 0x55, 0x60, 0xa, 0x5f, 0x55}
	STORE3CODE := []byte{0x60, 0x20, 0x5f, 0x55, 0x60, 0xa, 0x60, 0x1f, 0x55, 0x60, 0xa, 0x60, 0x2f, 0x55}
	STORE4CODE := []byte{0x60, 0x20, 0x5f, 0x55, 0x5f, 0x5f, 0x55, 0x60, 0xa, 0x60, 0x1f, 0x55, 0x60, 0xa, 0x60, 0x2f, 0x55} // Has gas refund, for zeroing slot 0
	evm.Code = STORECODE
	evm.Code = STORE2CODE
	evm.Code = STORE3CODE
	evm.Code = STORE4CODE

	MULCODE := []byte{0x60, 0x02, 0x60, 0x02, 0x02, 0x00}
	evm.Code = MULCODE

	LOG1CODE := []byte{0x60, 0x0a, 0x5f, 0x52, 0x60, 0x14, 0x60, 0x20, 0x52, 0x60, 0x1f, 0x60, 0x40, 0x5f, 0xa1, 0x60, 0x0a, 0x5f, 0x52, 0x60, 0x96, 0x60, 0x20, 0x52, 0x60, 0x2a, 0x60, 0x40, 0x5f, 0xa1}
	evm.Code = LOG1CODE

	// evm.Code = common.Hex2Bytes("602a601060011760005260206000f3600151600160005260206000f3")

	// real bytcode from a contract
	// uncomment to use it
	// evm.Code = common.Hex2Bytes("608060405234801561000f575f80fd5b506040518060400160405280600581526020017f4a657373650000000000000000000000000000000000000000000000000000008152505f9081610053919061029c565b506107e860018190555061036b565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806100dd57607f821691505b6020821081036100f0576100ef610099565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026101527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610117565b61015c8683610117565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6101a061019b61019684610174565b61017d565b610174565b9050919050565b5f819050919050565b6101b983610186565b6101cd6101c5826101a7565b848454610123565b825550505050565b5f90565b6101e16101d5565b6101ec8184846101b0565b505050565b5b8181101561020f576102045f826101d9565b6001810190506101f2565b5050565b601f82111561025457610225816100f6565b61022e84610108565b8101602085101561023d578190505b61025161024985610108565b8301826101f1565b50505b505050565b5f82821c905092915050565b5f6102745f1984600802610259565b1980831691505092915050565b5f61028c8383610265565b9150826002028217905092915050565b6102a582610062565b67ffffffffffffffff8111156102be576102bd61006c565b5b6102c882546100c6565b6102d3828285610213565b5f60209050601f831160018114610304575f84156102f2578287015190505b6102fc8582610281565b865550610363565b601f198416610312866100f6565b5f5b8281101561033957848901518255600182019150602085019450602081019050610314565b868310156103565784890151610352601f891682610265565b8355505b6001600288020188555050505b505050505050565b61070f806103785f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c806317d7de7c1461004e57806384da92a71461006c578063ef88a09214610088578063fd08921b146100a4575b5f80fd5b6100566100c2565b60405161006391906101e6565b60405180910390f35b61008660048036038101906100819190610343565b610151565b005b6100a2600480360381019061009d91906103bd565b610163565b005b6100ac61016d565b6040516100b991906103f7565b60405180910390f35b60605f80546100d09061043d565b80601f01602080910402602001604051908101604052809291908181526020018280546100fc9061043d565b80156101475780601f1061011e57610100808354040283529160200191610147565b820191905f5260205f20905b81548152906001019060200180831161012a57829003601f168201915b5050505050905090565b805f908161015f919061060a565b5050565b8060018190555050565b5f600154905090565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6101b882610176565b6101c28185610180565b93506101d2818560208601610190565b6101db8161019e565b840191505092915050565b5f6020820190508181035f8301526101fe81846101ae565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6102558261019e565b810181811067ffffffffffffffff821117156102745761027361021f565b5b80604052505050565b5f610286610206565b9050610292828261024c565b919050565b5f67ffffffffffffffff8211156102b1576102b061021f565b5b6102ba8261019e565b9050602081019050919050565b828183375f83830152505050565b5f6102e76102e284610297565b61027d565b9050828152602081018484840111156103035761030261021b565b5b61030e8482856102c7565b509392505050565b5f82601f83011261032a57610329610217565b5b813561033a8482602086016102d5565b91505092915050565b5f602082840312156103585761035761020f565b5b5f82013567ffffffffffffffff81111561037557610374610213565b5b61038184828501610316565b91505092915050565b5f819050919050565b61039c8161038a565b81146103a6575f80fd5b50565b5f813590506103b781610393565b92915050565b5f602082840312156103d2576103d161020f565b5b5f6103df848285016103a9565b91505092915050565b6103f18161038a565b82525050565b5f60208201905061040a5f8301846103e8565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061045457607f821691505b60208210810361046757610466610410565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026104c97fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261048e565b6104d3868361048e565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61050e6105096105048461038a565b6104eb565b61038a565b9050919050565b5f819050919050565b610527836104f4565b61053b61053382610515565b84845461049a565b825550505050565b5f90565b61054f610543565b61055a81848461051e565b505050565b5b8181101561057d576105725f82610547565b600181019050610560565b5050565b601f8211156105c2576105938161046d565b61059c8461047f565b810160208510156105ab578190505b6105bf6105b78561047f565b83018261055f565b50505b505050565b5f82821c905092915050565b5f6105e25f19846008026105c7565b1980831691505092915050565b5f6105fa83836105d3565b9150826002028217905092915050565b61061382610176565b67ffffffffffffffff81111561062c5761062b61021f565b5b610636825461043d565b610641828285610581565b5f60209050601f831160018114610672575f8415610660578287015190505b61066a85826105ef565b8655506106d1565b601f1984166106808661046d565b5f5b828110156106a757848901518255600182019150602085019450602081019050610682565b868310156106c457848901516106c0601f8916826105d3565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220171a252a9a14f80264325dafaebbdce726557b26ef8b20cf9ab8596851e5c44164736f6c63430008190033")

	// You can set calldata
	// evm.Calldata = common.Hex2Bytes("17d7de7c")
	evm.Run()
}
