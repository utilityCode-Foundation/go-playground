# Builder Pattern
```xml 
The Builder pattern helps us construct complex objects without directly instantiating their 
struct, or writing the logic they require. Imagine an object that could have dozens of fields 
that are more complex structs themselves. Now imagine that you have many objects with these 
characteristics, and you could have more. We don't want to write the logic to create all these 
objects in the package that just needs to use the objects.
```

### Objectives
- Abstract complex creations so that object creation is separated from the object
user
- Create an object step by step by filling its fields and creating the embedded
objects

### About the example
we are going to implement a car builder

### Acceptance criteria
- To be able to build a object step by step with needed property