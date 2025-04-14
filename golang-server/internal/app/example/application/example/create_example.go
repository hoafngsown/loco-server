package example_service

import (
	example_commands "rz-server/internal/app/example/application/example/commands"
	example_data "rz-server/internal/app/example/application/example/data"
	"rz-server/internal/common/errors/application_error"
)

func (u *ExampleService) CreateExample(command example_commands.CreateExampleCommand) (*example_data.ExampleData, *application_error.Error) {
	u.store.CreateExample("Hello")
	return nil, nil
}
