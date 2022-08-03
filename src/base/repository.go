package base

type IRepositoryBase interface {
	BeginTransaction()
	Commit()
	Rollback()
}
