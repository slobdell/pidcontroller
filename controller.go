package pidcontroller

type LinearPIDController struct {
	p                 *float64
	i                 *float64
	d                 *float64
	previousTimestamp *float64
	previousError     *float64
	integral          float64
}

// P, I, and D are represented as pointers so that these values can be externally tuned
func NewLinearPIDController(p, i, d *float64) *LinearPIDController {
	if p == nil || i == nil || d == nil {
		panic("Input values must be non nil")
	}

	return &LinearPIDController{
		p:                 p,
		i:                 i,
		d:                 d,
		previousTimestamp: nil,
		integral:          0,
	}
}

func (pid *LinearPIDController) GetOutput(timestamp, currentValue, targetValue float64) float64 {
	currentError := targetValue - currentValue

	deltaT := 0.0
	derivative := 0.0
	if pid.previousTimestamp != nil {
		deltaT = timestamp - *pid.previousTimestamp
		if deltaT == 0.0 {
			panic("No differential between previous timestamp and current")
		}
		pid.integral += currentError * deltaT
		derivative = (currentError - *pid.previousError) / deltaT
	} else {
		pid.previousError = new(float64)
		pid.previousTimestamp = new(float64)
	}

	*pid.previousError = currentError
	*pid.previousTimestamp = timestamp
	return *pid.p*currentError + *pid.i*pid.integral + *pid.d*derivative
}
