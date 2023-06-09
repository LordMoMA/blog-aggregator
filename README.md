# Blog Aggregation

![Blog Aggregator](./images/1.png "Blog Aggregator")

## Introduction

It's a web server that allows clients to:

- Add RSS feeds to be collected
- Follow and unfollow RSS feeds that other users have added
- Fetch all of the latest posts from the RSS feeds they follow
- RSS feeds are a way for websites to publish updates to their content.

You can use this project to keep up with your favorite blogs, news sites, podcasts, and more!

## Steps

I published each step of my design process in Medium, and it'll be helpful to follow along the details of this project with the articles if you want to learn to scrape:

1. [Build a Content Aggregator in Go(1-Servers)](https://medium.com/@lordmoma/build-a-content-aggregator-in-go-1-servers-a52388888386)
2. [Build a Content Aggregator in Go(2-PostgreSQL)](https://medium.com/@lordmoma/build-a-content-aggregator-in-go-2-postgresql-68d98b98f2af)
3. [Build a Content Aggregator in Go(3-Create Users)](https://medium.com/@lordmoma/build-a-content-aggregator-in-go-3-create-users-32cf84432fa6)
4. [Build a Content Aggregator in Go(4-API Key)](https://medium.com/@lordmoma/build-a-content-aggregator-in-go-4-api-key-181da1424e3a)
5. [Build a Content Aggregator in Go(5-Create a Feed)](https://medium.com/@lordmoma/build-a-content-aggregator-in-go-5-create-a-feed-1ffeba2aaf93)
6. [Build a Content Aggregator in Go(6-Get all feeds)](https://medium.com/@lordmoma/build-a-content-aggregator-in-go-6-get-all-feeds-17d6a78da83a)
7. [Build a Content Aggregator in Go(7-Feed Follows)](https://medium.com/@lordmoma/build-a-content-aggregator-in-go-7-feed-follows-50693c350cd1)
8. [Build a Content Aggregator in Go(8-Scraper)](https://medium.com/@lordmoma/build-a-content-aggregator-in-go-8-scraper-8a60593b52d0)
9. [Build a Content Aggregator in Go(9-Database POSTs Table)](https://medium.com/@lordmoma/build-a-content-aggregator-in-go-9-database-posts-table-5443ce8289a6)

## Improvements

The codebase needs refactoring and more optimizations, please feel free to contribute.

Some ideas:

- Support pagination of the endpoints that can return many items
- Support different options for sorting and filtering posts using query parameters
- Classify different types of feeds and posts (e.g. blog, podcast, video, etc.)
- Add a CLI client that uses the API to fetch and display posts, maybe it even allows you to read them in your terminal
- Scrape lists of feeds themselves from a third-party site that aggregates feed URLs
- Add support for other types of feeds (e.g. Atom, JSON, etc.)
- Add integration tests that use the API to create, read, update, and delete feeds and posts
- Add bookmarking or "liking" to posts
- Create a simple web UI that uses the backend API
