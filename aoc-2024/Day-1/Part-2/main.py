def find_first_occurrence(arr, target):
    low, high = 0, len(arr) - 1
    first_occurrence = -1
    while low <= high:
        mid = (low + high) // 2
        if arr[mid] == target:
            first_occurrence = mid
            high = mid - 1  # Continue searching in the left half
        elif arr[mid] < target:
            low = mid + 1
        else:
            high = mid - 1
    return first_occurrence

def find_last_occurrence(arr, target):
    low, high = 0, len(arr) - 1
    last_occurrence = -1
    while low <= high:
        mid = (low + high) // 2
        if arr[mid] == target:
            last_occurrence = mid
            low = mid + 1  # Continue searching in the right half
        elif arr[mid] < target:
            low = mid + 1
        else:
            high = mid - 1
    return last_occurrence

def count_occurrences(arr, target):
    first = find_first_occurrence(arr, target)
    if first == -1:
        return 0  # Target not found in the array
    last = find_last_occurrence(arr, target)
    return last - first + 1

with open("../input.txt", "r") as file:
    x = []
    y = []
    for line in file:
        [a, b] = line.split("   ")
        x.append(int(a))
        y.append(int(b))
    x.sort()
    y.sort()
    z = 0
    for i in range(len(x)):
        z += x[i] * count_occurrences(y, x[i])
    print(z)