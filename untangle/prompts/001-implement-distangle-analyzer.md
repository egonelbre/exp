<objective>
Create a Go static analysis tool using the golang.org/x/tools/go/analysis package that tracks field access patterns across structs marked with special comments. This analyzer will help identify which parts of tracked objects are actually used by different functions and methods, enabling better understanding of data flow and potential optimization opportunities.
</objective>

<context>
The distangle analyzer operates on Go codebases to track how specific structs (marked with `//distangle:track`) have their fields accessed throughout the codebase. This is particularly useful for:
- Understanding which fields of large structs are actually used by different code paths
- Identifying opportunities to split or refactor structs
- Documenting data flow through complex systems

The analyzer must work across package boundaries, tracking how tracked structs flow through function calls.

Review the test case to understand expected behavior:
@testdata/kitchen/sink/sink.go - Contains the tracked `World` struct and functions that access it
@testdata/kitchen/ext/bowl.go - Contains cross-package function accessing tracked struct
@testdata/kitchen/server/server.go - Contains methods marked to print access patterns
</context>

<requirements>
The analyzer must implement three distinct stages:

**Stage 1: Discover Tracked Structs**
- Find all struct type declarations preceded by `//distangle:track` comment
- Create object facts for each tracked struct that can be used in later passes
- Store metadata about the struct (name, fields, package location)

**Stage 2: Analyze Field Accesses**
- For every function and method in the codebase, analyze which fields of tracked structs are:
  - Read (accessed on right-hand side of expressions)
  - Written (assigned to)
- Track both direct field accesses (e.g., `world.Knives`) and indirect accesses through method calls (e.g., `world.GetIngredients()`)
- For function calls, read the callee's object fact and propagate its field accesses
- Respect `//distangle:ignore` comments - treat marked functions/methods as having zero field accesses
- Create object facts for each function/method storing its tracked field accesses
- Handle cross-package scenarios correctly using the analysis framework's fact system

**Stage 3: Report Access Patterns**
- Find all functions and methods marked with `//distangle:print` comment
- For each marked function/method, retrieve its object fact
- Print a report showing which tracked struct fields were accessed
- Format should be clear and useful for understanding data flow

The analyzer must:
- Use the analysis.Analyzer interface correctly with proper configuration
- Implement object facts using types.Object as keys for cross-package tracking
- Handle the multi-pass nature of analysis (facts from dependencies must be available)
- Work correctly with method receivers and function parameters
- Track transitive field accesses through function call chains
</requirements>

<implementation>
Create the analyzer implementation in `./distangle.go`.

**Architecture Guidance:**

1. **Analyzer Setup**: Define the analysis.Analyzer with:
   - Name: "distangle"
   - Doc: Clear description of purpose
   - Run: Main analysis function
   - Requires: Any required analyses (likely needs SSA or type info)
   - FactTypes: Register your custom fact types for tracked structs and function access patterns

2. **Object Fact Design**: 
   - Create a fact type for tracked structs (e.g., `TrackedStructFact`) containing field information
   - Create a fact type for function/method access patterns (e.g., `AccessPatternFact`) containing sets of accessed field names
   - Ensure facts are serializable (use basic types like maps, slices, strings)
   - Use types.Object as fact keys - for structs use the TypeName object, for functions use the Func object

3. **Stage 1 Implementation**:
   - Iterate through package syntax trees looking for type declarations
   - Check comment groups for `//distangle:track` directive
   - Export facts for tracked structs so later passes can identify them
   - Store field names and types for later validation

4. **Stage 2 Implementation**:
   - For each function/method, perform AST traversal to find field accesses
   - Use ast.Inspect or ast.Walk to traverse function bodies
   - Identify SelectorExpr nodes where the base expression has a tracked struct type
   - Distinguish reads from writes (check parent node context)
   - For method calls on tracked structs, look up the method's access pattern fact
   - For function calls passing tracked structs, look up the function's access pattern fact
   - Aggregate all direct and transitive accesses
   - Check for `//distangle:ignore` comment before creating access pattern facts
   - Export access pattern facts for each function/method

