// Generated by custom bridge script. DO NOT EDIT.

package bindings

var FunctionMapping = map[[4]byte]string{
	{0, 0, 0, 0}:         "transfer()",
	{96, 128, 96, 64}:    "deploy()",
	{53, 214, 144, 89}:   "AToken.allowMigration()",
	{221, 98, 237, 62}:   "AToken.allowance(address,address)",
	{9, 94, 167, 179}:    "AToken.approve(address,uint256)",
	{112, 160, 130, 49}:  "AToken.balanceOf(address)",
	{49, 60, 229, 103}:   "AToken.decimals()",
	{164, 87, 194, 215}:  "AToken.decreaseAllowance(address,uint256)",
	{111, 110, 190, 200}: "AToken.externalBurn(address,uint256)",
	{153, 249, 136, 152}: "AToken.externalMint(address,uint256)",
	{3, 92, 112, 153}:    "AToken.getLegacyTokenAddress()",
	{134, 83, 164, 101}:  "AToken.getMetamorphicContractAddress(bytes32,address)",
	{57, 80, 147, 81}:    "AToken.increaseAllowance(address,uint256)",
	{129, 41, 252, 28}:   "AToken.initialize()",
	{69, 75, 6, 8}:       "AToken.migrate(uint256)",
	{6, 253, 222, 3}:     "AToken.name()",
	{149, 216, 155, 65}:  "AToken.symbol()",
	{24, 22, 13, 221}:    "AToken.totalSupply()",
	{169, 5, 156, 187}:   "AToken.transfer(address,uint256)",
	{35, 184, 114, 221}:  "AToken.transferFrom(address,address,uint256)",
	{157, 194, 159, 172}: "ATokenBurner.burn(address,uint256)",
	{134, 83, 164, 101}:  "ATokenBurner.getMetamorphicContractAddress(bytes32,address)",
	{134, 83, 164, 101}:  "ATokenMinter.getMetamorphicContractAddress(bytes32,address)",
	{64, 193, 15, 25}:    "ATokenMinter.mint(address,uint256)",
	{18, 230, 191, 106}:  "AliceNetFactory.callAny(address,uint256,bytes)",
	{108, 15, 121, 182}:  "AliceNetFactory.contracts()",
	{71, 19, 238, 122}:   "AliceNetFactory.delegateCallAny(address,bytes)",
	{206, 155, 121, 48}:  "AliceNetFactory.delegator()",
	{39, 254, 24, 34}:    "AliceNetFactory.deployCreate(bytes)",
	{86, 242, 167, 97}:   "AliceNetFactory.deployCreate2(uint256,bytes32,bytes)",
	{57, 202, 180, 114}:  "AliceNetFactory.deployProxy(bytes32)",
	{250, 72, 29, 165}:   "AliceNetFactory.deployStatic(bytes32,bytes)",
	{23, 207, 242, 197}:  "AliceNetFactory.deployTemplate(bytes)",
	{170, 241, 15, 66}:   "AliceNetFactory.getImplementation()",
	{134, 83, 164, 101}:  "AliceNetFactory.getMetamorphicContractAddress(bytes32,address)",
	{207, 225, 11, 48}:   "AliceNetFactory.getNumContracts()",
	{225, 215, 168, 228}: "AliceNetFactory.initializeContract(address,bytes)",
	{243, 158, 193, 247}: "AliceNetFactory.lookup(bytes32)",
	{52, 138, 12, 220}:   "AliceNetFactory.multiCall(bytes[])",
	{141, 165, 203, 91}:  "AliceNetFactory.owner()",
	{131, 205, 156, 195}: "AliceNetFactory.setDelegator(address)",
	{215, 132, 212, 38}:  "AliceNetFactory.setImplementation(address)",
	{19, 175, 64, 53}:    "AliceNetFactory.setOwner(address)",
	{4, 60, 148, 20}:     "AliceNetFactory.upgradeProxy(bytes32,address,bytes)",
	{221, 98, 237, 62}:   "BToken.allowance(address,address)",
	{9, 94, 167, 179}:    "BToken.approve(address,uint256)",
	{199, 99, 96, 113}:   "BToken.bTokensToEth(uint256,uint256,uint256)",
	{112, 160, 130, 49}:  "BToken.balanceOf(address)",
	{179, 144, 192, 171}: "BToken.burn(uint256,uint256)",
	{155, 5, 114, 3}:     "BToken.burnTo(address,uint256,uint256)",
	{49, 60, 229, 103}:   "BToken.decimals()",
	{164, 87, 194, 215}:  "BToken.decreaseAllowance(address,uint256)",
	{0, 131, 129, 114}:   "BToken.deposit(uint8,address,uint256)",
	{228, 252, 107, 109}: "BToken.distribute()",
	{117, 168, 106, 126}: "BToken.ethToBTokens(uint256,uint256)",
	{110, 153, 96, 195}:  "BToken.getAdmin()",
	{159, 159, 185, 104}: "BToken.getDeposit(uint256)",
	{134, 83, 164, 101}:  "BToken.getMetamorphicContractAddress(bytes32,address)",
	{171, 215, 10, 162}:  "BToken.getPoolBalance()",
	{94, 206, 243, 175}:  "BToken.getTotalBTokensDeposited()",
	{57, 80, 147, 81}:    "BToken.increaseAllowance(address,uint256)",
	{129, 41, 252, 28}:   "BToken.initialize()",
	{160, 113, 45, 104}:  "BToken.mint(uint256)",
	{79, 35, 38, 40}:     "BToken.mintDeposit(uint8,address,uint256)",
	{68, 154, 82, 248}:   "BToken.mintTo(address,uint256)",
	{6, 253, 222, 3}:     "BToken.name()",
	{112, 75, 108, 2}:    "BToken.setAdmin(address)",
	{118, 123, 193, 191}: "BToken.setSplits(uint256,uint256,uint256,uint256)",
	{149, 216, 155, 65}:  "BToken.symbol()",
	{24, 22, 13, 221}:    "BToken.totalSupply()",
	{169, 5, 156, 187}:   "BToken.transfer(address,uint256)",
	{35, 184, 114, 221}:  "BToken.transferFrom(address,address,uint256)",
	{146, 23, 130, 120}:  "BToken.virtualMintDeposit(uint8,address,uint256)",
	{30, 19, 191, 187}:   "BTokenErrorCodes.BTOKEN_BURN_AMOUNT_EXCEEDS_SUPPLY()",
	{14, 25, 208, 36}:    "BTokenErrorCodes.BTOKEN_CONTRACTS_DISALLOWED_DEPOSITS()",
	{185, 73, 212, 133}:  "BTokenErrorCodes.BTOKEN_DEPOSIT_AMOUNT_ZERO()",
	{254, 74, 150, 154}:  "BTokenErrorCodes.BTOKEN_DEPOSIT_BURN_FAIL()",
	{163, 214, 0, 241}:   "BTokenErrorCodes.BTOKEN_INVALID_BALANCE()",
	{248, 126, 17, 78}:   "BTokenErrorCodes.BTOKEN_INVALID_BURN_AMOUNT()",
	{61, 212, 40, 22}:    "BTokenErrorCodes.BTOKEN_INVALID_DEPOSIT_ID()",
	{23, 190, 97, 50}:    "BTokenErrorCodes.BTOKEN_MARKET_SPREAD_TOO_LOW()",
	{66, 231, 69, 228}:   "BTokenErrorCodes.BTOKEN_MINIMUM_BURN_NOT_MET()",
	{252, 69, 244, 207}:  "BTokenErrorCodes.BTOKEN_MINIMUM_MINT_NOT_MET()",
	{40, 16, 209, 66}:    "BTokenErrorCodes.BTOKEN_MINT_INSUFFICIENT_ETH()",
	{199, 245, 60, 79}:   "BTokenErrorCodes.BTOKEN_SPLIT_VALUE_SUM_ERROR()",
	{218, 230, 129, 188}: "ETHDKG.accuseParticipantDidNotDistributeShares(address[])",
	{125, 242, 78, 233}:  "ETHDKG.accuseParticipantDidNotSubmitGPKJ(address[])",
	{4, 58, 111, 18}:     "ETHDKG.accuseParticipantDidNotSubmitKeyShares(address[])",
	{237, 190, 123, 247}: "ETHDKG.accuseParticipantDistributedBadShares(address,uint256[],uint256[2][],uint256[2],uint256[2])",
	{247, 44, 69, 182}:   "ETHDKG.accuseParticipantNotRegistered(address[])",
	{128, 0, 18, 100}:    "ETHDKG.accuseParticipantSubmittedBadGPKJ(address[],bytes32[],uint256[2][][],address)",
	{82, 46, 17, 119}:    "ETHDKG.complete()",
	{128, 185, 126, 1}:   "ETHDKG.distributeShares(uint256[],uint256[2][])",
	{50, 212, 213, 112}:  "ETHDKG.getBadParticipants()",
	{140, 132, 141, 50}:  "ETHDKG.getConfirmationLength()",
	{41, 88, 232, 28}:    "ETHDKG.getETHDKGPhase()",
	{225, 70, 55, 42}:    "ETHDKG.getMasterPublicKey()",
	{28, 103, 213, 118}:  "ETHDKG.getMasterPublicKeyHash()",
	{134, 83, 164, 101}:  "ETHDKG.getMetamorphicContractAddress(bytes32,address)",
	{236, 186, 219, 54}:  "ETHDKG.getMinValidators()",
	{208, 135, 210, 136}: "ETHDKG.getNonce()",
	{253, 71, 140, 169}:  "ETHDKG.getNumParticipants()",
	{191, 119, 134, 182}: "ETHDKG.getParticipantInternalState(address)",
	{192, 22, 186, 238}:  "ETHDKG.getParticipantsInternalState(address[])",
	{16, 109, 165, 125}:  "ETHDKG.getPhaseLength()",
	{162, 188, 156, 120}: "ETHDKG.getPhaseStartBlock()",
	{228, 163, 1, 22}:    "ETHDKG.initialize(uint256,uint256)",
	{87, 181, 28, 156}:   "ETHDKG.initializeETHDKG()",
	{43, 124, 103, 36}:   "ETHDKG.isETHDKGCompleted()",
	{67, 206, 213, 52}:   "ETHDKG.isETHDKGHalted()",
	{116, 123, 33, 124}:  "ETHDKG.isETHDKGRunning()",
	{8, 239, 207, 22}:    "ETHDKG.isMasterPublicKeySet()",
	{72, 144, 70, 90}:    "ETHDKG.migrateValidators(address[],uint256[],uint256[4][],uint8,uint256,uint256,uint256,uint256[4])",
	{52, 66, 175, 92}:    "ETHDKG.register(uint256[2])",
	{255, 62, 94, 69}:    "ETHDKG.setConfirmationLength(uint16)",
	{223, 141, 21, 123}:  "ETHDKG.setCustomAliceNetHeight(uint256)",
	{138, 60, 36, 204}:   "ETHDKG.setPhaseLength(uint16)",
	{16, 31, 73, 193}:    "ETHDKG.submitGPKJ(uint256[4])",
	{98, 166, 82, 62}:    "ETHDKG.submitKeyShare(uint256[2],uint256[2],uint256[4])",
	{232, 50, 50, 36}:    "ETHDKG.submitMasterPublicKey(uint256[4])",
	{101, 230, 43, 155}:  "ETHDKG.tryGetParticipantIndex(address)",
	{250, 15, 51, 185}:   "ETHDKGErrorCodes.ETHDKG_ACCUSED_DID_NOT_DISTRIBUTE_SHARES_IN_ROUND()",
	{3, 72, 245, 204}:    "ETHDKGErrorCodes.ETHDKG_ACCUSED_DID_NOT_PARTICIPATE_IN_GPKJ_SUBMISSION()",
	{172, 70, 131, 190}:  "ETHDKGErrorCodes.ETHDKG_ACCUSED_DID_NOT_SUBMIT_GPKJ_IN_ROUND()",
	{175, 28, 143, 88}:   "ETHDKGErrorCodes.ETHDKG_ACCUSED_DISTRIBUTED_GPKJ()",
	{40, 56, 237, 174}:   "ETHDKGErrorCodes.ETHDKG_ACCUSED_DISTRIBUTED_SHARES_IN_ROUND()",
	{76, 210, 145, 191}:  "ETHDKGErrorCodes.ETHDKG_ACCUSED_HAS_COMMITMENTS()",
	{225, 24, 121, 204}:  "ETHDKGErrorCodes.ETHDKG_ACCUSED_NOT_PARTICIPATING_IN_ROUND()",
	{245, 244, 110, 115}: "ETHDKGErrorCodes.ETHDKG_ACCUSED_NOT_VALIDATOR()",
	{85, 184, 60, 86}:    "ETHDKGErrorCodes.ETHDKG_ACCUSED_PARTICIPATING_IN_ROUND()",
	{178, 59, 131, 88}:   "ETHDKGErrorCodes.ETHDKG_ACCUSED_SUBMITTED_SHARES_IN_ROUND()",
	{229, 143, 4, 237}:   "ETHDKGErrorCodes.ETHDKG_COMMITMENT_NOT_ON_CURVE()",
	{129, 104, 127, 128}: "ETHDKGErrorCodes.ETHDKG_COMMITMENT_ZERO()",
	{214, 89, 21, 226}:   "ETHDKGErrorCodes.ETHDKG_DISPUTER_DID_NOT_DISTRIBUTE_SHARES_IN_ROUND()",
	{59, 43, 130, 69}:    "ETHDKGErrorCodes.ETHDKG_DISPUTER_DID_NOT_SUBMIT_GPKJ_IN_ROUND()",
	{118, 61, 249, 61}:   "ETHDKGErrorCodes.ETHDKG_DISPUTER_NOT_PARTICIPATING_IN_ROUND()",
	{245, 73, 128, 199}:  "ETHDKGErrorCodes.ETHDKG_DISTRIBUTED_SHARE_HASH_ZERO()",
	{196, 233, 203, 227}: "ETHDKGErrorCodes.ETHDKG_GPKJ_ZERO()",
	{77, 118, 41, 29}:    "ETHDKGErrorCodes.ETHDKG_INVALID_ARGS()",
	{248, 253, 121, 68}:  "ETHDKGErrorCodes.ETHDKG_INVALID_COMMITMENTS()",
	{221, 53, 215, 218}:  "ETHDKGErrorCodes.ETHDKG_INVALID_KEYSHARE_G1()",
	{132, 104, 192, 146}: "ETHDKGErrorCodes.ETHDKG_INVALID_KEYSHARE_G2()",
	{168, 82, 113, 63}:   "ETHDKGErrorCodes.ETHDKG_INVALID_KEY_OR_PROOF()",
	{51, 65, 188, 223}:   "ETHDKGErrorCodes.ETHDKG_INVALID_NONCE()",
	{11, 202, 114, 100}:  "ETHDKGErrorCodes.ETHDKG_INVALID_NUM_COMMITMENTS()",
	{232, 220, 214, 122}: "ETHDKGErrorCodes.ETHDKG_INVALID_NUM_ENCRYPTED_SHARES()",
	{47, 94, 152, 222}:   "ETHDKGErrorCodes.ETHDKG_INVALID_OR_DUPLICATED_PARTICIPANT()",
	{109, 30, 72, 24}:    "ETHDKGErrorCodes.ETHDKG_INVALID_SHARES_OR_COMMITMENTS()",
	{193, 105, 251, 196}: "ETHDKGErrorCodes.ETHDKG_KEYSHARE_PHASE_INVALID_NONCE()",
	{46, 248, 205, 110}:  "ETHDKGErrorCodes.ETHDKG_MASTER_PUBLIC_KEY_PAIRING_CHECK_FAILURE()",
	{100, 162, 6, 56}:    "ETHDKGErrorCodes.ETHDKG_MIGRATION_INPUT_DATA_MISMATCH()",
	{6, 166, 33, 223}:    "ETHDKGErrorCodes.ETHDKG_MIGRATION_INVALID_NONCE()",
	{126, 159, 57, 131}:  "ETHDKGErrorCodes.ETHDKG_MIN_VALIDATORS_NOT_MET()",
	{76, 76, 253, 117}:   "ETHDKGErrorCodes.ETHDKG_NOT_IN_DISPUTE_PHASE()",
	{83, 24, 127, 106}:   "ETHDKGErrorCodes.ETHDKG_NOT_IN_GPKJ_SUBMISSION_PHASE()",
	{135, 137, 202, 34}:  "ETHDKGErrorCodes.ETHDKG_NOT_IN_KEYSHARE_SUBMISSION_PHASE()",
	{48, 114, 227, 99}:   "ETHDKGErrorCodes.ETHDKG_NOT_IN_MASTER_PUBLIC_KEY_SUBMISSION_PHASE()",
	{196, 66, 55, 129}:   "ETHDKGErrorCodes.ETHDKG_NOT_IN_POST_GPKJ_DISPUTE_PHASE()",
	{109, 66, 158, 242}:  "ETHDKGErrorCodes.ETHDKG_NOT_IN_POST_GPKJ_SUBMISSION_PHASE()",
	{96, 152, 118, 70}:   "ETHDKGErrorCodes.ETHDKG_NOT_IN_POST_KEYSHARE_SUBMISSION_PHASE()",
	{115, 133, 219, 93}:  "ETHDKGErrorCodes.ETHDKG_NOT_IN_POST_REGISTRATION_PHASE()",
	{142, 37, 209, 225}:  "ETHDKGErrorCodes.ETHDKG_NOT_IN_POST_SHARED_DISTRIBUTION_PHASE()",
	{244, 237, 218, 86}:  "ETHDKGErrorCodes.ETHDKG_NOT_IN_REGISTRATION_PHASE()",
	{186, 80, 173, 32}:   "ETHDKGErrorCodes.ETHDKG_NOT_IN_SHARED_DISTRIBUTION_PHASE()",
	{131, 192, 105, 228}: "ETHDKGErrorCodes.ETHDKG_ONLY_VALIDATORS_ALLOWED()",
	{189, 241, 208, 168}: "ETHDKGErrorCodes.ETHDKG_PARTICIPANT_DISTRIBUTED_SHARES_IN_ROUND()",
	{91, 190, 177, 27}:   "ETHDKGErrorCodes.ETHDKG_PARTICIPANT_PARTICIPATING_IN_ROUND()",
	{196, 63, 19, 58}:    "ETHDKGErrorCodes.ETHDKG_PARTICIPANT_SUBMITTED_GPKJ_IN_ROUND()",
	{23, 35, 176, 132}:   "ETHDKGErrorCodes.ETHDKG_PARTICIPANT_SUBMITTED_KEYSHARES_IN_ROUND()",
	{41, 137, 111, 212}:  "ETHDKGErrorCodes.ETHDKG_PUBLIC_KEY_NOT_ON_CURVE()",
	{84, 173, 128, 142}:  "ETHDKGErrorCodes.ETHDKG_PUBLIC_KEY_ZERO()",
	{35, 39, 127, 135}:   "ETHDKGErrorCodes.ETHDKG_REQUISITES_INCOMPLETE()",
	{224, 193, 175, 216}: "ETHDKGErrorCodes.ETHDKG_SHARES_AND_COMMITMENTS_MISMATCH()",
	{121, 236, 130, 150}: "ETHDKGErrorCodes.ETHDKG_VARIABLE_CANNOT_BE_SET_WHILE_RUNNING()",
	{70, 81, 36, 134}:    "Governance.updateValue(uint256,uint256,bytes32)",
	{76, 233, 22, 77}:    "GovernanceErrorCodes.GOVERNANCE_ONLY_FACTORY_ALLOWED()",
	{223, 217, 76, 249}:  "MultipleProposalAccusation.accuseMultipleProposal(bytes,bytes,bytes,bytes)",
	{134, 83, 164, 101}:  "MultipleProposalAccusation.getMetamorphicContractAddress(bytes32,address)",
	{94, 119, 57, 103}:   "MultipleProposalAccusation.isAccused(bytes32)",
	{9, 94, 167, 179}:    "PublicStaking.approve(address,uint256)",
	{112, 160, 130, 49}:  "PublicStaking.balanceOf(address)",
	{66, 150, 108, 104}:  "PublicStaking.burn(uint256)",
	{234, 120, 90, 94}:   "PublicStaking.burnTo(address,uint256)",
	{137, 70, 92, 98}:    "PublicStaking.circuitBreakerState()",
	{42, 13, 139, 209}:   "PublicStaking.collectEth(uint256)",
	{190, 68, 67, 121}:   "PublicStaking.collectEthTo(address,uint256)",
	{227, 92, 62, 40}:    "PublicStaking.collectToken(uint256)",
	{136, 83, 185, 80}:   "PublicStaking.collectTokenTo(address,uint256)",
	{153, 168, 158, 204}: "PublicStaking.depositEth(uint8)",
	{129, 145, 245, 229}: "PublicStaking.depositToken(uint8,uint256)",
	{32, 234, 13, 72}:    "PublicStaking.estimateEthCollection(uint256)",
	{144, 89, 83, 237}:   "PublicStaking.estimateExcessEth()",
	{62, 237, 62, 255}:   "PublicStaking.estimateExcessToken()",
	{147, 197, 116, 141}: "PublicStaking.estimateTokenCollection(uint256)",
	{153, 120, 81, 50}:   "PublicStaking.getAccumulatorScaleFactor()",
	{8, 24, 18, 252}:     "PublicStaking.getApproved(uint256)",
	{84, 134, 82, 210}:   "PublicStaking.getEthAccumulator()",
	{9, 15, 112, 240}:    "PublicStaking.getMaxMintLock()",
	{134, 83, 164, 101}:  "PublicStaking.getMetamorphicContractAddress(bytes32,address)",
	{235, 2, 195, 1}:     "PublicStaking.getPosition(uint256)",
	{196, 124, 110, 20}:  "PublicStaking.getTokenAccumulator()",
	{133, 109, 232, 210}: "PublicStaking.getTotalReserveAToken()",
	{25, 184, 190, 47}:   "PublicStaking.getTotalReserveEth()",
	{213, 0, 47, 46}:     "PublicStaking.getTotalShares()",
	{129, 41, 252, 28}:   "PublicStaking.initialize()",
	{233, 133, 233, 197}: "PublicStaking.isApprovedForAll(address,address)",
	{228, 42, 103, 60}:   "PublicStaking.lockOwnPosition(uint256,uint256)",
	{12, 198, 93, 251}:   "PublicStaking.lockPosition(address,uint256,uint256)",
	{14, 78, 177, 91}:    "PublicStaking.lockWithdraw(uint256,uint256)",
	{160, 113, 45, 104}:  "PublicStaking.mint(uint256)",
	{43, 175, 42, 203}:   "PublicStaking.mintTo(address,uint256,uint256)",
	{6, 253, 222, 3}:     "PublicStaking.name()",
	{99, 82, 33, 30}:     "PublicStaking.ownerOf(uint256)",
	{66, 132, 46, 14}:    "PublicStaking.safeTransferFrom(address,address,uint256)",
	{184, 141, 79, 222}:  "PublicStaking.safeTransferFrom(address,address,uint256,bytes)",
	{162, 44, 180, 101}:  "PublicStaking.setApprovalForAll(address,bool)",
	{151, 27, 80, 91}:    "PublicStaking.skimExcessEth(address)",
	{122, 165, 7, 251}:   "PublicStaking.skimExcessToken(address)",
	{1, 255, 201, 167}:   "PublicStaking.supportsInterface(bytes4)",
	{149, 216, 155, 65}:  "PublicStaking.symbol()",
	{200, 123, 86, 221}:  "PublicStaking.tokenURI(uint256)",
	{35, 184, 114, 221}:  "PublicStaking.transferFrom(address,address,uint256)",
	{173, 253, 192, 63}:  "PublicStaking.tripCB()",
	{255, 7, 252, 14}:    "Snapshots.getAliceNetHeightFromLatestSnapshot()",
	{197, 232, 253, 225}: "Snapshots.getAliceNetHeightFromSnapshot(uint256)",
	{194, 234, 102, 3}:   "Snapshots.getBlockClaimsFromLatestSnapshot()",
	{69, 223, 197, 153}:  "Snapshots.getBlockClaimsFromSnapshot(uint256)",
	{52, 8, 228, 112}:    "Snapshots.getChainId()",
	{217, 193, 22, 87}:   "Snapshots.getChainIdFromLatestSnapshot()",
	{25, 247, 70, 105}:   "Snapshots.getChainIdFromSnapshot(uint256)",
	{2, 108, 43, 126}:    "Snapshots.getCommittedHeightFromLatestSnapshot()",
	{225, 140, 105, 122}: "Snapshots.getCommittedHeightFromSnapshot(uint256)",
	{117, 121, 145, 168}: "Snapshots.getEpoch()",
	{46, 238, 48, 206}:   "Snapshots.getEpochFromHeight(uint256)",
	{207, 232, 167, 59}:  "Snapshots.getEpochLength()",
	{213, 24, 242, 67}:   "Snapshots.getLatestSnapshot()",
	{134, 83, 164, 101}:  "Snapshots.getMetamorphicContractAddress(bytes32,address)",
	{66, 67, 141, 123}:   "Snapshots.getMinimumIntervalBetweenSnapshots()",
	{118, 241, 10, 208}:  "Snapshots.getSnapshot(uint256)",
	{209, 127, 204, 86}:  "Snapshots.getSnapshotDesperationDelay()",
	{124, 196, 204, 230}: "Snapshots.getSnapshotDesperationFactor()",
	{62, 204, 15, 94}:    "Snapshots.initialize(uint32,uint32)",
	{244, 95, 162, 70}:   "Snapshots.mayValidatorSnapshot(uint256,uint256,uint256,bytes32,uint256)",
	{174, 39, 40, 234}:   "Snapshots.migrateSnapshots(bytes[],bytes[])",
	{235, 124, 122, 254}: "Snapshots.setMinimumIntervalBetweenSnapshots(uint32)",
	{194, 232, 254, 242}: "Snapshots.setSnapshotDesperationDelay(uint32)",
	{63, 167, 161, 173}:  "Snapshots.setSnapshotDesperationFactor(uint32)",
	{8, 202, 31, 37}:     "Snapshots.snapshot(bytes,bytes)",
	{173, 203, 58, 68}:   "SnapshotsErrorCodes.SNAPSHOT_CALLER_NOT_ETHDKG_PARTICIPANT()",
	{141, 23, 222, 100}:  "SnapshotsErrorCodes.SNAPSHOT_CONSENSUS_RUNNING()",
	{12, 36, 85, 93}:     "SnapshotsErrorCodes.SNAPSHOT_INCORRECT_BLOCK_HEIGHT()",
	{63, 119, 44, 111}:   "SnapshotsErrorCodes.SNAPSHOT_INCORRECT_CHAIN_ID()",
	{152, 84, 189, 197}:  "SnapshotsErrorCodes.SNAPSHOT_MIGRATION_INPUT_DATA_MISMATCH()",
	{79, 42, 218, 238}:   "SnapshotsErrorCodes.SNAPSHOT_MIGRATION_NOT_ALLOWED()",
	{74, 30, 194, 238}:   "SnapshotsErrorCodes.SNAPSHOT_MIN_BLOCKS_INTERVAL_NOT_PASSED()",
	{131, 217, 149, 250}: "SnapshotsErrorCodes.SNAPSHOT_ONLY_VALIDATORS_ALLOWED()",
	{29, 68, 158, 209}:   "SnapshotsErrorCodes.SNAPSHOT_SIGNATURE_VERIFICATION_FAILED()",
	{133, 201, 219, 161}: "SnapshotsErrorCodes.SNAPSHOT_WRONG_MASTER_PUBLIC_KEY()",
	{33, 36, 29, 254}:    "ValidatorPool.CLAIM_PERIOD()",
	{97, 174, 225, 53}:   "ValidatorPool.MAX_INTERVAL_WITHOUT_SNAPSHOTS()",
	{156, 135, 227, 237}: "ValidatorPool.POSITION_LOCK_PERIOD()",
	{118, 156, 198, 149}: "ValidatorPool.claimExitingNFTPosition()",
	{201, 88, 224, 214}:  "ValidatorPool.collectProfits()",
	{143, 87, 153, 36}:   "ValidatorPool.completeETHDKG()",
	{156, 205, 248, 48}:  "ValidatorPool.getDisputerReward()",
	{217, 224, 220, 89}:  "ValidatorPool.getLocation(address)",
	{118, 32, 127, 156}:  "ValidatorPool.getLocations(address[])",
	{210, 153, 47, 84}:   "ValidatorPool.getMaxNumValidators()",
	{134, 83, 164, 101}:  "ValidatorPool.getMetamorphicContractAddress(bytes32,address)",
	{114, 37, 128, 182}:  "ValidatorPool.getStakeAmount()",
	{181, 216, 150, 39}:  "ValidatorPool.getValidator(uint256)",
	{192, 149, 20, 81}:   "ValidatorPool.getValidatorData(uint256)",
	{156, 125, 137, 97}:  "ValidatorPool.getValidatorsAddresses()",
	{39, 73, 130, 64}:    "ValidatorPool.getValidatorsCount()",
	{128, 216, 89, 17}:   "ValidatorPool.initialize(uint256,uint256,uint256)",
	{87, 181, 28, 156}:   "ValidatorPool.initializeETHDKG()",
	{32, 194, 133, 109}:  "ValidatorPool.isAccusable(address)",
	{200, 209, 165, 228}: "ValidatorPool.isConsensusRunning()",
	{228, 173, 117, 241}: "ValidatorPool.isInExitingQueue(address)",
	{24, 133, 87, 15}:    "ValidatorPool.isMaintenanceScheduled()",
	{250, 205, 116, 59}:  "ValidatorPool.isValidator(address)",
	{4, 141, 86, 199}:    "ValidatorPool.majorSlash(address,address)",
	{100, 192, 70, 28}:   "ValidatorPool.minorSlash(address,address)",
	{21, 11, 122, 2}:     "ValidatorPool.onERC721Received(address,address,uint256,bytes)",
	{30, 89, 117, 244}:   "ValidatorPool.pauseConsensus()",
	{188, 51, 187, 1}:    "ValidatorPool.pauseConsensusOnArbitraryHeight(uint256)",
	{101, 189, 145, 175}: "ValidatorPool.registerValidators(address[],uint256[])",
	{35, 128, 219, 26}:   "ValidatorPool.scheduleMaintenance()",
	{125, 144, 114, 132}: "ValidatorPool.setDisputerReward(uint256)",
	{130, 123, 251, 223}: "ValidatorPool.setLocation(string)",
	{108, 13, 160, 180}:  "ValidatorPool.setMaxNumValidators(uint256)",
	{67, 128, 140, 80}:   "ValidatorPool.setStakeAmount(uint256)",
	{151, 27, 80, 91}:    "ValidatorPool.skimExcessEth(address)",
	{122, 165, 7, 251}:   "ValidatorPool.skimExcessToken(address)",
	{238, 158, 73, 189}:  "ValidatorPool.tryGetTokenID(address)",
	{246, 68, 46, 36}:    "ValidatorPool.unregisterAllValidators()",
	{198, 232, 106, 214}: "ValidatorPool.unregisterValidators(address[])",
	{184, 230, 220, 77}:  "ValidatorPoolErrorCodes.VALIDATORPOOL_ADDRESS_ALREADY_VALIDATOR()",
	{14, 20, 77, 105}:    "ValidatorPoolErrorCodes.VALIDATORPOOL_ADDRESS_NOT_ACCUSABLE()",
	{76, 231, 141, 0}:    "ValidatorPoolErrorCodes.VALIDATORPOOL_ADDRESS_NOT_VALIDATOR()",
	{223, 111, 60, 38}:   "ValidatorPoolErrorCodes.VALIDATORPOOL_CALLER_NOT_VALIDATOR()",
	{142, 232, 55, 196}:  "ValidatorPoolErrorCodes.VALIDATORPOOL_CONSENSUS_RUNNING()",
	{76, 8, 175, 155}:    "ValidatorPoolErrorCodes.VALIDATORPOOL_DISHONEST_VALIDATOR_NOT_ACCUSABLE()",
	{145, 38, 175, 119}:  "ValidatorPoolErrorCodes.VALIDATORPOOL_ETHDKG_ROUND_RUNNING()",
	{130, 107, 113, 27}:  "ValidatorPoolErrorCodes.VALIDATORPOOL_ETH_BALANCE_CHANGED()",
	{82, 121, 145, 163}:  "ValidatorPoolErrorCodes.VALIDATORPOOL_FACTORY_SHOULD_OWN_POSITION()",
	{44, 32, 100, 122}:   "ValidatorPoolErrorCodes.VALIDATORPOOL_INSUFFICIENT_FUNDS_IN_STAKE_POSITION()",
	{56, 220, 230, 224}:  "ValidatorPoolErrorCodes.VALIDATORPOOL_INVALID_INDEX()",
	{251, 95, 106, 33}:   "ValidatorPoolErrorCodes.VALIDATORPOOL_MAX_VALIDATORS_MET()",
	{5, 127, 43, 41}:     "ValidatorPoolErrorCodes.VALIDATORPOOL_MINIMUM_STAKE_NOT_MET()",
	{8, 171, 238, 254}:   "ValidatorPoolErrorCodes.VALIDATORPOOL_MIN_BLOCK_INTERVAL_NOT_MET()",
	{117, 39, 169, 143}:  "ValidatorPoolErrorCodes.VALIDATORPOOL_ONLY_CONTRACTS_ALLOWED()",
	{12, 231, 212, 30}:   "ValidatorPoolErrorCodes.VALIDATORPOOL_PAYOUT_TOO_LOW()",
	{76, 69, 74, 151}:    "ValidatorPoolErrorCodes.VALIDATORPOOL_PROFITS_ONLY_CLAIMABLE_DURING_CONSENSUS()",
	{229, 181, 15, 230}:  "ValidatorPoolErrorCodes.VALIDATORPOOL_REGISTRATION_PARAMETER_LENGTH_MISMATCH()",
	{85, 239, 198, 88}:   "ValidatorPoolErrorCodes.VALIDATORPOOL_SENDER_NOT_IN_EXITING_QUEUE()",
	{50, 159, 163, 65}:   "ValidatorPoolErrorCodes.VALIDATORPOOL_TOKEN_BALANCE_CHANGED()",
	{124, 214, 45, 22}:   "ValidatorPoolErrorCodes.VALIDATORPOOL_VALIDATORS_GREATER_THAN_AVAILABLE()",
	{14, 45, 77, 85}:     "ValidatorPoolErrorCodes.VALIDATORPOOL_WAITING_PERIOD_NOT_MET()",
	{9, 94, 167, 179}:    "ValidatorStaking.approve(address,uint256)",
	{112, 160, 130, 49}:  "ValidatorStaking.balanceOf(address)",
	{66, 150, 108, 104}:  "ValidatorStaking.burn(uint256)",
	{234, 120, 90, 94}:   "ValidatorStaking.burnTo(address,uint256)",
	{137, 70, 92, 98}:    "ValidatorStaking.circuitBreakerState()",
	{42, 13, 139, 209}:   "ValidatorStaking.collectEth(uint256)",
	{190, 68, 67, 121}:   "ValidatorStaking.collectEthTo(address,uint256)",
	{227, 92, 62, 40}:    "ValidatorStaking.collectToken(uint256)",
	{136, 83, 185, 80}:   "ValidatorStaking.collectTokenTo(address,uint256)",
	{153, 168, 158, 204}: "ValidatorStaking.depositEth(uint8)",
	{129, 145, 245, 229}: "ValidatorStaking.depositToken(uint8,uint256)",
	{32, 234, 13, 72}:    "ValidatorStaking.estimateEthCollection(uint256)",
	{144, 89, 83, 237}:   "ValidatorStaking.estimateExcessEth()",
	{62, 237, 62, 255}:   "ValidatorStaking.estimateExcessToken()",
	{147, 197, 116, 141}: "ValidatorStaking.estimateTokenCollection(uint256)",
	{153, 120, 81, 50}:   "ValidatorStaking.getAccumulatorScaleFactor()",
	{8, 24, 18, 252}:     "ValidatorStaking.getApproved(uint256)",
	{84, 134, 82, 210}:   "ValidatorStaking.getEthAccumulator()",
	{9, 15, 112, 240}:    "ValidatorStaking.getMaxMintLock()",
	{134, 83, 164, 101}:  "ValidatorStaking.getMetamorphicContractAddress(bytes32,address)",
	{235, 2, 195, 1}:     "ValidatorStaking.getPosition(uint256)",
	{196, 124, 110, 20}:  "ValidatorStaking.getTokenAccumulator()",
	{133, 109, 232, 210}: "ValidatorStaking.getTotalReserveAToken()",
	{25, 184, 190, 47}:   "ValidatorStaking.getTotalReserveEth()",
	{213, 0, 47, 46}:     "ValidatorStaking.getTotalShares()",
	{129, 41, 252, 28}:   "ValidatorStaking.initialize()",
	{233, 133, 233, 197}: "ValidatorStaking.isApprovedForAll(address,address)",
	{228, 42, 103, 60}:   "ValidatorStaking.lockOwnPosition(uint256,uint256)",
	{12, 198, 93, 251}:   "ValidatorStaking.lockPosition(address,uint256,uint256)",
	{14, 78, 177, 91}:    "ValidatorStaking.lockWithdraw(uint256,uint256)",
	{160, 113, 45, 104}:  "ValidatorStaking.mint(uint256)",
	{43, 175, 42, 203}:   "ValidatorStaking.mintTo(address,uint256,uint256)",
	{6, 253, 222, 3}:     "ValidatorStaking.name()",
	{99, 82, 33, 30}:     "ValidatorStaking.ownerOf(uint256)",
	{66, 132, 46, 14}:    "ValidatorStaking.safeTransferFrom(address,address,uint256)",
	{184, 141, 79, 222}:  "ValidatorStaking.safeTransferFrom(address,address,uint256,bytes)",
	{162, 44, 180, 101}:  "ValidatorStaking.setApprovalForAll(address,bool)",
	{151, 27, 80, 91}:    "ValidatorStaking.skimExcessEth(address)",
	{122, 165, 7, 251}:   "ValidatorStaking.skimExcessToken(address)",
	{1, 255, 201, 167}:   "ValidatorStaking.supportsInterface(bytes4)",
	{149, 216, 155, 65}:  "ValidatorStaking.symbol()",
	{200, 123, 86, 221}:  "ValidatorStaking.tokenURI(uint256)",
	{35, 184, 114, 221}:  "ValidatorStaking.transferFrom(address,address,uint256)",
	{173, 253, 192, 63}:  "ValidatorStaking.tripCB()"}
