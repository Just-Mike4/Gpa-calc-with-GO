package main
import "fmt"
func to_points(x float64) float64{
	var value float64
	switch {
	case x<=39:
		value=0.0
	case x>=40 && x<=44:
		value=1.0
	case x>=45 && x<=49:
		value=2.0
	case x>=50 && x<=59:
		value=3.0
	case x>=60 && x<=69:
		value=4.0
	default:
		value=5.0
	}
	return value
}	
func main(){
	var Courses,score,loop,GPA,unit float64
	var Course string
	units_sum:=0.0
	CGP_sum:=0.0
	fmt.Println("Enter Number of courses:")
	fmt.Scanln(&Courses)
	var Course_name[]string
	var Units[]float64
	var Points[]float64
	var Scores[]float64
	var CGP[]float64
	for loop=0;loop<Courses;loop++{
		fmt.Println("Enter name of course:")
		fmt.Scanln(&Course)
		Course_name=append(Course_name,Course)

		fmt.Println("Enter unit of course:")
		fmt.Scanln(&unit)
		Units=append(Units,unit)

		fmt.Println("Enter score /100:")
		fmt.Scanln(&score)
		Scores=append(Scores,score)

		point:= to_points(score)
		Points=append(Points,point)
		CGP=append(CGP,point*unit)

	}
	for _,item:= range CGP{
		CGP_sum+=item

	}
	for _,items:= range Units{
		units_sum+=items
	}
	GPA=(CGP_sum/units_sum)
	fmt.Println("Courses are:",Course_name)
	fmt.Println("Units are:",Units)
	fmt.Println("Points are:",Points)
	fmt.Println("GPA is:",GPA)
}



