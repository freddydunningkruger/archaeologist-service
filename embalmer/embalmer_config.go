package embalmer

type EmbalmerConfig struct {
	EMBALMER_PRIVATE_KEY string
	ARCH_PRIVATE_KEY string
	RECIPIENT_PRIVATE_KEY string
	ETH_NODE string
	CONTRACT_ADDRESS string
	TOKEN_ADDRESS string
	RESURRECTION_TIME int64
	STORAGE_FEE string
	DIGGING_FEE string
	BOUNTY string
}