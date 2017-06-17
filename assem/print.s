#    mov r9, len
#    push r9
#    mov r8, msg
#    push r8
#    call print

print:
    push rbp
    mov rbp, rsp

    mov rax, 0x2000004
    mov rdi, 1
    mov rsi, [rbp+16]
    mov rdx, [rbp+24]
    syscall

    pop rbp
    ret


#    stack
#    [rbp] 返回地址
#    [rbp+8] rbp
#    [rbp+16] 参数1