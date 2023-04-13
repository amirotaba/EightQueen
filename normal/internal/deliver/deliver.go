package deliver

import (
	"chess_normal/internal/entity"
	"chess_normal/internal/usecase"
	"chess_normal/internal/utils"
	"fmt"
)

func Run() {
	usecase.Cal()
	entity.Data = utils.Unique(entity.Data)
	for _, d := range entity.Data {
		fmt.Println(d)
	}
	fmt.Println("total conditions: ", len(entity.Data))
}
