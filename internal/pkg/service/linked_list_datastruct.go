package service

type linkedList struct {
	head   *node
	length int
}

type node struct {
	value int64
	next  *node
}
