package main

func main() {
	count, amount := int64(100), int64(100)
	for i := int64(0); i < count; i++ {
		x := algo.SimpleRand(count, amount)
	}
}