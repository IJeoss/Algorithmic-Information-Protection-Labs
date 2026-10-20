#!/usr/bin/env python3
from __future__ import annotations

import argparse
import csv
import re
import statistics
from collections import defaultdict
from pathlib import Path

import matplotlib.pyplot as plt

BENCHMARK_PATTERN=re.compile(
    r"^BenchmarkMatrixMultiply/N=(\d+)/(reference|packed_naive|four_russians/k=(\d+))-\d+"
    r"\s+\d+\s+([\d.]+)\s+ns/op(?:\s|$)"
)


def parse_results(path: Path) -> dict[tuple[int, str, int | None], list[float]]:
    samples: dict[tuple[int, str, int | None], list[float]]=defaultdict(list)
    for line in path.read_text(encoding="utf-8").splitlines():
        match=BENCHMARK_PATTERN.match(line.strip())
        if not match:
            continue

        size=int(match.group(1))
        method_name=match.group(2)
        block_size=int(match.group(3)) if match.group(3) else None
        method="four_russians" if block_size is not None else method_name
        ns_per_op=float(match.group(4))
        samples[(size, method, block_size)].append(ns_per_op)

    if not samples:
        raise ValueError(
            f"В{path} не найдено строк BenchmarkMatrixMultiply. "
            "Проверь команду go test -bench и путь к файлу."
        )
    return samples


def write_csv(
    samples: dict[tuple[int, str, int | None], list[float]], output: Path
) -> dict[tuple[int, str, int | None], float]:
    medians: dict[tuple[int, str, int | None], float] ={}
    with output.open("w", encoding="utf-8", newline="") as f:
        writer=csv.writer(f, lineterminator="\n")
        writer.writerow(
            ["N", "method", "k", "runs", "median_ns_op", "min_ns_op", "max_ns_op"]
        )
        for (n, method, k), values in sorted(
            samples.items(), key=lambda item: (item[0][0], item[0][1], item[0][2] or 0)
        ):
            median=statistics.median(values)
            medians[(n, method, k)]=median
            writer.writerow(
                [n, method, "" if k is None else k, len(values), median, min(values), max(values)]
            )
    return medians


def series(
    medians: dict[tuple[int, str, int | None], float], method: str
) -> tuple[list[int], list[float]]:
    pairs=sorted(
        (n, value)
        for (n, name, k), value in medians.items()
        if name==method and k is None
    )
    return [n for n, _ in pairs], [value / 1_000_000 for _, value in pairs]


def best_for_each_size(
    medians: dict[tuple[int, str, int | None], float]
) -> dict[int, tuple[int, float]]:
    best: dict[int, tuple[int, float]] ={}
    for (n, method, k), value in medians.items():
        if method != "four_russians" or k is None:
            continue
        if n not in best or value < best[n][1]:
            best[n]=(k, value)
    return best


def save_plot(path: Path) -> None:
    plt.tight_layout()
    plt.savefig(path, dpi=180)
    plt.close()
    print(f"График:{path}")


def plot_timings(
    medians: dict[tuple[int, str, int | None], float],
    best: dict[int, tuple[int, float]],
    output_dir: Path,
) -> None:
    plt.figure(figsize=(9, 5.5))
    for method, label in (
        ("reference", "Скалярный O(N³)"),
        ("packed_naive", "Наивный, XOR строк"),
    ):
        ns, ms=series(medians, method)
        if ns:
            plt.plot(ns, ms, marker="o", label=label)

    ns=sorted(best)
    if ns:
        plt.plot(
            ns,
            [best[n][1] / 1_000_000 for n in ns],
            marker="o",
            label="Four Russians (лучший k)",
        )
    plt.xscale("log", base=2)
    plt.yscale("log")
    plt.xticks(ns, [str(n) for n in ns])
    plt.xlabel("Размер матрицы N × N")
    plt.ylabel("Время, мс/операцию (медиана)")
    plt.title("Сравнение времени умножения")
    plt.grid(True, which="both", alpha=0.3)
    plt.legend()
    save_plot(output_dir / "timings.png")


