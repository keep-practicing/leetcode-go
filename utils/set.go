package utils

// Exists 空结构体
var Exists = struct{}{}

// Set data structure implement by go.
type Set map[interface{}]struct{}

// NewSet 初始化set
func NewSet(items ...interface{}) Set {
	s := make(Set)
	s.Add(items...)
	return s
}

// Add 向set中添加元素
func (s Set) Add(items ...interface{}) error {
	for _, item := range items {
		s[item] = Exists
	}
	return nil
}

// Contains 查看set中是否存在item
func (s Set) Contains(item interface{}) bool {
	_, ok := s[item]
	return ok
}

// Size  集合的大小
func (s Set) Size() int {
	return len(s)
}

// Clear 清空集合
func (s Set) Clear() {
	s = make(Set)
}

// Equal 判断两个set是否相等
func (s Set) Equal(other Set) bool {
	if s.Size() != other.Size() {
		return false
	}
	for key := range s {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

//IsSubset 判断s是否是other的子集
func (s Set) IsSubset(other Set) bool {
	if s.Size() > other.Size() {
		return false
	}
	for key := range s {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}
