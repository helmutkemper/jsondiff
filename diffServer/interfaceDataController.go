package diffServer

type DataConfig interface {
	SetDataKeys(keys string)
	SetDeleteKeys(deleteKeys string)
	SetAmount(amount string)
	SetInteractions(interactions string)
	SetNumberOfKeys(numberOfKeys string)
	GetFieldWidth() (fieldWidth int)
}
