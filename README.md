# Rain CSV Parser

This project was developed as rain interview challenge , and its function is:
 * Read a csv file (one at time)
 * Standardize data following the format (specified in `config` file)  
 * Generate an output with correct and bad data, where each output in a csv file.

# Setup

### Prerequisites

The only requirement for this project is to have [Go 1.17+](https://go.dev/dl/) installed on your machine

### Installation

```shell
go mod download
```

### Run the app

```shell
go run main.go <csv_file_name>
```

> 🚩 **Note**
>
> The app will get files from `input` folder and write in `output` folder by default.
> If you want to change folders, you can do so by modifying `INPUT_PATH` and `OUTPUT_PATH` variables in `config/config.go` file.
