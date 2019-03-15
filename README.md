# warm-up

Warm up is a pseudo implementation of `golang.org/x/time/rate` with the biggest difference being that you can supply a
custom function with which to calculate the rate at every token. Please see the [Golang documentation](https://godoc.org/golang.org/x/time/rate) for details on the
terminology used here, especially _Limiter_, _tokens_, and _bucket_.

### Use cases:
- Sending lots of requests to a system, but you want to give the system a chance to auto-scale.
- Warming up a load balancer to be ready to handle a lot of requests.
- Gradually increasing load to a system.

## Example
```go
func main() {
	fmt.Println("starting...")

	// Define our custom function
	f := func(x float64) float64 {
		return math.Pow(x/10, 2)
	}

	// Create a new limiter that will increase in throughput according to (x/10)^2,
	// with a burst up to 0 tokens.
	l := warmup.NewLimiter(f, 0)

	i := 0
	for {
		ctx := context.Background()
		l.WaitN(ctx, 1)
		fmt.Println("i", i)
		i++
	}
}
```

This example will print out at each iteration, slowly increasing over time.

## License

Copyright © 2019 Trevor Hutto

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this work except in compliance with the License. You may obtain a copy of the License in the LICENSE file, or at:

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.