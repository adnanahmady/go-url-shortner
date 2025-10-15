# go-url-shortner

This is a simple in memory url shortner built with the `net/http` package itself,
by restarting the application all data is reset.

**Caution:** since this is a simple project, you may not find production level packages,
like env variables, or localization or advanced logging system packages like `zerolog`
or loggin on the files, etc.

# Index

* [Installation](#installation)

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
