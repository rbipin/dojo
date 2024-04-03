def print_grades(grades):
    for grade in grades:
        print grade

def grades_sum(grades):
    total = 0
    for grade in grades: 
        total += grade
    return total
    
def grades_average(grades):
    sum_of_grades = grades_sum(grades)
    average = sum_of_grades / float(len(grades))
    return average
    
def grades_variance(grades):
    variance=0
    average=grades_average(grades)
    for item in grades:
        variance=variance+ ((average-item)**2)
    return variance/len(grades)

print grades_variance(grades)

