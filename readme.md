# Internship Aggregator (Go)

##  Overview

This project is a system that collects internship offers from different websites and displays them in one place.

The goal is to automate internship search and make it easier to explore opportunities.

---

##  Features

* Scrape internship listings from external websites
* Store data in a database
* Provide REST API to access internships
* (Optional) Web UI for browsing

---

##  Tech Stack

* Go (Golang)
* Gin (HTTP framework)
* Colly (web scraping)
* PostgreSQL (database)


---

##  Project Structure

```
/cmd            # application entrypoints
/internal
    /scraper    # scraping logic
    /api        # HTTP handlers
    /models     # data models
    /storage    # database logic
/pkg            # reusable packages
```

## ▶ Running the project

```bash
go run cmd/main.go
```

---

