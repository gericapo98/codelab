# Dictionaries are mutable, unordered collections of key-value pairs.
# They are defined using curly braces {} and can contain various data types.
# Each key-value pair is separated by a colon (:), and pairs are separated by commas.
# The keys must be unique and immutable, while the values can be of any data type.
# Dictionaries are useful for storing related data and can be accessed using their keys.
# Example of a dictionary
# A dictionary is a collection of key-value pairs.
# It is mutable, meaning you can change it after creation.
# It is unordered, meaning the order of the items is not guaranteed.
# It is indexed, meaning you can access the items using their keys.
# It is iterable, meaning you can loop through the items.
# It is dynamic, meaning you can add or remove items at any time.
# It is a built-in data type in Python.
# It is a collection of key-value pairs.

customer = {
    "name": "John Doe",
    "age": 30,
    "is_verified": True,
}
print(customer.get("birthdate", "Jan 1, 1990"))  # Returns the value for the key "birthdate" or
# "Jan 1, 1990" if not found
