# go-url-shortner

This is a simple in memory url shortner built with the `net/http` package itself.

**Caution:** the persistance of data is handled by storing in a json file named `urls.json`,
for the first time the file doesn't exist.

**Caution:** since this is a simple project, you may not find production level packages,
like env variables, or localization or advanced logging system packages like `zerolog`
or loggin on the files, etc.

# Index

* [Installation](#installation)
* [Usage](#usage)
    * [Step 1: shorten url](#step-1-shorten-a-url-post-request)
    * [Step 2: redirect from short url](#step-2-redirect-from-a-short-url-get-request)
* [Commands](#commands)
    * [Run application](#run-application)
    * [Run automated tests](#run-tests)
    * [Fix code style](#fix-code-style)
    * [Build the application](#build-the-application)
* [Project structure](#project-structure)

## Installation

**Clone the repository:**

```bash
git clone https://github.com/adnanahmady/go-url-shortner.git
cd go-url-shortner
```

**Run the server:**
```bash
go run main.go // the server will listen on `http://localhost:5000`
```

## Usage

### Step 1: shorten a URL (POST request)

Send a POST request to `/shorten` with a form parameter `url` containing the long URL.

**Example usage `curl`:**

```bash
curl -X POST -d "url=https://www.golang.org/doc/install" http://localhost:5000/shorten

# Output: http://localhost:5000/aBcDef
```

### Step 2: redirect from a short URL (GET request)

Access the generated short URL in your browser or with `curl`.

**Example:**

```bash
# In your browser: http://localhost:5000/aBcDef
# Or using curl:
curl -L http://localhost:5000/aBcDef
```

## Commands

You can use prepared make commands for easier use

### Run application

Run the application using

```bash
make run
```

### Run tests

You can run the applicaiton tests by this command

```bash
make test
```

or alternatively

```bash
make t
```

### Fix code style

You can fix code style by this command

```bash
make lint
```

### Build the application

You can build the app by this command

```bash
make build
```

## Project Structure

Here is a view of the project structure

```bash
.
├── go.mod
├── go.sum
├── internal # the application itself (clean architecture)
│   ├── application # application layer
│   │   ├── create_short_url_use_case.go
│   │   ├── errors.go # application layer error variables
│   │   └── get_short_url_use_case.go
│   ├── domain # domain layer
│   │   ├── errors.go # domain layer errors
│   │   └── repositories.go # repository interfaces
│   ├── infra # infrastructure layer
│   │   └── memory_url_repository.go # url repository memory implementation
│   ├── integration # api tests (these tests evaluate the application behaviour and are not unit tests)
│   │   ├── create_short_code_test.go
│   │   └── redirect_to_url_test.go
│   ├── presentation # presentation layer
│   │   ├── v1_handlers.go
│   │   └── v1_routers.go
│   ├── wire_gen.go
│   └── wire.go # dependency injection manager
├── LICENSE
├── main.go
├── main_test.go
├── Makefile
├── pkg # shared pacakges of the application
│   ├── applog # application logger
│   │   ├── logger.go
│   │   └── logger_test.go
│   ├── reqeust # request utils
│   │   ├── context.go
│   │   ├── log.go # logger middleware
│   │   └── server.go # wrapper arround the net/http package
│   ├── store
│   │   ├── errors.go
│   │   └── memory.go
│   └── test # test helpers
│       ├── assert # test assertions
│       │   └── assertions.go
│       ├── requests.go # testing request utils
│       └── setup.go # test setup
├── README.md
└── urls.json # json storage file (generated and updated by the applciation)

13 directories, 30 files
```