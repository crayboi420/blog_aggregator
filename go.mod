module github.com/crayboi420/blog_aggregator

go 1.22.2

replace github.com/crayboi420/blog_aggregator/internal/database => ./internal/database

require (
	github.com/crayboi420/blog_aggregator/internal/database v0.0.0
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
)

require github.com/google/uuid v1.6.0
