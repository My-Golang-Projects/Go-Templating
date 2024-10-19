# Go-Templating

## Table of Contents
1. [Introduction](#introduction)
2. [Examples of Template Actions](#examples-of-template-actions)
3. [Template Functions](#template-functions)
4. [Whitespace Control in Templates](#whitespace-control-in-templates)
5. [Using Go Functions in Templates](#using-go-functions-in-templates)

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

These functions can be used within template actions to perform various operations on our data, allowing for more complex logic and formatting in our templates.

## Whitespace Control in Templates

Go templates provide a way to control whitespace output around actions using a dash (-). This is particularly useful for formatting the output of our templates.

1. Trimming Whitespace After an Action:
   ```go
   {{ .Value -}}
   ```
   The dash (-) at the end of the action will trim all whitespace (including newlines) immediately following the action.

2. Trimming Whitespace Before an Action:
   ```go
   {{- .Value }}
   ```
   The dash (-) at the beginning of the action will trim all whitespace (including newlines) immediately preceding the action.

3. Trimming Whitespace on Both Sides:
   ```go
   {{- .Value -}}
   ```
   Using dashes on both sides will trim whitespace before and after the action.

Example:
```go
Number of dogs: {{ . | len -}}
{{ range . }}
. . .
We could also have written {{ len . -}}.
```

In this example:
- The first line uses `{{- }}` to remove the newline after printing the number of dogs.
- The last line uses `{{ -}}` to remove the newline after printing the length.

Note: We should be cautious when using whitespace control, especially at the beginning of actions, as it can sometimes lead to unexpected results if not used carefully.

### When to Use Whitespace Control

1. Removing Unnecessary Newlines:
   Particularly useful in loops to prevent extra blank lines.

2. Formatting HTML or Other Structured Output:
   Helps in creating clean, properly indented output without extraneous whitespace.

3. Creating Single-Line Outputs:
   Useful when we want to ensure all output is on a single line, regardless of how the template is formatted.

Remember, readability of our template is important. Use whitespace control judiciously to balance between clean output and maintainable templates.

## Using Go Functions in Templates

Go templates allow you to use both built-in functions and custom functions defined in your Go code. This feature greatly enhances the capabilities of your templates.

### Built-in Functions

Go templates come with several built-in functions. Some common ones include:

1. `len`: Returns the length of a string, slice, map, or array
   ```go
   {{ len .Items }}
   ```

2. `printf`: Formats a string using the specified format
   ```go
   {{ printf "Hello, %s!" .Name }}
   ```

3. `index`: Returns the result of indexing its first argument by the following arguments
   ```go
   {{ index .Array 0 }}
   ```

4. `and`, `or`, `not`: Logical operators
   ```go
   {{ if and .IsAdmin (not .IsDeleted) }}Admin content{{ end }}
   ```

### Custom Functions

You can also define and use custom functions in your templates. Here's how:

1. Define your function in Go:
   ```go
   func add(a, b int) int {
       return a + b
   }
   ```

2. Create a template.FuncMap to hold your custom functions:
   ```go
   funcMap := template.FuncMap{
       "add": add,
   }
   ```

3. Create your template with the custom functions:
   ```go
   tmpl, err := template.New("myTemplate").Funcs(funcMap).Parse(templateString)
   ```

4. Use the custom function in your template:
   ```go
   The sum is: {{ add 5 3 }}
   ```

### Example: Using Custom Functions

Here's a complete example of defining and using a custom function in a template:

```go
package main

import (
    "os"
    "text/template"
)

func multiply(a, b int) int {
    return a * b
}

func main() {
    funcMap := template.FuncMap{
        "multiply": multiply,
    }

    const templateText = "{{.X}} times {{.Y}} is {{multiply .X .Y}}"

    tmpl, err := template.New("calc").Funcs(funcMap).Parse(templateText)
    if err != nil {
        panic(err)
    }

    data := struct {
        X, Y int
    }{3, 4}

    err = tmpl.Execute(os.Stdout, data)
    if err != nil {
        panic(err)
    }
}
```

This will output: "3 times 4 is 12"

### Best Practices

1. Keep template logic simple. If you find yourself needing complex logic, consider moving it to Go code.
2. Use descriptive names for your custom functions.
3. Remember that functions used in templates should return values that can be printed or used in template constructs.
4. Be cautious with functions that modify state, as templates are often used in concurrent environments.

By combining built-in functions, custom functions, and Go's powerful templating syntax, you can create flexible and dynamic templates for various use cases.
