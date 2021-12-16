# Rain CSV Parser

This project was developed as Rain interview challenge, and its function is:
 * Read csv files
 * Standardize data following the format (specified in `config` file)  
 * Generate an output with correct and bad data, where each output in a csv file.

# Setup

### Prerequisites

The only requirement for this project is to have [Go 1.17+](https://go.dev/dl/) installed on your machine

### Installation

```shell
go mod download
```

### Execution

```shell
go run main.go <csv_file_name>
```

> 🚩 **Note**
>
> The app will get files from `input` folder and write in `output` folder by default.
> If you want to change folders, you can do so by modifying `INPUT_FOLDER` and `OUTPUT_FOLDER` variables in `config/config.go` file.

### Tests
```shell
go test ./... -cover
```

# Folder architecture

The structure with the main project folders is specified below:
```
- config
- - config.go     // project and table formatting configs
- input           // files to be processed
- output          // file processing result
- src             // program implementation
- - commons       // useful common methos
- - constants     // application constant list
- - domain        // application domain classes
- - iostrategy    // IO file strategies implementation
- - - implementations
- - - - csv       // csv read and write implementation
- - pkg           // methods that can be outsourced to other projects (like libs) 
- - parser        // parser service
- - reader        // reader service
- - validator     // validator service
- - writer        // writer service
- main.go
```

All the algorithm's execution logic is based on the four implemented services, and follows the sequence:
1. `reader` service reads the specified file and returns an object of type `TableDomain`.
2. `parser` service formats the object, applying:
   * standardization rules for naming headers per match
   * selection in case of more than one match per header and
   * column grouping, returning a standardized table object.
3. `validator` service takes the standardized object and applies rules for `required` and `unique` fields, and returns two table objects, one with the correct and the other with the faulty data.
4. `writer` service receive the tables, one at time, and writes them to the output folder.

## Config structure

The read file format configuration, mandatory columns, possible column names, match selector and column grouper
are defined in the `config.go` file, allowing the adjustment of these parameters in a centralized and simple way.

#### `TableColumnSchema`

this object is responsible for mapping the name of the column, if it is required, unique and the list of possible synonymous words to find during column standardization.

```go
type TableColumnSchema struct {
    Name          string
    Unique        bool
    Required      bool
    PossibleWords []string
}
```
> If a column is not found within the `TableColumnSchema` list, it is kept in the output table as a non-standard column.

#### `ColumnMatcher`

this struct is responsible for selecting from a list of combinations for a column names, which column name should be kept.

```go
type ColumnMatcher struct {
    Matches  []string
    Selected string
}
```
> If a column has more than one match and this combination vector is not mapped in the `ColumnMatcher` list, an error is returned in application.

#### `ColumnGrouper`

this struct is responsible for mapping a set of column names that must be grouped into a new column.

```go
type ColumnGroup struct {
    Headers   []string
    GroupName string
    Separator string
}
```
> If only part of the columns in the `ColumnGroup` list is found, the grouping is not performed.

# Future works

For the evolution and improvement of the project, it would be interesting:

* Enable the processing of more than one input files at the same time, and allow grouping of all outputs into the same pair of tables.
* Read and aggregate the application's output tables in order to group all executions into a single output pattern, and do single-column validations based on the entire dataset.
* Implement the processing of other file structure types. The read and write architecture has already been designed thinking about new file formats, 
it is only necessary to respect the `IOStrategy` interface established.
* Present a parallel approach for column validation (required and unique fields), because these validations are not interdependent, 
which would simplify the implementation.
* Implement more table standardization rules, such as column deletion and data type validation in cells