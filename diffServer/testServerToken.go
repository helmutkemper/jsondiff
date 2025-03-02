package diffServer

type testServerToken struct {
	Token      Token
	TestServer TestServer
}

func (e *testServerToken) init() {
	//e.token.init()
	e.TestServer.Init()
}
