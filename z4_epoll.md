### 同步IO

> 进程也可以换成线程

+ 阻塞IO (问一次 + 傻等)

用户进程process在Blocking IO读recvfrom操作的两个阶段都是等待的。在数据没准备好的时候，process原地等待kernel准备数据。kernel准备好数据后，process继续等待kernel将数据copy到自己的buffer。在kernel完成数据的copy后process才会从recvfrom系统调用中返回。

+ 非阻塞IO (不停的催问 + 傻等)

process在NonBlocking IO读recvfrom操作的第一个阶段是不会block等待的，如果kernel数据还没准备好，那么recvfrom会立刻返回一个EWOULDBLOCK错误。当kernel准备好数据后，进入处理的第二阶段的时候，process会等待kernel将数据copy到自己的buffer，在kernel完成数据的copy后process才会从recvfrom系统调用中返回。


+ 多路复用 (一个家长拦截 + 谁好了谁干活)

IO多路复用，就是我们熟知的select、poll、epoll模型。在IO多路复用的时候，process在两个处理阶段都是block住等待的。初看好像IO多路复用没什么用，其实select、poll、epoll的优势在于可以以较少的代价来同时监听处理多个IO。


### 异步IO

异步IO要求process在recvfrom操作的两个处理阶段上都不能等待，也就是process调用recvfrom后立刻返回，kernel自行去准备好数据并将数据从kernel的buffer中copy到process的buffer在通知process读操作完成了，然后process在去处理。遗憾的是，linux的网络IO中是不存在异步IO的，linux的网络IO处理的第二阶段总是阻塞等待数据copy完成的。真正意义上的网络异步IO是Windows下的IOCP（IO完成端口）模型。


很多时候，我们比较容易混淆non-blocking IO和asynchronous IO，认为是一样的。但是通过上图，几种IO模型的比较，会发现non-blocking IO和asynchronous IO的区别还是很明显的，non-blocking IO仅仅要求处理的第一阶段不block即可，而asynchronous IO要求两个阶段都不能block住。








### Select


int select(int nfds, fd_set *readfds, fd_set *writefds, fd_set *exceptfds, struct timeval *timeout);

select 函数监视的文件描述符分3类，分别是writefds、readfds、和exceptfds。调用后select函数会阻塞，直到有描述副就绪（有数据 可读、可写、或者有except），或者超时（timeout指定等待时间，如果立即返回设为null即可），函数返回。当select函数返回后，可以 通过遍历fdset，来找到就绪的描述符。



当用户process调用select的时候，select会将需要监控的readfds集合拷贝到内核空间（假设监控的仅仅是socket可读），然后遍历自己监控的socket sk，挨个调用sk的poll逻辑以便检查该sk是否有可读事件，遍历完所有的sk后，如果没有任何一个sk可读，那么select会调用schedule_timeout进入schedule循环，使得process进入睡眠。如果在timeout时间内某个sk上有数据可读了，或者等待timeout了，则调用select的process会被唤醒，接下来select就是遍历监控的sk集合，挨个收集可读事件并返回给用户了

1. 需要拷贝到内核
2. fds集合有限制 1024
3. 遍历效率低


### poll

int poll(struct pollfd *fds, nfds_t nfds, int timeout);


poll没有解决性能问题, 只是解决了 1024 限制问题
所以是个鸡肋


select遗留的三个问题中，问题(1)是用法限制问题，问题(2)和(3)则是性能问题。poll和select非常相似，poll并没着手解决性能问题，poll只是解决了select的问题(1)fds集合大小1024限制问题。下面是poll的函数原型，poll改变了fds集合的描述方式，使用了pollfd结构而不是select的fd_set结构，使得poll支持的fds集合限制远大于select的1024。poll虽然解决了fds集合大小1024的限制问题，但是，它并没改变大量描述符数组被整体复制于用户态和内核态的地址空间之间，以及个别描述符就绪触发整体描述符集合的遍历的低效问题。poll随着监控的socket集合的增加性能线性下降，poll不适合用于大并发场景。



### epoll

