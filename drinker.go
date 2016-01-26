package main

import "fmt"

type stateFn func(d *drinker) stateFn

type drinker struct {
	money  int
	cap    int
	empty  int
	bottle int
	drink  int
	state  stateFn
}

func main() {
	tan := Drinker(10)
	fmt.Printf("%+v\n", tan.start())
}

func Drinker(money int) *drinker {
	return &drinker{
		money: money,
	}
}

func (d *drinker) start() *drinker {
	for d.state = Buy; d.state != nil; {
		d.state = d.state(d)
	}

	return d
}

func Drink(d *drinker) stateFn {
	d.cap += d.bottle
	d.empty += d.bottle
	d.drink += d.bottle
	d.bottle = 0

	return Buy(d)
}

func Buy(d *drinker) stateFn {
	switch {
	case d.money%2 == 0 && d.money != 0:
		return BuyWithMoney(d)
	case d.cap >= 4 && d.cap != 0:
		return BuyWithCap(d)
	case (d.empty >= 2) && d.empty != 0:
		return BuyWithEmptyBottle(d)
	default:
		return nil
	}
}

func BuyWithMoney(d *drinker) stateFn {
	d.bottle += 1
	d.money -= 2
	return Drink(d)
}

func BuyWithCap(d *drinker) stateFn {
	d.bottle += 1
	d.cap -= 4
	return Drink(d)
}

func BuyWithEmptyBottle(d *drinker) stateFn {
	d.bottle += 1
	d.empty -= 2
	return Drink(d)
}
