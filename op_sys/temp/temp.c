#include <stdlib.h>
#include <stdio.h>

int main(){
    int* c = malloc(16);
    printf("c is at memory address: %p\n", (c));
    printf("c is %i\n", *(c + 99999000090009999));


}
