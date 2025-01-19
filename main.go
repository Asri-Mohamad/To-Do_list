package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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
		fmt.Println(ch - 48)
		switch ch {
		case '1':

			addTask(&taskMaster)
		case '2':

			deleteTask(&taskMaster)
		case '3':

			taskMaster = editTask(taskMaster)

		case '4':

			showList(&taskMaster)
			fmt.Println("Press any key...")
			_ = Master_Function.CharGetKey()
		case '5':

			saveList(taskMaster)
		case '6':

			taskMaster = loadList(taskMaster)
		case '7':

			if len(taskMaster) > 0 {
				fmt.Printf("The list in not empty do you want Save change?(%v/%v)\n", Master_Function.ColorText("green", "Y"), Master_Function.ColorText("Read", "N"))
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

	fmt.Printf("%s) Add to list\n%s) Delete from list\n%s) Edit \n%s) Show list\n%s) Save to file\n%s) Load list\n%s) Exit\nAnswer: ",
		Master_Function.ColorText("green", "1"), Master_Function.ColorText("green", "2"), Master_Function.ColorText("green", "3"),
		Master_Function.ColorText("green", "4"), Master_Function.ColorText("green", "5"),
		Master_Function.ColorText("green", "6"), Master_Function.ColorText("green", "7"))

}

// ---------------------------------------------------
func addTask(tasks *[]taskStruct) {

	newTask := bufio.NewReader(os.Stdin)

	fmt.Println("Add New task.....")
	fmt.Printf("Enter new  %s :", Master_Function.ColorText("yellow", "Task"))
	readTask, _ := newTask.ReadString('\n')
	readTask = strings.TrimSpace(readTask)
	fmt.Printf("Enter date %s :", Master_Function.ColorText("blue", "Date"))
	readDate, _ := newTask.ReadString('\n')
	readDate = strings.TrimSpace(readDate)
	fmt.Printf("Enter time %s :", Master_Function.ColorText("mango", "Time"))
	readTime, _ := newTask.ReadString('\n')
	readTime = strings.TrimSpace(readTime)

	read := taskStruct{
		Task: readTask,
		Date: readDate,
		Time: readTime}
	*tasks = append(*tasks, read)
	fmt.Println("New task added...")

	_ = Master_Function.CharGetKey()
}

//-------------------------------------------------

func deleteTask(tasks *[]taskStruct) {
	var index int
	var getI string
	if len(*tasks) <= 0 {
		fmt.Println("List is empty...")
		_ = Master_Function.CharGetKey()
		return
	}
	showList(tasks)
	fmt.Print("Enter number of task for delete :")
out1:
	for {
		fmt.Scanln(&getI)
		if index, err := strconv.Atoi(getI); err == nil {

			if index >= 0 && index < len(*tasks) {
				break out1

			}
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
		fmt.Printf("%s  %s %s   %s %s   %s %s\n%s\n", Master_Function.ColorText("read", strconv.Itoa(key)+")"),
			Master_Function.ColorText("yellow", "Task :"), Master_Function.ColorText("green", data.Task),
			Master_Function.ColorText("yellow", "Date:"), Master_Function.ColorText("cyan", data.Date),
			Master_Function.ColorText("yellow", "Time:"), Master_Function.ColorText("blue", data.Time),
			Master_Function.ColorText("magenta", "-----------------------"))

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

		}

		fmt.Printf("\nSave is completly...%v data saved", len(tasks))
		_ = Master_Function.CharGetKey()
	}

}

// ---------------------------------------------------
func loadList(tasks []taskStruct) []taskStruct {
	if len(tasks) > 0 {
		fmt.Printf("\n%s", Master_Function.ColorText("yellow", "The list in not empty do you want over write?(Y/N)"))
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
		fmt.Printf("%s\n", Master_Function.ColorText("read", "I Can't open file ...."))
		_ = Master_Function.CharGetKey()
		return tasks
	} else {
		defer file.Close()
		decode := json.NewDecoder(file)
		err := decode.Decode(&tasks)
		if err != nil {
			fmt.Printf("\n%s", Master_Function.ColorText("read", "Decoding file problem...."))
			_ = Master_Function.CharGetKey()
			return tasks
		}

	}
	fmt.Printf("%s\n", Master_Function.ColorText("yellow", "The list loded OK...."))
	_ = Master_Function.CharGetKey()
	return tasks
}

// ---------------------------------------------------
func editTask(tasks []taskStruct) []taskStruct {
	var getI string
	var index int
	if len(tasks) <= 0 {
		fmt.Println("List is empty...")
		_ = Master_Function.CharGetKey()
		return tasks
	}
	showList(&tasks)
	fmt.Print("Enter number of task for Edit :")
out1:
	for {
		fmt.Scanln(&getI)
		if index, err := strconv.Atoi(getI); err == nil {

			if index >= 0 && index < len(tasks) {

				break out1

			}
		}
		fmt.Printf("\nError to Enter! please Enter betwin 0 to %v: ", len(tasks)-1)

	}

	fmt.Printf("\n%v) Task: %s   Date: %s   Time: %s\nAre you shore for Edit ?(Y/N)", index, tasks[index].Task,
		tasks[index].Date, tasks[index].Time)

outLoop:
	for {

		switch Master_Function.CharGetKey() {
		case 'Y', 'y':
			newTask := bufio.NewReader(os.Stdin)
			fmt.Printf("\nEnter new  %s :", Master_Function.ColorText("yellow", "Task"))
			readTask, _ := newTask.ReadString('\n')
			readTask = strings.TrimSpace(readTask)
			fmt.Printf("Enter date %s :", Master_Function.ColorText("blue", "Date"))
			readDate, _ := newTask.ReadString('\n')
			readDate = strings.TrimSpace(readDate)
			fmt.Printf("Enter time %s :", Master_Function.ColorText("mango", "Time"))
			readTime, _ := newTask.ReadString('\n')
			readTime = strings.TrimSpace(readTime)

			tasks[index] = taskStruct{readTask, readDate, readTime}
			fmt.Println("\nEdit Complet...")
			break outLoop

		case 'N', 'n':
			fmt.Println("N")
			break outLoop

		}
	}
	return tasks
}

// ---------------------------------------------------
