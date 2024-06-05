package sq

/*
设计一个自助结账系统，该系统需要通过一个队列来模拟顾客通过购物车的结算过程，需要实现的功能有：

get_max()：获取结算商品中的最高价格，如果队列为空，则返回 -1
add(value)：将价格为 value 的商品加入待结算商品队列的尾部
remove()：移除第一个待结算的商品价格，如果队列为空，则返回 -1
注意，为保证该系统运转高效性，以上函数的均摊时间复杂度均为 O(1)
*/

// Checkout 也是用降序队列实现
type Checkout struct {
	q        []int
	descendQ []int
}

func Constructor3() Checkout {
	return Checkout{
		q:        make([]int, 0),
		descendQ: make([]int, 0),
	}
}

func (checkout *Checkout) Get_max() int {
	if len(checkout.descendQ) == 0 {
		return -1
	}

	return checkout.descendQ[0]
}

func (checkout *Checkout) Add(value int) {
	checkout.q = append(checkout.q, value)
	for index, tempValue := range checkout.descendQ {
		if value > tempValue {
			checkout.descendQ = checkout.descendQ[:index]
			checkout.descendQ = append(checkout.descendQ, value)
			return
		}
	}

	checkout.descendQ = append(checkout.descendQ, value)
}

func (checkout *Checkout) Remove() int {
	if len(checkout.q) == 0 {
		return -1
	}

	val := checkout.q[0]
	checkout.q = checkout.q[1:]
	if val == checkout.descendQ[0] {
		checkout.descendQ = checkout.descendQ[1:]
	}
	return val

}
