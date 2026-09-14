import csv
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

plt.plot(n_values, karatsuba_ms, marker="o", label="Karatsuba")
plt.plot(n_values, toom3_ms, marker="o", label="Toom-3")

plt.xlabel("n")
plt.ylabel("Time (ms)")
plt.title("Polynomial multiplication performance")
plt.legend()
plt.grid(True)

plt.savefig("results_plot.png", dpi=200)
plt.show()