# Go Image File I/O Handler

This repository is a mini-project focused on handling file I/O operations for image files using the Go standard library. The project provides basic functionality to upload multiple images and serve images from local storage. The goal is to improve skills in managing file I/O in Go.

## Features
- Upload multiple image files.
- Uses only Go's standard library and gorilla/mux.
- Stores images locally in the repository directory.

## Requirements
- Go (latest stable version)

## Installation & Setup
```sh
# Clone the repository
git clone https://github.com/duckcoding00/multiple-file.git
cd multiple-file/cmd

# Run the application
go run .
```

## API Endpoints
### 1. Upload Images
- **Endpoint:** `POST /upload`
- **Description:** Uploads multiple image files and saves them locally.
- **Request:** Multipart form-data with image files.
- **Response:**
  ```json
  {
    "message": "CREATED",
    "data": "filepaths"
  }
  ```


## Project Purpose
This project is created to practice and improve skills in handling file I/O operations using Go's standard library. It focuses on working with local storage without relying on external packages.

