# Factory Pattern
```xml 
the singleton pattern is a software design pattern that restricts the instantiation of a class to one object.
```

### Objectives
As a general guide, we consider using the Singleton pattern when the following rule
applies:
- We need a single, shared value, of some particular type.
- We need to restrict object creation of some type to a single unit along the entire
program.

To translate that in Go, whenever any goroutine tries to access an instance of a variable, you should get the same variable.

There are various ways to implement the singleton pattern in Go, but it is also quite easy to get it wrong. Let’s look at the different ways you can implement a singleton pattern and why most of them are wrong.

GetInstance instead of New constructor
If you come from other programming languages, this is the most common way to implement the singleton pattern. However, there are inherent problems with thread-safety in this code.

package singleton

type singleton struct {}

var instance *singleton

func GetInstance() *singleton {
    if instance == nil {
        instance = &singleton{}
    }
    return instance
}
In the above code, let’s say multiple goroutines come to the if condition and check if the instance is nil at the same time, each goroutine would create its own reference of the variable. There is absolutely no guarantee which instance would be returned to which goroutine.

Moreover, if multiple goroutines are holding different references of the variable and passing it around, each variable would have a different state, leading to inconsistent behavior.

During development, it might look like everything is working fine. Usually, these kinds of bugs crop up only during production, when you are operating under heavy load and spawning thousands of goroutines.

Mutex locks
If you have a thread safety issue, it makes sense to put a lock around the code you want to protect. While it might work, it does lead to other complications.

Take a look at this code below.

var mut Sync.Mutex

func GetInstance() *singleton {
    // Lock and unlock the entire GetInstance function
    mut.Lock()
    defer mu.Unlock()

    if instance == nil {
        instance = &singleton{}
    }
    return instance
}
In the above code, as soon as you enter the function, you are using a mutex lock, and only then you are checking whether your instance exists or not. By making sure your code is covered by a Lock() and Unlock() statement, only one goroutine can execute this code.

However, do you need such aggressive locking and unlocking to every call of this function? The if block is going to be executed only once in the entire lifetime. Once the instance has been created, you wouldn’t need to lock. By using mutex locking, you have created an unnecessary bottleneck in your entire program.

Check-Lock-Check pattern
The next obvious solution is to check if you need to lock, lock, and then doing a second check again before performing the actual operation. To explain it in python,

if not instance_exists():
    lock()
    if not instance_exists():
        create_instance()
return instance
The first check is to determine whether we need to lock or not. Once we acquire the lock, we need to check once again if the instance exists, because between the time we did the first check and the lock statement, some other thread could have created the instance. If the second check also passes, we create the instance.

If the first check fails, we don’t have to wait for an expensive lock to acquire. We return the existing instance of the variable.

Doing the same logic in Golang:

func GetInstance() *singleton {
    if instance == nil {
        mu.Lock()
        defer mu.Unlock()

        if instance == nil {
            instance = &singleton{}
        }
    }
    return instance
}
Though this code is a lot better than the old code, there is still a teeny-tiny chance that it could fail. The if condition check is not an atomic check and though it might work fine for 99.99% of the times, it is still technically the wrong way to do it.

sync package to the rescue
The right way to implement a singleton pattern in Go is to use the sync package’s Once.Do() function. This function makes sure that your specified code is executed only once and never more than once.

The way to use the Once.Do() function is as below.

once.Do(func() {
    // This will be executed only once in the entire lifetime of the program
})
Here is a sample program which shows how even if you call it multiple times, it gets executed only once.

package main
import (
    "fmt"
    "sync"
)
var once sync.Once
func main() {
    for i := 0; i < 10; i++ {
        only_once()
    }
}
func only_once() {
    once.Do(func() { fmt.Println("Hi") })
}
Run in Playground

If you want to understand how it works, take a look at the source-code of the function.

You can see that it uses the sync.atomic package’s StoreUint32 and LoadUint32 functions to store/load an integer value and using that to perform the checks. The difference between the normal plain if we did earlier vs. this, there is a 100% guarantee that the variable read will be atomic.

Now if we use Once.Do() in our singleton package, we get the following code.

package singleton

import (
    "sync"
)

type singleton struct {}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
    once.Do(func() {
        instance = &singleton{}
    })
    return instance
}
This looks much simpler and easier to understand. And this also assures you that the reference to the variable you get out of this function call is always the same.


Copied from [here](https://progolang.com/how-to-implement-singleton-pattern-in-go/).