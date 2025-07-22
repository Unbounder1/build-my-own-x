#include <stdbool.h>
#include <stdio.h>

bool running = true;

typedef enum {
    PSH,
    ADD,
    POP,
    SET,
    HLT
} InstructionSet;

const int program[] = {
    PSH, 5,
    PSH, 6,
    ADD,
    POP,
    HLT
};

int ip = 0;
int sp = -1;
int stack[256];
int r[8];

int fetch();
void eval(int instr);
int pop();

int main() {
    while (running) {
        eval(fetch());
        ip++;
    }
    for (int i = 0; i <= sp; i++){
        printf("%i\n", stack[i]);
    }
    
}

int fetch() {
    return program[ip];
}

void eval(int instr) {
    switch (instr) {
    case HLT: {
        running = false;
        break;
    }
    case PSH: {
        sp++;
        stack[sp] = program[++ip];
        printf("%i\n", stack[sp]);
        break;
    }
    case ADD: {
        int r1 = pop();
        int r2 = pop();
        stack[++sp] = r1 + r2;
        printf("%i\n", stack[sp]);
        break;
    }
    case SET: {
        sp++;
        int ptr = program[++ip];
        r[ptr - 'A'] = program[++ip];
        break;
    }
    case POP: {
        pop();
        printf("%i\n", stack[sp]);
        break;
    }
    }
}

int pop(){
    return stack[sp--];
}