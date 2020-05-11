# Parallelism vs Concurrency

Parallelism means running multiple operations at the same time. For example, your program got two threads to run and your machine
has two core. So you can run those threads in parallel. 

A Parallel program doesn't mean to be concurrent. A program that needs to be run sequentially is hard to make parallelize.
Concurrency is all about breaking the code into sub program that can potentially run at the same time and at the end produce same 
result as if it would run as single unit. So, we cannot parallelise a concurrent program.  

## Synchronous program example

```xml
func main() {
	print("zero")
	print("msi")
}

func print(str string) {

	for i := 0; i < 10; i++ {
		fmt.Println(str)
		time.Sleep(500)
	}
}
```

## Concurrent Program example 1:


```xml

func main() {

	go printStr("zero")
	printStr("msi")

}

func printStr(str string) {
	for i := 0; i < 50; i++ {
		fmt.Println(strconv.Itoa(i) + " " + str)
		time.Sleep(time.Second * 1)
	}

}


```

In this concurrent program, there's 2 go routine, 1 main and another one that is declared. 

