package service

import "git.ultraware.nl/NiseVoid/qb/qbdb"

func Upsert(db qbdb.Target, s Service) {
	upsert(db, s)
}

func GetAll(db qbdb.Target) []Service {
	return getAll(db)
}

func GetByID(db qbdb.Target, id int) (Service, bool) {
	return getByID(db, id)
}

func CreateSubscription(db qbdb.Target, serviceID int) {
	createSubscription(db, serviceID)
}

func GetAllSubscriptions(db qbdb.Target) []Subscription {
	return getAllSubscriptions(db)
}

func DeleteSubscription(db qbdb.Target, id int) {
	deleteSubscription(db, id)
}
