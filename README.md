# Random Quote Generator

An API that returns random quotes.

# Building

## Prerequisites

To build this project you'll need Docker installed.

## Commands

Build the image

```
docker build -t random-quote-generator .
```

Run the container

```
docker run -p 3000:3000 -d random-quote-generator
```

View the quote in your browser

```
http://localhost:3000
```
