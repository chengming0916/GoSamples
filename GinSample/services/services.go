package services

import "GoSamples/GinSample/domain"

type UserManager struct {
}

func (p *UserManager) FindAll() {

}

func (p *UserManager) Create(user *domain.User) error {
	return domain.Create(user)
}
