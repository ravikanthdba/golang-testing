# golang-testing
Unit, Integration and Functional Testing in Golang

# Unit Test
    - White Box: Access to the source code (private functions and private variables)
    - Black Box: Test case is put outside the package, allowed to test only the public functions
    - Google recommends White box testing

# Integration Test
    - More than 1 function
    - There are 2 functions: a and b, and a calls b.
    - If we test on function b, it is a unit test
    - If we test on a, it is the integration test
    - The reason we do this is, if we have two functions a and b, independently, they run fine, but when we combine both of the functions, then it fails.

# Functional Test
    - Application is up and running
    - Given requirement is working fine

    Input -----> Black Box -----> Output

    - Send a POST request to the system and get the output

# Pyramid


               -------- Functional --------
       --------------- Integration ---------------
----------------------  Unit Test   ----------------------



# Test Cases should have:

1) Initialization
2) Execution
3) Validation