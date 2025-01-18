package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Asri-Mohamad/Master_Function"
)

type taskStruct struct {
	Task string `json:"task"`
	Date string `json:"data"`
	Time string `json:"time"`
}

func main() {
	var taskMaster []taskStruct

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
			saveList(taskMaster)
		case '5':
			fmt.Println("5")
			taskMaster = loadList(taskMaster)
		case '6':
			fmt.Println("6")
			if len(taskMaster) > 0 {
				fmt.Println("The list in not empty do you want Save change?(Y/N)")
				for {
					yn := Master_Function.CharGetKey()
					if yn == 'Y' || yn == 'y' || yn == 'n' || yn == 'N' {
						if yn == 'n' || yn == 'N' {
							break checkExit
						} else {
							saveList(taskMaster)
							fmt.Println("List save and exit ....")
							break checkExit
						}

					}
				}

			}

			break checkExit

		}
	}

}

// ---------------------------------------------------
func showMenu() {
	Master_Function.Cls()
	fmt.Print("1) Add to list .\n2)Delete from list.\n3)Show list.\n4)Save to file\n5)Load List\n6)Exit list.\nAnsser :")

}

// ---------------------------------------------------
func addTask(tasks *[]taskStruct) {

	newTask := bufio.NewReader(os.Stdin)

	fmt.Println("Add New task.....")
	fmt.Print("Enter new task :")
	readTask, _ := newTask.ReadString('\n')
	readTask = strings.TrimSpace(readTask)
	fmt.Print("Enter date :")
	readDate, _ := newTask.ReadString('\n')
	readDate = strings.TrimSpace(readDate)
	fmt.Print("Enter Time :")
	readTime, _ := newTask.ReadString('\n')
	readTime = strings.TrimSpace(readTime)

	read := taskStruct{
		Task: readTask,
		Date: readDate,
		Time: readTime}
	*tasks = append(*tasks, read)
	fmt.Println("New task added...")
	//fmt.Printf("%T - %T - %T", readTask, readDate, readTime)
	_ = Master_Function.CharGetKey()
}

//-------------------------------------------------

func deleteTask(tasks *[]taskStruct) {
	var index int
	if len(*tasks) <= 0 {
		fmt.Println("List is empty...")
		_ = Master_Function.CharGetKey()
		return
	}
	showList(tasks)
	fmt.Print("Enter number of task for delete :")
out1:
	for {
		fmt.Scanln(&index)
		if index >= 0 && index < len(*tasks) {
			break out1

		}
		fmt.Printf("\nError to Enter! please Enter betwin 0 to %v: ", len(*tasks)-1)

	}

	fmt.Printf("\n%v) Task: %s   Date: %s   Time: %s\nAre you shore for delete ?(Y/N)", index, (*tasks)[index].Task, (*tasks)[index].Date, (*tasks)[index].Time)

outLoop:
	for {

		switch Master_Function.CharGetKey() {
		case 'Y', 'y':
			*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
			fmt.Println("Y")
			fmt.Println("Deleted Complet...")
			break outLoop

		case 'N', 'n':
			fmt.Println("N")
			break outLoop

		}
	}
	fmt.Println()
	_ = Master_Function.CharGetKey()
}

// ---------------------------------------------------
func showList(tasks *[]taskStruct) {
	fmt.Println()
	for key, data := range *tasks {
		fmt.Printf("%v) Task: %s   Date: %s   Time: %s\n--------------\n", key, data.Task, data.Date, data.Time)

	}

}

// ---------------------------------------------------
func saveList(tasks []taskStruct) {

	showList(&tasks)
	print("Do you want to save this list(Y/N)?")

	for {
		yn := Master_Function.CharGetKey()
		if yn == 'Y' || yn == 'y' || yn == 'n' || yn == 'N' {
			if yn == 'n' || yn == 'N' {
				return
			}
			break
		}
	}

	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Create file have problem...")
		_ = Master_Function.CharGetKey()
		return

	} else {

		defer file.Close()

		encode := json.NewEncoder(file)
		err := encode.Encode(tasks)
		if err != nil {
			fmt.Println("Write file have problem....")
			_ = Master_Function.CharGetKey()
			return
		}

		fmt.Printf("Save is completly...%v data saved", len(tasks))
		_ = Master_Function.CharGetKey()
	}

}

// ---------------------------------------------------
func loadList(tasks []taskStruct) []taskStruct {
	if len(tasks) > 0 {
		fmt.Println("The list in not empty do you want over write?(Y/N)")
		for {
			yn := Master_Function.CharGetKey()
			if yn == 'Y' || yn == 'y' || yn == 'n' || yn == 'N' {
				if yn == 'n' || yn == 'N' {
					return tasks
				}
				break
			}
		}
	}
	file, err := os.Open("tasks.json")

	if err != nil {
		fmt.Println("I Can't open file ....")
		_ = Master_Function.CharGetKey()
		return tasks
	} else {
		defer file.Close()
		decode := json.NewDecoder(file)
		err := decode.Decode(&tasks)
		if err != nil {
			fmt.Println("Decoding file problem....")
			_ = Master_Function.CharGetKey()
			return tasks
		}

	}
	fmt.Println("The list loded OK....")
	_ = Master_Function.CharGetKey()
	return tasks
}

// ---------------------------------------------------
