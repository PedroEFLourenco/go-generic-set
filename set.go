type ISet[T comparable] interface {
 Contains(value T) bool
 ContainsAll(values []T) bool
 Add(value T)
 AddAll(values []T)
 GetSize() int
 Remove(value T)
 RemoveAll(values []T)
 ToSlice() []T
 ExtractMapPrimitive() map[T]bool
}

type Set[T comparable] map[T]bool

func NewSetOfType[T comparable]() Set[T] {
 return make(map[T]bool)
}

func (s Set[T]) Contains(value T) bool {
 return s[value]
}

func (s Set[T]) ContainsAll(values []T) bool {
 for _, entry := range values {
  if !s[entry] {
   return false
  }
 }
 return true
}

func (s Set[T]) Add(value T) {
 s[value] = true
}

func (s Set[T]) AddAll(values []T) {
 for _, entry := range values {
  s[entry] = true
 }
}

func (s Set[T]) GetSize() int {
 return len(s)
}

func (s Set[T]) Remove(value T) {
 delete(s, value)
}

func (s Set[T]) RemoveAll(values []T) {
 for _, entry := range values {
  delete(s, entry)
 }
}

func (s Set[T]) ToSlice() []T {
 var uuidSlice []T

 for key := range s {
  uuidSlice = append(uuidSlice, key)
 }

 return uuidSlice
}

func (s Set[T]) ExtractMapPrimitive() map[T]bool {
 return s
}
