__author__ = 'feng'

import string
import random
import time
import numpy as np
import matplotlib.pyplot as plt


def binary_search(arr, key):
    lo, hi = 0, (len(arr) - 1)
    while lo <= hi:
        mid = (lo + hi) >> 1
        k = arr[mid]
        if k == key:
            return mid
        elif k < key:
            lo = mid + 1 # search upper subarray
        else:
            hi = mid - 1 # search lower subarray
            # (-(insertion point) - 1)
    return -(lo + 1)


def random_string(size, chars=string.digits + string.ascii_letters):
    return ''.join(random.choice(chars) for x in range(size))


COUNT = 1000 * 200
KEY_COUNTS = list(range(2000, COUNT + 1, COUNT / 10))
LOOKUP = 1500


def compare_dict_bs(data, name, array_type=None):
    all_keys = data.keys()
    items = data.items()

    rets = []
    for count in KEY_COUNTS:
        map = dict(items[:count])
        keys_arr = sorted(map.keys())

        lookup_keys = [random.choice(keys_arr) for i in range(LOOKUP / 2)] + \
                      [random.choice(all_keys) for i in range(LOOKUP / 2)]

        if array_type:
            import array

            keys_arr = array.array(array_type, keys_arr)

        map_counter = 0
        start = time.time()
        for key in lookup_keys:
            if map.get(key, None):
                map_counter += 1
        map_time = time.time() - start

        bs_counter = 0
        start = time.time()
        for key in lookup_keys:
            if binary_search(keys_arr, key) >= 0:
                bs_counter += 1

        bs_time = time.time() - start

        if bs_counter != map_counter:
            raise Exception("hash map and binary search's result conflict")

        print name, count, map_time, bs_time
        rets.append((count, map_time * 1000, bs_time * 1000))
    return rets


def plot_result(numbers, name, output):
    ind = np.arange(len(numbers))
    width = 0.35

    fig = plt.figure(figsize=(7, 5))
    ax = fig.add_subplot(111)

    rects1 = ax.bar(ind, [mt for c, mt, bt in numbers], width, color='r')
    binary_times = [bt for c, mt, bt in numbers]
    rects2 = ax.bar(ind + width, binary_times, width, color='y')

    ax.legend((rects1[0], rects2[0]), ('dict', "binary-search"))

    ax.set_ybound(upper=max(binary_times) * 1.3)
    ax.set_ylabel('time (ms)')
    ax.set_xlabel('items count')
    ax.set_title('%s: dict vs binary-search, %s lookups' % (name, LOOKUP))
    ax.set_xticks(ind + width)
    ax.set_xticklabels(["%dk" % (c / 1000) for c, mt, bt in numbers])

    if output:
        plt.savefig(output)
    else:
        plt.show()


def main():
    key_str = {}
    for i in range(COUNT):
        key_str[random_string(random.randint(2, 15))] = True

    results = compare_dict_bs(key_str, "string")
    plot_result(results, 'string', "/tmp/string.png")

    key_int = {}
    for i in range(COUNT):
        key_int[random.randint(0, 1 << 32)] = True
    results = compare_dict_bs(key_int, "int(array)", array_type='I')
    plot_result(results, 'int(array)', "/tmp/int_array.png")

    results = compare_dict_bs(key_int, "int")
    plot_result(results, 'int', "/tmp/int.png")

if __name__ == "__main__":
    main()

