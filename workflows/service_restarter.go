package workflows

import (
	"github.com/stelligent/mu/common"
)

// NewServiceRestarter create a new workflow for a rolling restart
func NewServiceRestarter(ctx *common.Context, environmentName string, serviceName string, batchSize int) Executor {

	workflow := new(serviceWorkflow)

	return newPipelineExecutor(
		workflow.serviceInput(ctx, serviceName),
		workflow.serviceRestarter(environmentName),
	)
}

func (workflow *serviceWorkflow) serviceRestarter(environmentName string) Executor {
	return nil
}
