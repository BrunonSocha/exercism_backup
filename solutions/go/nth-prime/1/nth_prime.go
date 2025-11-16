package prime
import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("negative integers and 0 are not valid input values.")
	}
	var primes []int
	primes = append(primes, 2)
	i := 3
	for {
		if len(primes) == n {
			break
		}
		prime := true
		for j := 2; j < i; j++ {
			if i % j == 0 {
				prime = false
			}
		}
		if prime {
			primes = append(primes, i)
		}
		i++
	}
	return primes[n-1], nil
}
