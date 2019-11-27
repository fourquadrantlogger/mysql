package innodb

type TableMeta struct {
	Name    string
	Fields  []Field
	Indexs  []Index
	MainKey Index
}
