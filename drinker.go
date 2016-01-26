package main

import (
	"fmt"
	"os"
)

// stateFn represents the state of an alcoholic as a function that returns the next state
type stateFn func(d *drinker) stateFn

// drinker holds the state of an alcoholic.
type drinker struct {
	money  int
	cap    int
	empty  int
	bottle int
	drank  int
	state  stateFn
}

func main() {

	var ringgit int

	fmt.Print("Berapa Ringgit ada? (RM): ")
	_, err := fmt.Scanf("%d", &ringgit)

	if err != nil {
		fmt.Fprintf(os.Stderr, "drinker: %v", err)
	}

	fmt.Printf("%+v\n", alcoholic(ringgit).start())
}

// alcohlic creates a new alcoholic, with money
func alcoholic(money int) *drinker {
	return &drinker{
		money: money,
	}
}

// start runs the state machine for the drinker.
func (d *drinker) start() *drinker {
	for d.state = buy; d.state != nil; {
		d.state = d.state(d)
	}

	return d
}

// yamSheng to drink all of the available bottle.
func yamSheng(d *drinker) stateFn {
	d.cap += d.bottle
	d.empty += d.bottle
	d.drank += d.bottle
	d.bottle = 0

	return buy(d)
}

// buy bottle with money, cap, empty bottle.
func buy(d *drinker) stateFn {
	switch {
	case d.money%2 == 0 && d.money != 0:
		return buyWithMoney(d)
	case d.cap >= 4 && d.cap != 0:
		return buyWithCap(d)
	case (d.empty >= 2) && d.empty != 0:
		return buyWithEmptyBottle(d)
	default:
		return nil
	}
}

// buyWithMoney exchanges RM2 with 1 bottle.
func buyWithMoney(d *drinker) stateFn {
	d.money -= 2
	d.bottle++
	return yamSheng(d)
}

// buyWithCap exchanges 4 caps with 1 bottle.
func buyWithCap(d *drinker) stateFn {
	d.cap -= 4
	d.bottle++
	return yamSheng(d)
}

// buyWithEmptyBottle exchanges 2 empty bottle with 1 bottle.
func buyWithEmptyBottle(d *drinker) stateFn {
	d.empty -= 2
	d.bottle++
	return yamSheng(d)
}
