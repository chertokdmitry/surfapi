package services

import (
	"gitlab.com/chertokdmitry/surfapi/src/domain/subscriptions"
	"gitlab.com/chertokdmitry/surfapi/src/utils/errors"
)

var (
	SubscriptionService subscriptionServiceInterface = &subscriptionService{}
)

type subscriptionService struct {
}

type subscriptionServiceInterface interface {
	GetByChatId(chatId int64) ([]subscriptions.Sub, *errors.RestErr)
	Delete(subId int64) *errors.RestErr
	Insert(sub *subscriptions.Sub) *errors.RestErr
	//GetAll() ([]spots.Spot, *errors.RestErr)
}

func (ss *subscriptionService) Insert(s *subscriptions.Sub) *errors.RestErr {
	return s.Insert()
}

func (ss *subscriptionService) Delete(subId int64) *errors.RestErr {
	s := &subscriptions.Sub{Id: subId}
	return s.Delete()
}

func (ss *subscriptionService) GetByChatId(chatId int64) ([]subscriptions.Sub, *errors.RestErr) {
	s := &subscriptions.Sub{}
	result, err := s.GetByChatId(chatId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//func (ss *spotsService) GetAll() ([]spots.Spot, *errors.RestErr) {
//	s := &spots.Spot{}
//	result, err := s.GetAll()
//	if err != nil {
//		return nil, err
//	}
//
//	return result, nil
//}
