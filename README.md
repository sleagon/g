# G package



commonly used functions in go.package g.

## Usage

```go
func main() {
  g.Contains([]int{1, 2, 3}, 1)
  g.Reverse([]int{1, 2, 3})
}
```

## All functions

```
// import "github.com/sleagon/g"


FUNCTIONS

func Contains[T comparable](values []T, v T) bool
    Contains check v is contained by values.

func Filter[T any](values []T, f func(T, int) bool) []T
    Filter is a filter operator, like java's filter.

func Flat2[T any](values [][]T) []T
    Flat2 flatten 2-dimensional array to 1-dimensional array.

func Flat3[T any](values [][][]T) []T
    Flat3 flatten 3-dimensional array to 1-dimensional array.

func GroupBy[T any, K comparable](values []T, f func(T, int) K) map[K][]T
    GroupBy is a group by operator, like guava's group by.

func IndexOf[T comparable](values []T, v T) int
    IndexOf find the index of v in values, return -1 if element not found.

func IndexOfBy[T comparable](values []T, f func(T, int) bool) int
    IndexOfBy find the index of v in values, return -1 if element not found.

func Keys[K comparable, V any](values map[K]V) []K
    Keys returns the keys of the map, order is not guaranteed.

func LastIndexOf[T comparable](values []T, v T) int
    LastIndexOf find the last index of v in values, return -1 if element not
    found.

func LastIndexOfBy[T comparable](values []T, f func(T, int) bool) int
    LastIndexOfBy find the last index of v in values, return -1 if element not
    found.

func Map[F any, T any](values []F, f func(F, int) T) []T
    Map is a map operator, like js/java's map.

func MapEntries[K comparable, V any, T any](values map[K]V, f func(K, V) T) []T
    MapEntries returns the f(entries) of the map, order is not guaranteed.

func MapKeys[K comparable, V any, T any](values map[K]V, f func(K) T) []T
    MapKeys returns the f(key)s of the map, order is not guaranteed.

func MapValues[K comparable, V any, T any](values map[K]V, f func(V) T) []T
    MapValues returns the f(value)s of the map, order is not guaranteed.

func Max[T Numberic](values []T) T
    Max returns the maximum value in the given slice.

func MaxBy[T any](values []T, less func(T, T) bool) T
    MaxBy returns the maximum value in the given slice, using the provided less
    function to compare values.

func MaxKey[K Numberic, V any](values map[K]V) K
    MaxKey returns the maximum key of the map.

func MaxKeyBy[K comparable, V any](values map[K]V, less func(K, K) bool) K
    MaxKeyBy returns the maximum key of the map based on less method.

func MaxValue[K comparable, V Numberic](values map[K]V) V
    MaxValue returns the maximum value of the map.

func MaxValueBy[K comparable, V any](values map[K]V, less func(V, V) bool) V
    MaxValueBy returns the maximum value of the map based on less method.

func Min[T Numberic](values []T) T
    Min returns the minimum value in the given slice.

func MinBy[T any](values []T, less func(T, T) bool) T
    MinBy returns the minimum value in the given slice, using the provided less
    function to compare values.

func MinKey[K Numberic, V any](values map[K]V) K
    MinKey returns the minimum key of the map.

func MinKeyBy[K comparable, V any](values map[K]V, less func(K, K) bool) K
    MinKeyBy returns the minimum key of the map based on less method.

func MinValue[K comparable, V Numberic](values map[K]V) V
    MinValue returns the minimum value of the map.

func MinValueBy[K comparable, V any](values map[K]V, less func(V, V) bool) V
    MinValueBy returns the minimum value of the map based on less method.

func Must[T any](v T, err error) T
    Must assert a condition, and if it fails, panic.

func Must0(err error)
    Must0 assert err is nil

func Must1[T any](v T, err error) T
    MustN assert n conditions, and if it fails, panic.

func Must2[T1 any, T2 any](v1 T1, v2 T2, err error) (T1, T2)
    MustN assert n conditions, and if it fails, panic.

func Must3[T1 any, T2 any, T3 any](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3)
    MustN assert n conditions, and if it fails, panic.

func Must4[T1 any, T2 any, T3 any, T4 any](v1 T1, v2 T2, v3 T3, v4 T4, err error) (T1, T2, T3, T4)
    MustN assert n conditions, and if it fails, panic.

func Partition[T any](values []T, size int) [][]T
    Partition is a partition operator, like java's partition.

func PutAll[K comparable, V any](target map[K]V, elements map[K]V) map[K]V
    PutAll add all elements to exist map, if the key is already exists, the
    value will be overridden.

func Reduce[F any, T any](values []F, f func(T, F, int) T) T
    Reduce is a reduce operator, like js's reduce.

func Reverse[T any](values []T) []T
    Reverse reverse the order of elements in values.

func ReverseS(s string) string
    ReverseS returns a string with the reverse order of the given string.

func Slice2Map[K comparable, V any](values []V, f func(V, int) K) map[K]V
    Slice2Map transfer slice to a map base on given func.

func Ternary[T any](cond bool, v1 T, v2 T) T
    Ternary is a ternary operator, like C's ?:.

func TernaryF[T any](cond bool, f1 func() T, f2 func() T) T
    If is a if-else operator, like C's if-else.

func Uniq[T comparable](values []T) []T
    Uniq remove duplicated elements from values

func ValueOf[T any](ptr *T) T
    ValueOf returns the value of the given pointer.

func Values[K comparable, V any](values map[K]V) []V
    Values returns the values of the map, order is not guaranteed.


TYPES

type Numberic interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}
    Numberic is a numeric type, which can use ><= operators.

```