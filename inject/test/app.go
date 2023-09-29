package test

import "github.com/katsuokaisao/gorm/usecase/gorm"

type Application struct {
	gormUseCase gorm.GormUseCase
}

func NewApplication(gormUseCase gorm.GormUseCase) *Application {
	return &Application{
		gormUseCase: gormUseCase,
	}
}

func (a *Application) Run() {
	a.gormUseCase.Test()
}