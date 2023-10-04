package repositories

type Registry interface {
	Database() DatabaseRepo
	Cache() CacheRepo
}

type impl struct {
	cacheRepo CacheRepo
	dbRepo    DatabaseRepo
}

func New() Registry {
	return &impl{}
}

func (i impl) Database() DatabaseRepo {
	return i.dbRepo
}

func (i impl) Cache() CacheRepo {
	return i.cacheRepo
}
