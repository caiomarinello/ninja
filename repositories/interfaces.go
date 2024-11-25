package repositories

// Registrar is the interface that wraps the basic Register method.
// Register is supposed to take a component and save it into a database.
type Registrar[T interface{}] interface {
	Register(component T) error
}

// Fetcher is the interface that wraps the methods FetchById and FetchAll.
type Fetcher[T interface{}] interface {
	FetchById(id int) (T, error)
	FetchAll() ([]T, error)
}

// Updater is the interface that wraps the basic Update method.
// Update is supposed to update a given component in the database.
type Updater[T interface{}] interface {
	Update(component T) error
}

// Deleter is the interface that wraps the basic Delete method.
type Deleter[T interface{}] interface {
	Delete(id int) error
}
