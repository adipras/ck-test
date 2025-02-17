# CK-Test Project

This project provides two main functionalities:
1. Group Cities by Country
2. Root Word Replacement

## Table of Contents
- [Features](#features)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Program Details](#program-details)
  - [1. Group Cities by Country](#1-group-cities-by-country)
  - [2. Root Word Replacement](#2-root-word-replacement)

## Features
- Interactive menu-driven interface
- Multiple program selection
- CSV file processing
- Text manipulation
- JSON output generation
- HTML results page with download options

## Project Structure
project/
├── main.go
├── menu/
│ └── menu.go
├── programs/
│ ├── cities_processor.go
│ └── root_words.go
└── utils/
└── fileutils.go

## Prerequisites
- Go 1.16 or higher
- CSV file containing city data (for Program 1)


## Setup

1. Clone the repository to your local machine:

    ```sh
    git clone <repository-url>
    cd ck-test
    ```

2. Ensure you have the [cities.csv](https://raw.githubusercontent.com/dr5hn/countries-states-cities-database/refs/heads/master/csv/cities.csv) file in the root directory of the project.

3. Initialize the Go module:

    ```sh
    go mod tidy
    ```

## Running the Project

To run the project, execute the following command in the root directory:

```sh
go run main.go
```

## Usage
After running the project, you will see a menu with the following options:

```sh
Available Programs:
1. Group Cities by Country
2. Root Word Replacement
0. Exit

Enter program number to run (0 to exit):
```