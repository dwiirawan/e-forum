package server

func (app *WebServer) UseModules(appModules ...AppModule) {
	for _, appModule := range appModules {
		appModule.Init(app)
	}
}

type AppModule interface {
	Init(*WebServer)
}
