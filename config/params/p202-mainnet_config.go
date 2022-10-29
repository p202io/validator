package params

import (
	//eth1Params "github.com/ethereum/go-ethereum/params"
)

// UseP202MainnetNetworkConfig uses the Project202 Mainnet specific network config.
func UseP202MainnetNetworkConfig() {
	cfg := BeaconNetworkConfig().Copy()
	cfg.ContractDeploymentBlock = 0
	cfg.BootstrapNodes = []string{}
	OverrideBeaconNetworkConfig(cfg)
}

// RopstenConfig defines the config for the Project202 Mainnet.
func P202MainnetConfig() *BeaconChainConfig {
	cfg := MainnetConfig().Copy()

	// Constants (Non-configurable)
	cfg.GenesisDelay = 120

	// Misc constant.
	cfg.MinGenesisActiveValidatorCount = 202
	cfg.MinGenesisTime = 1667028600 // P202_TODO

	// Gwei value constants.
	cfg.MaxEffectiveBalance = 202 * 1e9
	cfg.EjectionBalance = 101 * 1e9

	// Time parameter constants.
	//cfg.Eth1FollowDistance = 1024
	cfg.SecondsPerETH1Block = 3
	cfg.SecondsPerSlot = 3

	// Ethereum PoW parameters.
	cfg.DepositChainID = 202 //eth1Params.P202MainnetChainConfig.ChainID.Uint64()
	cfg.DepositNetworkID = 202 //eth1Params.P202MainnetChainConfig.ChainID.Uint64()
	cfg.DepositContractAddress = "0x4242424242424242424242424242424242424242"

	// Prysm constants.
	cfg.ConfigName = P202MainnetName
	cfg.PresetBase = P202MainnetName

	// Fork related values.
	cfg.GenesisForkVersion = []byte{0x00, 0x00, 0x02, 0x02}
	cfg.AltairForkEpoch = 1
	cfg.AltairForkVersion = []byte{0x01, 0x00, 0x02, 0x02}
	cfg.BellatrixForkEpoch = 2
	cfg.BellatrixForkVersion = []byte{0x02, 0x00, 0x02, 0x02}
	//cfg.CapellaForkEpoch = 
	cfg.CapellaForkVersion = []byte{0x03, 0x00, 0x02, 0x02}
	//cfg.ShardingForkEpoch = 
	cfg.ShardingForkVersion = []byte{0x04, 0x00, 0x02, 0x02}

	// Bellatrix
	cfg.TerminalTotalDifficulty = "12000"

	cfg.InitializeForkSchedule()
	return cfg
}
