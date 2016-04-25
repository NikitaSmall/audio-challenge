# audio-challenge
Simple web-helper for uawebchallenge. Speak with it and see the task done!

### usual heroku note:
Because of free heroku node requests can take near 10 seconds. Be patient, please.

### about messages:
- General advise: speak calm and clearly.
- To fire `pizza` task you need to start (or your message should contain) words like `закажи пиццу`, `купи пиццу`.
 Any other tasks or unrecognisable murmuring will be ignored.
 The recognition engine is pretty good when using general words with good pronounce.
- To set time you need to say `через два часа` or `завтра`.
 If you want to complete the task right now you may omit any time words.
- To set address you need to say `по улице Пушкина, 14` or `проспект Шевченко, 5Б квартира 23`.
 In case of omitting address (or if you tell it not clearly) just empty message will appear.
- To set order list you need to name all that you want, e.g. `пицца маргарита` or `пеперони`.
 Engine is not good with hard to recognise names such as `Madrigaldia` or `Some-unknown-word`.
 You may see the list of supported and recognisable pizzas at `lists/food.txt`. Other foods will be ignored.
- To set the target pizzeria you need to pronounce it calm.
 You may see the list of supported pizzerias at `lists/pizzeria.txt`
- You may provide your name. For this purpose you may say something like `... на имя Евгения Борисовича Конопатько` or
 `для Елены Игнатьевны Брудько`. I hope you have recognisable names, but only first name is used for real.

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
5. Gorilla/session as a simple session interface
6. Yandex speech cloud as a voice recognition tool
7. mgo as a mongoDB adapter
- Bunch of small tools and libs that could be seen inside of source code at `import` statement

## Comment about language and structure:
I used Golang in this project due to my love to this language and according the task:
Go is fast, good at multi-threading (I use go-routines to ask for tasks), simple language.

Every package and every source code file filled with a go-doc comments about package
and about every function or method (even unexported functions have comments due to make your work a little bit easier).

I broke up the functionality into the small packages. Their purposes are obvious but here is a list:
- `router` - a package with handlers and router config. Handlers are kind a simple actions,
while each file (`handlers.go`, `socket_handlers.go`) are kind of controllers.
- `socket` - an util package to provide socket connection. No business logic here, only data-transfer.
- `util` - an util package to make some basic tasks e.g. xml-unmarshaling.
- `config` - an util package to load configs, env variables and so on.
- `task` - a package that provides business logic around tasks.
- `user` - a package that provides basic auth logic.
