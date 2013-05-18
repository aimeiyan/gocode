#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int precedence(char c) {
    switch (c) {
    case '-':
    case '+':
        return 1;
    case '*':
    case '/':
        return 2;
    case ')':
        return 0;
        // return 3;
    case '(':
        return 3;
    default:
        printf("not expected %c in precedence\n", c);
        return -100;
    }
}

char* in2post(char* in) {
    // TODO error checking, overflow
    static char buf[8000];
    char stack[100];
    int stack_size = 0;
    char *dest = buf;

    for(;*in;in++){
        switch(*in){
        case ' ':
            break;
        case '(':
            //     stack[stack_size ++] = *in;
            //     break;
        case ')':
            //     stack_size --;
            //     stack_size --;
            //     *dest++ = stack[stack_size];
            //     break;
        case '-':
        case '+':
        case '*':
        case '/':
            // printf("----------char %c, stack_size: %d\n", *in, stack_size);
            while(stack_size > 0) {
                char top = stack[stack_size - 1];
                // current is equal lower
                if (precedence(*in) > precedence(top)) {
                    // printf("push stack %c\n", *in);
                    stack[stack_size ++] = *in;
                    break;
                } else {
                    if (top != '(') {
                        *dest++ = top;
                    }
                    // printf("top %c, dest: %s, ---\n", top, buf);
                    stack_size -= 1;
                }
            }

            if (stack_size == 0 && *in != ')') { // empty stack
                stack[stack_size ++] = *in;
                // printf("emtpy stack, push %c\n", *in);
            }
            break;
        default:
            *dest++ = *in;
        }
    }
    while (stack_size > 0) {
        *dest++ = stack[--stack_size];
    }
    *dest = '\0';
    return buf;
}

int main(int argc, char** argv) {
    char *ins[] = {
        "(a + b * c) + d" , "abc*+d+",
        "a+b*c-d"         , "abc*+d-",
        "a+b*c/d -c*d"    , "abc*d/+cd*-",
        "(3 + 2) / 1"     , "32+1/",
        "3 + 2 / 1"       , "321/+",
    };

    int total = sizeof(ins) / sizeof(*ins), fail = 0;

    for(int  i = 0; i < total; i+=2 ){
        char *s = ins[i];
        char *result = in2post(s);
        char *expected = ins[i+1];
        if(strcmp(result, expected) != 0) {
            printf("ERROR, %15s => %15s, expect %15s\n", s, result, expected);
            fail += 1;
        } else {
            printf("OK, %-15s => %-15s\n", s, result);
        }
    }

    printf("TOTAL: %d, FAIL: %d\n", total, fail);
    return 0;
}
