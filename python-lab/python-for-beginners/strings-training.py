course = "Python's course for beginnners"
print(course)

course2 = '''
    Hi Geri,
    This is a string that spans multiple lines.
'''
print(course2)

print(course[0]) # the indexes start at 0
print(course[-1]) # the indexes start at -1, it is the last character
print(course[0:3]) # the slice starts at 0 and ends at 3 (not included)
print(course[0:]) # the slice starts at 0 and goes to the end
print(course[:3]) # the slice starts at the beginning and goes to 3 (not included) 
print(course[0:5:2]) # the slice starts at 0 and ends at 5 (not included) and takes every second character
print(course[::2]) # the slice starts at the beginning and goes to the end and takes every second character
print(course[::-1]) # the slice starts at the end and goes to the beginning (reverses the string)

name = 'Jennifer'
print(name[1:-1]) # the slice starts at 1 and ends at -1 (not included)
print(name[1:]) # the slice starts at 1 and goes to the end
print(name[:]) # the slice starts at the beginning and goes to the end
print(name[-1:]) # the slice starts at -1 and goes to the end
print(name[-2:]) # the slice starts at -2 and goes to the end
print(name[-2:-1]) # the slice starts at -2 and ends at -1 (not included)
print(name[-2:]) # the slice starts at -2 and goes to the end
print(name[-2:-1]) # the slice starts at -2 and ends at -1 (not included)
print(name[::-1])
print(name[::]) 
print(name[::2])


first = 'John'
last = 'Smith'

message = first + ' [' + last + '] is a coder'
msg = f'{first} [{last}] is a coder'
print(message) 
print(msg)

course = "Python's course for beginnners"
print(len(course))
print(course.upper()) # converts the string to uppercase
print(course.lower()) # converts the string to lowercase
print(course.find('c')) # finds the first occurrence of 'P' in the string
print(course.replace('P', 'J')) # replaces 'P' with 'J' in the string
print('Python' in course) # checks if 'Python' is in the string
print('Python' not in course) # checks if 'Python' is not in the string
print(course.title()) # converts the first character of each word to uppercase
print(course.strip()) # removes leading and trailing whitespace
print(course.split()) # splits the string into a list of words
course.upper()
course.lower()
course.title()
course.strip()

