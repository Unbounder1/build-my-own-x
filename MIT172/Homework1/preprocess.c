#define ONE 1

// Macros can behave similar to inline functions
// Parentheses around parameters are required to preserve order of operations
#define MIN(a,b) ((a) < (b) ? (a) : (b))

int c = ONE, d = ONE + 5;
int e = MIN(c,d);

#ifndef NDEBUG
// This code will only be compiled when the macro NDEBUG is not defined
// IF clang is passed DNDEBUG on the command line then NDEBUG will be defined
if (something){}
#endif
