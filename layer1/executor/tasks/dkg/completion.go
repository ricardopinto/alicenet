package dkg

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/alicenet/alicenet/layer1/executor/tasks"
	"github.com/alicenet/alicenet/layer1/executor/tasks/dkg/state"
	"github.com/alicenet/alicenet/layer1/executor/tasks/dkg/utils"
)

// CompletionTask contains required state for safely complete ETHDKG
type CompletionTask struct {
	*tasks.BaseTask
	// variables that are unique only for this task
	StartBlockHash common.Hash `json:"startBlockHash"`
}

// asserting that CompletionTask struct implements interface tasks.Task
var _ tasks.Task = &CompletionTask{}

// NewCompletionTask creates a background task that attempts to call Complete on ethdkg
func NewCompletionTask(start uint64, end uint64) *CompletionTask {
	return &CompletionTask{
		BaseTask: tasks.NewBaseTask(start, end, false, nil),
	}
}

// Prepare prepares for work to be done in the CompletionTask
func (t *CompletionTask) Prepare(ctx context.Context) *tasks.TaskErr {
	logger := t.GetLogger().WithField("method", "Prepare()")
	logger.Debug("preparing task")

	dkgState, err := state.GetDkgState(t.GetDB())
	if err != nil {
		return tasks.NewTaskErr(fmt.Sprintf(tasks.ErrorLoadingDkgState, err), false)
	}

	if dkgState.Phase != state.DisputeGPKJSubmission {
		return tasks.NewTaskErr("it's not in DisputeGPKJSubmission phase", false)
	}

	// setup leader election
	block, err := t.GetClient().GetBlockByNumber(ctx, big.NewInt(int64(t.GetStart())))
	if err != nil {
		return tasks.NewTaskErr(fmt.Sprintf("CompletionTask.Prepare(): error getting block by number: %v", err), true)
	}

	t.StartBlockHash = block.Hash()

	return nil
}

// Execute executes the task business logic
func (t *CompletionTask) Execute(ctx context.Context) (*types.Transaction, *tasks.TaskErr) {
	logger := t.GetLogger().WithField("method", "Execute()")
	logger.Debug("initiate execution")

	dkgState, err := state.GetDkgState(t.GetDB())
	if err != nil {
		return nil, tasks.NewTaskErr(fmt.Sprintf(tasks.ErrorLoadingDkgState, err), false)
	}

	client := t.GetClient()
	// submit if I'm a leader for this task
	if !utils.AmILeading(client, ctx, logger, int(t.GetStart()), t.StartBlockHash.Bytes(), dkgState.NumberOfValidators, dkgState.Index) {
		return nil, tasks.NewTaskErr("not leading Completion yet", true)
	}

	c := t.GetContractsHandler().EthereumContracts()
	txnOpts, err := client.GetTransactionOpts(ctx, dkgState.Account)
	if err != nil {
		return nil, tasks.NewTaskErr(fmt.Sprintf(tasks.FailedGettingTxnOpts, err), true)
	}

	// Complete ETHDKG
	logger.Info("Trying to complete ETHDKG")
	txn, err := c.Ethdkg().Complete(txnOpts)
	if err != nil {
		return nil, tasks.NewTaskErr(fmt.Sprintf("completion failed: %v", err), true)
	}

	return txn, nil
}

// ShouldExecute checks if it makes sense to execute the task
func (t *CompletionTask) ShouldExecute(ctx context.Context) (bool, *tasks.TaskErr) {
	logger := t.GetLogger().WithField("method", "ShouldExecute()")
	logger.Debug("should execute task")

	eth := t.GetClient()
	c := t.GetContractsHandler().EthereumContracts()

	callOpts, err := eth.GetCallOpts(ctx, eth.GetDefaultAccount())
	if err != nil {
		return false, tasks.NewTaskErr(fmt.Sprintf(tasks.FailedGettingCallOpts, err), true)
	}
	phase, err := c.Ethdkg().GetETHDKGPhase(callOpts)
	if err != nil {
		return false, tasks.NewTaskErr(fmt.Sprintf("error getting ethdkg phases in completion BaseTask: %v", err), true)
	}

	if phase == uint8(state.Completion) {
		logger.Debugf("completion already ocurred: %v", phase)
		return false, nil
	}

	return true, nil
}
