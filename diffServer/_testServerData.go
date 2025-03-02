package diffServer

type TestServerData struct {
	//Data       Data
	TestServer TestServer
}

func (e *TestServerData) Init() {
	e.TestServer.Init()
	//e.Data.Init(10)
}
