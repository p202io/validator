package params

import (
	//eth1Params "github.com/p202io/go-ethereum/params"
)

// UseP202TestnetNetworkConfig uses the Project202 Testnet specific network config.
func UseP202TestnetNetworkConfig() {
	cfg := BeaconNetworkConfig().Copy()
	cfg.ContractDeploymentBlock = 0
	cfg.BootstrapNodes = []string{}
	OverrideBeaconNetworkConfig(cfg)
}

// P202TestnetConfig defines the config for the Project202 Testnet.
func P202TestnetConfig() *BeaconChainConfig {
	cfg := MainnetConfig().Copy()

	// Constants (Non-configurable)
	cfg.GenesisDelay = 120

	// Misc constant.
	cfg.TargetCommitteeSize = 4
	cfg.MaxValidatorsPerCommittee = 2048
	cfg.ChurnLimitQuotient = 32
	cfg.ShuffleRoundCount = 10
	cfg.MinGenesisActiveValidatorCount = 64
	cfg.MinGenesisTime = 1666764000 // P202_TODO

	// Gwei value constants.
	cfg.MaxEffectiveBalance = 202 * 1e9
	cfg.EjectionBalance = 101 * 1e9

	// Time parameter constants.
	cfg.SecondsPerETH1Block = 3
	cfg.SecondsPerSlot = 3
	cfg.SlotsPerEpoch = 8
	cfg.SqrRootSlotsPerEpoch = 2
	cfg.EpochsPerEth1VotingPeriod = 4
	cfg.SlotsPerHistoricalRoot = 64
	cfg.ShardCommitteePeriod = 64
	cfg.Eth1FollowDistance = 16
	cfg.SafeSlotsToUpdateJustified = 2

	// State list length constants.
	cfg.EpochsPerHistoricalVector = 64
	cfg.EpochsPerSlashingsVector = 64

	// Reward and penalty quotients
	cfg.InactivityPenaltyQuotient = 33554432
	cfg.MinSlashingPenaltyQuotient = 64
	cfg.ProportionalSlashingMultiplier = 2

	// Misc values.
	cfg.SyncCommitteeSize = 32
	cfg.EpochsPerSyncCommitteePeriod = 8

	// Ethereum PoW parameters.
	cfg.DepositChainID = 10202 //eth1Params.P202TestnetChainConfig.ChainID.Uint64()
	cfg.DepositNetworkID = 10202 //eth1Params.P202TestnetChainConfig.ChainID.Uint64()
	cfg.DepositContractAddress = "0x4242424242424242424242424242424242424242"

	// Prysm constants.
	cfg.ConfigName = P202TestnetName
	cfg.PresetBase = P202TestnetName

	// Fork related values.
	cfg.GenesisForkVersion = []byte{0x00, 0x01, 0x02, 0x02}
	cfg.AltairForkEpoch = 1
	cfg.AltairForkVersion = []byte{0x01, 0x01, 0x02, 0x02}
	cfg.BellatrixForkEpoch = 2
	cfg.BellatrixForkVersion = []byte{0x02, 0x01, 0x02, 0x02}
	cfg.CapellaForkEpoch = 3
	cfg.CapellaForkVersion = []byte{0x03, 0x01, 0x02, 0x02}
	cfg.ShardingForkVersion = []byte{0x04, 0x01, 0x02, 0x02}

	// Bellatrix
	cfg.TerminalTotalDifficulty = "1000"

	cfg.InitializeForkSchedule()
	return cfg
}
