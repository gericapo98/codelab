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