5. **Stage 3 Implementation**:
   - Find functions/methods with `//distangle:print` comment
   - Import their access pattern facts
   - Use pass.Reportf or pass.Report to output findings
   - Format output to clearly show which fields were accessed

**Critical Implementation Details:**

- Use `pass.TypesInfo.Uses` and `pass.TypesInfo.Defs` to resolve identifiers to types.Object
- Check if a type is a tracked struct by importing facts for its TypeName object
- Handle pointers correctly - `*World` and `World` both need to be recognized
- For method receivers, extract the base struct type even if it's a pointer receiver
- Consider both `obj.Field` direct accesses and method call propagation
- The `//distangle:ignore` comment should be checked on the function/method declaration's doc comment
- Ensure proper fact import/export - facts are only available across packages if properly exported

**Why these constraints matter:**
- Object facts are required (not package facts) because we need to track individual structs and functions across packages
- Multi-pass execution is inherent to go/analysis - dependencies are analyzed first, their facts become available to dependents
- Accurate type resolution via TypesInfo is critical - AST alone doesn't provide type information for determining if something is a tracked struct
</implementation>

<testing>
Create test infrastructure in `./distangle_test.go`:

1. Use `analysistest.Run` to test against the kitchen testdata
2. Expected behavior for the test package:
   - `sink.World` should be identified as a tracked struct
   - `sink.API()` should have no recorded accesses (marked with `//distangle:ignore`)
   - `sink.Chop()` should show accesses to: Knives, Ingredients
   - `sink.Mix()` should show accesses to: Bowls, Spoons, Ingredients
   - `ext.CleanBowl()` should show accesses to: Bowls
   - Server methods should print their transitive access patterns based on what they call

3. Verify that cross-package tracking works correctly
4. Test that method calls propagate field accesses
5. Test that ignore directives work as expected
</testing>

<output>
Create these files:
- `./distangle.go` - Main analyzer implementation with all three stages
- `./distangle_test.go` - Test suite using analysistest package
- `./go.mod` - Update if needed to include golang.org/x/tools dependency
- `./cmd/distangle/main.go` - Optional: standalone command-line tool using singlechecker

The analyzer should be importable as a package and usable with standard analysis drivers.
</output>

<verification>
Before declaring complete, verify:

1. Run `go test ./...` and ensure all tests pass
2. Run the analyzer on the kitchen testdata and manually verify output
3. Check that `//distangle:print` annotated functions produce readable reports
4. Verify cross-package fact propagation works (ext.CleanBowl sees sink.World as tracked)
5. Confirm that `//distangle:ignore` prevents fact creation
6. Test with `go vet -vettool=$(which distangle)` if you create the standalone tool
</verification>

<success_criteria>
- Analyzer correctly identifies all `//distangle:track` marked structs
- Field access tracking works for direct field references
- Method calls on tracked structs propagate access information
- Function calls propagate transitive access patterns
- `//distangle:ignore` directive correctly suppresses fact creation
- `//distangle:print` directive triggers readable output reports
- All tests pass
- Cross-package analysis works correctly via object facts
- Code is well-structured following go/analysis best practices
</success_criteria>

<reference>
Key go/analysis package concepts:
- `pass.Report` / `pass.Reportf` - Report findings
- `pass.TypesInfo` - Type information for resolution
- `pass.ExportObjectFact` / `pass.ImportObjectFact` - Cross-package fact sharing
- `types.Object` - Keys for object facts (use TypeName for types, Func for functions)
- `ast.Inspect` - Convenient AST traversal
- Comment map access via `ast.File.Comments` or function/type spec doc comments

Thoroughly analyze the test data to understand the expected behavior patterns. Consider edge cases like:
- Nested field accesses (tracked.Field.SubField)
- Field accesses in different expression contexts (assignments, comparisons, function arguments)
- Method receivers vs parameters
- Pointer vs value types
</reference>
