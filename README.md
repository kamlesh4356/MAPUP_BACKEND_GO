


# Go Server for Sorting Arrays

## Objective
Develop a Go server with two endpoints (`/process-single` and `/process-concurrent`) to demonstrate your skills in sequential and concurrent processing. The server should sort arrays provided in the request and return the time taken to execute the sorting in both sequential and concurrent manners.

## Requirements

### 1. Server Setup
- Create a Go server listening on port 8000.
- Implement two endpoints: `/process-single` for sequential processing and `/process-concurrent` for concurrent processing.

### 2. Input Format
- The server should accept a JSON payload with the following structure:
```json
{
  "to_sort": [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
}
```
- Each element in the `to_sort` array is a sub-array that needs to be sorted.

### 3. Task Implementation
- For the `/process-single` endpoint, sort each sub-array sequentially.
- For the `/process-concurrent` endpoint, sort each sub-array concurrently using Go's concurrency features (goroutines, channels).

### 4. Response Format
- Both endpoints should return a JSON response with the following structure:
```json
{
  "sorted_arrays": [[sorted_sub_array1], [sorted_sub_array2], ...],
  "time_ns": "<time_taken_in_nanoseconds>"
}
```

### 5. Performance Measurement
- Measure the time taken to sort all sub-arrays in each endpoint in nanoseconds using Go's time package.

### 6. Dockerization
- Containerize the Go server using Docker.
- Provide a Dockerfile for building the Docker image.
```
