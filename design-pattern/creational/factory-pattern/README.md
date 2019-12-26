# Factory Pattern
```xml 
When using the Factory method design pattern, we gain an extra layer of encapsulation so that our program can grow in a controlled environment. With the Factory method, we delegate the creation of families of objects to a different package or object to abstract us from the knowledge of the pool of possible objects we could use. Imagine that you want to
organize your holidays using a trip agency. You don't deal with hotels and traveling and you just tell the agency the destination you are interested in so that they provide you with everything you need. The trip agency represents a Factory of trips.
```

### Objectives
- Delegating the creation of new instances of structures to a different part of the
program
- Working at the interface level instead of with concrete implementations
- Grouping families of objects to obtain a family object creator

### About the example
we are going to implement a payments method Factory, which is going to
provide us with different ways of paying at a shop.

### Acceptance criteria
- To have a common method for every payment method called Pay
- To be able to delegate the creation of payments methods to the Factory
- To be able to add more payment methods to the library by just adding it to the
factory method