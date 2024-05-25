#define _OPEN_THREADS
#include <pthread.h>
#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>

void *thread(void *arg) {
    pthread_t self = pthread_self();
    int i = 0;

    for (i = 0; i < 10000; i++) {
        flockfile(stdout);
        printf("%ld\n", self);
        fflush(stdout);
        funlockfile(stdout);
    }
    pthread_exit(0);
}

int main() {
    pthread_t thid[100];
    int i;
    void *ret;

    for (i = 0; i < 100; i++) {
        if (pthread_create(&thid[i], NULL, thread, NULL) != 0) {
            perror("pthread_create() error");
            exit(1);
        }
    }

    for (i = 0; i < 100; i++) {
        if (pthread_join(thid[i], NULL) != 0) {
            perror("pthread_create() error");
            exit(3);
        }
    }
}