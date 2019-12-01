#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <sys/time.h>
#include <omp.h>


#define PROBLEM_SIZE 8192


struct timeval startTime;
struct timeval finishTime;
double timeIntervalLength;


int main(int argc, char* argv[])
{
	int i,j;
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
#	pragma omp parallel default(none) shared(yt, c, y, yout, k1, k2, k3, k4, h, pow, totalSum) \
			private(i, j)
	{
		#pragma omp for schedule(guided)
		for (i = 0; i < PROBLEM_SIZE; i++)
		{
			yt[i] = 0.0;
			for (j = 0; j < PROBLEM_SIZE; j++)
			{
				yt[i] += c[i][j]*y[j];
			}
			k1[i] = h*(pow[i]-yt[i]);
		}

		#pragma omp for schedule(guided)
		for (i = 0; i < PROBLEM_SIZE; i++)
		{
			yt[i] = 0.0;
			for (j = 0; j < PROBLEM_SIZE; j++)
			{
				yt[i] += c[i][j]*(y[j]+0.5*k1[j]);
			}
			k2[i] = h*(pow[i]-yt[i]);
		}

		#pragma omp for schedule(guided)
		for (i = 0; i < PROBLEM_SIZE; i++)
		{
			yt[i] = 0.0;
			for (j = 0; j < PROBLEM_SIZE; j++)
			{
				yt[i] += c[i][j]*(y[j]+0.5*k2[j]);
			}
			k3[i] = h*(pow[i]-yt[i]);
		}

		#pragma omp for schedule(guided) reduction(+: totalSum)
		for (i =0; i < PROBLEM_SIZE; i++)
		{
			yt[i]=0.0;
			for (j = 0; j < PROBLEM_SIZE; j++)
			{
				yt[i] += c[i][j]*(y[j]+k3[j]);
			}
			k4[i] = h*(pow[i]-yt[i]);

			yout[i] = y[i] + (k1[i] + 2*k2[i] + 2*k3[i] + k4[i])/6.0;
			totalSum+=yout[i];
		}
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
