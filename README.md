# groupie-tracker

Groupie Trackers
Overview

Groupie Trackers is a web application designed to receive, manipulate, and display data from a provided API related to music artists, their concert locations, and dates. This project focuses on creating an engaging user experience through various data visualizations while ensuring robust client-server communication.
Features

    Artist Information: Display detailed profiles of artists, including names, images, years of activity, first album release dates, and member details.
    Concert Locations: List of past and upcoming concert venues for each artist.
    Concert Dates: A timeline of past and future concert dates.
    Data Relations: Use a relational model to connect artists, concert dates, and locations seamlessly.
    User-Friendly Visualizations: Implement a variety of visual representations (cards, tables, lists) to present information clearly and attractively.
    Event Handling: Trigger client-server interactions to fetch and display data dynamically based on user actions.

API Structure

The provided API consists of four main parts:

    Artists: Contains information about bands and artists.
    Fields: name, image, year_active, first_album_date, members.

    Locations: Lists the venues for past and upcoming concerts.

    Dates: Contains dates of past and upcoming concerts.

    Relations: Links artists with their corresponding concert dates and locations.

## Getting Started
### Prerequisites

    Go installed on your machine.
    Familiarity with RESTful APIs and JSON data handling.
    Basic understanding of HTML/CSS for front-end development.

## Installation

    Clone the Repository:

```bash

git clone <repository-url>
cd groupie-trackers
```
Install Dependencies: (Only standard Go packages are allowed, so this step may be omitted)

Run the Go Server:
``` bash

    go run main.go
```
    Access the Web Application: Open your browser and navigate to http://localhost:8080.

## Usage

    Explore Artists: View detailed profiles of various artists.
    Concert Information: Check concert locations and dates.
    Dynamic Interactions: Engage with the site to trigger API calls and receive updated information.

## Testing

Run tests to ensure the integrity of your application. Make sure to create test files for unit testing the functionality of your backend.

```bash

go test ./...
```

## Contributing

If you would like to contribute to the project, please fork the repository and submit a pull request with your changes. Ensure your code adheres to the best practices outlined above.
License

This project is licensed under the MIT License. Feel free to use and modify the code as you see fit.

## Authors
* Wambita Fana
* Franklyne Namayi
* Tomlee Abila