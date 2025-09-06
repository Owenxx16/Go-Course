package student

import (
	"fmt"

	"github.com/Owenxx16/utils"
)

var students []Student

func addStudent() {
	var scores []float64
	var id int
	fmt.Println("==========THEM SINH VIEN==========")
	for {
		id := utils.GetPositiveInt("Nhap ID sinh vien: ")
		if utils.CheckIdDuplicate(id, students) {
			break
		}
		fmt.Println("ID da ton tai, vui long nhap ID khac.")
	}
	name := utils.GetNonEmptyString("Nhap ten sinh vien: ")
	class := utils.GetNonEmptyString("Nhap lop cua sinh vien: ")
	totalScores := utils.GetPositiveInt("Nhap so diem cua sinh vien: ")
	for i := 0; i < totalScores; i++ {
		score := utils.GetPositiveFloat(fmt.Sprintf("Nhap diem thu %d: ", i+1))
		scores = append(scores, float64(score))
	}
	student := Student{
		Id:    id,
		Name:  name,
		Class: class,
		Score: scores,
	}
	students = append(students, student)
	fmt.Println("Them sinh vien thanh cong!")
}
func deleteStudent() {
	fmt.Println("==========XOA SINH VIEN==========")
	id := utils.GetPositiveInt("Nhap ID sinh vien can xoa: ")
	for i, student := range students {
		if student.Id == id {
			students = append(students[:i], students[i+1:]...)
			fmt.Println("Xoa sinh vien thanh cong!")
			return
		}
	}
	fmt.Println("Khong tim thay sinh vien voi ID da cho.")
}
func updateStudent() {
	id := utils.GetPositiveInt("Nhap ID sinh vien can sua: ")
	for _, student := range students {
		if student.Id == id {
			fmt.Println("Nhap thong tin moi (de trong neu khong muon thay doi):")
			name := utils.GetOptionalString("Nhap ten sinh vien: ", student.Name)
			class := utils.GetOptionalString("Nhap lop cua sinh vien: ", student.Class)

			newScore := make([]float64, len(student.Score))
			for i, score := range student.Score {
				prompt := fmt.Sprintf("Nhap diem thu %d (de trong neu khong muon thay doi): ", i+1)
				newScore[i] = utils.GetOptionalPositiveFloat(prompt, score)
			}

			students := Student{
				Id:    student.Id,
				Name:  name,
				Class: class,
				Score: newScore,
			}
			fmt.Println("Sua thong tin sinh vien thanh cong!")
			return
		}
	}
	fmt.Println("Khong tim thay sinh vien voi ID da cho.")
}
func listStudents() {
	fmt.Println("==========DANH SACH SINH VIEN==========")
	if len(students) == 0 {
		fmt.Println("Khong co sinh vien nao trong danh sach.")
		return
	}
	for _, student := range students {
		fmt.Printf("ID: %d \n Ten: %s \n Lop: %s \n Diem: %v\n", student.Id, student.Name, student.Class, student.Score)
	}
}
func searchStudent() {
	fmt.Println("==========TIM KIEM SINH VIEN==========")
	id := utils.GetPositiveInt("Nhap ID sinh vien can tim: ")
	for _, s := range students {
		if s.Id == id {
			fmt.Println("Da tim thay sinh vien:", s.getInfo())
			return
		}
	}
	fmt.Println("Khong tim thay sinh vien voi ID da cho.")
}
func StudentMenu() {
	for {
		utils.ClearScreen()
		fmt.Println("==========QUAN LY SINH VIEN==========")
		fmt.Println("1. Them sinh vien")
		fmt.Println("2. Xoa sinh vien")
		fmt.Println("3. Sua sinh vien")
		fmt.Println("4. Danh sach sinh vien")
		fmt.Println("5. Tim kiem sinh vien")
		fmt.Println("6. Tro ve menu chinh")

		choice := utils.GetPositiveInt("Nhap lua chon cua ban: ")

		switch choice {
		case 1:
			addStudent()
		case 2:
			deleteStudent()
		case 3:
			updateStudent()
		case 4:
			listStudents()
		case 5:
			searchStudent()
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
