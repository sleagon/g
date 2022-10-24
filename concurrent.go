package g

import (
	"context"
	"fmt"
)

type resultType[T any] struct {
	value T
	idx   int
}

func Go0[T any](funcs ...func(context.Context) (T, error)) func(context.Context) ([]T, error) {

	return func(ctx context.Context) ([]T, error) {

		rc, ec := make(chan resultType[T], len(funcs)), make(chan error, len(funcs))

		for i := range funcs {

			go func(idx int, vf func(context.Context) (T, error)) {
				defer func() {
					if err := recover(); err != nil {
						ec <- fmt.Errorf("panic: %v", err)
					}
				}()

				ret, e := vf(ctx)
				if e != nil {
					ec <- e
				} else {
					rc <- resultType[T]{value: ret, idx: idx}
				}
			}(i, funcs[i])
		}

		result := make([]T, len(funcs))
		for i := 0; i < len(funcs); i++ {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case err := <-ec:
				return nil, err
			case ret := <-rc:
				result[ret.idx] = ret.value
			}
		}

		return result, nil
	}
}

func Go[S any, T any](funcs ...func(context.Context, S) (T, error)) func(context.Context, []S) ([]T, error) {
	return Go1(funcs...)

}

func Go1[S any, T any](funcs ...func(context.Context, S) (T, error)) func(context.Context, []S) ([]T, error) {

	return func(ctx context.Context, values []S) ([]T, error) {

		rc, ec := make(chan resultType[T], len(values)), make(chan error, len(values))

		for i := range funcs {

			go func(idx int, vi S, vf func(context.Context, S) (T, error)) {
				defer func() {
					if err := recover(); err != nil {
						ec <- fmt.Errorf("panic: %v", err)
					}
				}()

				ret, e := vf(ctx, vi)
				if e != nil {
					ec <- e
				} else {
					rc <- resultType[T]{value: ret, idx: idx}
				}
			}(i, values[i], funcs[i])
		}

		result := make([]T, len(values))
		for i := 0; i < len(values); i++ {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case e := <-ec:
				return nil, e
			case r := <-rc:
				result[r.idx] = r.value
			}
		}

		return result, nil
	}

}

func Go2[S1 any, S2 any, T any](funcs ...func(context.Context, S1, S2) (T, error)) func(
	context.Context, []S1, []S2) ([]T, error) {

	return func(ctx context.Context, values1 []S1, values2 []S2) ([]T, error) {

		if len(values1) != len(values2) {
			return nil, fmt.Errorf("length of values1 and values2 must be equal")
		}

		rc, ec := make(chan resultType[T], len(values1)), make(chan error, len(values1))

		for i := range funcs {

			go func(idx int, vi1 S1, vi2 S2, vf func(context.Context, S1, S2) (T, error)) {
				defer func() {
					if err := recover(); err != nil {
						ec <- fmt.Errorf("panic: %v", err)
					}
				}()

				ret, e := vf(ctx, vi1, vi2)
				if e != nil {
					ec <- e
				} else {
					rc <- resultType[T]{value: ret, idx: idx}
				}
			}(i, values1[i], values2[i], funcs[i])
		}

		result := make([]T, len(values1))
		for i := 0; i < len(values1); i++ {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case e := <-ec:
				return nil, e
			case r := <-rc:
				result[r.idx] = r.value
			}
		}

		return result, nil
	}
}

func Go3[S1 any, S2 any, S3 any, T any](funcs ...func(context.Context, S1, S2, S3) (T, error)) func(
	context.Context, []S1, []S2, []S3) ([]T, error) {

	return func(ctx context.Context, values1 []S1, values2 []S2, values3 []S3) ([]T, error) {

		if len(values1) != len(values2) || len(values1) != len(values3) {
			return nil, fmt.Errorf("length of values1, values2 and values3 must be equal")
		}

		rc, ec := make(chan resultType[T], len(values1)), make(chan error, len(values1))

		for i := range funcs {

			go func(idx int, vi1 S1, vi2 S2, vi3 S3, vf func(context.Context, S1, S2, S3) (T, error)) {
				defer func() {
					if err := recover(); err != nil {
						ec <- fmt.Errorf("panic: %v", err)
					}
				}()

				ret, e := vf(ctx, vi1, vi2, vi3)
				if e != nil {
					ec <- e
				} else {
					rc <- resultType[T]{value: ret, idx: idx}
				}
			}(i, values1[i], values2[i], values3[i], funcs[i])
		}

		result := make([]T, len(values1))

		for i := 0; i < len(values1); i++ {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case e := <-ec:
				return nil, e
			case r := <-rc:
				result[r.idx] = r.value
			}
		}

		return result, nil
	}
}
