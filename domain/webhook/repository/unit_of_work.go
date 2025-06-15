package repository

type UnitOfWork interface {
	Begin() (UnitOfWorkProcessor, error)
}

type UnitOfWorkProcessor interface {
	PaymentWriter() PaymentWriter
	OutboxWriter() OutboxWriter

	Rollback() error
	Commit() error
}
