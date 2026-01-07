# untangle

A proof-of-concept Go static analysis tool that tracks field access patterns for structs. Mark structs with `//untangle:track` to monitor which fields are accessed throughout your codebase, and use `//untangle:print` on functions to report their access patterns.

The analyzer builds a call graph, detects recursive function cycles using Tarjan's algorithm, and performs topological sorting to ensure dependencies are analyzed in the correct order.
