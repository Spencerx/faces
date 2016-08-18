package execution

import lexec "github.com/reconquest/lexec-go"

type Operation struct {
	*lexec.Execution
}

func (operation *Operation) Start() error {
	return operation.Execution.Start()
}

func (operation *Operation) Run() error {
	err := operation.Start()
	if err != nil {
		return err
	}

	err = operation.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (operation *Operation) Wait() error {
	err := operation.Execution.Wait()
	if err != nil {
		return expandError(
			err, operation.Execution.GetStreamsData(),
		)
	}

	return nil
}

func (operation *Operation) Output() ([]byte, []byte, error) {
	stdout, stderr, err := operation.Execution.Output()
	if err != nil {
		return nil, nil, expandError(
			err, operation.Execution.GetStreamsData(),
		)
	}

	return stdout, stderr, nil
}

func (operation *Operation) Kill() error {
	return operation.Execution.Process().Kill()
}
