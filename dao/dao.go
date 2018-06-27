package dao

// Dao handles read/write access to a given data source
type Dao interface {
	Find(args ...interface{}) ([]string, error)
	Close() error
}
