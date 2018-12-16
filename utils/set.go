package utils

// Exists 空结构体
var Exists = struct{}{}

// Set is the main interface
type Set struct {
	// struct为结构体类型的变量
	M map[interface{}]struct{}
}

// NewSet 初始化set
func NewSet(items ...interface{}) *Set {
	s := &Set{}
	// 声明map类型的数据结构
	s.M = make(map[interface{}]struct{})
	s.Add(items...)
	return s
}

// Add 向set中添加元素
func (s *Set) Add(items ...interface{}) error {
	for _, item := range items {
		s.M[item] = Exists
	}
	return nil
}

// Contains 查看set中是否存在item
func (s *Set) Contains(item interface{}) bool {
	_, ok := s.M[item]
	return ok
}

// Size  集合的大小
func (s *Set) Size() int {
	return len(s.M)
}

// Clear 清空集合
func (s *Set) Clear() {
	s.M = make(map[interface{}]struct{})
}

// Equal 判断两个set是否相等
func (s *Set) Equal(other *Set) bool {
	if s.Size() != other.Size() {
		return false
	}
	for key := range s.M {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

//IsSubset 判断s是否是other的子集
func (s *Set) IsSubset(other *Set) bool {
	if s.Size() > other.Size() {
		return false
	}
	for key := range s.M {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}
