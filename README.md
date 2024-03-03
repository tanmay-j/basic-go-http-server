# Go Http Server

## Introduction

Recently, I wanted to get into learning Distributed Systems. I realised that the majority of applications and courses uses Go. I not being familiar with Go tried to YOLO my way through the course but soon realise Go is a little different than the programming languages I know. So I tried to follow an advice that I once heard someone say  "If you want to learn a new programming language, Write a basic HTTP server in that language." So that it what I set out to do. I soon realised that I can't jump into this with no direction. So I asked chatGPT to write me an assignment to write a basic http server in go that just serves files. 

If you want to have a look at the assignment, you can find it in the same directory under the name assignment.md.

## "Actual" README (One of the requirements to get marks xD)

### How to run the server
- navigate to the root directory of this project and run `go run .\main.go`

### Test scenarios
- I have a couple of users set in users.txt.
- I have seeded the public folder with `index.html` and `hehe.go`
- Feel free to add/remove users and files while testing the functionality.

### Test scenarios I have tried
- Wrong username/password
- No or different type of auth
- Non admin user trying to access .go files.

## Note to everyone
Hi, if you went through the repo, thanks! Also I feel, I might be missing a few security aspects while doing this project, feel free to point them out. A  scenario I can think of are can through this API the can the user access files outside public directory? 