package diffServer

var KMethod = []string{"GET", "POST"} //, "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}

const (
	KStatusWaitUserAction = "Esperando usuário"
)

const (
	KColorRed    = "\033[31m" // Vermelho
	KColorGreen  = "\033[32m" // Verde
	KColorYellow = "\033[33m" // Amarelo
	KColorBlue   = "\033[34m" // Azul
	KColorNormal = "\033[0m"  // Reset para cor padrão
)
