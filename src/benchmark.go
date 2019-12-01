package main

import (
  "fmt"
  "os"
  "os/exec"
  "strconv"
  "strings"
)

const NUM_ITERATIONS = 10
const THREAD_STEP = 1
const MAX_THREAD = 40

// Compare rk4 and mmult in GO and in C for various number of threads
// Save results as CSV files in the results folder
func main() {
  fmt.Println("Benchmarking MMULT in C Pthreads")
  benchmarkC("mmult_pthread")
  // fmt.Println("Benchmarking MMULT in OMP...")
  // benchmarkOMP("mmult")
  // fmt.Println("Benchmarking MMULT in go...")
  // benchmarkGo("mmult_parallel.go")
  fmt.Println("Benchmarking RK4 in C Pthreads...")
  benchmarkC("rk4_pthread")
  // fmt.Println("Benchmarking RK4 in OMP...")
  // benchmarkOMP("rk4")
  // fmt.Println("Benchmarking RK4 in go...")
  // benchmarkGo("rk4_parallel.go")

}

// Benchmark the executable OMP program passed in
func benchmarkC(program string) {
  // create results file in RW mode
  out_file, file_err := os.OpenFile(fmt.Sprintf("../results/%s_c_benchmark.csv", program),
                  os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if file_err != nil {
    fmt.Println("Errored opening the results file")
    fmt.Println(file_err)
  }
  // Delete everything in the file (results from previous runs are destroyed)
  out_file.Truncate(0)
  results := make([]float64, 0)

  // This is the main work loop
  for i := 1; i <= MAX_THREAD; i += THREAD_STEP {
    thread_results := make([]float64, NUM_ITERATIONS)
    var current_min float64 = 0.0
    // for each iteration over number of threads, need to set environment variable
    env_err := os.Setenv("OMP_NUM_THREADS", strconv.Itoa(i))
    if env_err != nil {
      fmt.Println("Setting env went wrong")
    }
    fmt.Println("Currently working with", i, "thread(s)")
    for j := 0; j < NUM_ITERATIONS; j++ {
      // cmd is the command that will run in the form ../c_files/mmult
      cmd := exec.Command(fmt.Sprintf("../c_files/%s", program), fmt.Sprintf("%d", i))
      // This gets STDOUT and puts it in cmd_out
      cmd_out, cmd_err := cmd.Output()
      if cmd_err != nil {
        fmt.Println("Running command errored")
        fmt.Println(cmd_err)
        return
      }
      // Need to get useful time info out of the stdout
      time, _ := strconv.ParseFloat(string(cmd_out), 64)
      thread_results[j] = time
      // Keeps track of minimum time for final result
      if current_min == 0.0 || current_min > time {
        current_min = time
      }
    }
    results = append(results, current_min)
  }

  // Append all data in CSV format
  out_file.Write([]byte("threads, time \n"))
  for i := 1; i <= MAX_THREAD; i += THREAD_STEP {
    curr := results[0]
    results = results[1:]
    out_file.Write([]byte(fmt.Sprintf("%d, %f\n", i, curr)))
  }
  out_file.Close()
}

// Benchmark the executable OMP program passed in
func benchmarkOMP(program string) {
  // create results file in RW mode
  out_file, file_err := os.OpenFile(fmt.Sprintf("../results/%s_c_benchmark.csv", program),
                  os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if file_err != nil {
    fmt.Println("Errored opening the results file")
    fmt.Println(file_err)
  }
  // Delete everything in the file (results from previous runs are destroyed)
  out_file.Truncate(0)
  results := make([]float64, 0)

  // This is the main work loop
  for i := 1; i <= MAX_THREAD; i += THREAD_STEP {
    thread_results := make([]float64, NUM_ITERATIONS)
    var current_min float64 = 0.0
    // for each iteration over number of threads, need to set environment variable
    env_err := os.Setenv("OMP_NUM_THREADS", strconv.Itoa(i))
    if env_err != nil {
      fmt.Println("Setting env went wrong")
    }
    fmt.Println("Currently working with", i, "thread(s)")
    for j := 0; j < NUM_ITERATIONS; j++ {
      // cmd is the command that will run in the form ../c_files/mmult
      cmd := exec.Command(fmt.Sprintf("../c_files/%s", program))
      // This gets STDOUT and puts it in cmd_out
      cmd_out, cmd_err := cmd.Output()
      if cmd_err != nil {
        fmt.Println("Running command errored")
        fmt.Println(cmd_err)
        return
      }
      // Need to get useful time info out of the stdout
      time, _ := strconv.ParseFloat(string(cmd_out), 64)
      thread_results[j] = time
      // Keeps track of minimum time for final result
      if current_min == 0.0 || current_min > time {
        current_min = time
      }
    }
    results = append(results, current_min)
  }

  // Append all data in CSV format
  out_file.Write([]byte("threads, time \n"))
  for i := 1; i <= MAX_THREAD; i += THREAD_STEP {
    curr := results[0]
    results = results[1:]
    out_file.Write([]byte(fmt.Sprintf("%d, %f\n", i, curr)))
  }
  out_file.Close()
}

// Will benchmark a GO file provided. I tried to call the functions themselves,
// but I had trouble with GO finding the package they were in and all that.
// Therefore, this method just calls 'go run ' on 'program'
func benchmarkGo(program string) {
  const NANO_IN_MILLI = 1000000.0
  // create results file in RW mode
  out_file, file_err := os.OpenFile(fmt.Sprintf("../results/%s_go_benchmark.csv", program),
                  os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if file_err != nil {
    fmt.Println("Errored opening the results file")
    fmt.Println(file_err)
  }
  out_file.Truncate(0)
  results := make([]float64, 0)
  // Test with different number of threads
  for i := 1; i <= MAX_THREAD; i += THREAD_STEP {
    thread_results := make([]float64, NUM_ITERATIONS)
    var current_min float64 = 0.0
    fmt.Println("Currently working with", i, "threads")
    for j := 0; j < NUM_ITERATIONS; j++ {
      // for each iteration over number of threads, need to set cmd-line args correctly
      cmd := exec.Command("go", "run", program, strconv.Itoa(i))
      cmd_out, cmd_err := cmd.Output()
      if cmd_err != nil {
        fmt.Println("Running command errored")
        fmt.Println(cmd_err)
        return
      }
      // Turns out that our GO implementation prints a time.Duration type, which
      // is in nanoseconds if you convert it to an int64.
      time, time_err := strconv.ParseInt(strings.TrimSpace(string(cmd_out)), 10, 64)
      if time_err != nil {
        fmt.Println(time_err)
      }
      time_ms := float64(time) / NANO_IN_MILLI
      fmt.Println("time is", time_ms)
      thread_results[j] = time_ms
      if current_min == 0.0 || current_min > time_ms {
        current_min = time_ms
      }
    }
    results = append(results, current_min)
  }
  fmt.Println(results)

  out_file.Write([]byte("threads, time \n"))
  for i := 1; i <= MAX_THREAD; i += THREAD_STEP {
    curr := results[0]
    results = results[1:]
    out_file.Write([]byte(fmt.Sprintf("%d, %f\n", i, curr)))
  }
  out_file.Close()
}
