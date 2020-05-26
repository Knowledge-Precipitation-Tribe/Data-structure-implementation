# 基本概念

## 定义

线性表是具有**相同特性**数据元素的一个**有限**序列。该序列中所含元素的个数叫做线性表的长度，用n\(n&gt;=0\)表示。注意，n可以等于零，表示线性表是一个空表。

线性表可以是有序的也可以是无序的。

![](../.gitbook/assets/image%20%2840%29.png)

与线性表对应的就是非线性表

![](../.gitbook/assets/image%20%2839%29.png)

## 逻辑特性

线性表只有一个表头元素，一个表尾元素，表头元素没有前驱，表尾元素没有后继，除表头和表尾元素之外，其他元素只有一个直接前驱，也只有一个直接后继。

## 存储结构

### 顺序表

顺序表就是把线性表中的所有元素按照其逻辑顺序，以此存储到从指定存储位置开始的一块连续的存储空间中。也就是数组存储。

### 链表

在链表存储中，每个结点不仅包含所存元素的信息，还包含元素之间逻辑关系的信息，如单链表中前驱结点包含后继结点的地址信息，这样就可以通过前驱结点中的地址信息找到后继结点的位置。

### 两种存储结构的比较

因为数组可以通过下标位置进行访问，所以数组有一个关键特性就是---随机访问。数组是如何实现随机访问的呢？

我们拿一个长度为 10 的 int 类型的数组 int\[\] a = new int\[10\]来举例。在我画的这个图中，计算机给数组 a\[10\]，分配了一块连续内存空间 1000～1039，其中，内存块的首地址为 base\_address = 1000。

![](../.gitbook/assets/image%20%2838%29.png)

我们知道，计算机会给每个内存单元分配一个地址，计算机通过地址来访问内存中的数据。当计算机需要随机访问数组中的某个元素时，它会首先通过下面的寻址公式，计算出该元素存储的内存地址：

```c
a[i]_address = base_address + i * data_type_size
```

其中 data\_type\_size 表示数组中每个元素的大小。我们举的这个例子里，数组中存储的是 int 类型数据，所以 data\_type\_size 就为 4 个字节。这个公式非常简单，我就不多做解释了。

这里我要特别纠正一个“错误”。我在面试的时候，常常会问数组和链表的区别，很多人都回答说，“链表适合插入、删除，时间复杂度 O\(1\)；数组适合查找，查找时间复杂度为 O\(1\)”。

实际上，这种表述是不准确的。数组是适合查找操作，但是查找的时间复杂度并不为 O\(1\)。即便是排好序的数组，你用二分查找，时间复杂度也是 O\(logn\)。所以，正确的表述应该是，数组支持随机访问，根据下标随机访问的时间复杂度为 O\(1\)。

因为链表是根据指针进行连接，所以不支持随机访问，而且链表结点中要包含指针，所以链表的存储空间利用率较顺序表稍低一点。

而且数组要求利用连续的存储空间，而链表的存储空间可以实现动态分配。

### 数组的插入与删除

数组为了保持内存数据的连续性，会导致插入、删除这两个操作比较低效。现在我们就来详细说一下，究竟为什么会导致低效？又有哪些改进方法呢？

#### 插入操作

假设数组的长度为 n，现在，如果我们需要将一个数据插入到数组中的第 k 个位置。为了把第 k 个位置腾出来，给新来的数据，我们需要将第 k～n 这部分的元素都顺序地往后挪一位。那插入操作的时间复杂度是多少呢？你可以自己先试着分析一下。

如果在数组的末尾插入元素，那就不需要移动数据了，这时的时间复杂度为 O\(1\)。但如果在数组的开头插入元素，那所有的数据都需要依次往后移动一位，所以最坏时间复杂度是 O\(n\)。 因为我们在每个位置插入元素的概率是一样的，所以平均情况时间复杂度为 \(1+2+…n\)/n=O\(n\)。

