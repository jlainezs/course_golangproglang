# Hands-on exercise #73 - wrapper

the `TimedFunction` acts as a wrapper function that measures the
elapsed time taken by `MyFunction` to execute. It captures the start time, calls `MyFunction`,
calculates the elapsed time, and then prints it. By using the wrapper, you can easily add
timing functionality to multiple functions without modifying their implementation.