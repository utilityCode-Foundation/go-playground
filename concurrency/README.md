# Parallelism vs Concurrency

Parallelism means running two operations at the same time. For example, your program got two threads to run and your machine
has two core. So you can run those threads in parallel. 

A Parallel programming doesn't mean to be concurrent. Concurrency means running chunks of a program as individual
executing tasks and producing right result. 

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



