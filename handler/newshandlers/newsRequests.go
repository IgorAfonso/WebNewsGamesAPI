package newsHandler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateNewsObject struct {
	Title         string `json:"title"`
	FirstContent  string `json:"firstContent"`
	SecondContent string `json:"secondContent"`
	ThirdContent  string `json:"thirdContent"`
}

type UpdateNewsObject struct {
	Title         string `json:"title"`
	FirstContent  string `json:"firstContent"`
	SecondContent string `json:"secondContent"`
	ThirdContent  string `json:"thirdContent"`
}