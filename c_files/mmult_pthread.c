#include <stdio.h>
#include <time.h>
#include <sys/time.h>
#include <pthread.h>
#include <strings.h>

#define		NROW	1024
#define		NCOL	NROW


#define TEST_RESULTS
int NUMTHREADS;

//Input Array A
int inputArrayA  [NROW][NCOL];
//Input Array B
int inputArrayB  [NROW][NCOL];
//Output Array C
long outputArrayC [NROW][NCOL];

struct timeval startTime;
struct timeval finishTime;
double timeIntervalLength;

void *threadMult(void *rank) {
	long myRank = (long)rank;
	int my_first_i = myRank * NROW/NUMTHREADS;
	int my_last_i = my_first_i + NROW/NUMTHREADS;

	int i, j, k;
	for(i=my_first_i; i<my_last_i && i < NROW; i++)
	{
		for(j=0; j<NCOL; j++)
		{
			long sum = 0;
			for(k=0;k<NROW;k++)
			{
				sum+=inputArrayA[i][k]*inputArrayB[k][j];
			}
			outputArrayC[i][j] = sum;
		}
	}
}

int main(int argc, char* argv[])
{
	int i,j,k;
	double totalSum;
	int num_threads = atoi(argv[1]);
	NUMTHREADS = num_threads;

	//INITIALIZE ARRAYS
	for(i=0;i<NROW;i++)
	{
		for(j=0;j<NCOL;j++)
		{
			inputArrayA[i][j]= i*NCOL+j;
			inputArrayB[i][j]= j*NCOL+i;
			outputArrayC[i][j]= 0;
		}
	}


	//Get the start time
	gettimeofday(&startTime, NULL); /* START TIME */

   	pthread_t tid[NUMTHREADS];
	void* status;

	for(i = 0; i < NUMTHREADS; i++) {
		pthread_create(&tid[i], NULL, threadMult, (void*)i);
	}

	for(i = 0; i < NUMTHREADS; i++) {
		pthread_join(tid[i], &status);
	}

	//Get the end time
	gettimeofday(&finishTime, NULL);  /* END TIME */



	#ifdef TEST_RESULTS
		//CALCULATE TOTAL SUM
		//[Just for verification]
		totalSum=0;
		//
		for(i=0;i<NROW;i++)
		{
			for(j=0;j<NCOL;j++)
			{
				totalSum+=(double)outputArrayC[i][j];
			}
		}

		//printf("\nTotal Sum = %g\n",totalSum);
	#endif

	//Calculate the interval length
	timeIntervalLength = (double)(finishTime.tv_sec-startTime.tv_sec) * 1000000
	                     + (double)(finishTime.tv_usec-startTime.tv_usec);
	timeIntervalLength=timeIntervalLength/1000;

	//Print the interval length
	printf("%g", timeIntervalLength);

	return 0;
}
