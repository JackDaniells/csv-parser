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
> If you want to change folders, you can do so by modifying `INPUT_FOLDER` and `OUTPUT_FOLDER` variables in `config/config.go` file.

# Folder architecture

```
- config
- - config.go     // project and table formatting configs
- input           // files to be processed
- output          // file processing result
- src             // program implementation
- - iostrategy    // file reading and writing strategies implementation
- - - implementations
- - - - csv       // csv implementation
- - parser        // service responsible for standardizing tables
- - reader        // service responsible for read tables
- - validator     // service responsible for validating column rules
- - writer        // service responsible for write tables
- main.go
```

The read file format configuration, mandatory columns, possible column names, match selector and column grouper
are in the `config/config.go` file too, allowing the adjustment of these parameters in a centralized and simple way.