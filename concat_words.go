package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
)

func main() {
    // Define command-line flags
    prefixPtr := flag.String("prefix", "", "Prefix to prepend to each word")
    delimiterPtr := flag.String("delimiter", "", "Delimiter to insert between prefix and word (optional)")
    wordListPtr := flag.String("wordlist", "", "Path to the text file containing the list of words, one per line")
    outputFilePtr := flag.String("output", "output.txt", "Path to the output text file where the new wordlist will be written")

    // Custom usage function for the flags
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
        fmt.Fprintf(os.Stderr, "  %s [OPTIONS] args...\n", os.Args[0])
        fmt.Fprintf(os.Stderr, "Options:\n")
        flag.PrintDefaults()
    }

    // Parse the flags
    flag.Parse()

    // Show usage menu if the prefix or wordlist file path is not provided
    if *prefixPtr == "" || *wordListPtr == "" {
        flag.Usage()
        os.Exit(1)
    }

    // Open the file containing the list of words
    file, err := os.Open(*wordListPtr)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
        os.Exit(1)
    }
    defer file.Close()

    // Create a scanner to read the file
    scanner := bufio.NewScanner(file)

    // Slice to hold the words from the file
    var words []string

    // Read words from the file and append to the slice
    for scanner.Scan() {
        words = append(words, scanner.Text())
    }

    // Check for scanner errors
    if err := scanner.Err(); err != nil {
        fmt.Fprintf(os.Stderr, "Error reading from file: %v\n", err)
        os.Exit(1)
    }

    // Open the output file for writing
    outputFile, err := os.Create(*outputFilePtr)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error creating output file: %v\n", err)
        os.Exit(1)
    }
    defer outputFile.Close()

    // Create a writer for the output file
    writer := bufio.NewWriter(outputFile)

    // Loop through the list of words and concatenate with prefix and delimiter
    for _, word := range words {
        newWord := *prefixPtr + *delimiterPtr + word
        // Write the new word to the output file, followed by a newline
        if _, err := writer.WriteString(newWord + "\n"); err != nil {
            fmt.Fprintf(os.Stderr, "Error writing to output file: %v\n", err)
            os.Exit(1)
        }
    }

    // Flush the writer to ensure all data is written to the file
    if err := writer.Flush(); err != nil {
        fmt.Fprintf(os.Stderr, "Error flushing output file: %v\n", err)
        os.Exit(1)
    }

    fmt.Printf("Concatenated words have been written to %s\n", *outputFilePtr)
}
