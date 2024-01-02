# Parking Lot Application

## Overview

 This application manages the rental and return of cars in rental company. it allows user to:
   - Create a new car, list all cars, get a car by id, update a car, delete a car
   - View car information, rent a car, return a car
   - Manage rental records, including start date, end date, kilometers driven, and associated car details

## Technologies

- **Backend:** Go
- **Database:** MySQL

## Setup Guide

### Prerequisites

- Go version 1.16 or higher
- MySQL version 8.0 or higher

### Installation

1. Clone the repository

```bash
    git clone https://github.com/natnael-meresa/parking-lot
```

2. Create a database named `parking_lot`

```bash
    mysql -u root -p
    CREATE DATABASE parking_lot;
```
 
 - you can follow the steps in the link if you are using windows
   https://www.w3schools.com/mysql/mysql_install_windows.asp#:~:text=The%20simplest%20and%20recommended%20method,%2Dinstaller%2Dcommunity%2D8.0.
 - you can follow the steps in the link if you are using mac
    https://www.geeksforgeeks.org/how-to-install-mysql-on-macos/
 - you can follow the steps in the link if you are using linux
    https://www.digitalocean.com/community/tutorials/how-to-install-mysql-on-ubuntu-22-04

3. Update the ./config/config.yaml file with your database information

4. Install the dependencies

```bash
    go mod download
```
5. Run the application

```bash
    go run main.go
```

## API Documentation

### Create a new car

| Method | Endpoint | Description | Sample Request | Sample Response |
|---|---|---|---|---|
| POST | /cars | Create a new car | { "registration": "[Car registration]", "model": "[Car model]", "mileage": [Car mileage] } | { "id": [Car ID], "registration": "[Car registration]", "model": "[Car model]", "mileage": [Car mileage] } |
| GET | /cars | List all cars | | [ { "id": [Car ID], "registration": "[Car registration]", ... }, ... ] |
| POST | /cars/:registration/rent | Rent a car | ... | .. |
| PUT | /cars/:registration/return | Return a car | {"kilometersDriven": [Kilometers driven]} | ... |