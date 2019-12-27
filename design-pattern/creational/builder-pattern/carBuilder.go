package builder_pattern


type CarBuilder interface {
	TopSpeed(int) CarBuilder
	Paint(string) CarBuilder
	Type(string) CarBuilder
	Build() Car
}

type carBuilder struct {
	speedOption int
	color       string
	carType string
}

func (cb *carBuilder) Type(carType string) CarBuilder {
	cb.carType=carType
	return cb
}

func (cb *carBuilder) TopSpeed(speed int) CarBuilder {
	cb.speedOption = speed
	return cb
}

func (cb *carBuilder) Paint(color string) CarBuilder {
	cb.color = color
	return cb
}

func (cb *carBuilder) Build() Car {
	return &car{
		topSpeed: cb.speedOption,
		color:    cb.color,
		carType: cb.carType,
	}
}

func New() CarBuilder {
	return &carBuilder{}
}


