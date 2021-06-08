package main

type machineGun struct {
	On    bool
	Ammo  int
	Power int
}

func (S *machineGun) Shoot() bool {
	if S.On == false {
		return false
	} else if S.Ammo != 0 {
		S.Ammo--
		return true
	} else {
		return false
	}
}

func (RB *machineGun) RideBike() bool {
	if RB.On == false {
		return false
	} else if RB.Power != 0 {
		RB.Power--
		return true
	} else {
		return false
	}
}

func main() {
	mg := new(machineGun)

}
