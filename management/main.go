package main

import (
	"fmt"

	"github.com/Owenxx16/student"
	"github.com/Owenxx16/teacher"
	"github.com/Owenxx16/utils"
)

func main() {

	// Main Menu for options

	for {
		utils.ClearScreen()
		fmt.Println("CHUONG TRINH QUAN LY")
		fmt.Println("1. Quan ly sinh vien")
		fmt.Println("2. Quan ly giang vien")
		fmt.Println("3. Thoat")
		//fmt.Println("Chon chuc nang: ")

		choice := utils.GetPositiveInt("Nhap lua chon cua ban: ")

		switch choice {
		case 1:
			student.StudentMenu()
		case 2:
			teacher.TeacherMenu()
		case 3:
			fmt.Println("Thoat chuong trinh")
			return
		default:
			fmt.Println("=========================================")
			fmt.Println("Lua chon khong hop le, vui long chon lai")
			fmt.Println("=========================================")
		}

		utils.ReadInput("Nhan Enter de tiep tuc...")
	}

}
