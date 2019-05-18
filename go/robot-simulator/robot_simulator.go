package robot

import "fmt"

const (
	N Dir = 0
	E Dir = 90
	S Dir = 180
	W Dir = 270
)

func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case E:
		Step1Robot.X++
	case S:
		Step1Robot.Y--
	case W:
		Step1Robot.X--
	}
}

func Left() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = W
	case E:
		Step1Robot.Dir = N
	case S:
		Step1Robot.Dir = E
	case W:
		Step1Robot.Dir = S
	}
}

func Right() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Dir = E
	case E:
		Step1Robot.Dir = S
	case S:
		Step1Robot.Dir = W
	case W:
		Step1Robot.Dir = N
	}
}

func (d Dir) String() string {
	return fmt.Sprintf("%d", d)
}

type Action func(Dir, Pos) (Dir, Pos)

func A(dir Dir, pos Pos) (Dir, Pos) {
	newPos := pos
	switch dir {
	case N:
		newPos.Northing++
	case E:
		newPos.Easting++
	case S:
		newPos.Northing--
	case W:
		newPos.Easting--
	}
	return dir, newPos
}

func L(dir Dir, pos Pos) (Dir, Pos) {
	var newDir Dir
	switch dir {
	case N:
		newDir = W
	case E:
		newDir = N
	case S:
		newDir = E
	case W:
		newDir = S
	}
	return newDir, pos
}

func R(dir Dir, pos Pos) (Dir, Pos) {
	var newDir Dir
	switch dir {
	case N:
		newDir = E
	case E:
		newDir = S
	case S:
		newDir = W
	case W:
		newDir = N
	}
	return newDir, pos
}

func StartRobot(commands chan Command, actions chan Action) {
	defer close(actions)
	for c := range commands {
		switch c {
		case 'A':
			actions <- A
		case 'L':
			actions <- L
		case 'R':
			actions <- R
		}
	}
}

func Room(extent Rect, robot Step2Robot, act chan Action, rep chan Step2Robot) {
	for action := range act {
		newDir, newPos := action(robot.Dir, robot.Pos)
		if extent.IsValid(newPos) {
			robot.Pos = newPos
		}
		robot.Dir = newDir
	}
	rep <- robot
}

func (r Rect) IsValid(pos Pos) bool {
	return pos.Easting >= r.Min.Easting && pos.Easting <= r.Max.Easting &&
		pos.Northing >= r.Min.Northing && pos.Northing <= r.Max.Northing
}

type Action3 int

func StartRobot3(name, script string, action chan Action3, log chan string) {
	_ = name
	_ = script
	_ = action
	_ = log
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, report chan []Step3Robot, log chan string) {
	report <- robots
	_ = extent
	_ = action
	_ = log
}
