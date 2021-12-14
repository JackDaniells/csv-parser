Rain CSV Parser

Our company needs to parse CSV files coming from different employers to create a list of eligible employees to sign up for Rain.

Although files need to have key pieces of data, Rain does not control the overall structure of the files. For example, column names and order can be different from file to file.

Your job is to create a parser written in Golang that can read these files and standardize the data to be processed later. You can use the 4 samples files we provide (roster1.csv, roster2.csv, roster3.csv, roster4.csv) as a starting point, but you should support additional file structures.

Requirements:
•	Parse input files one at a time and generate two files as output, one for correct data, one for bad data
•	Minimal validations (you are free to create additional validations):
•	Required data:
•	Employee name
•	Employee salary
•	Employee email
•	Employee ID
•	Employee email must be unique
•	Employee ID must be unique
•	Output a summary of the processing steps in the console
•	Create unit tests
•	Project must be implemented in Golang 1.16+
•	Project must have a clear folder structure, scalable and simple (not a single file script)
•	In the README file, you should cover at least:
    - How to run the project
    - Explain your chosen architecture. 
    - Why do you think it is a good fit for this problem?
    - How you would evolve your submitted code

Notes:
•	IMPORTANT: Don’t post your project on public online repositories
•	You can contact us in case of questions
•	Useful links:
•	https://golang.org/doc/effective_go.html
•	https://www.quora.com/What-is-the-best-software-development-checklist?share=1
