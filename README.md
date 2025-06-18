# README.md

# News Fluss

News Fluss is a RESTful API built with Go and the Gin framework. This project provides endpoints to manage news articles, sources, and top headlines, allowing users to create, search, and retrieve news data.

## Project Structure

```
news-fluss
├── main.go                   # Entry point of the application
├── go.mod                    # Module dependencies
├── go.sum                    # Module dependency checksums
├── README.md                 # Project documentation
├── controllers               # HTTP request handlers
│   ├── news_controller.go
│   ├── sources_controller.go
│   └── headlines_controller.go
├── routes                    # API route definitions
│   └── news_routes.go
├── models                    # Data models for search parameters
│   ├── news.go               # News search parameters (POST /news)
│   └── source.go             # Source search parameters (POST /sources)
└── config                    # Configuration settings
    └── config.go
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd news-fluss
   ```

2. **Initialize Go modules (if not already initialized):**
   ```
   go mod init
   ```

3. **Install dependencies:**
   ```
   go get -u github.com/gin-gonic/gin
   go mod tidy
   ```

4. **Run the application:**
   ```
   go run main.go
   ```

## API Endpoints

All available endpoints are defined in [`routes/news_routes.go`](routes/news_routes.go):

### News Endpoints

- `GET /news`  
  Retrieve a list of news articles.

- `POST /news`  
  Search for news articles.  
  The request body should match the structure in `models/news.go` (see **Models** below).

### Top Headlines Endpoints

- `GET /topHeadlines`  
  Retrieve top headlines.

- `GET /topHeadlines/:category`  
  Retrieve top headlines by category.

### Sources Endpoints

- `GET /sources`  
  Retrieve a list of news sources.

- `GET /sources/:category`  
  Retrieve news sources by category.

- `POST /sources`  
  Search for news sources.  
  The request body should match the structure in `models/source.go` (see **Models** below).

## Models

Data models for search parameters are defined in the `models` folder:

- `models/news.go` – News search parameters for `POST /news`
- `models/source.go` – Source search parameters for `POST /sources`

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License.