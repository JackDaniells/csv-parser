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

The structure with the main project folders is specified below:
```
- config
- - config.go     // project and table formatting configs
- input           // files to be processed
- output          // file processing result
- src             // program implementation
- - domain        // application domain classes
- - iostrategy    // IO file strategies implementation
- - - implementations
- - - - csv       // csv implementation
- - parser        // parser service
- - reader        // reader service
- - validator     // validator service
- - writer        // writer service
- main.go
```

All the algorithm's execution logic is based on the four implemented services, and follows the sequence:
* `reader` service reads the specified file and returns an object of type `TableDomain`.
* `parser` service formats the object, applying:
  * standardization rules for naming headers per match
  * selection in case of more than one match per header and
  * column grouping, returning a standardized table object.
* `validator` service takes the standardized object and applies rules for `required` and `unique` fields, and returns two table objects, one with the correct and the other with the faulty data.
* `writer` service receive the tables, one at time, and writes them to the output folder.

The read file format configuration, mandatory columns, possible column names, match selector and column grouper
are defined in the `config.go` file, allowing the adjustment of these parameters in a centralized and simple way.

### `TableColumnSchema`

this object is responsible for mapping the name of the column, if it is required, unique and the list of possible synonymous words to find during column standardization.

```go
type TableColumnSchema struct {
    Name          string
    Unique        bool
    Required      bool
    PossibleWords []string
}
```

### `ColumnMatcher`

this struct is responsible for selecting from a list of combinations for a column names, which column name should be kept.

```go
type ColumnMatcher struct {
    Matches  []string
    Selected string
}
```

### `ColumnGrouper`

this struct is responsible for mapping a set of column names that must be grouped into a new column.

```go
type ColumnGroup struct {
    Headers   []string
    GroupName string
    Separator string
}
```