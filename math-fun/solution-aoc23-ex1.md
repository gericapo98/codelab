In case you don't like my formalization:  
<br>
<img src="https://media1.tenor.com/m/UXp3rdEKVeAAAAAd/mai-sakurajima.gif" alt="Mai Sakurajima" width="300" />
<br>

# AoC-2023-Ex1: Calibration Value Extraction

Determine the sum of "calibration values" for each line in a dataset where calibration values are derived from the first and last **natural numbers** in each string.

---

## Dataset Definition
Given an indexed dataset \( D \) of strings:
$D = \{s_1, s_2, \ldots, s_n\}$

Each string s contains a mix of characters and **natural numbers** $\mathbb{N} = \{1, 2, 3, \ldots\}$.  

We write:

$s_1 = (a_1, a_2, a_3, \ldots)$
and so on.

---

To extract all natural numbers from a string s, we define $\forall a_i \in s$ the function $\delta(a)$:
$\delta(a) = \begin{cases}
1 & \text{if } a \in \mathbb{N}, \\
0 & \text{otherwise.}
\end{cases}$

Using $\delta(a)$, construct a sequence of all digits in s :
$(a)_s = \{a_i \mid \delta(a_i) = 1, a_i \in s\}$

For a string s: \
   Extract the **minimum** and **maximum** values from this sequence:  
   $x = \min((a)), \quad y = \max((a))$

Combine x and y to form the calibration value:
   $v(s) = 10x + y$

---

Sum all v(s) for $s \in D$: 

$z = \sum_{s \in D} v(s).$



# AoC-2023-Ex1: Calibration Value Extraction

## Problem Statement
Determine the sum of "calibration values" for each line in a dataset \( D \), where calibration values are derived from the smallest and largest **natural numbers** present in each string.

---

## Definitions

### Dataset Representation
Given a dataset \( D \) of strings:
\[
D = \{s_1, s_2, \dots, s_n\}
\]

Each string \( s \) contains a mix of **natural numbers** and non-numeric characters.

---

### Characteristic Function
To identify digits in \( s \), we define the characteristic function:
\[
\delta(c) =
\begin{cases}
1 & \text{if } c \text{ is a digit (0-9)} \\
0 & \text{otherwise.}
\end{cases}
\]

---

### Extraction of Numbers
From the characteristic function \( \delta(c) \), we define a sequence \( T(s) \), consisting of all natural numbers in \( s \):
\[
T(s) = \{n_1, n_2, \dots, n_k\}, \quad \text{where } n_i \in \mathbb{N}.
\]

---

### Calibration Value
For each string \( s \):
1. Extract the smallest (\( x \)) and largest (\( y \)) numbers:
   \[
   x = \min T(s), \quad y = \max T(s).
   \]
2. Combine to form:
   \[
   v(s) = 10x + y.
   \]

---

### Total Calibration Value
The sum of all calibration values for \( s \in D \) is:
\[
z = \sum_{s \in D} v(s).
\]

---
