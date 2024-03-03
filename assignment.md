# Assignment: Basic Secure HTML Server in GO
## Overview
In this assignment, you will write a basic secure HTML server in GO that can serve static web pages and handle HTTPS requests. You will learn how to use the net/http package, and how to implement basic authentication and authorization.

## Specifications
Your server should listen on port 8080 and serve files from the public directory in the same folder as your executable. Your server should support HTTP.

Your server should implement basic authentication using a username and password. The username and password should be stored in a file named users.txt in the same folder as your executable. The file should have one line per user, with the username and password separated by a colon (:). For example:
```
alice:secret
bob:1234
```

Your server should only allow authenticated users to access the files in the public directory. If a user tries to access a file without providing a valid username and password, your server should respond with a 401 Unauthorized status code and a WWW-Authenticate header. For example:
```
HTTP/1.1 401 Unauthorized
WWW-Authenticate: Basic realm="Secure Server"
Content-Type: text/plain
Content-Length: 12

Unauthorized
```
Your server should also check the authorization level of the user based on the file extension. Your server should only allow users with the admin role to access files with the .go extension. The role of the user should be stored in the users.txt file after the password, separated by a comma (,). For example:
```
alice:secret,admin
bob:1234
```
If a user tries to access a file with the .go extension without having the admin role, your server should respond with a 403 Forbidden status code and a plain text message. For example:

```
HTTP/1.1 403 Forbidden
Content-Type: text/plain
Content-Length: 9

Forbidden 
```

Your server should handle any errors or exceptions gracefully and respond with appropriate status codes and messages. For example, if a user requests a file that does not exist, your server should respond with a 404 Not Found status code and a plain text message. For example:

```
HTTP/1.1 404 Not Found
Content-Type: text/plain
Content-Length: 9

Not Found
```
## Grading Policy
Your assignment will be graded based on the following criteria:

- **Functionality**: Your server should meet all the specifications and work correctly with different requests and scenarios. (50%)
- **Code Quality**: Your code should be well-structured, readable, and follow the GO coding conventions. You should use meaningful variable names, proper indentation, and comments where necessary. (25%)
- **Documentation**: You should provide a README file that explains how to run and test your server, and any assumptions or design decisions you made. You should also the users file, and some sample HTML files in the public directory for testing purposes. (25%)

## Submission
You should submit your assignment as a single ZIP file that contains your GO source code file, the README file, the certificate and key files, the users file, and the public directory with the sample HTML files. The ZIP file should be named as yourname_server.zip, where yourname is your name or student ID. You should upload your ZIP file to the course website or email it to the instructor before the deadline. Late submissions will incur a penalty of 10% per day.