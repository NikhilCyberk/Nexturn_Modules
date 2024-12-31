/*
Exercise 1: Employee Management System

Topics Covered: Go Conditions, Go Loops, Go Constants, Go Functions, Go Arrays, Go Strings, Go Errors

Case Study:

A company wants to manage its employees' data in memory. Each employee has an ID, name, age, and department. You need to build a small application that performs the following:

1. Add Employee: Accept input for employee details and store them in an array of structs. Validate the input:

	o ID must be unique.
	o Age should be greater than 18. If validation fails, return custom error messages.

2. Search Employee: Search for an employee by ID or name using conditions. Return the details if found, or return an error if not found.

3. List Employees by Department: Use loops to filter and display all employees in a given department.

4. Count Employees: Use constants to define a department (e.g., "HR", "IT"), and display the count of employees in that department.

Bonus:

Refactor the repetitive code using functions, and add error handling for invalid operations like searching for a non-existent employee.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Department constants
const (
	HR      = "HR"
	IT      = "IT"
	FINANCE = "FINANCE"
)

// Custom errors
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

// validateEmployee checks if employee data is valid
func validateEmployee(emp Employee, employees []Employee) error {
	// Validate ID
	if emp.ID <= 0 {
		return &ValidationError{Field: "ID", Message: "must be positive"}
	}
	for _, e := range employees {
		if e.ID == emp.ID {
			return &ValidationError{Field: "ID", Message: "must be unique"}
		}
	}

	// Validate Name
	if strings.TrimSpace(emp.Name) == "" {
		return &ValidationError{Field: "Name", Message: "cannot be empty"}
	}

	// Validate Age
	if emp.Age <= 18 {
		return &ValidationError{Field: "Age", Message: "must be greater than 18"}
	}

	// Validate Department
	dept := strings.ToUpper(emp.Department)
	if dept != HR && dept != IT && dept != FINANCE {
		return &ValidationError{Field: "Department", Message: "must be HR, IT, or FINANCE"}
	}

	return nil
}

// readInput reads user input with proper error handling
func readInput(prompt string) (string, error) {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return "", scanner.Err()
	}
	return strings.TrimSpace(scanner.Text()), nil
}

// addEmployee adds a new employee with validation
func addEmployee(employees *[]Employee) error {
	fmt.Println("\n=== Adding New Employee ===")
	
	// Read ID
	idStr, err := readInput("Enter employee ID: ")
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return &ValidationError{Field: "ID", Message: "must be a number"}
	}

	// Read Name
	name, err := readInput("Enter employee name: ")
	if err != nil {
		return err
	}

	// Read Age
	ageStr, err := readInput("Enter employee age: ")
	if err != nil {
		return err
	}
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return &ValidationError{Field: "Age", Message: "must be a number"}
	}

	// Read Department
	dept, err := readInput("Enter employee department (HR/IT/FINANCE): ")
	if err != nil {
		return err
	}

	newEmployee := Employee{
		ID:         id,
		Name:       name,
		Age:        age,
		Department: strings.ToUpper(dept),
	}

	if err := validateEmployee(newEmployee, *employees); err != nil {
		return err
	}

	*employees = append(*employees, newEmployee)
	fmt.Println("✅ Employee added successfully!")
	return nil
}

// searchEmployee searches for an employee by ID or name
func searchEmployee(employees []Employee) {
	query, err := readInput("\nEnter employee ID or name to search: ")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	found := false
	fmt.Println("\n=== Search Results ===")
	
	for _, emp := range employees {
		if strconv.Itoa(emp.ID) == query || strings.Contains(strings.ToLower(emp.Name), strings.ToLower(query)) {
			displayEmployee(emp)
			found = true
		}
	}

	if !found {
		fmt.Println("❌ No employees found matching your search.")
	}
}

// listEmployeesByDepartment lists all employees in a department
func listEmployeesByDepartment(employees []Employee) {
	dept, err := readInput("\nEnter department (HR/IT/FINANCE): ")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	dept = strings.ToUpper(dept)
	count := 0
	fmt.Printf("\n=== Employees in %s Department ===\n", dept)

	for _, emp := range employees {
		if emp.Department == dept {
			displayEmployee(emp)
			count++
		}
	}

	if count == 0 {
		fmt.Printf("❌ No employees found in %s department.\n", dept)
	}
}

// countEmployeesByDepartment counts employees in each department
func countEmployeesByDepartment(employees []Employee) {
	counts := make(map[string]int)
	
	for _, emp := range employees {
		counts[emp.Department]++
	}

	fmt.Println("\n=== Department Counts ===")
	fmt.Printf("HR: %d employees\n", counts[HR])
	fmt.Printf("IT: %d employees\n", counts[IT])
	fmt.Printf("FINANCE: %d employees\n", counts[FINANCE])
}

// displayEmployee shows employee details in a formatted way
func displayEmployee(emp Employee) {
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("ID: %d\n", emp.ID)
	fmt.Printf("Name: %s\n", emp.Name)
	fmt.Printf("Age: %d\n", emp.Age)
	fmt.Printf("Department: %s\n", emp.Department)
}

// showAllEmployees displays all employees
func showAllEmployees(employees []Employee) {
	if len(employees) == 0 {
		fmt.Println("\n❌ No employees in the system.")
		return
	}

	fmt.Println("\n=== All Employees ===")
	for _, emp := range employees {
		displayEmployee(emp)
	}
}

// fill sample data
func fillSampleData(employees *[]Employee) {
	*employees = []Employee{
		{ID: 1, Name: "John Doe", Age: 25, Department: HR},
		{ID: 2, Name: "Jane Smith", Age: 30, Department: IT},
		{ID: 3, Name: "Bob Johnson", Age: 35, Department: FINANCE},
	}
}


func main() {
	employees := make([]Employee, 0)
	fillSampleData(&employees)	

	for {
		fmt.Println("\n=== Employee Management System ===")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. List Employees by Department")
		fmt.Println("4. Count Employees by Department")
		fmt.Println("5. Show All Employees")
		fmt.Println("6. Exit")
		
		choice, err := readInput("Enter your choice (1-6): ")
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}

		switch choice {
		case "1":
			if err := addEmployee(&employees); err != nil {
				fmt.Printf("❌ Error adding employee: %v\n", err)
			}
		case "2":
			searchEmployee(employees)
		case "3":
			listEmployeesByDepartment(employees)
		case "4":
			countEmployeesByDepartment(employees)
		case "5":
			showAllEmployees(employees)
		case "6":
			fmt.Println("\n👋 Goodbye!")
			return
		default:
			fmt.Println("❌ Invalid choice. Please try again.")
		}
	}
}