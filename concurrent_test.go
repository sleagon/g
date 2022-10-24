package g

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGo0(t *testing.T) {

	is := assert.New(t)

	// 直接返回
	{
		results, err := Go0(func(ctx context.Context) (string, error) {
			return "hello", nil
		}, func(ctx context.Context) (string, error) {
			return "world", nil
		})(context.Background())

		is.NoError(err)
		is.ElementsMatch([]string{"hello", "world"}, results)
	}

	// 返回错误
	{
		_, err := Go0(func(ctx context.Context) (string, error) {
			return "hello", nil
		}, func(ctx context.Context) (string, error) {
			return "", errors.New("world")
		})(context.Background())

		is.Error(err)
	}

	// context超时
	{
		ctx, cancel := context.WithTimeout(context.Background(), 0)
		defer cancel()

		_, err := Go0(func(ctx context.Context) (string, error) {
			return "hello", nil
		}, func(ctx context.Context) (string, error) {
			return "world", nil
		})(ctx)

		is.Error(err)
	}

	// panic
	{
		_, err := Go0(func(ctx context.Context) (string, error) {
			panic("hello")
		}, func(ctx context.Context) (string, error) {
			return "world", nil
		})(context.Background())

		is.Error(err)
	}

	// 休眠几秒
	{
		now := time.Now()

		results, err := Go0(func(ctx context.Context) (string, error) {
			time.Sleep(time.Second * 3)
			return "hello", nil
		}, func(ctx context.Context) (string, error) {
			time.Sleep(time.Second)
			return "world", nil
		})(context.Background())

		is.NoError(err)
		is.ElementsMatch([]string{"hello", "world"}, results)
		is.True(time.Since(now) >= time.Second*3)
		is.True(time.Since(now) < time.Second*4)
	}
}

func TestGo1(t *testing.T) {

	is := assert.New(t)

	// 直接返回
	{
		result, err := Go1(func(ctx context.Context, vi int) (string, error) {
			return "hello" + fmt.Sprint(vi), nil
		}, func(ctx context.Context, vi int) (string, error) {
			return "world" + fmt.Sprint(vi), nil
		}, func(ctx context.Context, vi int) (string, error) {
			return "!" + fmt.Sprint(vi), nil
		})(context.Background(), []int{1, 2, 3})

		is.NoError(err)
		is.ElementsMatch([]string{"hello1", "world2", "!3"}, result)
	}

	// 返回错误
	{
		_, err := Go1(func(ctx context.Context, vi int) (string, error) {
			return "hello" + fmt.Sprint(vi), nil
		}, func(ctx context.Context, vi int) (string, error) {
			return "world" + fmt.Sprint(vi), nil
		}, func(ctx context.Context, vi int) (string, error) {
			return "", errors.New("!")
		})(context.Background(), []int{1, 2, 3})

		is.Error(err)
	}

	// context超时
	{
		ctx, cancel := context.WithTimeout(context.Background(), 0)
		defer cancel()

		_, err := Go1(func(ctx context.Context, vi int) (string, error) {
			return "hello" + fmt.Sprint(vi), nil
		}, func(ctx context.Context, vi int) (string, error) {
			return "world" + fmt.Sprint(vi), nil
		}, func(ctx context.Context, vi int) (string, error) {
			return "!" + fmt.Sprint(vi), nil
		})(ctx, []int{1, 2, 3})

		is.Error(err)
	}

	// panic
	{
		_, err := Go1(func(ctx context.Context, vi int) (string, error) {
			panic("hello")
		}, func(ctx context.Context, vi int) (string, error) {
			return "world" + fmt.Sprint(vi), nil
		}, func(ctx context.Context, vi int) (string, error) {
			return "!" + fmt.Sprint(vi), nil
		})(context.Background(), []int{1, 2, 3})

		is.Error(err)
	}

	// 休眠几秒
	{
		now := time.Now()

		result, err := Go1(func(ctx context.Context, vi int) (string, error) {
			time.Sleep(time.Second * 3)
			return "hello" + fmt.Sprint(vi), nil
		}, func(ctx context.Context, vi int) (string, error) {
			time.Sleep(time.Second)
			return "world" + fmt.Sprint(vi), nil
		}, func(ctx context.Context, vi int) (string, error) {
			time.Sleep(time.Second)
			return "!" + fmt.Sprint(vi), nil
		})(context.Background(), []int{1, 2, 3})

		is.NoError(err)
		is.ElementsMatch([]string{"hello1", "world2", "!3"}, result)
		is.True(time.Since(now) >= time.Second*3)
		is.True(time.Since(now) < time.Second*4)
	}

}

