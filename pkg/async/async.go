package async

import "context"

type Promise interface {
	Await() (<-chan interface{}, error)
}

type promise struct {
	await func(ctx context.Context) (<-chan interface{}, error)
}

func (f promise) Await() (<-chan interface{}, error) {
	return f.await(context.Background())
}

// Exec executes the async function
func Exec(f func(ctx context.Context) <-chan interface{}) Promise {
	var result <-chan interface{}
	result = make(<-chan interface{})
	c := make(chan struct{})
	go func() {
		defer close(c)
		result = f(context.Background())
	}()
	return promise{
		await: func(ctx context.Context) (<-chan interface{}, error) {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()

			case <-c:
				return result, nil
			}
		},
	}
}
