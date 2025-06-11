package repository

type UnitOfWork interface {
	Begin() (UnitOfWorkProcessor, error)
}

type UnitOfWorkProcessor interface {
	PaymentWriter() PaymentWriter

	Rollback() error
	Commit() error
}
