# Go CDN API

## About

This is a basic CDN written in Golang. Currently it only supports image uploads. This is my first Go project so don't shid on me.

## Dev Installation Instructions

- Copy the `.env.example` and rename it to `.env`
- Run the program: `go run .`
- Program will be running on port `3001` by default (corresponding to .env)

## Usage

### `/api/ping`

**Response**

```
pong
```

### `/api/upload`

\
**Body** \
Form data with image named "file".

**Response**

```json
{
  "file_url": "https://example.com/download/uuid.png"
}
```

### `/api/download/{file_name}`

**Response**
\
`image/~`
