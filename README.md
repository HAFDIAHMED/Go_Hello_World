# Go_Hello_World


# Get started with Go (on Windows 10)

<i>Install Go (if you haven't already).</i>

<i>Write some simple "Hello, world" code.</i><br>

<i>Use the go command to run your code.</i><br>

<i>Use the Go package discovery tool to find packages you can use in your own code.</i><br>

<i>Call functions of an external module.</i><br>

# Install Go

Just use the [Download and install](https://golang.org/doc/install) steps.

# Write some code 

Get started with Hello, World.

Open a command prompt and cd to your home directory.

### `cd %HOMEPATH%`

Create a hello directory for your first Go source code.<br>
For example, use the following commands:

###  `mkdir hello` <br>`cd hello`

#Enable dependency tracking for your code.
When your code imports packages contained in other modules, you manage those dependencies through your code's own module. That module is defined by a go.mod file that tracks the modules that provide those packages. That go.mod file stays with your code, including in your source code repository.<br>

To enable dependency tracking for your code by creating a go.mod file, run the go mod init command, giving it the name of the module your code will be in. The name is the module's module path. In most cases, this will be the repository location where your source code will be kept, such as github.com/mymodule. If you plan to publish your module for others to use, the module path must be a location from which Go tools can download your module.<br>

For the purposes of this tutorial, just use example.com/hello.




