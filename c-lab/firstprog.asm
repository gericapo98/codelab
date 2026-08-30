section .data
    message db "Hello, world!", 0xA
    msg_len equ $ - message

section .bss

section .text
global _start

_start:
    mov ecx, 10
.loop:
    mov eax, 4
    mov ebx, 1
    mov edx, msg_len
    lea esi, [message]
    int 0x80

    dec ecx
    jnz .loop

    mov eax, 1
    xor ebx, ebx
    int 0x80

