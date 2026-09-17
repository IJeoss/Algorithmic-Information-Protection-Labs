import csv
import math
import matplotlib.pyplot as plt

n_values=[]
degrees=[]
karatsuba_ms=[]
toom3_ms=[]

with open("results.csv", "r", newline="") as file:
    reader=csv.DictReader(file)

    for row in reader:
        n_values.append(int(row["n"]))
        degrees.append(int(row["degree"]))
        karatsuba_ms.append(float(row["karatsuba_ms"]))
        toom3_ms.append(float(row["toom3_ms"]))

plt.figure(figsize=(10, 6))

plt.plot(n_values, karatsuba_ms, marker="o", label="Карацуба")
plt.plot(n_values, toom3_ms, marker="o", label="Тоом-Кук-3")

plt.xlabel("n")
plt.ylabel("Время (мс)")
plt.title("Производительность умножения полиномов")
plt.legend()
plt.grid(True)

plt.savefig("results_plot.png", dpi=200)

ordinary_theory=[n**2 for n in n_values]
karatsuba_theory=[n**math.log(3, 2) for n in n_values]
toom3_theory=[n**math.log(5, 3) for n in n_values]

ordinary_theory=[x/ordinary_theory[0] for x in ordinary_theory]
karatsuba_theory=[x/karatsuba_theory[0] for x in karatsuba_theory]
toom3_theory=[x/toom3_theory[0] for x in toom3_theory]

plt.figure(figsize=(10, 6))

plt.plot(n_values, ordinary_theory, label="Наивный O(n^2)")
plt.plot(n_values, karatsuba_theory, label="Карацуба O(n^log2_3)")
plt.plot(n_values, toom3_theory, label="Тоом-Кук3 O(n^log3_5)")

plt.xlabel("n")
plt.ylabel("Нормированный рост")
plt.title("Теоретическая сложность")
plt.legend()
plt.grid(True)

plt.savefig("complexity_plot.png", dpi=200)
plt.show()
