package entities

type StorageInforTbl struct {
	Date     string `json:"date"`
	Type     string `json:"type"`
	LineID   int    `json: "line_id"`
	HashCode string `json:"hash_code"`
}

func NewInfor(Date string, Type string, LineID int, HashCode string) StorageInforTbl {
	return StorageInforTbl{Date, Type, LineID, HashCode}
}
