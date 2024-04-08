package models

type CarInfo struct {
	RegNum       string `json:"regNum"`
	Mark         string `json:"mark"`
	Model        string `json:"model"`
	Year         int    `json:"year"`
	CurrentOwner Owner  `json:"owner"`
}

type Owner struct {
	Id         int    `json:"-"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type Car struct {
	RegNum  string `json:"regNum"`
	Mark    string `json:"mark"`
	Model   string `json:"model"`
	Year    int    `json:"year"`
	OwnerId int    `json:"owner"`
}

type RegNumCar struct {
	RegNum string
}

type SaveNums struct {
	RegNums []string `json:"regNums"`
}
