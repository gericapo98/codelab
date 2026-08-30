has_high_income = True
has_good_credit = False

if has_high_income and has_good_credit:
    print("Eligible for loan")

if has_high_income or has_good_credit:
    print("maybe Eligible for loan")

temperature = 30
if temperature > 30:
    print("It's a hot day")
elif temperature > 20:
    print("It's a nice day")
elif temperature > 10:
    print("It's a bit chilly")
else:
    print("It's cold")
    