```
int epoll_create(int size)；//创建一个epoll的句柄，size用来告诉内核这个监听的数目一共有多大  新版本用红黑树,这个参数意义不大了

int epoll_ctl(int epfd, int op, int fd, struct epoll_event *event)；

函数是对指定描述符fd执行op操作。
- epfd：是epoll_create()的返回值。
- op：表示op操作，用三个宏来表示：添加EPOLL_CTL_ADD，删除EPOLL_CTL_DEL，修改EPOLL_CTL_MOD。分别添加、删除和修改对fd的监听事件。
- fd：是需要监听的fd（文件描述符）
- epoll_event：是告诉内核需要监听什么事，struct epoll_event结构如下：

struct epoll_event {
  __uint32_t events;  /* Epoll events */
  epoll_data_t data;  /* User data variable */
};

//events可以是以下几个宏的集合：
EPOLLIN ：表示对应的文件描述符可以读（包括对端SOCKET正常关闭）；
EPOLLOUT：表示对应的文件描述符可以写；
EPOLLPRI：表示对应的文件描述符有紧急的数据可读（这里应该表示有带外数据到来）；
EPOLLERR：表示对应的文件描述符发生错误；
EPOLLHUP：表示对应的文件描述符被挂断；
EPOLLET： 将EPOLL设为边缘触发(Edge Triggered)模式，这是相对于水平触发(Level Triggered)来说的。
EPOLLONESHOT：只监听一次事件，当监听完这次事件之后，如果还需要继续监听这个socket的话，需要再次把这个socket加入到EPOLL队列里




int epoll_wait(int epfd, struct epoll_event * events, int maxevents, int timeout);

等待epfd上的io事件，最多返回maxevents个事件。
参数events用来从内核得到事件的集合，maxevents告之内核这个events有多大，这个maxevents的值不能大于创建epoll_create()时的size，参数timeout是超时时间（毫秒，0会立即返回，-1将不确定，也有说法说是永久阻塞）。该函数返回需要处理的事件数目，如返回0表示已超时。
```

* 拷贝问题的解决

epoll引入了epoll_ctl系统调用，将高频调用的epoll_wait和低频的epoll_ctl隔离开。同时，epoll_ctl通过(EPOLL_CTL_ADD、EPOLL_CTL_MOD、EPOLL_CTL_DEL)三个操作来分散对需要监控的fds集合的修改，做到了有变化才变更，将select或poll高频、大块内存拷贝(集中处理)变成epoll_ctl的低频、小块内存的拷贝(分散处理)，避免了大量的内存拷贝。


同时，对于高频epoll_wait的可读就绪的fd集合返回的拷贝问题，epoll通过内核与用户空间mmap(内存映射)同一块内存来解决。mmap将用户空间的一块地址和内核空间的一块地址同时映射到相同的一块物理内存地址（不管是用户空间还是内核空间都是虚拟地址，最终要通过地址映射映射到物理地址），使得这块物理内存对内核和对用户均可见，减少用户态和内核态之间的数据交换。


* 循环问题解决

epoll引入了2个中间层，一个双向链表(ready_list)，一个单独的睡眠队列(single_epoll_wait_list)，

---



epoll巧妙的引入一个中间层解决了大量监控socket的无效遍历问题。细心的同学会发现，epoll在中间层上为每个监控的socket准备了一个单独的回调函数epoll_callback_sk，而对于select/poll，所有的socket都公用一个相同的回调函数。正是这个单独的回调epoll_callback_sk使得每个socket都能单独处理自身，当自己就绪的时候将自身socket挂入epoll的ready_list。

---

同时，epoll引入了一个睡眠队列single_epoll_wait_list，分割了两类睡眠等待。process不再睡眠在所有的socket的睡眠队列上，而是睡眠在epoll的睡眠队列上，在等待”任意一个socket可读就绪”事件。而中间wait_entry_sk则代替process睡眠在具体的socket上，当socket就绪的时候，它就可以处理自身了。



### epoll LT ET

　epoll对文件描述符的操作有两种模式：LT（level trigger）和ET（edge trigger）。LT模式是默认模式，LT模式与ET模式的区别如下：

