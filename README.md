# groupie-tracker-Search

## Overview

**Groupie Tracker Search** is a web application designed to receive, manipulate, and display data from a provided API related to music artists, their concert locations, and dates. This project focuses on creating an engaging user experience through various data visualizations while ensuring robust client-server communication.

## Features

1. **Artist Information:** Display detailed profiles of artists, including names, images, years of activity, first album release dates, and member details.

2. **Concert Locations:** List of past and upcoming concert venues for each artist.
    Concert Dates: A timeline of past and future concert dates.
3. **Data Relations:** Use a relational model to connect artists, concert dates, and locations seamlessly.
4.  **User-Friendly Visualizations:** Implement a variety of visual representations (cards, tables, lists) to present information clearly and attractively.
5. **Event Handling:** Trigger client-server interactions to fetch and display data dynamically based on user actions like clicking  and searching.
6. **Search Functionality:** Implement search feature that allows the user to search for an  artist,  a member , dates, years and locations

## API Structure

The provided API consists of four main parts:

* **Artists:** Contains information about bands and artists.

* **Locations:** Lists the venues for past and upcoming concerts.

* **Dates:** Contains dates of past and upcoming concerts.

* **Relations:** Links artists with their corresponding concert dates and locations.

## Getting Started

### Prerequisites

[Go](https://go.dev/doc/install) installed on your machine.
    Familiarity with [RESTful APIs](https://aws.amazon.com/what-is/restful-api/) and [JSON](https://developer.mozilla.org/en-US/docs/Learn/JavaScript/Objects/JSON) data handling.
Basic understanding of HTML/CSS/JAVASCRIPT for front-end development.

## Installation

* Clone the Repository:
```bash

git clone https://learn.zone01kisumu.ke/git/tabila/groupie-tracker-search-bar

cd groupie-tracker
```
Install Dependencies: (Only standard [Go](https://go.dev/doc/install) packages are allowed, so this step may be omitted)

Run the Go Server:
``` bash
go run .
```
Access the Web Application: Open your browser and navigate to http://localhost:8080.

To use another port
```sh
export PORT=3000
go run .
 ```
 _this will set the port to a diffrent port that is **3000**_

## Testing

Run tests to ensure the integrity of your application.
```bash
go test -cover ./...
```

## Contributing

If you would like to contribute to the project, please fork the repository and submit a pull request with your changes. Ensure your code adheres to the best practices outlined above.

## License

This project is licensed under the [MIT License](LICENSE). 

## Authors
* [Wambita Fana]()
* [Franklyne Namayi]()
* [Tomlee Abila]()