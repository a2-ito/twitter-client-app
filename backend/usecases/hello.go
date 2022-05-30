package usecases

import (
	"context"
	"fmt"
)

func (u *appUseCase) Hello(ctx context.Context) (err error) {
	fmt.Println("useCase Hello")
	return nil
}
