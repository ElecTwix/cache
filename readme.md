# Cache
   


## Usage

### With Any type

```go
abc := cache.NewCache[string, any](*time.NewTicker(1 * time.Second))
abc.Set("a", "b", 1*time.Second)
raw, ok := abc.Get("a") // raw is any
_, castok := raw.(string) // can cast to string
```


### With Data type

```go
abc := cache.NewCache[string, int](*time.NewTicker(1 * time.Second))
abc.Set("a", 15, 1*time.Second)
data, ok := abc.Get("a")  // data is int
```

### Also can use key with int

```go
abc := cache.NewCache[int, int](*time.NewTicker(1 * time.Second))
abc.Set(1, 15, 1*time.Second)
data, ok := abc.Get(1)  // data is int
```

### With custom type

```go

type User struct {
    Name string
    Age  int
}

abc := cache.NewCache[User, string](*time.NewTicker(1 * time.Second))
abc.Set("a", User{Name: "abc", Age: 15}, 1*time.Second)
data, ok := abc.Get("a")  // data is User
```


## License

[MIT](LICENSE)

## Benchmark

```bash

goos: linux
goarch: amd64
pkg: github.com/ElecTwix/cache
cpu: Intel(R) Core(TM) i7-6600U CPU @ 2.60GHz
BenchmarkSet-4                  14385724                71.32 ns/op            0 B/op          0 allocs/op
BenchmarkSetGetWithAny-4        15412664                72.94 ns/op            0 B/op          0 allocs/op
BenchmarkSetGetWithInt-4        14484638                69.35 ns/op            0 B/op          0 allocs/op
```

