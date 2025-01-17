package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Asri-Mohamad/Master_Function"
)

// ---------------------------------------------------
func addTask(tasks *[]string) {

	newTask := bufio.NewReader(os.Stdin)

	fmt.Println("Add New task.....")
	fmt.Print("Enter new task :")
	read, _ := newTask.ReadString('\n')

	*tasks = append(*tasks, read)
	fmt.Println("New task added...")
	_ = Master_Function.CharGetKey()
}

//-------------------------------------------------

func deleteTask(tasks *[]string) {
	var index int
	if len(*tasks) <= 0 {
		fmt.Println("List is empty...")
		_ = Master_Function.CharGetKey()
		return
	}
	showList(*&tasks)
	fmt.Print("Enter number of task for delete :")
out1:
	for {
		fmt.Scanln(&index)
		if index >= 0 && index < len(*tasks) {
			break out1

		}
		fmt.Printf("\nError to Enter! please Enter betwin 0 to %v: ", len(*tasks)-1)

	}

	fmt.Printf("Are you shore for delete %v is %s task?(Y/N)", index, (*tasks)[index])

outLoop:
	for {

		switch Master_Function.CharGetKey() {
		case 'Y', 'y':
			*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
			fmt.Println("Y")
			fmt.Println("index is deleted...")
			break outLoop

		case 'N', 'n':
			fmt.Println("N")
			break outLoop

		}
	}
	fmt.Println()

}

// ---------------------------------------------------
func showMenu() {
	Master_Function.Cls()
	fmt.Print("1) Add to list .\n2)Delete from list.\n3)Show list.\n4)Exit list.\nAnsser :")

}

// ---------------------------------------------------
func showList(tasks *[]string) {
	fmt.Println()
	for key, data := range *tasks {
		fmt.Printf("%v) %s", key, data)

	}

}

// ---------------------------------------------------
func main() {
	var taskMaster []string

checkExit:
	for {
		showMenu()
		ch := Master_Function.CharGetKey()

		switch ch {
		case '1':
			fmt.Println("1")
			addTask(&taskMaster)
		case '2':
			fmt.Println("2")
			deleteTask(&taskMaster)
		case '3':
			fmt.Println("3")
			showList(&taskMaster)
			fmt.Println("Press any key...")
			_ = Master_Function.CharGetKey()
		case '4':
			fmt.Println("4")
			break checkExit

		}
	}

}
