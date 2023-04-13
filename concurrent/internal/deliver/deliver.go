package deliver

import (
	"ChessConcurrent/internal/usecase"
	"ChessConcurrent/internal/utils"
	"sync"
)

func Run() {
	var wg sync.WaitGroup
	wg.Add(8)
	go usecase.AddFirst(&wg)
	go usecase.AddSecond(&wg)
	go usecase.AddThird(&wg)
	go usecase.AddFourth(&wg)
	go usecase.AddFifth(&wg)
	go usecase.AddSixth(&wg)
	go usecase.AddSeventh(&wg)
	go usecase.AddEighth(&wg)
	wg.Wait()
	utils.PrintConditions()
}