def plot_speedups(
    medians: dict[tuple[int, str, int | None], float],
    best: dict[int, tuple[int, float]],
    output_dir: Path,
) -> None:
    ns=sorted(n for n in best if (n, "packed_naive", None) in medians)
    speedups=[medians[(n, "packed_naive", None)] / best[n][1] for n in ns]
    plt.figure(figsize=(9, 5.5))
    plt.plot(ns, speedups, marker="o")
    plt.axhline(1.0, linestyle="--", linewidth=1, label="Одинаковая скорость")
    plt.xscale("log", base=2)
    plt.xticks(ns, [str(n) for n in ns])
    plt.xlabel("Размер матрицы N × N")
    plt.ylabel("Ускорение относительно packed_naive, раз")
    plt.title("Ускорение Four Russians с лучшим k")
    plt.grid(True, alpha=0.3)
    plt.legend()
    save_plot(output_dir / "speedups.png")


def plot_k_sweep(
    medians: dict[tuple[int, str, int | None], float],
    best: dict[int, tuple[int, float]],
    output_dir: Path,
) -> None:
    plt.figure(figsize=(9, 5.5))
    for n in sorted(best):
        baseline=medians.get((n, "packed_naive", None))
        if baseline is None:
            continue
        values=sorted(
            (k, baseline / t)
            for (size, method, k), t in medians.items()
            if size==n and method=="four_russians" and k is not None
        )
        if values:
            plt.plot(
                [k for k, _ in values],
                [speedup for _, speedup in values],
                marker="o",
                label=f"N={n}",
            )
    plt.axhline(1.0, linestyle="--", linewidth=1)
    plt.xlabel("Размер блока k")
    plt.ylabel("Ускорение относительно packed_naive, раз")
    plt.title("Зависимость ускорения Four Russians от k")
    plt.grid(True, alpha=0.3)
    plt.legend(ncol=2)
    save_plot(output_dir / "block_sizes.png")


def print_summary(
    medians: dict[tuple[int, str, int | None], float],
    best: dict[int, tuple[int, float]],
) -> None:
    print("\n N    лучший k    Four Russians, нс/op    ускорение к packed_naive")
    print("-" * 70)
    first_win=None
    for n in sorted(best):
        k, best_time=best[n]
        baseline=medians.get((n, "packed_naive", None))
        speedup=baseline / best_time if baseline else None
        if speedup is not None and speedup > 1 and first_win is None:
            first_win=n
        ratio=f"{speedup:>8.2f}×" if speedup is not None else "      н/д"
        print(f"{n:>3}{k:>10}{best_time:>23.1f}{ratio:>27}")

    print("\nПервый измеренный размер, где Four Russians быстрее packed_naive:",
          first_win if first_win is not None else "не найден")
    print("Это результаты конкретной машины/серии замеров, не универсальный порог.")


def main() -> None:
    default_results=Path(__file__).resolve().parent.parent / "results"
    parser=argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "input", nargs="?", type=Path, default=default_results / "bench.txt",
        help="файл с выводом go test -bench",
    )
    parser.add_argument(
        "--output-dir", type=Path, default=default_results,
        help="каталог для CSV и графиков",
    )
    args=parser.parse_args()
    args.output_dir.mkdir(parents=True, exist_ok=True)
    samples=parse_results(args.input)
    medians=write_csv(samples, args.output_dir / "timings.csv")
    best=best_for_each_size(medians)
    if not best:
        raise ValueError("Нет результатов Four Russians. Проверь полноту бенчмарка.")
    plot_timings(medians, best, args.output_dir)
    plot_speedups(medians, best, args.output_dir)
    plot_k_sweep(medians, best, args.output_dir)
    print(f"CSV:{args.output_dir / 'timings.csv'}")
    print_summary(medians, best)


if __name__=="__main__":
    main()
