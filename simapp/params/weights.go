package params

// Default simulation operation weights for messages and gov proposals
const (
	DefaultWeightMsgSend                        int = 100
	DefaultWeightMsgMultiSend                   int = 10
	DefaultWeightMsgSetWithdrawAddress          int = 50
	DefaultWeightMsgWithdrawDelegationReward    int = 50
	DefaultWeightMsgWithdrawValidatorCommission int = 50
	DefaultWeightMsgFundCommunityPool           int = 50
	DefaultWeightMsgDeposit                     int = 100
	DefaultWeightMsgVote                        int = 67
	DefaultWeightMsgVoteWeighted                int = 33
	DefaultWeightMsgUnjail                      int = 100
	DefaultWeightMsgCreateValidator             int = 100
	DefaultWeightMsgEditValidator               int = 5
	DefaultWeightMsgDelegate                    int = 100
	DefaultWeightMsgUndelegate                  int = 100
	DefaultWeightMsgBeginRedelegate             int = 100

	DefaultWeightCommunitySpendProposal int = 5
	DefaultWeightTextProposal           int = 5
	DefaultWeightParamChangeProposal    int = 5

	// feegrant
	DefaultWeightGrantAllowance  int = 100
	DefaultWeightRevokeAllowance int = 100

	// farming
	// TODO: determine the weights
	DefaultWeightMsgCreateFixedAmountPlan int = 0
	DefaultWeightMsgCreateRatioPlan       int = 0
	DefaultWeightMsgStake                 int = 0
	DefaultWeightMsgUnstake               int = 0
	DefaultWeightMsgClaim                 int = 0
)
