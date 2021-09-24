package entities

type StorageInforTbl struct {
	Date     string
	Type     string
	LineID   int
	HashCode string
}

func NewInfor(Date string, Type string, LineID int, HashCode string) StorageInforTbl {
	return StorageInforTbl{Date, Type, LineID, HashCode}
}
