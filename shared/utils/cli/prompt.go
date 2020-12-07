package cli

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)


// Prompt for user input
func Prompt(initialPrompt string, expectedFormat string, incorrectFormatPrompt string) string {

    // Print initial prompt
    fmt.Println(initialPrompt)

    // Get valid user input
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan(); !regexp.MustCompile(expectedFormat).MatchString(scanner.Text()); scanner.Scan() {
        fmt.Println("")
        fmt.Println(incorrectFormatPrompt)
    }
    fmt.Println("")

    // Return user input
    return scanner.Text()

}


// Prompt for confirmation
func Confirm(initialPrompt string) bool {
    response := Prompt(fmt.Sprintf("%s [y/n]", initialPrompt), "(?i)^(y|yes|n|no)$", "Please answer 'y' or 'n'")
    return (strings.ToLower(response[:1]) == "y")
}


// Prompt for user selection
func Select(initialPrompt string, options []string) (int, string) {

    // Get prompt
    prompt := initialPrompt
    for i, option := range options {
        prompt += fmt.Sprintf("\n%d: %s", (i + 1), option)
    }

    // Get expected response format
    optionNumbers := []string{}
    for i, _ := range options {
        optionNumbers = append(optionNumbers, strconv.Itoa(i + 1))
    }
    expectedFormat := fmt.Sprintf("^(%s)$", strings.Join(optionNumbers, "|"))

    // Prompt user
    response := Prompt(prompt, expectedFormat, "Please enter a number corresponding to an option")

    // Get selected option
    index, _ := strconv.Atoi(response)
    selectedIndex := index - 1
    selectedOption := options[selectedIndex]

    // Return
    return selectedIndex, selectedOption

}

