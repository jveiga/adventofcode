package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Instruction struct {
	Name    string
	Operand int
	Visited bool
}

func (i *Instruction) toggle() {
}

func main() {
	lines := strings.Split(string(input), "\n")
	// lines := strings.Split(string(example), "\n")

	jmps := make([]int, 0)
	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		split := strings.Split(line, " ")
		if len(split) > 2 {
			fmt.Println(split)
			continue
		}
		operand, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println(err)
			continue
		}
		instruction := Instruction{
			Name:    split[0],
			Operand: operand,
		}
		if split[0] == "jmp" || split[0] == "nop" {
			jmps = append(jmps, i)
		}
		instructions[i] = instruction
	}

	currentInstructions := make([]Instruction, len(instructions))
	acc := 0
outer:
	for _, jmp := range jmps {
		fmt.Println("=====", jmp)
		acc = 0
		copy(currentInstructions, instructions)
		instructionPointer := 0
		var name string
		if currentInstructions[jmp].Name == "nop" {
			name = "jmp"
		}
		if currentInstructions[jmp].Name == "jmp" {
			name = "nop"
		}
		instruction := Instruction{
			Name:    name,
			Operand: currentInstructions[jmp].Operand,
		}
		currentInstructions[jmp] = instruction
		fmt.Println(currentInstructions)

		// current program
		visited := 0
		for {
			visited++
			instruction := currentInstructions[instructionPointer]
			fmt.Println(instructionPointer, instruction, acc)
			if instruction.Visited {
				break
			}
			instruction.Visited = true
			fmt.Println(instructionPointer, instruction, acc)
			currentInstructions[instructionPointer] = instruction
			switch instruction.Name {
			case "nop":
				instructionPointer += 1
				if instructionPointer > len(currentInstructions)-1 {
					log.Fatal("instructionPointer out of bounds", (acc))
				}
				continue
			case "acc":
				acc += instruction.Operand
				instructionPointer += 1
				if instructionPointer > len(currentInstructions)-1 {
					log.Fatal("instructionPointer out of bounds", (acc))
				}
			case "jmp":
				instructionPointer += instruction.Operand
				if instructionPointer > len(currentInstructions)-1 {
					log.Fatal("instructionPointer out of bounds", acc)
				}
			default:
				panic("how did we get here")
			}
			fmt.Println(visited, len(instructions))
			if visited == len(instructions) {
				fmt.Println("=================> ", acc)
				break outer
			}
		}
		fmt.Println("total", acc)
	}
	fmt.Println(acc)
}

var example = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

