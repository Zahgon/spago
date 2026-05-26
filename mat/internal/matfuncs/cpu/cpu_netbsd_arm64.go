// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

// Minimal copy of functionality from x/sys/unix so the cpu package can call
// sysctl without depending on x/sys/unix.

const (
	_CTL_QUERY = -2

	_SYSCTL_VERS_1 = 0x1000000
)

var _zero uintptr

func sysctl(mib []int32, old *byte, oldlen *uintptr, new *byte, newlen uintptr) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type sysctlNode struct {
	Flags          uint32
	Num            int32
	Name           [32]int8
	Ver            uint32
	__rsvd         uint32
	Un             [16]byte
	_sysctl_size   [8]byte
	_sysctl_func   [8]byte
	_sysctl_parent [8]byte
	_sysctl_desc   [8]byte
}

func sysctlNodes(mib []int32) ([]sysctlNode, error) {
	_ = "STUB: not implemented"

	// Get a list of all sysctl nodes below the given MIB by performing
	// a sysctl for the given MIB with CTL_QUERY appended.
	return nil, nil
}

// Now that we know the size, get the actual nodes.

func nametomib(name string) ([]int32, error) {
	_ = "STUB: not implemented"
	// Split name into components.
	return nil, nil
}

// Discover the nodes and construct the MIB OID.

// aarch64SysctlCPUID is struct aarch64_sysctl_cpu_id from NetBSD's <aarch64/armreg.h>
type aarch64SysctlCPUID struct {
	midr      uint64 /* Main ID Register */
	revidr    uint64 /* Revision ID Register */
	mpidr     uint64 /* Multiprocessor Affinity Register */
	aa64dfr0  uint64 /* A64 Debug Feature Register 0 */
	aa64dfr1  uint64 /* A64 Debug Feature Register 1 */
	aa64isar0 uint64 /* A64 Instruction Set Attribute Register 0 */
	aa64isar1 uint64 /* A64 Instruction Set Attribute Register 1 */
	aa64mmfr0 uint64 /* A64 Memory Model Feature Register 0 */
	aa64mmfr1 uint64 /* A64 Memory Model Feature Register 1 */
	aa64mmfr2 uint64 /* A64 Memory Model Feature Register 2 */
	aa64pfr0  uint64 /* A64 Processor Feature Register 0 */
	aa64pfr1  uint64 /* A64 Processor Feature Register 1 */
	aa64zfr0  uint64 /* A64 SVE Feature ID Register 0 */
	mvfr0     uint32 /* Media and VFP Feature Register 0 */
	mvfr1     uint32 /* Media and VFP Feature Register 1 */
	mvfr2     uint32 /* Media and VFP Feature Register 2 */
	pad       uint32
	clidr     uint64 /* Cache Level ID Register */
	ctr       uint64 /* Cache Type Register */
}

func sysctlCPUID(name string) (*aarch64SysctlCPUID, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func doinit() { _ = "STUB: not implemented"; return }
