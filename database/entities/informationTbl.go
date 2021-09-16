package entities

type InformationTbl struct {
	Date        string
	Type        string
	LineID      int
	Information string
}

func NewInformationTbl(Date string, Type string, LineID int, Information string) InformationTbl {
	return InformationTbl{Date, Type, LineID, Information}
}
