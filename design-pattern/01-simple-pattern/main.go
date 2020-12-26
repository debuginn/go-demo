package _1_simple_pattern

import (
	"sync/atomic"
)

type Mutex struct {
	state int32
	sema  uint32
}

type RWMutex struct {
	w           Mutex  // 互斥锁解决多个 writer 的竞争
	writerSem   uint32 // writer 信号量
	readerSem   uint32 // reader 信号量
	readerCount int32  // reader 的数量，记录当前 reader 等待的数量
	readerWait  int32  // writer 等待完成的 reader的 数量
}

const rwmutexMaxReaders = 1 << 30

func (rw *RWMutex) RLock() {
	// 对 reader 计数 +1，readerCount 会出现负数
	// 1、没有 writer 竞争或者持有锁的时候，readerCount 充当计数器存在
	// 2、如果有 writer 竞争锁或者持有锁时，那么，readerCount 不仅仅承担着 reader 的计数功能，还能够标识当前是否有 writer 竞争或持有锁
	if atomic.AddInt32(&rw.readerCount, 1) < 0 {
		// rw.readerCount 是负值的时候，意味着此时有 writer 等待请求锁
		// 因为writer优先级高，所以把后来的 reader 阻塞休眠
		runtime_SemacquireMutex(&rw.readerSem, false, 0)
	}
}

func (rw *RWMutex) RUnlock() {
	// 将 reader 计数 -1
	if r := atomic.AddInt32(&rw.readerCount, -1); r < 0 {
		// 如果为 负数，代表着当前有 writer 在竞争锁，检查是不是所有的 reader 都将锁释放
		// 若释放了就让 writer 获取到锁进行写操作
		rw.rUnlockSlow(r) // 有等待的writer
	}
}

func (rw *RWMutex) rUnlockSlow(r int32) {
	// rUnlockSlow 将持有锁的 reader 计数 -1 的时候；
	// 会检查既有的 reader 是不是都已经释放了锁；
	// 如果都释放了锁，就会唤醒 writer，让 writer 持有锁。
	if atomic.AddInt32(&rw.readerWait, -1) == 0 {
		// 最后一个reader了，writer终于有机会获得锁了
		runtime_Semrelease(&rw.writerSem, false, 1)
	}
}

func runtime_Semrelease(u *uint32, b bool, i int) {

}

func runtime_SemacquireMutex(u *uint32, b bool, i int) {

}
