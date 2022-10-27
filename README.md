# Illegal Char Demo

Demo for issues related to running `go_test` targets from within IJ using the Bazel plugin.

### Version Info:

Goland: 2022.2.2 Build #GO-222.3739.57, built on August 16, 2022
Bazel Version: 5.3.2 (via Bazelisk v1.11.0)
Bazel Plugin: com.google.idea.bazel.ijwb (2022.10.04.0.1-api-version-222)


### Examples:

#### Running from the command line

Everything works as expected.

* [Without Failing Test](screenshots/terminal-passing.png)
* [With Failing Test](screenshots/terminal-failing.png)

#### Running from Goland

* [Without Failing Test](screenshots/run-ij-passing.png) - the tests pass, but on the left side, 
  the message "No tests were found" appears, even with "show passing" selected.
* [With Failing Test](screenshots/run-ij-failing.png) - the tests fail as expected, but on the 
  left side, only the failing test shows up, even with "show passing" selected.

#### Debugging from Goland
* [Debugging](screenshots/debug-ij.png) - the tests don't even run, as an illegal character
is found in the generated executable.



