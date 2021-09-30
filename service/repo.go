package service

import (
	"monitor/internal/model"

	"git.ultraware.nl/NiseVoid/qb"
	"git.ultraware.nl/NiseVoid/qb/qbdb"
	"git.ultraware.nl/NiseVoid/qb/qc"
	"git.ultraware.nl/NiseVoid/qb/qf"
)

func upsert(db qbdb.Target, service Service) {
	s := model.Service()

	uq := s.Update().
		Set(s.State, qf.Excluded(s.State))

	iq := s.Insert(
		s.Name,
		s.ObjectPath,
		s.State,
	).Values(
		service.Name,
		service.ObjectPath,
		service.State,
	).Upsert(uq, s.Name)

	db.MustExec(iq)
}

func getAll(db qbdb.Target) []Service {
	s := model.Service()

	q := s.Select(
		s.ID,
		s.Name, s.ObjectPath,
		s.State,
	).OrderBy(qb.Asc(s.ID))

	rows := db.MustQuery(q)
	defer rows.Close() //nolint: errcheck

	services := []Service{}
	for rows.Next() {
		service := Service{}
		rows.MustScan(
			&service.ID,
			&service.Name, &service.ObjectPath,
			&service.State,
		)

		services = append(services, service)
	}

	return services
}

func getByID(db qbdb.Target, id int) (Service, bool) {
	s := model.Service()

	q := s.Select(
		s.ID,
		s.Name, s.ObjectPath,
		s.State,
	).
		Where(qc.Eq(s.ID, id))

	service := Service{}
	return service, db.QueryRow(q).MustScan(
		&service.ID,
		&service.Name, &service.ObjectPath,
		&service.State,
	)
}

func createSubscription(db qbdb.Target, id int) {
	s := model.Subscription()

	q := s.
		Insert(s.ServiceID).
		Values(id).
		IgnoreConflict(s.ServiceID)

	db.MustExec(q)
}

func getAllSubscriptions(db qbdb.Target) []Subscription {
	s := model.Subscription()

	q := s.Select(
		s.ID,
		s.ServiceID,
	)

	rows := db.MustQuery(q)
	defer rows.Close()

	subscriptions := []Subscription{}
	for rows.Next() {
		subscription := Subscription{}
		rows.MustScan(
			&subscription.ID, &subscription.ServiceID,
		)

		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions
}

func deleteSubscription(db qbdb.Target, id int) {
	s := model.Subscription()

	q := s.Delete(qc.Eq(s.ID, id))

	db.MustExec(q)
}