func TestGo2(t *testing.T) {

	is := assert.New(t)

	// 直接返回
	{
		result, err := Go2(func(ctx context.Context, vi int, sj string) (string, error) {
			return "hello" + fmt.Sprint(vi) + sj, nil
		}, func(ctx context.Context, vi int, sj string) (string, error) {
			return "world" + fmt.Sprint(vi) + sj, nil
		}, func(ctx context.Context, vi int, sj string) (string, error) {
			return "!" + fmt.Sprint(vi) + sj, nil
		})(context.Background(), []int{1, 2, 3}, []string{"a", "b", "c"})

		is.NoError(err)
		is.ElementsMatch([]string{"hello1a", "world2b", "!3c"}, result)
	}

	// 返回错误
	{
		_, err := Go2(func(ctx context.Context, vi int, sj string) (string, error) {
			return "hello" + fmt.Sprint(vi) + sj, nil
		}, func(ctx context.Context, vi int, sj string) (string, error) {
			return "world" + fmt.Sprint(vi) + sj, nil
		}, func(ctx context.Context, vi int, sj string) (string, error) {
			return "", errors.New("!")
		})(context.Background(), []int{1, 2, 3}, []string{"a", "b", "c"})

		is.Error(err)
	}

	// context超时
	{
		ctx, cancel := context.WithTimeout(context.Background(), 0)
		defer cancel()

		_, err := Go2(func(ctx context.Context, vi int, sj string) (string, error) {
			return "hello" + fmt.Sprint(vi) + sj, nil
		}, func(ctx context.Context, vi int, sj string) (string, error) {
			return "world" + fmt.Sprint(vi) + sj, nil
		}, func(ctx context.Context, vi int, sj string) (string, error) {
			return "!" + fmt.Sprint(vi) + sj, nil
		}, func(ctx context.Context, vi int, sj string) (string, error) {
			return "!" + fmt.Sprint(vi) + sj, nil
		})(ctx, []int{1, 2, 3, 4}, []string{"a", "b", "c", "d"})

		is.Error(err)
	}

	// panic
	{
		_, err := Go2(func(ctx context.Context, vi int, sj string) (string, error) {
			panic("hello")
		}, func(ctx context.Context, vi int, sj string) (string, error) {
			return "world" + fmt.Sprint(vi) + sj, nil
		})(context.Background(), []int{1, 2, 3}, []string{"a", "b", "c"})

		is.Error(err)
	}

	// 休眠几秒
	{
		now := time.Now()

		result, err := Go2(func(ctx context.Context, vi int, sj string) (string, error) {
			time.Sleep(time.Second * 3)
			return "hello" + fmt.Sprint(vi) + sj, nil
		}, func(ctx context.Context, vi int, sj string) (string, error) {
			time.Sleep(time.Second)
			return "world" + fmt.Sprint(vi) + sj, nil
		}, func(ctx context.Context, vi int, sj string) (string, error) {
			time.Sleep(time.Second)
			return "!" + fmt.Sprint(vi) + sj, nil
		}, func(ctx context.Context, vi int, sj string) (string, error) {
			time.Sleep(time.Second)
			return "!" + fmt.Sprint(vi) + sj, nil
		})(context.Background(), []int{1, 2, 3, 4}, []string{"a", "b", "c", "d"})

		is.NoError(err)
		is.ElementsMatch([]string{"hello1a", "world2b", "!3c", "!4d"}, result)
		is.True(time.Since(now) >= time.Second*3)
		is.True(time.Since(now) < time.Second*4)

	}

}

