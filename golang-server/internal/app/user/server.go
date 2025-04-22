package example_server_app

import (
	"database/sql"
	auth_api "rz-server/internal/app/user/api/auth"
	user_api "rz-server/internal/app/user/api/user"
	auth_service "rz-server/internal/app/user/application/auth"
	user_service "rz-server/internal/app/user/application/user"
	"rz-server/internal/app/user/domain/auth"
	"rz-server/internal/app/user/domain/user"
	sql_store "rz-server/internal/app/user/infra/store/sql"
	auth_sql_store "rz-server/internal/app/user/infra/store/sql/auth"
	metadata_sql_store "rz-server/internal/app/user/infra/store/sql/metadata"
	user_sql_store "rz-server/internal/app/user/infra/store/sql/user"
	"rz-server/internal/common/interfaces"
)

var _ interfaces.ServerApp = (*ServerApp)(nil)

type ServerApp struct {
	server interfaces.Server
	event  <-chan interfaces.Event
	util   *interfaces.Util
	sqlDB  *sql.DB
}

func New(cmd *interfaces.CMD) *ServerApp {
	return &ServerApp{
		server: cmd.Server,
		event:  cmd.ConsumeEvent,
		util:   cmd.Util,
		sqlDB:  cmd.SqlDB,
	}
}

func (userApp *ServerApp) RegisterAPI() error {
	repository := sql_store.NewRepository(userApp.sqlDB, userApp.util)

	auth_store := auth_sql_store.New(repository)
	user_store := user_sql_store.New(repository)
	metadata_store := metadata_sql_store.New(repository)

	auth_domain := auth.New()
	authService := auth_service.New(auth_store, user_store, auth_domain)

	auth_api.New(
		userApp.server,
		authService,
		userApp.util,
	).Register()

	user_domain := user.New()
	user_service := user_service.New(user_store, metadata_store, user_domain)
	user_api.New(
		userApp.server,
		user_service,
		userApp.util,
	).Register()

	return nil
}

func (exampleApp *ServerApp) RegisterDomainEvent() error {
	return nil
}