var input = `acc -5
nop +333
acc +45
jmp +288
acc -9
jmp +1
acc +27
jmp +464
acc +34
jmp +478
jmp +356
acc +10
acc +20
acc +29
acc -10
jmp +359
acc +29
acc +31
acc +36
acc +42
jmp +502
acc +14
jmp +45
jmp +499
acc -19
acc +4
acc +24
nop +460
jmp +465
acc +29
acc +6
acc +25
jmp +355
acc -10
acc +50
jmp -27
acc +46
acc +2
acc -18
acc +8
jmp +85
nop +264
acc +44
jmp +310
acc +23
acc -15
acc -12
jmp +290
acc -5
acc +4
acc -7
jmp +248
acc +23
jmp +434
nop -6
jmp +239
jmp +1
acc -19
jmp +67
acc +40
acc +21
acc +24
jmp +366
acc +38
acc +15
acc -2
jmp +542
acc +27
acc +21
acc +44
acc +31
jmp -60
jmp -51
acc +14
jmp +254
acc +43
acc -3
acc +30
jmp -5
acc +12
jmp +330
acc +4
jmp +81
nop +107
acc -12
nop +98
jmp +467
jmp +111
acc +0
acc +48
acc -4
jmp +184
jmp +310
acc +5
acc +1
acc +49
jmp +477
jmp +279
acc +12
acc -7
nop +51
jmp +125
jmp +1
acc -6
acc -19
acc -10
jmp +109
acc +28
acc +4
jmp +422
acc +12
jmp +152
jmp -71
acc -8
nop +252
nop +303
jmp -4
acc +1
jmp +200
acc -2
jmp +453
jmp +443
acc +5
acc -18
jmp +304
jmp +414
acc +15
jmp +271
acc +22
jmp +371
acc +29
acc +29
acc -17
jmp +166
jmp +49
acc -4
acc -6
jmp +461
acc +21
acc +49
jmp +458
acc +27
acc +3
acc -12
jmp +7
jmp +216
nop +385
acc +0
acc +11
acc +13
jmp +343
nop +273
acc +38
acc -2
jmp +61
nop +8
jmp +135
acc +13
acc +46
jmp +239
acc +38
acc +6
jmp +225
nop +337
jmp +66
acc +49
acc +10
jmp +167
acc -18
acc +32
jmp +107
nop +195
acc +39
jmp +391
acc +13
jmp +227
jmp +71
nop +340
acc +30
jmp +19
acc +42
acc +34
jmp +349
acc +46
jmp -130
nop +383
acc +45
acc -17
acc +13
jmp +354
acc +39
acc +26
nop +55
jmp -100
acc -16
acc +13
acc -1
jmp +395
acc +33
nop +106
acc -14
acc -7
jmp -74
acc +0
acc -8
jmp -28
nop +265
acc +27
acc +30
acc +23
jmp -112
acc +22
acc +7
acc +2
jmp +71
acc -6
acc +15
nop -89
acc +24
jmp +92
jmp +353
jmp -104
acc +19
acc +12
acc +12
jmp -132
acc +20
acc +27
jmp -60
jmp -170
acc +13
acc +15
jmp +114
acc +3
acc +13
acc -16
acc +50
jmp +124
acc +28
acc -10
acc +0
acc +21
jmp +192
acc +2
acc +17
acc +18
jmp +318
acc +41
acc +34
acc +0
acc -5
jmp +17
nop -131
acc +29
acc +46
nop +238
jmp +172
jmp +1
acc +14
acc +32
acc -15
jmp +331
jmp +209
jmp +189
acc +1
nop +163
acc +46
jmp -77
acc +0
jmp -131
acc +21
acc +8
acc +26
acc +12
jmp +72
jmp +258
jmp +183
acc +17
acc -12
acc +15
jmp +5
nop +85
acc +23
acc +40
jmp +53
acc +24
jmp +257
acc -10
acc +34
acc +49
jmp -178
acc +28
jmp +164
jmp +250
acc -14
acc +32
acc +13
jmp +96
jmp -290
acc -19
jmp -276
acc -15
nop +271
acc -16
jmp +264
acc +36
jmp -172
acc -11
acc +33
nop -261
jmp +77
nop +70
acc +6
nop -98
acc +8
jmp -158
acc +37
jmp +17
acc -2
acc -13
jmp -52
acc -9
jmp -116
acc +18
acc +9
acc -12
jmp +8
jmp +226
jmp +284
acc -2
nop +44
nop -320
nop -259
jmp -312
jmp -286
acc +18
jmp +224
acc +46
acc +29
jmp +1
acc +18
jmp +123
acc +31
nop -209
nop -39
jmp -171
acc -12
jmp -53
acc +19
acc +42
nop -317
acc -6
jmp -122
jmp -90
jmp +1
jmp -248
acc +0
jmp -34
acc +33
acc -8
jmp -312
acc +47
acc -10
acc -4
jmp -281
jmp +106
acc -17
acc +0
acc +17
nop +229
jmp +54
acc +31
acc +23
jmp -88
jmp -301
jmp +105
jmp -264
acc -9
acc -17
acc +25
jmp +120
jmp -274
jmp +140
nop +35
jmp -146
acc +31
jmp -63
acc +14
acc +45
acc +48
acc +7
jmp -246
jmp +108
acc +18
acc -8
jmp -12
acc +6
jmp +163
nop -80
nop -340
jmp -63
jmp -126
acc +26
acc +6
acc +5
acc -2
jmp +111
acc +47
acc +20
jmp +39
acc +38
acc +3
acc -13
acc +27
jmp -263
nop +9
acc +50
jmp -149
acc -18
jmp -245
acc +2
acc +40
acc -4
jmp -302
acc +10
acc +20
jmp -220
jmp -93
nop +43
acc -18
acc +6
jmp +1
jmp -307
jmp +75
jmp -177
acc +8
acc +31
acc +47
jmp -321
acc +22
jmp +1
acc +9
acc +32
jmp -56
acc -13
nop -140
acc +24
jmp -368
jmp -285
acc +38
acc +32
jmp +1
jmp -205
acc +47
acc +21
jmp -304
jmp -17
acc +9
jmp -399
nop -233
acc +18
jmp -63
acc +45
jmp -335
acc +35
acc +9
acc -12
jmp -19
acc +8
acc +48
jmp -179
acc +37
acc +15
jmp -182
acc +2
acc +22
acc +7
jmp -271
jmp -288
jmp -345
acc +21
nop -107
acc +17
jmp -462
acc +41
jmp +1
jmp -158
nop -310
acc +38
nop +28
acc +24
jmp -32
jmp -375
acc +20
acc +15
acc +11
acc -3
jmp -186
acc -15
jmp -40
acc +38
acc +27
acc +50
acc +8
jmp -406
acc +15
acc +39
jmp -409
nop -396
acc -14
acc -5
jmp -40
nop -156
acc +3
acc -3
acc +22
jmp -16
acc +9
jmp +68
nop -109
acc +18
jmp -198
nop -455
nop -195
jmp +1
nop -3
jmp -46
acc +40
acc +26
acc +47
jmp -509
jmp -92
jmp -166
nop -335
acc +6
acc +1
acc +28
jmp +44
jmp -79
acc -18
acc +13
jmp -10
jmp +66
acc +29
acc +34
jmp +1
acc +44
jmp -129
acc -5
acc +41
acc +48
acc +28
jmp +16
acc -1
acc +30
acc -4
jmp +52
acc +1
acc +37
jmp -312
acc +14
nop -340
jmp -341
jmp -55
nop -366
acc +14
jmp -185
jmp -450
acc -4
acc -4
acc +37
jmp -93
jmp -170
jmp +1
acc -13
acc +47
acc +29
jmp -456
acc -12
acc -9
jmp -397
acc -11
acc +6
jmp -207
acc +18
jmp -387
nop -268
acc +40
acc +26
jmp -21
acc +47
jmp -91
acc -15
jmp -227
nop -466
acc +4
acc -19
jmp -231
acc +29
acc +15
acc +0
acc +35
jmp -303
acc +28
acc +36
acc +34
acc -11
jmp -168
acc +48
jmp -521
acc +28
jmp -25
acc +47
acc +16
acc -13
acc +11
jmp +1`