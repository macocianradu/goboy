package cpu

import (
	"radu.macocian.me/goboy/cpu/operations"
)

/*
	0x10 STOP Execution of a STOP instruction stops both the system clock and oscillator circuit. STOP mode is entered and the LCD controller also stops. However, the status of the internal RAM register ports remains unchanged.

STOP mode can be cancelled by a reset signal.
If the RESET terminal goes LOW in STOP mode, it becomes that of a normal reset status.
The following conditions should be met before a STOP instruction is executed and stop mode is entered:

	All interrupt-enable (IE) flags are reset.
	Input to P10-P13 is LOW for all.
*/
func STOP(context op_context) {
	//TODO
	return
}
