__author__ = 'feng'

import string
import random
import time
import numpy as np
import matplotlib.pyplot as plt
import subprocess

def main(path):
    p = subprocess.Popen("go run binary_hash.go", stdin=subprocess.PIPE,
                         stdout=subprocess.PIPE,shell=True)
    output = p.stdout.readlines()

    for line in output:
        print line.strip()

    lables = output[0].split()
    datas = []
    for line in output[1:]:
        datas.append([int(i) for i in line.split()])



    ind = np.arange(len(datas))
    width = 0.17

    fig = plt.figure(figsize=(8, 6))
    ax = fig.add_subplot(111)

    rects = []

    colors = ['b', 'g', 'r', 'c', 'm']

    upper = 0

    for i in range(len(lables) - 1):
        d = [line[i+1] / 1000 for line in datas]
        upper = max(upper, max(d))
        rect = ax.bar(ind + i * width, d, width, color=colors[i])
        rects.append(rect)

    ax.set_ybound(upper=upper*1.6)
    ax.set_ylabel("time (us)")
    ax.set_xlabel("items count")
    ax.set_xticks(ind + 3*width)
    ax.set_xticklabels(["%dk" % (line[0] / 1000) for line in datas])
    print ["%dk" % (line[0] / 1000) for line in datas]

    ax.legend([r[0] for r in rects], lables[1:])
    if output:
        plt.savefig(path)
    else:
        plt.show()

    # print datas, lables

if __name__ == "__main__":
    main("/tmp/string.png")
