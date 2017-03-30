package pidcontroller_test

import (
	"pidcontroller"
	"testing"
)

func TestP(t *testing.T) {
	p := 1.0
	i := 0.0
	d := 0.0
	var expected float64
	var output float64
	controller := pidcontroller.NewLinearPIDController(&p, &i, &d)

	output = controller.GetOutput(1.0, 0.0, 5.0)
	expected = 5.0
	if output != expected {
		t.Error("Values not equal", output, expected)
	}

	output = controller.GetOutput(2.0, 0.0, 5.0)
	expected = 5.0
	if output != expected {
		t.Error("Values not equal", output, expected)
	}
}

func TestPD(t *testing.T) {
	p := 1.0
	i := 0.0
	d := 1.0
	var expected float64
	var output float64
	controller := pidcontroller.NewLinearPIDController(&p, &i, &d)

	output = controller.GetOutput(1.0, 0.0, 5.0)
	expected = 5.0
	if output != expected {
		t.Error("Values not equal", output, expected)
	}

	output = controller.GetOutput(2.0, 0.0, 5.0)
	expected = 5.0
	if output != expected {
		t.Error("Values not equal", output, expected)
	}

	output = controller.GetOutput(3.0, 1.0, 5.0)
	expected = 3.0
	if output != expected {
		t.Error("Values not equal", output, expected)
	}
}

func TestPI(t *testing.T) {
	p := 1.0
	i := 1.0
	d := 0.0
	var expected float64
	var output float64
	controller := pidcontroller.NewLinearPIDController(&p, &i, &d)

	output = controller.GetOutput(1.0, 0.0, 5.0)
	expected = 5.0
	if output != expected {
		t.Error("Values not equal", output, expected)
	}

	output = controller.GetOutput(2.0, 0.0, 5.0)
	expected = 10.0
	if output != expected {
		t.Error("Values not equal", output, expected)
	}

	output = controller.GetOutput(3.0, 0.0, 5.0)
	expected = 15.0
	if output != expected {
		t.Error("Values not equal", output, expected)
	}
}

func TestPID(t *testing.T) {
	p := 1.0
	i := 1.0
	d := 1.0
	var expected float64
	var output float64
	controller := pidcontroller.NewLinearPIDController(&p, &i, &d)

	output = controller.GetOutput(1.0, 0.0, 5.0)
	expected = 5.0
	if output != expected {
		t.Error("Values not equal", output, expected)
	}

	output = controller.GetOutput(2.0, 0.0, 5.0)
	expected = 10.0
	if output != expected {
		t.Error("Values not equal", output, expected)
	}

	output = controller.GetOutput(3.0, 1.0, 5.0)
	expected = 12.0
	if output != expected {
		t.Error("Values not equal", output, expected)
	}
}

func TestGlobalTuning(t *testing.T) {
	p := 1.0
	i := 1.0
	d := 1.0
	var expected float64
	var output float64
	controller := pidcontroller.NewLinearPIDController(&p, &i, &d)

	output = controller.GetOutput(1.0, 0.0, 5.0)
	expected = 5.0
	if output != expected {
		t.Error("Values not equal", output, expected)
	}

	p = 2.0
	output = controller.GetOutput(2.0, 0.0, 5.0)
	expected = 15.0
	if output != expected {
		t.Error("Values not equal", output, expected)
	}
}
