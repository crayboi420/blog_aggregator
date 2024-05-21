# Blog Aggregator

Blog aggregator from [bootdev](https://www.boot.dev/learn/build-blog-aggregator)

## Introduction

This is a go project using postgres databases. The idea is to have users with RSS feeds stored in the database. The feeds themselves are scraped in intervals and all posts from that feed are stored. The posts can then be viewed per user.

## Usage

The code requires setting some environment variables up in the .env file. These are ```PORT``` : the port on localhost to run on and ```CONN```: The connection string to a postgres database.

The setup also requires installing goose and migrating the database to the latest version.

After that, the code can be run by building:

```
go build . && ./blog_aggregator
```

## Contributing

You can help by raising any issues that you find in the program
