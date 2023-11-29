### CDN server

## Features

- **File Upload:** Upload files with a maximum size of 10 GB.
- **File Retrieval:** Retrieve uploaded files using their filenames.
- **File Deletion:** Delete files with filename.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) installed on your machine.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/your-repository.git
   ```

2. Change into the project directory:

   ```bash
   cd your-repository
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

   The server will start on [http://localhost:5051](http://localhost:5051).

## Usage

### Uploading Files

- To upload a file, make a POST request to the root endpoint (`/`) with the file in the request body.

  ```bash
  curl -X POST -F "file=@/path/to/your/file" http://localhost:5051/
  ```

  Note: Replace `/path/to/your/file` with the actual path to the file you want to upload.

- The response will include a message indicating whether the upload was successful and the filename.

### Retrieving Files

- To retrieve an uploaded file, make a GET request to the endpoint `/:filename`, where `:filename` is the name of the file you want to retrieve.

  ```bash
  curl http://localhost:5051/your_filename
  ```

  Note: Replace `your_filename` with the actual filename you want to retrieve.

- The file will be served to the client.

### Delete File

- To delete an uploaded file, make a DELETE request to the endpoint `/:filename`, where `:filename` is the name of the file you want to delete.

  ```bash
  curl -X DELETE http://localhost:5051/your_filename
  ```

  Note: Replace `your_filename` with the actual filename you want to delete.

- The file will be deleted

### Malicious File Check

- The service checks for malicious file types and prevents the upload of executable files with certain extensions. If a file with a restricted extension is detected, the service will respond with an error message.

### Start Server with Docker compose

```
docker compose up
```

### Run Dockerfile

```
docker build -t cdn-server .
docker run -d -p 5051:5051 cdn-server
```

### Run from docker image

```
docker run -d -v /path/to/your/bucket:/app/bucket -p 5051:5051 --restart always 01822531439/cdn
```
