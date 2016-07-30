#include <stdio.h>

long long int mo = 1000000007;

int main() {

    long long int two_power[10000];
    two_power[0] = 1;

    for(int i = 1; i < 10000; i++) {
        two_power[i] = ((two_power[i-1]*2)%mo);
    }

    long long int ans[100000];

    int count = 0;
    for(int i = 1; i < 10000; i++) {
        for(int j = 0; j <= i-1; j++) {
            if(count >= 100000) {
                i = 1000000;
                break;
            }
            long long int temp = ((two_power[i] + two_power[j])%mo);
            if(count == 0) {
                ans[0] = temp;
            }else {
                ans[count] = ((ans[count-1] + temp)%mo);
            }
            count++;
        }
    }

    int t,n;
    scanf("%d", &t);

    for(int i=0; i < t; i++) {
           scanf("%d", &n);
                  printf("%lld\n", ans[n-1]);
    }
    return 0;
}
