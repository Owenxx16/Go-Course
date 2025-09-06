package teacher

import (
	"fmt"

	"github.com/Owenxx16/utils"
)

var teachers []Teacher

func addTeacher() {
	fmt.Println("==========THEM SINH VIEN==========")
	var id int
	for {
		id := utils.GetPositiveInt("Nhap ID giang vien: ")
		if utils.CheckIdDuplicate(id, teachers) {
			break
		}
		fmt.Println("ID da ton tai, vui long nhap ID khac.")
	}
	name := utils.GetNonEmptyString("Nhap ten giang vien: ")
	subject := utils.GetNonEmptyString("Nhap mon day cua giang vien: ")
	baseSalary := utils.GetPositiveFloat("Nhap luong khoi diem cua giang vien: ")
	bonus := utils.GetPositiveFloat("Nhap so tien thuong cua sinh vien: ")

	teacher := Teacher{
		Id:         id,
		Name:       name,
		Subject:    subject,
		BaseSalary: baseSalary,
		Bonus:      bonus,
	}
	teachers = append(teachers, teacher)
	fmt.Println("Them giang vien thanh cong!")
}
func deleteTeacher() {}
func updateTeacher() {}
func listTeachers() {
	fmt.Println("==========DANH SACH GIANG VIEN==========")
	if len(teachers) == 0 {
		fmt.Println("Khong co giang vien nao trong danh sach.")
		return
	}
	for _, teacher := range teachers {
		fmt.Printf("ID: %d \n Ten: %s \n Lop: %s \n Diem: %v\n", teacher.Id, teacher.Name, teacher.Subject, teacher.BaseSalary, teacher.Bonus)
	}
}
func searchTeacher() {}
func TeacherMenu() {
	for {
		utils.ClearScreen()
		fmt.Println("==========QUAN LY SINH VIEN==========")
		fmt.Println("1. Them giang vien")
		fmt.Println("2. Xoa giang vien")
		fmt.Println("3. Sua giang vien")
		fmt.Println("4. Danh sach giang vien")
		fmt.Println("5. Tim kiem giang vien")
		fmt.Println("6. Tro ve menu chinh")

		choice := utils.GetPositiveInt("Nhap lua chon cua ban: ")

		switch choice {
		case 1:
			addTeacher()
		case 2:
			deleteTeacher()
		case 3:
			updateTeacher()
		case 4:
			listTeachers()
		case 5:
			searchTeacher()
		case 6:
			return
		default:
			fmt.Println("=========================================")
			fmt.Println("Lua chon khong hop le, vui long chon lai")
			fmt.Println("=========================================")
		}

		utils.ReadInput("Nhan Enter de tiep tuc...")
	}
}
