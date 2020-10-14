# 基本概念

## 什么是图

图是一种非线性的数据结构，图中每个元素称其为顶点，顶点和顶点之间存在的关系称其为边。如下图所示。

![&#x56FE;](https://github.com/Knowledge-Precipitation-Tribe/Graph-neural-network/raw/master/images/graph.png)

图是一种数据结构，由顶点集![](https://cdn.nlark.com/yuque/__latex/5206560a306a2e085a437fd258eb57ce.svg)和边集![](https://cdn.nlark.com/yuque/__latex/3a3ea00cfc35332cedf6e5e9a32e94da.svg)组成,表示为![](https://cdn.nlark.com/yuque/__latex/9e9992d6bf50b7580f971487c466a8cb.svg)，其中![](https://cdn.nlark.com/yuque/__latex/e466bb2d1248188656a05c2218505259.svg)_代表每个顶点，_![](https://cdn.nlark.com/yuque/__latex/71b06be0d54453dd06c616e7cd552337.svg)代表每条边。

## 有向图与无向图

### 有向图

如果图中的边存在方向性，则称这样的边为有向边，包含有向边的图成为有向图。

![&#x6709;&#x5411;&#x56FE;](https://cdn.nlark.com/yuque/0/2020/png/194664/1594604164898-3a7968ca-3ff6-4fab-9397-de25f3adbc44.png)

### 无向图

无向图中的边都是没有方向的。

![&#x65E0;&#x5411;&#x56FE;](https://github.com/Knowledge-Precipitation-Tribe/Graph-neural-network/raw/master/images/graph.png)

## 加权图与非加权图

### 加权图

图中的每条边都有一个实数与之相对应，称这样的图为加权图。在世纪场景中，权重可以代表距离或者两人之间的相似程度等，一般情况下认为各边上的权重代表定点之间的连接强度。

![image.png](https://cdn.nlark.com/yuque/0/2020/png/194664/1594604352104-7ac15eb2-31a3-4c2a-ae79-098686b9dbb3.png)

### 非加权图

非加权图与加权图相对应，非加权图认为各边上的权重是一样的。

![](../.gitbook/assets/image%20%28163%29.png)

## 连通图与非连通图

### 连通图

如果图中不存在任何孤立的结点，称这样的图为连通图。

### 非连通图

如果图中存在孤立的结点，称这样的图为非连通图。

## 二部图/二分图

![image.png](https://cdn.nlark.com/yuque/0/2020/png/194664/1594604423911-465d3051-8671-4e1b-94a9-836b9b417016.png)

二部图是一类特殊的图，设![](https://cdn.nlark.com/yuque/__latex/9e9992d6bf50b7580f971487c466a8cb.svg)是一个无向图，如果顶点![](https://cdn.nlark.com/yuque/__latex/5206560a306a2e085a437fd258eb57ce.svg)可分割为两个互不相交的子集![](https://cdn.nlark.com/yuque/__latex/31da6f87f19d9cd2264061a0afc2cbb1.svg)，并且图中的任意一条边![](https://cdn.nlark.com/yuque/__latex/6fd64a8eafc5224488e3523dd225bb7b.svg)均有![](https://cdn.nlark.com/yuque/__latex/2446e766815d367c2319d598aca88ff8.svg)，![](https://cdn.nlark.com/yuque/__latex/df8285ef215e525fab98788a7d060ad8.svg)或者![](https://cdn.nlark.com/yuque/__latex/eccc16d404a5648128dda8dfc0e53a0e.svg)，![](https://cdn.nlark.com/yuque/__latex/14ec1b13d65a53a5919f2d43d22a4a64.svg)，则称图![](https://cdn.nlark.com/yuque/__latex/dfcf28d0734569a6a693bc8194de62bf.svg)为一个二分图。二部图是一种十分常见的图数据结构，描述两类对象之间的交互关系，比如：用户和商品，作者与论文

## 邻居

如果存在一条边连接定点和，则称是的邻居，反之亦然。

## 顶点的度

与顶点![](https://cdn.nlark.com/yuque/__latex/5206560a306a2e085a437fd258eb57ce.svg)相关的边的条数称为顶点![](https://cdn.nlark.com/yuque/__latex/5206560a306a2e085a437fd258eb57ce.svg)的度。

## 图的表现形式

### 邻接矩阵

邻接矩阵是表示顶点之间相邻关系的矩阵。设![](https://cdn.nlark.com/yuque/__latex/9e9992d6bf50b7580f971487c466a8cb.svg)是具有![](https://cdn.nlark.com/yuque/__latex/7b8b965ad4bca0e41ab51de7b31363a1.svg)个顶点的图，顶点序号依次为![](https://cdn.nlark.com/yuque/__latex/570958b11ab796665e13fe079316779b.svg),则![](https://cdn.nlark.com/yuque/__latex/dfcf28d0734569a6a693bc8194de62bf.svg)的邻接矩阵是具有如下定义的![](https://cdn.nlark.com/yuque/__latex/7b8b965ad4bca0e41ab51de7b31363a1.svg)阶方阵![](https://cdn.nlark.com/yuque/__latex/7fc56270e7a70fa81a5935b72eacbe29.svg)：

* ![](https://cdn.nlark.com/yuque/__latex/5f2c14e7b2856bf0fe1356c34fe9cc5d.svg)表示顶点![](https://cdn.nlark.com/yuque/__latex/865c0c0b4ab0e063e5caa3387c1a8741.svg)与顶点![](https://cdn.nlark.com/yuque/__latex/363b122c528f54df4a0446b6bab05515.svg)邻接，即![](https://cdn.nlark.com/yuque/__latex/865c0c0b4ab0e063e5caa3387c1a8741.svg)与![](https://cdn.nlark.com/yuque/__latex/363b122c528f54df4a0446b6bab05515.svg)之间存在边
* ![](https://cdn.nlark.com/yuque/__latex/9d68bf4460774c527622c2880f3273b4.svg)表示顶点![](https://cdn.nlark.com/yuque/__latex/865c0c0b4ab0e063e5caa3387c1a8741.svg)与顶点![](https://cdn.nlark.com/yuque/__latex/363b122c528f54df4a0446b6bab05515.svg)不邻接，即![](https://cdn.nlark.com/yuque/__latex/865c0c0b4ab0e063e5caa3387c1a8741.svg)与![](https://cdn.nlark.com/yuque/__latex/363b122c528f54df4a0446b6bab05515.svg)之间没有边

![&#x90BB;&#x63A5;&#x77E9;&#x9635;](https://github.com/Knowledge-Precipitation-Tribe/Graph-neural-network/raw/master/images/adjMatrix.png)

### 邻接表

![](../.gitbook/assets/image%20%28162%29.png)

### 度矩阵

度矩阵其中包含的信息为的每一个顶点的度，度矩阵![](https://cdn.nlark.com/yuque/__latex/f623e75af30e62bbd73d6df5b50bb7b5.svg):

![&#x5EA6;&#x77E9;&#x9635;](https://github.com/Knowledge-Precipitation-Tribe/Graph-neural-network/raw/master/images/degreeMatrix.png)

### 拉普拉斯矩阵

给定一个有![](https://cdn.nlark.com/yuque/__latex/7b8b965ad4bca0e41ab51de7b31363a1.svg)个顶点的图![](https://cdn.nlark.com/yuque/__latex/9e9992d6bf50b7580f971487c466a8cb.svg)，其拉普拉斯矩阵被定义为![](https://cdn.nlark.com/yuque/__latex/e4ade126cf9eebe72b791053b0fdcdd8.svg)，![](https://cdn.nlark.com/yuque/__latex/f623e75af30e62bbd73d6df5b50bb7b5.svg)其中为图的度矩阵，![](https://cdn.nlark.com/yuque/__latex/7fc56270e7a70fa81a5935b72eacbe29.svg)为图的邻接矩阵，拉普拉斯矩阵![](https://cdn.nlark.com/yuque/__latex/d20caec3b48a1eef164cb4ca81ba2587.svg):

![&#x62C9;&#x666E;&#x62C9;&#x65AF;&#x77E9;&#x9635;](https://github.com/Knowledge-Precipitation-Tribe/Graph-neural-network/raw/master/images/laplaMatrix.png)

### 关联矩阵

![&#x5173;&#x8054;&#x77E9;&#x9635;](https://cdn.nlark.com/yuque/0/2020/png/194664/1594605467401-22263018-1637-4124-9b8f-655faeec3d8b.png)

关联矩阵![](https://cdn.nlark.com/yuque/__latex/9d5ed678fe57bcca610140957afab571.svg)的定义如下

![](https://cdn.nlark.com/yuque/__latex/3db52bb227d9a56fbe4804c6218df1f4.svg)

用关联矩阵存储图时，我们需要两个一维数组分别表示定点集合和边集合，需要一个二维数组表示关联矩阵。

