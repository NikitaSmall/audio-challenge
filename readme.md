# audio-challenge
Simple web-helper for uawebchallenge. Speak with it and see the task done!

## Setup:
Application written in golang. You need to create a `.env` file in order to compile and run the app.
The paths of inner packages was changed due to stay anonymous member of challenge.
You need to change it from `github.com/challenge/audio-challenge/config` (as example) to `github.com/{insert your nickname}/audio-challenge/config`.
Also usual golang environment is required.

## Tools:
1. Golang as a language
2. gin-gonic/Gin as http-router/micro-framework
3. Recorderjs as audio record tool
4. Gorilla/websocket as a websocket implementation in Go
- Bunch of small tools and libs that could be seen inside of source code at `import` statement

## Comment about language and structure:
I used Golang in this project due to my love to this language and according the task:
Go is fast, good at multi-threading (I use go-routines to ask for tasks), simple language.

I broke up the functionality into the small packages. Their purposes are obvious but here is a list:
- `router` - a package with handlers and router config. Handlers are kind a simple actions,
while each file (`handlers.go`, `socket_handlers.go`) are kind of controllers.
- `socket` - an util package to provide socket connection. No business logic here, only data-transfer.
- `util` - an util package to make some basic tasks e.g. audio-file saving.
- `config` - an util package to load configs, env variables and so on.
- `task` - a package that provides business logic around tasks.
