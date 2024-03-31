# stanford-algorithms-go

### Karatsuba Algorithm

## formula without Karatsuba Algorithm

| a formula to multiple two string inputted digits efficiently

- x and y are passed
- x equals to a \* 10 \*\* n/2 + b
- y equals to c \* 10 \*\* n/2 + d
- x and y are sometimes odd, sometimes even
- x and y are determined a random string digits.
- it sometimes does not match total length. its half of total length is called n.
- a first half of x length(if it is an odd, x length + 1 ) times 10 square n/2 equals to A, if it is an even.
- a second half of x length minus A length equals to B.
- a first half of y times 10 square n/2 equals to C
- a second half of y equals to D
- multiplying result has obtained by progressing the following step

```
1 (10 ** 2/n * a) * c
2 b * d
3 ((10 ** 2/n *a) + b) \* ((10 ** 2/n *c) + d)
4 3 - 2 - 1

x * y
= 1 + 2 + 4
= ac * 10 ** n + (10 ** n/2) * (ad + bc) + bd ← here is the star algorithm

```

## base

return the input digits if digits are equal to 1

## formla with Karatsuba Algorithm

```
ac * 10 ** n + (10 ** n/2) * (ad + bc) + bd

1 recursively compute ac
2 recursively compute bd
3 recursively compute (a + b)(c + d) ac ad bc bd
4 3 - 1 - 2 = ad + bc
5 10 ** n * 4 + 2 + 10 ** n/2 * 1
5 ac * 10 ** n + 10 ** n/2 * (ad + bc) + bd
6720000
2652
284000


```
