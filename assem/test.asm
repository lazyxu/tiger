SECTION .data
    msg: db "hello MeteorKL!", 0x0a
    len: equ $-msg
    msg1: db "123", 0x0a
    len1: equ $-msg1

SECTION .text
    global start

print:
    push rbp
    mov rbp, rsp

    mov rax, 0x2000004
    mov rdi, 1
    mov rsi, [rbp-0x-10]
    mov rdx, [rbp+24]
    syscall

    pop rbp
    ret

start:
	mov r8, [ebp+0x4]
	mov r8-0x8, msg
    push len
    mov r8, msg
    push r8
    call print

    push len1
    mov r8, msg1
    push r8
    call print

    mov rax, 0x2000001
    mov rdi, 0
    syscall

    