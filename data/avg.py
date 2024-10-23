import csv
from collections import defaultdict
import numpy as np
import matplotlib.pyplot as plt
import pandas as pd

def compute_averages(csv_file):
    df = pd.read_csv(csv_file)

    average_results = df.groupby('n').mean()[[
       "sortTime", "seqRuntime", "p1Runtime", "p3Runtime", "p6Runtime", "seqCounter", "p1Counter", "p3Counter", "p6Counter", "bckCounterSeq", "bckCounterP1", "bckCounterP3", "bckCounterP6", "bridgeCounterSeq", "bridgeCounterP1", "bridgeCounterP3", "bridgeCounterP6"
    ]]

    return average_results

def plot_averages_PC(averages, fileName):
    n_values = averages.index

    scale = np.array([n + np.log2(np.log2(n)) for n in n_values])

    plt.plot(n_values, scale, label=r"$n + \log^2(n)$", linestyle='--', color='black')
    
    plt.plot(scale, averages['seqCounter'], label='Sequential', color='blue', linestyle='--', zorder=2, dashes=(4,3))
    plt.plot(scale, averages['p1Counter'], label='P1', color='orange', zorder=1)
    plt.plot(scale, averages['p3Counter'], label='P3', color='green', zorder=1)
    plt.plot(scale, averages['p6Counter'], label='P6', color='red', zorder=1)
    
    plt.xlabel(r"$n + \log^2(n)$")
    plt.ylabel("Average Program Counter Value")
    plt.title("Average Program Counter Value vs " + r"$n + \log^2(n)$" + " for Different p")
    plt.legend()
    plt.grid(True)

    plt.savefig(f'cg/{fileName}/{fileName}_PC.png')
    plt.close()

def plot_averages_RT(averages, fileName):
    n_values = averages.index

    scale = np.array([n + np.log2(np.log2(n)) for n in n_values])

    plt.plot(n_values, scale, label=r"$n + \log^2(n)$", linestyle='--', color='black')
    
    plt.plot(scale, averages['seqRuntime'], label='Sequential', color='blue', linestyle='--', zorder=2, dashes=(4,3))
    plt.plot(scale, averages['p1Runtime'], label='P1', color='orange', zorder=1)
    plt.plot(scale, averages['p3Runtime'], label='P3', color='green', zorder=1)
    plt.plot(scale, averages['p6Runtime'], label='P6', color='red', zorder=1)
    
    plt.xlabel(r"$n + \log^2(n)$")
    plt.ylabel("Average Program Runtime (ns)")
    plt.title("Average Program Runtime (ns) vs " + r"$n + \log^2(n)$" + " for Different p")
    plt.legend()
    plt.grid(True)

    plt.savefig(f'cg/{fileName}/{fileName}_RT.png')
    plt.close()

fileNames = ["square", "circle", "curve"]
for fileName in fileNames:
    averages = compute_averages(f'cg/{fileName}/{fileName}.csv')
    plot_averages_PC(averages, fileName)
    plot_averages_RT(averages, fileName)
