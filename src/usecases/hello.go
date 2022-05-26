package usecases

import (
  "fmt"
  "context"
)

func (u *appUseCase) Hello(ctx context.Context) (err error) {
  fmt.Println("useCase Hello")
  return nil
}
