#!/usr/bin/env python3

import numpy

candidates = []
with open('input.txt', 'r') as fp:
    rows = fp.read()
    rows = rows.splitlines()
    for row in rows:
        items = rows.copy()
        items.remove(row)
        for item in items:
            sum = int(item) + int(row)
            if sum == 2020:
                candidates.append(int(item))
                candidates.append(int(row))
        rows.remove(row)

print('Candidates: %s' % candidates)
print('Solution: %s' % numpy.prod(candidates))

