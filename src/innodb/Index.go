package innodb

type Index struct {
	Name            string
	Type            string
	IsPrimary       bool
	IsNotnull       bool
	IsUnique        bool
	IsAutoIncrement bool
}
