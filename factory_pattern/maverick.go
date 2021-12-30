package main

type maverick struct {
	gun
}

func newMaverick() iGun {
	return &maverick{
		gun: gun{
			name:  "maverick 5",
			power: 5,
		},
	}
}
