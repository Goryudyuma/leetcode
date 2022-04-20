package main

/*
 * @lc app=leetcode id=2241 lang=golang
 *
 * [2241] Design an ATM Machine
 */

// @lc code=start
type ATM struct {
	coin20  int
	coin50  int
	coin100 int
	coin200 int
	coin500 int
}

func Constructor() ATM {
	return ATM{}
}

func (this *ATM) Deposit(banknotesCount []int) {
	this.coin20 += banknotesCount[0]
	this.coin50 += banknotesCount[1]
	this.coin100 += banknotesCount[2]
	this.coin200 += banknotesCount[3]
	this.coin500 += banknotesCount[4]
}

func (this *ATM) Withdraw(amount int) []int {
	f := true
	coin500 := amount / 500
	if this.coin500 < coin500 {
		coin500 = this.coin500
	}
	this.coin500 -= coin500
	defer func() {
		if !f {
			this.coin500 += coin500
		}
	}()
	amount -= coin500 * 500

	coin200 := amount / 200
	if this.coin200 < coin200 {
		coin200 = this.coin200
	}
	this.coin200 -= coin200
	defer func() {
		if !f {
			this.coin200 += coin200
		}
	}()
	amount -= coin200 * 200

	coin100 := amount / 100
	if this.coin100 < coin100 {
		coin100 = this.coin100
	}
	this.coin100 -= coin100
	defer func() {
		if !f {
			this.coin100 += coin100
		}
	}()
	amount -= coin100 * 100

	coin50 := amount / 50
	if this.coin50 < coin50 {
		coin50 = this.coin50
	}
	this.coin50 -= coin50
	defer func() {
		if !f {
			this.coin50 += coin50
		}
	}()
	amount -= coin50 * 50

	coin20 := amount / 20
	if this.coin20 < coin20 {
		coin20 = this.coin20
	}
	this.coin20 -= coin20
	defer func() {
		if !f {
			this.coin20 += coin20
		}
	}()
	amount -= coin20 * 20

	if amount != 0 {
		f = false
		return []int{-1}
	}
	return []int{coin20, coin50, coin100, coin200, coin500}
}

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */
// @lc code=end
