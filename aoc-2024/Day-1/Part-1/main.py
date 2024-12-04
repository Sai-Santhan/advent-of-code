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
        z += abs(x[i] - y[i])
    print(z)