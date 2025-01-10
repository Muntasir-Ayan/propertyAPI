# Property API

Property API is a RESTful API built with Beego and PostgreSQL, designed to manage property listings and their details.

## Table of Contents

- [Project Description](#project-description)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Database Schema](#database-schema)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Project Description

The Property API allows you to manage property listings, including details such as hotel name, rating, review count, type, and more. The API is built using the Beego framework and PostgreSQL for data storage.

## Prerequisites

Before you begin, ensure you have met the following requirements:
- https://github.com/Muntasir-Ayan/hotel-renta **This project should be in running**
- Docker and Docker Compose installed on your machine
- Go programming language installed (if you plan to run the application without Docker)
- PostgreSQL database

## Installation

To install and run the Property API, follow these steps:

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Muntasir-Ayan/propertyAPI.git
   cd propertyAPI
   ```
2. **Run**:
     ```bash
     bee run
     ```
## Api Endpoints:
- http://localhost:8080/v1/property/list
- http://localhost:8080/v1/property/details