func TestGo3(t *testing.T) {

	is := assert.New(t)

	// 直接返回
	{
		result, err := Go3(func(ctx context.Context, vi int, sj string, sk bool) (string, error) {
			return "hello" + fmt.Sprint(vi) + sj + fmt.Sprint(sk), nil
		}, func(ctx context.Context, vi int, sj string, sk bool) (string, error) {
			return "world" + fmt.Sprint(vi) + sj + fmt.Sprint(sk), nil
		}, func(ctx context.Context, vi int, sj string, sk bool) (string, error) {
			return "!" + fmt.Sprint(vi) + sj + fmt.Sprint(sk), nil
		})(context.Background(), []int{1, 2, 3}, []string{"a", "b", "c"}, []bool{true, false, true})

		is.NoError(err)
		is.ElementsMatch([]string{"hello1atrue", "world2bfalse", "!3ctrue"}, result)
	}

	// 返回错误
	{
		_, err := Go3(func(ctx context.Context, vi int, sj string, sk bool) (string, error) {
			return "hello" + fmt.Sprint(vi) + sj + fmt.Sprint(sk), nil
		}, func(ctx context.Context, vi int, sj string, sk bool) (string, error) {
			return "world" + fmt.Sprint(vi) + sj + fmt.Sprint(sk), nil
		}, func(ctx context.Context, vi int, sj string, sk bool) (string, error) {
			return "", errors.New("!")
		})(context.Background(), []int{1, 2, 3}, []string{"a", "b", "c"}, []bool{true, false, true})

		is.Error(err)
	}

	// context超时
	{
		ctx, cancel := context.WithTimeout(context.Background(), 0)
		defer cancel()

		_, err := Go3(func(ctx context.Context, vi int, sj string, sk bool) (string, error) {
			return "hello" + fmt.Sprint(vi) + sj + fmt.Sprint(sk), nil
		}, func(ctx context.Context, vi int, sj string, sk bool) (string, error) {
			return "world" + fmt.Sprint(vi) + sj + fmt.Sprint(sk), nil
		}, func(ctx context.Context, vi int, sj string, sk bool) (string, error) {
			return "!" + fmt.Sprint(vi) + sj + fmt.Sprint(sk), nil
		})(ctx, []int{1, 2, 3}, []string{"a", "b", "c"}, []bool{true, false, true})

		is.Error(err)
	}

	// panic
	{
		_, err := Go3(func(ctx context.Context, vi int, sj string, sk bool) (string, error) {
			panic("hello")
		}, func(ctx context.Context, vi int, sj string, sk bool) (string, error) {
			return "world" + fmt.Sprint(vi) + sj + fmt.Sprint(sk), nil
		}, func(ctx context.Context, vi int, sj string, sk bool) (string, error) {
			return "!" + fmt.Sprint(vi) + sj + fmt.Sprint(sk), nil
		})(context.Background(), []int{1, 2, 3}, []string{"a", "b", "c"}, []bool{true, false, true})

		is.Error(err)
	}

	// 休眠几秒
	{
		now := time.Now()

		result, err := Go3(func(ctx context.Context, vi int, sj string, sk bool) (string, error) {
			time.Sleep(time.Second * 3)
			return "hello" + fmt.Sprint(vi) + sj + fmt.Sprint(sk), nil
		}, func(ctx context.Context, vi int, sj string, sk bool) (string, error) {
			time.Sleep(time.Second)
			return "world" + fmt.Sprint(vi) + sj + fmt.Sprint(sk), nil
		}, func(ctx context.Context, vi int, sj string, sk bool) (string, error) {
			time.Sleep(time.Second)
			return "!" + fmt.Sprint(vi) + sj + fmt.Sprint(sk), nil
		})(context.Background(), []int{1, 2, 3}, []string{"a", "b", "c"}, []bool{true, false, true})

		is.NoError(err)
		is.ElementsMatch([]string{"hello1atrue", "world2bfalse", "!3ctrue"}, result)
		is.True(time.Since(now) > time.Second*3)
		is.True(time.Since(now) < time.Second*4)
	}

}
