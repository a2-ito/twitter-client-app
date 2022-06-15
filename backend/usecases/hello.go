package usecases

import (
	"context"
	"fmt"
)

// テスト用API 削除可能
func (u *appUseCase) Hello(ctx context.Context) (err error) {
	fmt.Println("useCase Hello")
	return nil
}
