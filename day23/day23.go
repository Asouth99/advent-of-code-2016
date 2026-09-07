package day23

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
	file := "./day23/input.txt"
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

func printState(reg map[string]int, ptr int, programLines [][]string, logger *log.Logger) {
	str := "|"
	for _, r := range []string{"a", "b", "c", "d"} {
		str += fmt.Sprintf(" %s:%d |", r, reg[r])
	}
	str += fmt.Sprintf("\nptr: %d - %v\n", ptr, programLines[ptr])
	logger.Print(str)
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

	// Initialise machine state
	registers := map[string]int{"a": 7, "b": 0, "c": 0, "d": 0}
	programPtr := 0

	// Loop over program
	for programPtr < len(programLines) {
		programLine := programLines[programPtr]

		printState(registers, programPtr, programLines, logger)

		op := programLine[0]

		switch op {
		case "tgl":
			n := registers[programLine[1]] + programPtr
			if n >= 0 && n < len(programLines) {
				inst := programLines[n]
				if len(inst) <= 2 {
					if inst[0] == "inc" {
						programLines[n][0] = "dec"
					} else {
						programLines[n][0] = "inc"
					}
				} else {
					if inst[0] == "jnz" {
						programLines[n][0] = "cpy"
					} else {
						programLines[n][0] = "jnz"
					}
				}
			}
		case "cpy":
			if _, ok := registers[programLine[2]]; ok {
				if _, ok := registers[programLine[1]]; ok {
					registers[programLine[2]] = registers[programLine[1]]
				} else {
					val, err := strconv.Atoi(programLine[1])
					if err != nil {
						logger.Fatal(err)
					}
					registers[programLine[2]] = val
				}
			}
		case "inc":
			registers[programLine[1]]++
		case "dec":
			registers[programLine[1]]--
		case "jnz":
			val, ok := registers[programLine[1]]
			if !ok {
				val, err = strconv.Atoi(programLine[1])
				if err != nil {
					logger.Fatal(err)
				}
			}
			if val != 0 {
				jump, ok := registers[programLine[2]]
				if !ok {
					jump, err = strconv.Atoi(programLine[2])
					if err != nil {
						logger.Fatal(err)
					}
				}
				programPtr += jump
				continue
			}
		default:
			logger.Fatalf("Unknown operand found in instruction '%s'", programLine)
		}

		programPtr++
	}

	answer := registers["a"]
	return answer
}

type VirtualMachine struct {
	Registers map[string]int
	Program   [][]string
	Ptr       int
	logger    *log.Logger
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

	vm.logger.Printf("Detecting loop from %d to %d\n", loopStartPtr, loopEndPtr)

	// 1. Clone state into sandbox VM
	loopVM := VirtualMachine{
		Ptr:       loopStartPtr,
		Registers: make(map[string]int, len(vm.Registers)),
		Program:   make([][]string, len(vm.Program)),
		logger:    vm.logger,
	}
	for k, v := range vm.Registers {
		loopVM.Registers[k] = v
	}
	for i, inst := range vm.Program {
		loopInst := make([]string, len(inst))
		copy(loopInst, inst)
		loopVM.Program[i] = loopInst
	}

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

			vm.logger.Printf("Accelerated loop [%d..%d] x %d iterations\n", loopStartPtr, loopEndPtr, iterations)

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
	vm.printState()

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
	}
}

func (vm *VirtualMachine) Run(logger *log.Logger) {

	for vm.Ptr >= 0 && vm.Ptr < len(vm.Program) {
		vm.step()
	}

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
	vm.Run(logger)

	answer := vm.Registers["a"]
	return answer
}
