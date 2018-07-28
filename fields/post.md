# Manipulating JSON Fields

On reddit there was an interesting question:

<SNIP>

Before diving into answering it's always useful to dive into the question rather than trying to immediately answer. It's easy to lose track of the end-goal and problem. Creating a better underlying model we often can simplify or eliminate the problem. After a few questions I got that the input JSON looks like this (slightly simplified):

```
{
  "Fields": [
    {
        "Name": "blah blah blah",
        "Type": "float", // could also be uint or int
        "Val": 15 // this will be parsed according to the type
    },
    ... // more fields
  ]
}
```

The goal of the program is to write code that computes with these values according to names.

# Interface

The initial implementation was based on using interfaces. Interfaces is indeed one of the clearest solution for variability in type.

We can see similar implementation in Ivy. However, due to interfaces we end up with a lot of different casting to make this happen. We lose all of the type-safety that our language gives us.

This can be a pretty good solution when we need to mix multiple types and the "math" comes from the end-user and not the programmer. In that scenario we cannot rely on the language type-safety anyways.

# Reflection

However, since the field data manipulation code is written directly in Go, we can do better.
Often the better solution is to get rid of variability and use concrete types rather than working on arbitrary types. We can try implementing something similar to JSON unmarshaling on arbitrary types.


# Composed serialization

However, using reflection is often frowned upon. Rightfully so, because it is often slower and code is hard to follow. We also lose a lot of type-safety.
We can avoid reflection in some cases by writing more noisy and longer code when using the parsing. I call this approach composed serialization.


# Conclusion