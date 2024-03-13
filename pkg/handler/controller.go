package handler

//控制层
type Controller struct {
	AnimalManager *AnimalManager
}

func NewController() *Controller {
	return &Controller{
		AnimalManager: &AnimalManager{},
	}
}
