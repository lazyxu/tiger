; http://blog.csdn.net/u013737447/article/details/49154509
;rdi                  ;pSrc1
;rsi                  ;pSrc2
;rdx                  ;pDst1
;rcx                  ;pDst2
;r8                   ;Stride1
;r9                   ;Stride2
;mov rax, [rsp + 8]    ;round1
SECTION .data
    msg: db "hello MeteorKL!", 0x0a, 0
    len: equ $-msg
    msg1: db "123", 0x0a, 0
    len1: equ $-msg1

SECTION .text
    global start

print_s:
    push rbp
    mov rbp, rsp

    mov rax, 0x2000004
    mov rdi, 1
    mov rsi, [rbp+16]
    mov rdx, [rbp+24]
    syscall

    pop rbp
    ret 0x10

print:
    push rbp
    mov rbp, rsp

    mov rax, [rbp+16]
    mov rbx, 0
    mov rcx, 0

print_loop:
    mov cl, [rax+rbx]
    cmp cl, 0
    je print_loop_exit
    inc bl
    jmp print_loop

print_loop_exit:
    push rbx
    push rax
    call print_s

    pop rbp
    ret

exit:
    mov rax, 0x2000001
    mov rdi, 0
    syscall
    ret

start:

    mov r8, msg
    push r8
    call print

    call exit

    