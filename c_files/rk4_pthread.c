#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <sys/time.h>
#include <pthread.h>
#include <strings.h>


#define PROBLEM_SIZE 8192
int NUMTHREADS;

struct timeval startTime;
struct timeval finishTime;
double timeIntervalLength;

	double h=0.3154;
	double*  y;
	double*  yt;
	double*  k1;
	double*  k2;
	double*  k3;
	double*  k4;
	double*  pow;
	double*  yout;
	double** c;

	double totalSum=0.0;

int counter1 = 0;
pthread_mutex_t barrier1;
int counter2 = 0;
pthread_mutex_t barrier2;
int counter3 = 0;
pthread_mutex_t barrier3;

pthread_mutex_t mutex1;

void* threadWork(void* rank)
{
	long my_rank = (long)rank;
	int my_first_i = my_rank * PROBLEM_SIZE/NUMTHREADS;
	int my_last_i = my_first_i + PROBLEM_SIZE/NUMTHREADS;
	double my_sum = 0.0;

	int i, j;
	for (i = my_first_i; i < my_last_i && i < PROBLEM_SIZE; i++)
	{
		yt[i] = 0.0;
		for (j = 0; j < PROBLEM_SIZE; j++)
		{
			yt[i] += c[i][j]*y[j];
		}
		k1[i] = h*(pow[i]-yt[i]);
	}

	pthread_mutex_lock(&barrier1);
	counter1++;
	pthread_mutex_unlock(&barrier1);
	while(counter1 < NUMTHREADS);

	for (i = my_first_i; i < my_last_i && i < PROBLEM_SIZE; i++)
	{
		yt[i] = 0.0;
		for (j = 0; j < PROBLEM_SIZE; j++)
		{
			yt[i] += c[i][j]*(y[j]+0.5*k1[j]);
		}
		k2[i] = h*(pow[i]-yt[i]);
	}

	pthread_mutex_lock(&barrier2);
	counter2++;
	pthread_mutex_unlock(&barrier2);
	while(counter2 < NUMTHREADS);

	for (i = my_first_i; i < my_last_i && i < PROBLEM_SIZE; i++)
	{
		yt[i] = 0.0;
		for (j = 0; j < PROBLEM_SIZE; j++)
		{
			yt[i] += c[i][j]*(y[j]+0.5*k2[j]);
		}
		k3[i] = h*(pow[i]-yt[i]);
	}

	pthread_mutex_lock(&barrier3);
	counter3++;
	pthread_mutex_unlock(&barrier3);
	while(counter3 < NUMTHREADS);

	for (i = my_first_i; i < my_last_i && i < PROBLEM_SIZE; i++)
	{
		yt[i]=0.0;
		for (j = 0; j < PROBLEM_SIZE; j++)
		{
			yt[i] += c[i][j]*(y[j]+k3[j]);
		}
		k4[i] = h*(pow[i]-yt[i]);

		yout[i] = y[i] + (k1[i] + 2*k2[i] + 2*k3[i] + k4[i])/6.0;
		my_sum+=yout[i];
	}

	pthread_mutex_lock(&mutex1);
	totalSum += my_sum;
	pthread_mutex_unlock(&mutex1);
}

int main(int argc, char* argv[])
{
	int i,j;

	int num_threads = atoi(argv[1]);
	NUMTHREADS = num_threads;

	y    = (double* )malloc(PROBLEM_SIZE*sizeof(double));
	yt   = (double* )malloc(PROBLEM_SIZE*sizeof(double));
	k1   = (double* )malloc(PROBLEM_SIZE*sizeof(double));
	k2   = (double* )malloc(PROBLEM_SIZE*sizeof(double));
	k3   = (double* )malloc(PROBLEM_SIZE*sizeof(double));
	k4   = (double* )malloc(PROBLEM_SIZE*sizeof(double));
	pow  = (double* )malloc(PROBLEM_SIZE*sizeof(double));
	yout = (double* )malloc(PROBLEM_SIZE*sizeof(double));
	c    = (double**)malloc(PROBLEM_SIZE*sizeof(double*));

	for (i=0;i<PROBLEM_SIZE;i++)
	{
		c[i]=(double*)malloc(PROBLEM_SIZE*sizeof(double));
	}

	for (i = 0; i < PROBLEM_SIZE; i++)
	{
		y[i]=i*i;
		pow[i]=i+i;
		for (j = 0; j < PROBLEM_SIZE; j++)
		{
			c[i][j]=i*i+j;
		}
	}


	//Get the start time
	gettimeofday(&startTime, NULL);  /* START TIME */

	pthread_mutex_init(&barrier1, NULL);
	pthread_mutex_init(&barrier2, NULL);
	pthread_mutex_init(&barrier3, NULL);
	pthread_mutex_init(&mutex1, NULL);

	pthread_t tid[NUMTHREADS];
	void* status;
	for(i = 0; i < NUMTHREADS; i++)
	{
		pthread_create(&tid[i], NULL, threadWork, (void*)i);
	}

	for(i = 0; i < NUMTHREADS; i++)
	{
		pthread_join(tid[i], &status);
	}

	//Get the end time
	gettimeofday(&finishTime, NULL);  /* END TIME */

	//printf("\n\ntotalSum=%g\n\n",totalSum);


	//Calculate the interval length
	timeIntervalLength = (double)(finishTime.tv_sec-startTime.tv_sec) * 1000000
	                     + (double)(finishTime.tv_usec-startTime.tv_usec);
	timeIntervalLength=timeIntervalLength/1000;

	//Print the interval length
	printf("%g", timeIntervalLength);




	return 0;
}