如果数组中的数据是有序的，我们在某个位置插入一个新的元素时，就必须按照刚才的方法搬移 k 之后的数据。但是，如果数组中存储的数据并没有任何规律，数组只是被当作一个存储数据的集合。在这种情况下，如果要将某个数据插入到第 k 个位置，**为了避免大规模的数据搬移，我们还有一个简单的办法就是，直接将第 k 位的数据搬移到数组元素的最后，把新的元素直接放入第 k 个位置。**

为了更好地理解，我们举一个例子。假设数组 a\[10\]中存储了如下 5 个元素：a，b，c，d，e。

我们现在需要将元素 x 插入到第 3 个位置。我们只需要将 c 放入到 a\[5\]，将 a\[2\]赋值为 x 即可。最后，数组中的元素如下： a，b，x，d，e，c。

![](../.gitbook/assets/image%20%2843%29.png)

利用这种处理技巧，在特定场景下，在第 k 个位置插入一个元素的时间复杂度就会降为 O\(1\)。这个处理思想在快排中也会用到，我会在排序那一节具体来讲，这里就说到这儿。

#### 删除操作

跟插入数据类似，如果我们要删除第 k 个位置的数据，为了内存的连续性，也需要搬移数据，不然中间就会出现空洞，内存就不连续了。

和插入类似，如果删除数组末尾的数据，则最好情况时间复杂度为 O\(1\)；如果删除开头的数据，则最坏情况时间复杂度为 O\(n\)；平均情况时间复杂度也为 O\(n\)。

**实际上，在某些特殊场景下，我们并不一定非得追求数组中数据的连续性。如果我们将多次删除操作集中在一起执行，删除的效率是不是会提高很多呢？**

我们继续来看例子。数组 a\[10\]中存储了 8 个元素：a，b，c，d，e，f，g，h。现在，我们要依次删除 a，b，c 三个元素。

![](../.gitbook/assets/image%20%2841%29.png)

为了避免 d，e，f，g，h 这几个数据会被搬移三次，我们可以先记录下已经删除的数据。**每次的删除操作并不是真正地搬移数据，只是记录数据已经被删除。当数组没有更多空间存储数据时，我们再触发执行一次真正的删除操作，这样就大大减少了删除操作导致的数据搬移。**

如果你了解 JVM，你会发现，**这不就是 JVM 标记清除垃圾回收算法的核心思想吗？**没错，数据结构和算法的魅力就在于此，很多时候我们并不是要去死记硬背某个数据结构或者算法，而是要学习它背后的思想和处理技巧，这些东西才是最有价值的。如果你细心留意，不管是在软件开发还是架构设计中，总能找到某些算法和数据结构的影子。

### 容器能否完全替代数组？

针对数组类型，很多语言都提供了容器类，比如 Java 中的 ArrayList、C++ STL 中的 vector。在项目开发中，什么时候适合用数组，什么时候适合用容器呢？

这里我拿 Java 语言来举例。如果你是 Java 工程师，几乎天天都在用 ArrayList，对它应该非常熟悉。那它与数组相比，到底有哪些优势呢？

我个人觉得，ArrayList 最大的优势就是可以将很多数组操作的细节封装起来。比如前面提到的数组插入、删除数据时需要搬移其他数据等。另外，它还有一个优势，就是支持动态扩容。

数组本身在定义的时候需要预先指定大小，因为需要分配连续的内存空间。如果我们申请了大小为 10 的数组，当第 11 个数据需要存储到数组中时，我们就需要重新分配一块更大的空间，将原来的数据复制过去，然后再将新的数据插入。

如果使用 ArrayList，我们就完全不需要关心底层的扩容逻辑，ArrayList 已经帮我们实现好了。每次存储空间不够的时候，它都会将空间自动扩容为 1.5 倍大小。

不过，这里需要注意一点，因为扩容操作涉及内存申请和数据搬移，是比较耗时的。所以，如果事先能确定需要存储的数据大小，最好在创建 ArrayList 的时候事先指定数据大小。

比如我们要从数据库中取出 10000 条数据放入 ArrayList。我们看下面这几行代码，你会发现，相比之下，事先指定数据大小可以省掉很多次内存申请和数据搬移操作。

```c
ArrayList<User> users = new ArrayList(10000);
for (int i = 0; i < 10000; ++i) {
  users.add(xxx);
}
```
