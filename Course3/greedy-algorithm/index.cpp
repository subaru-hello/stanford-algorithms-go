
#include <algorithm>  // for sorting.
#include <iostream>
#include <utility>  // for pair
using namespace std;

using Job = pair<int, int>;  // weight, completion time

// diffが同じ場合、weightが高い方を選択
// void template(const Job &a, const Job &b){
//     if(a == b){
//         return
//     }
// }

int main() {
  int N;
  vector<long long> W, T;
  cin >> N;
  vector<Job> jobs;

  // W[i] = jobiの重み, T[i] = jobiを終わらせるのにかかる時間
  for (int i = 0; i <= N; ++i) cin >> W[i] >> T[i];
  for (int i = 0; i <= N; ++i) {
    jobs.push_back(make_pair(W[i], T[i]));
  }

  // Example output: print each job's weight and completion time
  for () {
    cout << "Weight: " << job.first << ", Completion Time: " << job.second
         << endl;
  }
  return 0;
}