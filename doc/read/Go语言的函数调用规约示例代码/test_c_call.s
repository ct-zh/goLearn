	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 15, 0	sdk_version 15, 5
	.globl	_my_c_function                  ; -- Begin function my_c_function
	.p2align	2
_my_c_function:                         ; @my_c_function
	.cfi_startproc
; %bb.0:
	sub	sp, sp, #32
	.cfi_def_cfa_offset 32
	str	w0, [sp, #28]
	str	w1, [sp, #24]
	str	w2, [sp, #20]
	str	w3, [sp, #16]
	str	w4, [sp, #12]
	str	w5, [sp, #8]
	str	w6, [sp, #4]
	ldr	w8, [sp, #28]
	ldr	w9, [sp, #24]
	add	w8, w8, w9
	ldr	w9, [sp, #20]
	add	w8, w8, w9
	ldr	w9, [sp, #16]
	add	w8, w8, w9
	ldr	w9, [sp, #12]
	add	w8, w8, w9
	ldr	w9, [sp, #8]
	add	w8, w8, w9
	ldr	w9, [sp, #4]
	add	w0, w8, w9
	add	sp, sp, #32
	ret
	.cfi_endproc
                                        ; -- End function
	.globl	_main                           ; -- Begin function main
	.p2align	2
_main:                                  ; @main
	.cfi_startproc
; %bb.0:
	sub	sp, sp, #32
	stp	x29, x30, [sp, #16]             ; 16-byte Folded Spill
	add	x29, sp, #16
	.cfi_def_cfa w29, 16
	.cfi_offset w30, -8
	.cfi_offset w29, -16
	mov	w8, #0                          ; =0x0
	str	w8, [sp, #8]                    ; 4-byte Folded Spill
	stur	wzr, [x29, #-4]
	mov	w0, #17                         ; =0x11
	mov	w1, #34                         ; =0x22
	mov	w2, #51                         ; =0x33
	mov	w3, #68                         ; =0x44
	mov	w4, #85                         ; =0x55
	mov	w5, #102                        ; =0x66
	mov	w6, #119                        ; =0x77
	bl	_my_c_function
	ldr	w0, [sp, #8]                    ; 4-byte Folded Reload
	ldp	x29, x30, [sp, #16]             ; 16-byte Folded Reload
	add	sp, sp, #32
	ret
	.cfi_endproc
                                        ; -- End function
.subsections_via_symbols
