package diffServer

type ConnectDataForm interface {
	SetUrl(url string)
	SetMethod(method string)
	SetContent(content string)
	SetHeader(key, value string)
	SetParam(key, value string)

	GetHeader() []keyValue
	GetParam() []keyValue
	GetUrl() string
	GetMethod() string
	GetContent() string
	GetFieldWidth() int
}
