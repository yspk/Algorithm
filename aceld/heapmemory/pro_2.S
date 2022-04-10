"".foo STEXT size=264 args=0x8 locals=0x48 funcid=0x0
	0x0000 00000 (main.go:3)	TEXT	"".foo(SB), ABIInternal, $72-8
	0x0000 00000 (main.go:3)	CMPQ	SP, 16(R14)
	0x0004 00004 (main.go:3)	PCDATA	$0, $-2
	0x0004 00004 (main.go:3)	JLS	244
	0x000a 00010 (main.go:3)	PCDATA	$0, $-1
	0x000a 00010 (main.go:3)	SUBQ	$72, SP
	0x000e 00014 (main.go:3)	MOVQ	BP, 64(SP)
	0x0013 00019 (main.go:3)	LEAQ	64(SP), BP
	0x0018 00024 (main.go:3)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0018 00024 (main.go:3)	FUNCDATA	$1, gclocals·9fb7f0986f647f17cb53dda1484e0f7a(SB)
	0x0018 00024 (main.go:3)	FUNCDATA	$5, "".foo.arginfo1(SB)
	0x0018 00024 (main.go:3)	MOVQ	AX, "".arg_val+80(SP)
	0x001d 00029 (main.go:5)	MOVQ	$11, "".foo_val1+48(SP)
	0x0026 00038 (main.go:6)	MOVQ	$12, "".foo_val2+40(SP)
	0x002f 00047 (main.go:7)	LEAQ	type.int(SB), AX
	0x0036 00054 (main.go:7)	PCDATA	$1, $0
	0x0036 00054 (main.go:7)	CALL	runtime.newobject(SB)
	0x003b 00059 (main.go:7)	MOVQ	AX, "".&foo_val3+56(SP)
	0x0040 00064 (main.go:7)	MOVQ	$13, (AX)
	0x0047 00071 (main.go:8)	MOVQ	$14, "".foo_val4+32(SP)
	0x0050 00080 (main.go:9)	MOVQ	$15, "".foo_val5+24(SP)
	0x0059 00089 (main.go:9)	XORL	CX, CX
	0x005b 00091 (main.go:9)	NOP
	0x0060 00096 (main.go:15)	JMP	224
	0x0065 00101 (main.go:15)	MOVQ	CX, "".i+16(SP)
	0x006a 00106 (main.go:16)	PCDATA	$1, $1
	0x006a 00106 (main.go:16)	CALL	runtime.printlock(SB)
	0x006f 00111 (main.go:16)	LEAQ	"".arg_val+80(SP), AX
	0x0074 00116 (main.go:16)	CALL	runtime.printpointer(SB)
	0x0079 00121 (main.go:16)	CALL	runtime.printsp(SB)
	0x007e 00126 (main.go:16)	LEAQ	"".foo_val1+48(SP), AX
	0x0083 00131 (main.go:16)	CALL	runtime.printpointer(SB)
	0x0088 00136 (main.go:16)	CALL	runtime.printsp(SB)
	0x008d 00141 (main.go:16)	LEAQ	"".foo_val2+40(SP), AX
	0x0092 00146 (main.go:16)	CALL	runtime.printpointer(SB)
	0x0097 00151 (main.go:16)	CALL	runtime.printsp(SB)
	0x009c 00156 (main.go:16)	MOVQ	"".&foo_val3+56(SP), AX
	0x00a1 00161 (main.go:16)	CALL	runtime.printpointer(SB)
	0x00a6 00166 (main.go:16)	CALL	runtime.printsp(SB)
	0x00ab 00171 (main.go:16)	LEAQ	"".foo_val4+32(SP), AX
	0x00b0 00176 (main.go:16)	CALL	runtime.printpointer(SB)
	0x00b5 00181 (main.go:16)	CALL	runtime.printsp(SB)
	0x00ba 00186 (main.go:16)	LEAQ	"".foo_val5+24(SP), AX
	0x00bf 00191 (main.go:16)	NOP
	0x00c0 00192 (main.go:16)	CALL	runtime.printpointer(SB)
	0x00c5 00197 (main.go:16)	CALL	runtime.printnl(SB)
	0x00ca 00202 (main.go:16)	CALL	runtime.printunlock(SB)
	0x00cf 00207 (main.go:15)	MOVQ	"".i+16(SP), AX
	0x00d4 00212 (main.go:15)	LEAQ	1(AX), CX
	0x00d8 00216 (main.go:20)	MOVQ	"".&foo_val3+56(SP), AX
	0x00dd 00221 (main.go:20)	NOP
	0x00e0 00224 (main.go:15)	CMPQ	CX, $5
	0x00e4 00228 (main.go:15)	JLT	101
	0x00ea 00234 (main.go:20)	PCDATA	$1, $-1
	0x00ea 00234 (main.go:20)	MOVQ	64(SP), BP
	0x00ef 00239 (main.go:20)	ADDQ	$72, SP
	0x00f3 00243 (main.go:20)	RET
	0x00f4 00244 (main.go:20)	NOP
	0x00f4 00244 (main.go:3)	PCDATA	$1, $-1
	0x00f4 00244 (main.go:3)	PCDATA	$0, $-2
	0x00f4 00244 (main.go:3)	MOVQ	AX, 8(SP)
	0x00f9 00249 (main.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x00fe 00254 (main.go:3)	MOVQ	8(SP), AX
	0x0103 00259 (main.go:3)	PCDATA	$0, $-1
	0x0103 00259 (main.go:3)	JMP	0
	0x0000 49 3b 66 10 0f 86 ea 00 00 00 48 83 ec 48 48 89  I;f.......H..HH.
	0x0010 6c 24 40 48 8d 6c 24 40 48 89 44 24 50 48 c7 44  l$@H.l$@H.D$PH.D
	0x0020 24 30 0b 00 00 00 48 c7 44 24 28 0c 00 00 00 48  $0....H.D$(....H
	0x0030 8d 05 00 00 00 00 e8 00 00 00 00 48 89 44 24 38  ...........H.D$8
	0x0040 48 c7 00 0d 00 00 00 48 c7 44 24 20 0e 00 00 00  H......H.D$ ....
	0x0050 48 c7 44 24 18 0f 00 00 00 31 c9 0f 1f 44 00 00  H.D$.....1...D..
	0x0060 e9 7b 00 00 00 48 89 4c 24 10 e8 00 00 00 00 48  .{...H.L$......H
	0x0070 8d 44 24 50 e8 00 00 00 00 e8 00 00 00 00 48 8d  .D$P..........H.
	0x0080 44 24 30 e8 00 00 00 00 e8 00 00 00 00 48 8d 44  D$0..........H.D
	0x0090 24 28 e8 00 00 00 00 e8 00 00 00 00 48 8b 44 24  $(..........H.D$
	0x00a0 38 e8 00 00 00 00 e8 00 00 00 00 48 8d 44 24 20  8..........H.D$ 
	0x00b0 e8 00 00 00 00 e8 00 00 00 00 48 8d 44 24 18 90  ..........H.D$..
	0x00c0 e8 00 00 00 00 e8 00 00 00 00 e8 00 00 00 00 48  ...............H
	0x00d0 8b 44 24 10 48 8d 48 01 48 8b 44 24 38 0f 1f 00  .D$.H.H.H.D$8...
	0x00e0 48 83 f9 05 0f 8c 7b ff ff ff 48 8b 6c 24 40 48  H.....{...H.l$@H
	0x00f0 83 c4 48 c3 48 89 44 24 08 e8 00 00 00 00 48 8b  ..H.H.D$......H.
	0x0100 44 24 08 e9 f8 fe ff ff                          D$......
	rel 50+4 t=15 type.int+0
	rel 55+4 t=7 runtime.newobject+0
	rel 107+4 t=7 runtime.printlock+0
	rel 117+4 t=7 runtime.printpointer+0
	rel 122+4 t=7 runtime.printsp+0
	rel 132+4 t=7 runtime.printpointer+0
	rel 137+4 t=7 runtime.printsp+0
	rel 147+4 t=7 runtime.printpointer+0
	rel 152+4 t=7 runtime.printsp+0
	rel 162+4 t=7 runtime.printpointer+0
	rel 167+4 t=7 runtime.printsp+0
	rel 177+4 t=7 runtime.printpointer+0
	rel 182+4 t=7 runtime.printsp+0
	rel 193+4 t=7 runtime.printpointer+0
	rel 198+4 t=7 runtime.printnl+0
	rel 203+4 t=7 runtime.printunlock+0
	rel 250+4 t=7 runtime.morestack_noctxt+0
"".main STEXT size=274 args=0x0 locals=0x50 funcid=0x0
	0x0000 00000 (main.go:23)	TEXT	"".main(SB), ABIInternal, $80-0
	0x0000 00000 (main.go:23)	CMPQ	SP, 16(R14)
	0x0004 00004 (main.go:23)	PCDATA	$0, $-2
	0x0004 00004 (main.go:23)	JLS	264
	0x000a 00010 (main.go:23)	PCDATA	$0, $-1
	0x000a 00010 (main.go:23)	SUBQ	$80, SP
	0x000e 00014 (main.go:23)	MOVQ	BP, 72(SP)
	0x0013 00019 (main.go:23)	LEAQ	72(SP), BP
	0x0018 00024 (main.go:23)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0018 00024 (main.go:23)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0018 00024 (main.go:24)	MOVQ	$666, "".arg_val+56(SP)
	0x0021 00033 (<unknown line number>)	NOP
	0x0021 00033 (main.go:5)	MOVQ	$11, "".foo_val1+48(SP)
	0x002a 00042 (main.go:6)	MOVQ	$12, "".foo_val2+40(SP)
	0x0033 00051 (main.go:7)	MOVQ	$13, "".foo_val3+32(SP)
	0x003c 00060 (main.go:8)	MOVQ	$14, "".foo_val4+24(SP)
	0x0045 00069 (main.go:9)	MOVQ	$15, "".foo_val5+16(SP)
	0x004e 00078 (main.go:9)	XORL	AX, AX
	0x0050 00080 (main.go:15)	JMP	198
	0x0052 00082 (main.go:15)	MOVQ	AX, "".i+8(SP)
	0x0057 00087 (main.go:16)	PCDATA	$1, $0
	0x0057 00087 (main.go:16)	CALL	runtime.printlock(SB)
	0x005c 00092 (main.go:16)	LEAQ	"".arg_val+56(SP), AX
	0x0061 00097 (main.go:16)	CALL	runtime.printpointer(SB)
	0x0066 00102 (main.go:16)	CALL	runtime.printsp(SB)
	0x006b 00107 (main.go:16)	LEAQ	"".foo_val1+48(SP), AX
	0x0070 00112 (main.go:16)	CALL	runtime.printpointer(SB)
	0x0075 00117 (main.go:16)	CALL	runtime.printsp(SB)
	0x007a 00122 (main.go:16)	LEAQ	"".foo_val2+40(SP), AX
	0x007f 00127 (main.go:16)	NOP
	0x0080 00128 (main.go:16)	CALL	runtime.printpointer(SB)
	0x0085 00133 (main.go:16)	CALL	runtime.printsp(SB)
	0x008a 00138 (main.go:16)	LEAQ	"".foo_val3+32(SP), AX
	0x008f 00143 (main.go:16)	CALL	runtime.printpointer(SB)
	0x0094 00148 (main.go:16)	CALL	runtime.printsp(SB)
	0x0099 00153 (main.go:16)	LEAQ	"".foo_val4+24(SP), AX
	0x009e 00158 (main.go:16)	NOP
	0x00a0 00160 (main.go:16)	CALL	runtime.printpointer(SB)
	0x00a5 00165 (main.go:16)	CALL	runtime.printsp(SB)
	0x00aa 00170 (main.go:16)	LEAQ	"".foo_val5+16(SP), AX
	0x00af 00175 (main.go:16)	CALL	runtime.printpointer(SB)
	0x00b4 00180 (main.go:16)	CALL	runtime.printnl(SB)
	0x00b9 00185 (main.go:16)	CALL	runtime.printunlock(SB)
	0x00be 00190 (main.go:15)	MOVQ	"".i+8(SP), AX
	0x00c3 00195 (main.go:15)	INCQ	AX
	0x00c6 00198 (main.go:15)	CMPQ	AX, $5
	0x00ca 00202 (main.go:15)	JLT	82
	0x00cc 00204 (main.go:26)	MOVQ	"".foo_val3+32(SP), AX
	0x00d1 00209 (main.go:26)	MOVQ	AX, ""..autotmp_16+64(SP)
	0x00d6 00214 (main.go:26)	CALL	runtime.printlock(SB)
	0x00db 00219 (main.go:26)	MOVQ	""..autotmp_16+64(SP), AX
	0x00e0 00224 (main.go:26)	CALL	runtime.printint(SB)
	0x00e5 00229 (main.go:26)	CALL	runtime.printsp(SB)
	0x00ea 00234 (main.go:26)	LEAQ	"".foo_val3+32(SP), AX
	0x00ef 00239 (main.go:26)	CALL	runtime.printpointer(SB)
	0x00f4 00244 (main.go:26)	CALL	runtime.printnl(SB)
	0x00f9 00249 (main.go:26)	CALL	runtime.printunlock(SB)
	0x00fe 00254 (main.go:27)	MOVQ	72(SP), BP
	0x0103 00259 (main.go:27)	ADDQ	$80, SP
	0x0107 00263 (main.go:27)	RET
	0x0108 00264 (main.go:27)	NOP
	0x0108 00264 (main.go:23)	PCDATA	$1, $-1
	0x0108 00264 (main.go:23)	PCDATA	$0, $-2
	0x0108 00264 (main.go:23)	CALL	runtime.morestack_noctxt(SB)
	0x010d 00269 (main.go:23)	PCDATA	$0, $-1
	0x010d 00269 (main.go:23)	JMP	0
	0x0000 49 3b 66 10 0f 86 fe 00 00 00 48 83 ec 50 48 89  I;f.......H..PH.
	0x0010 6c 24 48 48 8d 6c 24 48 48 c7 44 24 38 9a 02 00  l$HH.l$HH.D$8...
	0x0020 00 48 c7 44 24 30 0b 00 00 00 48 c7 44 24 28 0c  .H.D$0....H.D$(.
	0x0030 00 00 00 48 c7 44 24 20 0d 00 00 00 48 c7 44 24  ...H.D$ ....H.D$
	0x0040 18 0e 00 00 00 48 c7 44 24 10 0f 00 00 00 31 c0  .....H.D$.....1.
	0x0050 eb 74 48 89 44 24 08 e8 00 00 00 00 48 8d 44 24  .tH.D$......H.D$
	0x0060 38 e8 00 00 00 00 e8 00 00 00 00 48 8d 44 24 30  8..........H.D$0
	0x0070 e8 00 00 00 00 e8 00 00 00 00 48 8d 44 24 28 90  ..........H.D$(.
	0x0080 e8 00 00 00 00 e8 00 00 00 00 48 8d 44 24 20 e8  ..........H.D$ .
	0x0090 00 00 00 00 e8 00 00 00 00 48 8d 44 24 18 66 90  .........H.D$.f.
	0x00a0 e8 00 00 00 00 e8 00 00 00 00 48 8d 44 24 10 e8  ..........H.D$..
	0x00b0 00 00 00 00 e8 00 00 00 00 e8 00 00 00 00 48 8b  ..............H.
	0x00c0 44 24 08 48 ff c0 48 83 f8 05 7c 86 48 8b 44 24  D$.H..H...|.H.D$
	0x00d0 20 48 89 44 24 40 e8 00 00 00 00 48 8b 44 24 40   H.D$@.....H.D$@
	0x00e0 e8 00 00 00 00 e8 00 00 00 00 48 8d 44 24 20 e8  ..........H.D$ .
	0x00f0 00 00 00 00 e8 00 00 00 00 e8 00 00 00 00 48 8b  ..............H.
	0x0100 6c 24 48 48 83 c4 50 c3 e8 00 00 00 00 e9 ee fe  l$HH..P.........
	0x0110 ff ff                                            ..
	rel 88+4 t=7 runtime.printlock+0
	rel 98+4 t=7 runtime.printpointer+0
	rel 103+4 t=7 runtime.printsp+0
	rel 113+4 t=7 runtime.printpointer+0
	rel 118+4 t=7 runtime.printsp+0
	rel 129+4 t=7 runtime.printpointer+0
	rel 134+4 t=7 runtime.printsp+0
	rel 144+4 t=7 runtime.printpointer+0
	rel 149+4 t=7 runtime.printsp+0
	rel 161+4 t=7 runtime.printpointer+0
	rel 166+4 t=7 runtime.printsp+0
	rel 176+4 t=7 runtime.printpointer+0
	rel 181+4 t=7 runtime.printnl+0
	rel 186+4 t=7 runtime.printunlock+0
	rel 215+4 t=7 runtime.printlock+0
	rel 225+4 t=7 runtime.printint+0
	rel 230+4 t=7 runtime.printsp+0
	rel 240+4 t=7 runtime.printpointer+0
	rel 245+4 t=7 runtime.printnl+0
	rel 250+4 t=7 runtime.printunlock+0
	rel 265+4 t=7 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
go.info."".foo$abstract SDWARFABSFCN dupok size=106
	0x0000 04 2e 66 6f 6f 00 01 01 11 61 72 67 5f 76 61 6c  ..foo....arg_val
	0x0010 00 00 00 00 00 00 0c 66 6f 6f 5f 76 61 6c 31 00  .......foo_val1.
	0x0020 05 00 00 00 00 0c 66 6f 6f 5f 76 61 6c 32 00 06  ......foo_val2..
	0x0030 00 00 00 00 0c 66 6f 6f 5f 76 61 6c 33 00 07 00  .....foo_val3...
	0x0040 00 00 00 0c 66 6f 6f 5f 76 61 6c 34 00 08 00 00  ....foo_val4....
	0x0050 00 00 0c 66 6f 6f 5f 76 61 6c 35 00 09 00 00 00  ...foo_val5.....
	0x0060 00 0c 69 00 0f 00 00 00 00 00                    ..i.......
	rel 0+0 t=23 type.*int+0
	rel 0+0 t=23 type.int+0
	rel 18+4 t=31 go.info.int+0
	rel 33+4 t=31 go.info.int+0
	rel 48+4 t=31 go.info.int+0
	rel 63+4 t=31 go.info.int+0
	rel 78+4 t=31 go.info.int+0
	rel 93+4 t=31 go.info.int+0
	rel 101+4 t=31 go.info.int+0
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
"".foo.arginfo1 SRODATA static dupok size=3
	0x0000 00 08 ff                                         ...
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
