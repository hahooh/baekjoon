package main

import "fmt"

func main() {
	studentNum := 0
	fmt.Scanf("%d", &studentNum)

	var abovePercentage []float32

	for i := 0; i < studentNum; i++ {
		studentGradeNum := 0
		fmt.Scanf("%d", &studentGradeNum)
		totalGrade := float32(0)

		var studentGrades []float32
		for j := 0; j < studentGradeNum; j++ {
			var grade float32
			fmt.Scanf("%f", &grade)
			studentGrades = append(studentGrades, grade)
			totalGrade += grade
		}
		avg := totalGrade / float32(studentGradeNum)
		abovePercentage = append(abovePercentage, getPercentageOfAboveAvg(studentGrades, avg))
	}

	for _, avg := range abovePercentage {
		fmt.Printf("%.3f%%\n", avg)
	}
}

func getPercentageOfAboveAvg(studentGrades []float32, avg float32) float32 {
	above := float32(0)
	for _, grade := range studentGrades {
		if grade > avg {
			above++
		}
	}
	return above / float32(len(studentGrades)) * 100
}
