// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mips64

import (
	"cmd/compile/internal/gc"
	"cmd/compile/internal/ssa"
	"cmd/internal/obj"
	"cmd/internal/obj/mips"
)

func Main() {
	gc.Thearch.LinkArch = &mips.Linkmips64
	if obj.GOARCH == "mips64le" {
		gc.Thearch.LinkArch = &mips.Linkmips64le
	}
	gc.Thearch.REGSP = mips.REGSP
	gc.Thearch.REGCTXT = mips.REGCTXT
	gc.Thearch.MAXWIDTH = 1 << 50

	gc.Thearch.Defframe = defframe
	gc.Thearch.Proginfo = proginfo

	gc.Thearch.SSARegToReg = ssaRegToReg
	gc.Thearch.SSAMarkMoves = func(s *gc.SSAGenState, b *ssa.Block) {}
	gc.Thearch.SSAGenValue = ssaGenValue
	gc.Thearch.SSAGenBlock = ssaGenBlock

	gc.Main()
	gc.Exit(0)
}