　　LT模式：当epoll_wait检测到描述符事件发生并将此事件通知应用程序，应用程序可以不立即处理该事件。下次调用epoll_wait时，会再次响应应用程序并通知此事件。

　　ET模式：当epoll_wait检测到描述符事件发生并将此事件通知应用程序，应用程序必须立即处理该事件。如果不处理，下次调用epoll_wait时，不会再次响应应用程序并通知此事件。



1. LT模式
LT(level triggered)是缺省的工作方式，并且同时支持block和no-block socket.在这种做法中，内核告诉你一个文件描述符是否就绪了，然后你可以对这个就绪的fd进行IO操作。如果你不作任何操作，内核还是会继续通知你的。


2. ET模式
ET(edge-triggered)是高速工作方式，只支持no-block socket。在这种模式下，当描述符从未就绪变为就绪时，内核通过epoll告诉你。然后它会假设你知道文件描述符已经就绪，并且不会再为那个文件描述符发送更多的就绪通知，直到你做了某些操作导致那个文件描述符不再为就绪状态了(比如，你在发送，接收或者接收请求，或者发送接收的数据少于一定量时导致了一个EWOULDBLOCK 错误）。但是请注意，如果一直不对这个fd作IO操作(从而导致它再次变成未就绪)，内核不会发送更多的通知(only once)

ET模式在很大程度上减少了epoll事件被重复触发的次数，因此效率要比LT模式高。epoll工作在ET模式的时候，必须使用非阻塞套接口，以避免由于一个文件句柄的阻塞读/阻塞写操作把处理多个文件描述符的任务饿死。




当使用epoll的ET模型来工作时，当产生了一个EPOLLIN事件后，
读数据的时候需要考虑的是当recv()返回的大小如果等于请求的大小，那么很有可能是缓冲区还有数据未读完，也意味着该次事件还没有处理完，所以还需要再次读取：(阻塞其他的)




### 总结

（1）select，poll实现需要自己不断轮询所有fd集合，直到设备就绪，期间可能要睡眠和唤醒多次交替。而epoll其实也需要调用epoll_wait不断轮询就绪链表，期间也可能多次睡眠和唤醒交替，但是它是设备就绪时，调用回调函数，把就绪fd放入就绪链表中，并唤醒在epoll_wait中进入睡眠的进程。虽然都要睡眠和交替，但是select和poll在“醒着”的时候要遍历整个fd集合，而epoll在“醒着”的时候只要判断一下就绪链表是否为空就行了，这节省了大量的CPU时间。这就是回调机制带来的性能提升。

（2）select，poll每次调用都要把fd集合从用户态往内核态拷贝一次，并且要把current往设备等待队列中挂一次，而epoll只要一次拷贝，而且把current往等待队列上挂也只挂一次（在epoll_wait的开始，注意这里的等待队列并不是设备等待队列，只是一个epoll内部定义的等待队列）。这也能节省不少的开销。


### 心得

所以，如果处理的连接数不是很高的话，使用select/epoll的web server不一定比使用multi-threading + blocking IO的web server性能更好，可能延迟还更大。select/epoll的优势并不是对于单个连接能处理得更快，而是在于能处理更多的连接。）



如果没有大量的idle -connection或者dead-connection，epoll的效率并不会比select/poll高很多，但是当遇到大量的idle- connection，就会发现epoll的效率大大高于select/poll。



在IO multiplexing Model中，实际中，对于每一个socket，一般都设置成为non-blocking，但是，如上图所示，整个用户的process其实是一直被block的。只不过process是被select这个函数block，而不是被socket IO给block。



有人会说，non-blocking IO并没有被block啊。这里有个非常“狡猾”的地方，定义中所指的”IO operation”是指真实的IO操作，就是例子中的recvfrom这个system call。non-blocking IO在执行recvfrom这个system call的时候，如果kernel的数据没有准备好，这时候不会block进程。但是，当kernel中数据准备好的时候，recvfrom会将数据从kernel拷贝到用户内存中，这个时候进程是被block了，在这段时间内，进程是被block的。





