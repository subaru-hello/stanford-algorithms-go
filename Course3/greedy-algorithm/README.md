run the greedy algorithm that schedules in decreasing order of the difference(weight - length).

if two jobs are equal difference(weight - length), you should schedule the job with higher weight first.

report sum of weighted completion times(cummulative times of job length) of the resulting schedule -- a positive integer --

[numbers of jobs]
[first job weight] [first job length]
[second job weight] [second job length]

### pseudo code

in general greedy algorithm focuses on easiness and efficiency of solving a problem, this scheduling problem deals with calculating a total completion times which in other words say, select as much as jobs. Selecting jobs as much as possible denote that we should select smaller job length but bigger job weight.

Taking the information given so far consideration, 3 steps are required.
1 sort jobs with descneding the difference(weight - length) order.
2 take as much as jobs(job length) until all the given numbers of job are seen through, if the difference are equal for comparing one, take the higher weight one.
3 return total completion times which we select jobs so far.

Basically, it has phases, processing, querying.

#### processing

receive inputs. N is numbers of jobs, Wi = i th job weight, Li = i th job length, completionTime = sum of weighted completion time.
