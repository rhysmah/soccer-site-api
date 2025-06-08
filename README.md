# Soccer Site - API Backend

## Structure

### HTTP Server
The HTTP server acts as a message send-and-receive system.

This is what exists within `http.ListenAndServe()`. This receives and parses HTTP requests and sends responses.

### Application Server Struct
This Server struct acts as a kind of info management system.

Within `./internal/server`, there is a Server struct. This is designed to capture information about the HTTP Server, such as the version, the server health, and so on.

The Server struct isn't technically required, but without it, we'd have global variables -- variables that other parts of the application could conceivably access and change. This is poor design.

As per _Learning Go: An Idiomatic Approach to Real-World Go Programming_:

> ...package-level state should be effectively immutable. Anytime your logic depends on values that are configured at startup or changed while your program is running, those values should be stored in a struct, and that logic should be implemented as a method. (Bodner, 150)

Thus, I have created the Server struct, with methods for accessing `startTime` and `version`. This data is then accessibly via the appropriate methods. 

### Health Check API
An endpoint that checks your API and alerts you when something's wrong. It's essentially a diagnostic tool.

#### Error-Handling the Health Check


##### Resources
- https://testfully.io/blog/api-health-check-monitoring/