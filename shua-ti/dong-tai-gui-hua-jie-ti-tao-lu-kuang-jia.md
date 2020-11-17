# 动态规划解题套路框架

## 来源

{% embed url="https://labuladong.gitbook.io/algo/di-ling-zhang-bi-du-xi-lie/dong-tai-gui-hua-xiang-jie-jin-jie" %}

## 什么问题适合用动态规划算法求解



## 解题思路

**明确 base case -&gt; 明确「状态」-&gt; 明确「选择」 -&gt; 定义 dp 数组/函数的含义**。

## 代码框架

```python
# 初始化 base case
dp[0][0][...] = base
# 进行状态转移
for 状态1 in 状态1的所有取值：
    for 状态2 in 状态2的所有取值：
        for ...
            dp[状态1][状态2][...] = 求最值(选择1，选择2...)
```

## 相关题目

* [斐波那契数列](https://leetcode-cn.com/problems/fibonacci-number/)
* [零钱兑换](https://leetcode-cn.com/problems/coin-change/)

