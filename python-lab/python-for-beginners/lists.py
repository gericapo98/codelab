names = ['Alice', 'Bob', 'Charlie', 'David', 'Eve']
print (names)
print (names[0])
print (names[1])
print (names[2:5])

# ------------------------------------------ #

numbers = [3, 6, 2, 8, 4, 10]
max = numbers[0]
for number in numbers:
    if number > max:
        max = number
print(max)

# ------------------------------------------ #
matrix = [
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9]  
]
print(matrix[0][1])  # Output: 1
for row in matrix:
    print()
    for item in row:
        print(item)

#------------------------------------------- #

numbers = [5, 2, 1, 7, 4]
numbers.insert(0, 10)  # Insert 10 at index 0
numbers.append(20)  # Append 20 to the end of the list
numbers.remove(2)  # Remove the first occurrence of 2
numbers.pop()  # Remove the last element
numbers.sort()  # Sort the list in ascending order
numbers.reverse()  # Reverse the order of the list
numbers.clear()  # Clear the list
# numbers = []  # Create an empty list 
# numbers = list()  # Create an empty list
numbers = numbers * 2
print(numbers)
print(numbers.__contains__(2))
print(numbers.count(2))  # Count occurrences of 2
print(numbers.append(2))  # Append 2 to the list
print(numbers.index(2))  # Get the index of the first occurrence of 2
print(numbers[0:3])  # Slice the list from index 0 to 2
print(numbers[0:])  # Slice the list from index 0 to the end
print(numbers[:3])  # Slice the list from the beginning to index 2
print(numbers[0:5:2])  # Slice the list from index 0 to 4 with a step of 2
print(numbers[::2])  # Slice the list from the beginning to the end with a step of 2
#------------------------------------------- #
numbers = [3, 3, 3, 4, 5]
numbers.sort()
print(numbers)
numbers.sort(reverse=True)
print(numbers)
#------------------------------------------- #
numbers = [2,2,4,6,3,4,6,1]
uniques = []
for number in numbers:
    if number not in uniques:
        uniques.append(number)
#-------------------------------------------- #

numbers = [1, 2, 3, 4, 5]
x,y,z = numbers[0], numbers[1], numbers[2]

