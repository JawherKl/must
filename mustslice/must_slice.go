package mustslice

import (
	"slices"

	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Equals checks if two slices are equal, panics if not.
// Equals 检查两个切片是否相等，不相等则触发 panic。
func Equals[V comparable](a, b []V) {
	if !slices.Equal(a, b) {
		zaplog.ZAPS.P1.LOG.Panic("expect slice equals while not equals", zap.Int("len_a", len(a)), zap.Int("len_b", len(b)))
	}
}

// In checks if an element exists in a slice, panics if not.
// In 检查某个元素是否存在于切片中，不存在则触发 panic。
func In[T comparable](v T, a []T) {
	if !slices.Contains(a, v) {
		zaplog.ZAPS.P1.LOG.Panic("expect value in slice while not in", zap.Any("v", v), zap.Int("len", len(a)))
	}
}

// Contains checks if a slice contains a specific element, panics if not.
// Contains 检查切片是否包含某个特定元素，不包含则触发 panic。
func Contains[T comparable](a []T, v T) {
	if !slices.Contains(a, v) {
		zaplog.ZAPS.P1.LOG.Panic("expect slice contains value while not contains", zap.Int("len", len(a)), zap.Any("v", v))
	}
}

// Have ensures the slice is not empty, panics if it is.
// Have 确保切片不为空，为空则触发 panic。
func Have[T any](a []T) {
	if len(a) == 0 {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH > 0 while got an none slice")
	}
}

// Nice ensures the slice is not empty, and returns it if it contains elements.
// Nice 确保切片不为空，若切片有元素则返回它。
func Nice[T any](a []T) []T {
	if len(a) == 0 {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH > 0 while got an none slice")
	}
	return a
}

// Length checks if the slice's length equals the expected value, panics if not.
// Length 检查切片的长度是否等于期望值，不等则触发 panic。
func Length[T any](a []T, n int) {
	if len(a) != n {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH = n while not equals", zap.Int("len", len(a)), zap.Int("n", n))
	}
}

// Len checks if the slice's length equals the expected value, panics if not.
// Len 检查切片的长度是否等于期望值，不等则触发 panic。
func Len[T any](a []T, n int) {
	if len(a) != n {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH = n while not equals", zap.Int("len", len(a)), zap.Int("n", n))
	}
}