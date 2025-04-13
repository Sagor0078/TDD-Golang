# Why use context?

    1.To cancel operations when a client disconnects or request times out.

    2.To pass request-scoped values (user ID, auth tokens, etc.).

    3.To avoid goroutine leaks in long-running processes.