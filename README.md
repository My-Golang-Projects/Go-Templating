# Go-Templating

## Table of Contents
1. [Introduction](#introduction)
2. [Examples of Template Actions](#examples-of-template-actions)
3. [Template Functions](#template-functions)

## Introduction

- A template is a file that contains static text and/or placeholders for dynamic content.
- It contains actions, which are instructions for the template engine that tells it how to walk through the data that was passed in and what to include in the output.
- Actions are enclosed in double curly braces {{ }} and they access the data via the *cursor*, denoted by a dot (`.`)
  - The cursor is not a pointer to the data, it's just a way to access the data.

## Examples of Template Actions

1. Variable Printing:
   ```go
   {{ .Name }}
   ```

2. Conditional Statements:
   ```go
   {{ if .IsLoggedIn }}
     Welcome, {{ .Username }}!
   {{ else }}
     Please log in.
   {{ end }}
   ```

3. Looping:
   ```go
   {{ range .Items }}
     <li>{{ .Title }}</li>
   {{ end }}
   ```

   This example demonstrates how to loop through a slice or array of items:
   - The `range` keyword iterates over each element in `.Items`.
   - For each iteration, the cursor (`.`) is set to the current item.
   - Inside the loop, we can access properties of the current item, like `.Title`.

   Example:
   ```go
   data := struct {
       Items []struct {
           Title string
       }
   }{
       Items: []struct{ Title string }{
           {Title: "First Item"},
           {Title: "Second Item"},
           {Title: "Third Item"},
       },
   }
   ```
   Output:
   ```html
   <li>First Item</li>
   <li>Second Item</li>
   <li>Third Item</li>
   ```

4. Function Calls:
   ```go
   {{ len .Array }}
   ```

   This example shows how to call built-in or custom functions in templates:
   - `len` is a built-in function that returns the length of arrays, slices, maps, or strings.
   - `.Array` is the argument passed to the `len` function.

   Example:
   ```go
   data := struct {
       Array []int
   }{
       Array: []int{1, 2, 3, 4, 5},
   }
   ```
   Output: `5`

5. Pipelines:
   ```go
   {{ .Name | printf "Hello, %s!" }}
   ```

   Pipelines allow you to chain operations, similar to Unix pipes:
   - The output of one operation becomes the last argument to the next operation.
   - In this case, `.Name` is passed as the last argument to `printf`.

   Example:
   ```go
   data := struct {
       Name string
   }{
       Name: "Alice",
   }
   ```
   Output: `Hello, Alice!`

   Pipelines can be more complex, chaining multiple operations:
   ```go
   {{ .Name | lower | printf "hello, %s!" | upper }}
   ```
   This would:
   1. Get the value of `.Name` (e.g., "Alice")
   2. Convert it to lowercase ("alice")
   3. Format it with `printf` ("hello, alice!")
   4. Convert the result to uppercase

   Final output: `HELLO, ALICE!`

6. With Action:
   ```go
   {{ with .User }}
     Name: {{ .Name }}
     Email: {{ .Email }}
   {{ end }}
   ```

   The `with` action allows you to change the scope of the cursor:
   - It checks if the value after `with` is non-empty (not nil, zero, or an empty collection).
   - If true, it executes the block with the cursor (`.`) set to that value.
   - If false, it skips the block entirely.

   Example:
   ```go
   data := struct {
       User struct {
           Name  string
           Email string
       }
   }{
       User: struct {
           Name  string
           Email string
       }{
           Name:  "John Doe",
           Email: "john@example.com",
       },
   }
   // `*data.User.Name*`
   ```
   Output:
   ```
   Name: John Doe
   Email: john@example.com
   ```

7. Template Definition and Inclusion:
   ```go
   {{ define "user-info" }}
     <p>Name: {{ .Name }}</p>
     <p>Email: {{ .Email }}</p>
   {{ end }}

   {{ template "user-info" .User }}
   ```

   This example shows how to define and include named templates:
   - The `define` action creates a named template.
   - The `template` action includes the named template, passing it the specified data.

8. Block Action:
   ```go
   {{ block "content" . }}
     <p>Default content</p>
   {{ end }}
   ```

   The `block` action defines a template and immediately uses it:
   - It's equivalent to defining a template and then including it.
   - Useful for providing default content that can be overridden.

9. Variable Assignment:
   ```go
   {{ $var := .SomeValue }}
   {{ $var }}
   ```

   You can create and use variables within templates:
   - Variables are prefixed with `$`.
   - They can store values for later use in the template.

10. Comparison Operators:
    ```go
    {{ if eq .Value 42 }}
      The answer is 42!
    {{ end }}
    ```

    Go templates support various comparison operators:
    - `eq`: Equal
    - `ne`: Not equal
    - `lt`: Less than
    - `le`: Less than or equal
    - `gt`: Greater than
    - `ge`: Greater than or equal

11. Logical Operators:
    ```go
    {{ if and .IsAdmin (not .IsDeleted) }}
      <p>Admin content</p>
    {{ end }}
    ```

    Logical operators allow for complex conditions:
    - `and`: Logical AND
    - `or`: Logical OR
    - `not`: Logical NOT

## Template Functions

Go templates provide several built-in functions that can be used within template actions. Here are some commonly used functions:

1. Comparison Functions:
   - `eq`: Equal
   - `ne`: Not equal
   - `lt`: Less than
   - `le`: Less than or equal
   - `gt`: Greater than
   - `ge`: Greater than or equal

   Example:
   ```go
   {{ if eq .Value 42 }}The answer is 42!{{ end }}
   {{ if lt .Age 18 }}You are under 18.{{ end }}
   ```

2. Logical Functions:
   - `and`: Returns the boolean AND of its arguments
   - `or`: Returns the boolean OR of its arguments
   - `not`: Returns the boolean NOT of its argument

   Example:
   ```go
   {{ if and .IsAdmin (not .IsDeleted) }}Admin content{{ end }}
   ```

3. String Functions:
   - `printf`: Formats a string using the specified format
   - `println`: Formats using the default format for its operands and writes to standard output
   - `html`: Returns a string with HTML special characters escaped
   - `js`: Returns a string with JavaScript special characters escaped
   - `urlquery`: Returns a string with URL query special characters escaped

   Example:
   ```go
   {{ printf "Hello, %s!" .Name }}
   {{ .Comment | html }}
   ```

4. Numeric Functions:
   - `add`: Returns the sum of its arguments
   - `sub`: Returns the difference between its arguments
   - `mul`: Returns the product of its arguments
   - `div`: Returns the quotient of its arguments

   Example:
   ```go
   {{ add .X .Y }}
   {{ div .Total .Count | printf "%.2f" }}
   ```

5. Slice/Array Functions:
   - `len`: Returns the length of a string, slice, array, or map
   - `index`: Returns the result of indexing its first argument by the following arguments

   Example:
   ```go
   {{ len .Items }}
   {{ index .Array 0 }}
   ```

6. Type Conversion Functions:
   - `int`: Converts its argument to an int
   - `float`: Converts its argument to a float64

   Example:
   ```go
   {{ add (int .X) (int .Y) }}
   ```

These functions can be used within template actions to perform various operations on your data, allowing for more complex logic and formatting in your templates.
