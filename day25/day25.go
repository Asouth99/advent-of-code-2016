package day25

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solve(part int, logger *log.Logger, inputFile ...string) (any, error) {
	file := "./day25/input.txt"
	if len(inputFile) > 0 {
		file = inputFile[0]
	}

	switch part {
	case 1:
		return SolvePart1(file, logger), nil
	case 2:
		return SolvePart2(file, logger), nil
	default:
		return -1, errors.New("incorrect part number recieved")
	}
}

type VirtualMachine struct {
	Registers map[string]int
	Program   [][]string
	Ptr       int
	logger    *log.Logger
	clock     []int
}

func (vm *VirtualMachine) isReg(s string) bool {
	for k, _ := range vm.Registers {
		if s == k {
			return true
		}
	}
	return false
}

func (vm *VirtualMachine) getVal(arg string) int {
	if val, ok := vm.Registers[arg]; ok {
		return val
	}
	v, _ := strconv.Atoi(arg)
	return v
}

func (vm *VirtualMachine) printState() {
	str := fmt.Sprintf("ptr: %d %v |", vm.Ptr, vm.Program[vm.Ptr])
	for _, r := range []string{"a", "b", "c", "d"} {
		str += fmt.Sprintf(" %s:%d |", r, vm.Registers[r])
	}
	vm.logger.Print(str + "\n")
}

func (vm *VirtualMachine) toggle(index int) {
	inst := vm.Program[index]
	if len(inst) <= 2 {
		if inst[0] == "inc" {
			inst[0] = "dec"
		} else {
			inst[0] = "inc"
		}
	} else {
		if inst[0] == "jnz" {
			inst[0] = "cpy"
		} else {
			inst[0] = "jnz"
		}
	}
}

func (vm *VirtualMachine) jnz(arg1 string, arg2 string) {
	condVal := vm.getVal(arg1)
	jumpVal := vm.getVal(arg2)

	if condVal != 0 {
		if jumpVal < 0 {
			if vm.detectLoop(arg1, jumpVal) {
				// Fast-forward succeeded! Registers updated instantly,
				// so we don't jump back into the loop.
				vm.Ptr++
				return
			}
		}

		vm.Ptr += jumpVal
	} else {
		vm.Ptr++
	}
}

func (vm *VirtualMachine) detectLoop(condReg string, jumpVal int) bool {
	loopStartPtr := vm.Ptr + jumpVal
	loopEndPtr := vm.Ptr

	// vm.logger.Printf("Detecting loop from %d to %d\n", loopStartPtr, loopEndPtr)

	// 1. Clone state into sandbox VM
	loopVM := VirtualMachine{
		Ptr:       loopStartPtr,
		Registers: make(map[string]int, len(vm.Registers)),
		Program:   make([][]string, len(vm.Program)),
		logger:    vm.logger,
		clock:     make([]int, len(vm.clock)),
	}
	for k, v := range vm.Registers {
		loopVM.Registers[k] = v
	}
	for i, inst := range vm.Program {
		loopInst := make([]string, len(inst))
		copy(loopInst, inst)
		loopVM.Program[i] = loopInst
	}
	copy(loopVM.clock, vm.clock)

	// 2. Simulate one full iteration of the loop
	maxSteps := (loopEndPtr - loopStartPtr + 1) * 2
	steps := 0

	for steps < maxSteps {
		// Stop when sandbox reaches the loop boundary again
		if loopVM.Ptr == loopEndPtr {
			// Calculate net changes per iteration
			deltas := make(map[string]int)
			for reg, newVal := range loopVM.Registers {
				deltas[reg] = newVal - vm.Registers[reg]
			}

			// Verify the control variable decreases toward 0
			condDelta, exists := deltas[condReg]
			if !exists || condDelta >= 0 || vm.Registers[condReg] <= 0 {
				return false
			}

			// Calculate total remaining full iterations
			iterations := vm.Registers[condReg] / (-condDelta)
			if iterations <= 0 {
				return false
			}

			// Fast-forward real VM registers
			for reg, delta := range deltas {
				vm.Registers[reg] += delta * iterations
			}

			// vm.logger.Printf("Accelerated loop [%d..%d] x %d iterations\n", loopStartPtr, loopEndPtr, iterations)

			return true
		}

		// Abort simulation if execution strays or encounters 'tgl'
		if loopVM.Ptr < loopStartPtr || loopVM.Ptr > loopEndPtr {
			return false
		}
		if loopVM.Program[loopVM.Ptr][0] == "tgl" {
			return false
		}

		loopVM.step()
		steps++
	}

	return false
}

// step fetches and executes a single instruction at the current Ptr.
func (vm *VirtualMachine) step() {
	// vm.printState()

	if vm.Ptr < 0 || vm.Ptr >= len(vm.Program) {
		return // Out of bounds
	}

	inst := vm.Program[vm.Ptr]
	op := inst[0]

	switch op {
	case "cpy":
		// Only copy if the target destination is a valid register name (not an integer)
		if vm.isReg(inst[2]) {
			vm.Registers[inst[2]] = vm.getVal(inst[1])
		}
		vm.Ptr++

	case "inc":
		if vm.isReg(inst[1]) {
			vm.Registers[inst[1]]++
		}
		vm.Ptr++

	case "dec":
		if vm.isReg(inst[1]) {
			vm.Registers[inst[1]]--
		}
		vm.Ptr++

	case "jnz":
		vm.jnz(inst[1], inst[2])
	case "tgl":
		target := vm.Ptr + vm.getVal(inst[1])
		if target >= 0 && target < len(vm.Program) {
			vm.toggle(target)
		}
		vm.Ptr++
	case "out":
		val := vm.getVal(inst[1])
		// expected := len(vm.clock) % 2
		// if val != expected {
		// 	vm.Ptr = -1 // Force VM to terminate early
		// 	return
		// }
		vm.clock = append(vm.clock, val)
		vm.Ptr++
	default:
		vm.logger.Fatalf("Unknown instruction found %v", inst)
	}
}

func (vm *VirtualMachine) Run() bool {
	for vm.Ptr >= 0 && vm.Ptr < len(vm.Program) {
		vm.step()

		if len(vm.clock) >= 100 {
			return true
		}
	}
	return false
}

func SolvePart1(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	// Read program
	programLines := [][]string{}
	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		lineSplit := strings.Fields(line)
		programLines = append(programLines, lineSplit)
	}

	// Initialise machine
	a := 180
	for true {
		vm := VirtualMachine{
			Registers: map[string]int{"a": a, "b": 0, "c": 0, "d": 0},
			Program:   programLines,
			Ptr:       0,
			logger:    logger,
		}
		if vm.Run() {
			logger.Printf("a = %d : clock = %v\n", a, vm.clock)
			break
		}
		logger.Printf("a = %d : clock = %v\n", a, vm.clock)
		a++
		break
	}

	answer := a
	return answer
}

func SolvePart2(inputFile string, logger *log.Logger) int {
	f, err := os.Open(inputFile)
	if err != nil {
		logger.Fatalf("error opening file: %v\n", err)
	}
	defer f.Close()

	// Read program
	programLines := [][]string{}
	scanner := bufio.NewScanner(f)
	i := -1
	for scanner.Scan() {
		i++
		line := scanner.Text()
		lineSplit := strings.Fields(line)
		programLines = append(programLines, lineSplit)
	}

	// Initialise machine
	vm := VirtualMachine{
		Registers: map[string]int{"a": 12, "b": 0, "c": 0, "d": 0},
		Program:   programLines,
		Ptr:       0,
		logger:    logger,
	}
	vm.Run()

	answer := vm.Registers["a"]
	return answer
}
