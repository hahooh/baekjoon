package main

import "fmt"

func main() {
	studentNum := 0
	fmt.Scanf("%d", &studentNum)

	var studentsGrades [][]float32
	for i := 0; i < studentNum; i++ {
		studentGradeNum := 0
		fmt.Scanf("%d", &studentGradeNum)
		var studentGrades []float32
		for j := 0; j < studentGradeNum; j++ {
			var grade float32
			fmt.Scanf("%f", &grade)
			studentGrades = append(studentGrades, grade)
		}
		studentsGrades = append(studentsGrades, studentGrades)
	}

	for _, studentGrades := range studentsGrades {
		totalGrade := float32(0)
		totalStudents := float32(len(studentGrades))
		for _, grade := range studentGrades {
			totalGrade += float32(grade)
		}
		avg := totalGrade / totalStudents
		fmt.Printf("%.3f%%\n", getPercentageOfAboveAVG(studentGrades, avg))
	}
}

func getPercentageOfAboveAVG(studentGrades []float32, avg float32) float32 {
	above := float32(0)
	for _, grade := range studentGrades {
		if grade > avg {
			above++
		}
	}
	return above / float32(len(studentGrades)) * 100
}
