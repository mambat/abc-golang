package popcount

// pc[i] is the population count of i.
//var pc [256]byte
// 初始化包变量方式一：使用 init 初始化函数（可以有多个），init 其不能被调用或引用外
//func init() {
//    for i := range pc {
//        pc[i] = pc[i/2] + byte(i&1)
//    }
//}

// 初始化包变量方式二：匿名函数
// pc[i] is the population count of i.
var pc [256]byte = func() (pc [256]byte) {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
    return
}()

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
    return int(pc[byte(x>>(0*8))] +
        pc[byte(x>>(1*8))] +
        pc[byte(x>>(2*8))] +
        pc[byte(x>>(3*8))] +
        pc[byte(x>>(4*8))] +
        pc[byte(x>>(5*8))] +
        pc[byte(x>>(6*8))] +
        pc[byte(x>>(7*8))])
}
