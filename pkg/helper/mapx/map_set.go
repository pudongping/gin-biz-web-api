// 集合
// 使用示例：
//
//	s := make(Set)
//	s.Add("Alex")
//	s.Add("Harry")
//	fmt.Println(s.Has("Alex")) ==> true
//	fmt.Println(s.Has("Jack"))  ==> false
package mapx

type Set map[string]struct{}

func (s Set) Add(key string) {
	s[key] = struct{}{}
}

func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s Set) Delete(key string) {
	delete(s, key)
}

func (s Set) AddAll(key ...string) {
	for _, v := range key {
		s.Add(v)
	}
}

func (s Set) HasAll(key ...string) (r []string, is bool) {
	var t = make(Set)

	for _, v := range key {
		if s.Has(v) {
			t.Add(v)
			r = append(r, v)
		}
	}

	is = len(key) == len(t)

	return
}

func (s Set) DeleteAll(key ...string) {
	for _, v := range key {
		s.Delete(v)
	}
}
