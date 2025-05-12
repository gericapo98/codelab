i = 1
while i <= 5:
    print('*' * i)
    i += 1
print("Done")

secret_number = 9
guess_count = 0
guess_limit = 3
i = 0
# ---------------------------------------- #

# ---------------------------------------- #

for item in 'Python':
    print(item)

for item in ['Mosh']:
    print(item)

for item in range(5, 10, 2):
    print(item)

# ---------------------------------------- #

prices = [10, 20, 30]
for index in range(len(prices)):
    prices[index] = prices[index] * 0.9
print(prices)

# ---------------------------------------- #
numbers = [2, 2, 2, 2, 2]
total = 0
for number in numbers:
    total += number
print(total)