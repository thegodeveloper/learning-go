# Language Mechanics On Memory Profiling

## Benchmark

```commandline
go test -run none -bench AlgorithmOne -benchtime 3s -benchmem
```

### Result

```commandline
goos: linux
goarch: amd64
pkg: github.com/williammunozr/learning-go/language_mechanics/memory_profiling
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
BenchmarkAlgorithmOne-8   	 2801496	      1293 ns/op	      53 B/op	       2 allocs/op
PASS
ok  	github.com/williammunozr/learning-go/language_mechanics/memory_profiling	4.924s
```
## Profiling AlgorithmOne

```commandline
go test -run none -bench AlgorithmOne -benchtime 3s -benchmem -memprofile mem.out
```

### Run the pprof

```commandline
go tool pprof -alloc_space memory_profiling.test mem.out                                                                               ─╯
File: memory_profiling.test
Type: alloc_space
Time: Mar 8, 2023 at 10:36am (-05)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof)
```

### Inspect the algOne function

```commandline
(pprof) list algOne
Total: 194.51MB
ROUTINE ======================== github.com/williammunozr/learning-go/language_mechanics/memory_profiling.algOne in /home/william/go/src/github.com/williammunozr/learning-go/language_mechanics/memory_profiling/main.go
   12.50MB   194.51MB (flat, cum)   100% of Total
         .          .     12:func algOne(data []byte, find []byte, repl []byte, output *bytes.Buffer) {
         .          .     13:	// use a bytes buffer to provide a stream to process.
         .   182.01MB     14:	input := bytes.NewReader(data)
         .          .     15:
         .          .     16:	// the number of bytes we are looking for.
         .          .     17:	size := len(find)
         .          .     18:
         .          .     19:	// declare the buffers we need to process the stream.
   12.50MB    12.50MB     20:	buf := make([]byte, size)
         .          .     21:	end := size - 1
         .          .     22:
         .          .     23:	// read in an initial number of bytes we need to get started.
         .          .     24:	if n, err := io.ReadFull(input, buf[:end]); err != nil {
         .          .     25:		output.Write(buf[:n])
(pprof)
```

### List Benchmark

```commandline
(pprof) list Benchmark
Total: 194.51MB
ROUTINE ======================== github.com/williammunozr/learning-go/language_mechanics/memory_profiling.BenchmarkAlgorithmOne in /home/william/go/src/github.com/williammunozr/learning-go/language_mechanics/memory_profiling/benchmark_test.go
         0   194.51MB (flat, cum)   100% of Total
         .          .      8:func BenchmarkAlgorithmOne(b *testing.B) {
         .          .      9:	var output bytes.Buffer
         .          .     10:	in := assembleInputStream()
         .          .     11:	find := []byte("elvis")
         .          .     12:	repl := []byte("Elvis")
         .          .     13:
         .          .     14:	b.ResetTimer()
         .          .     15:
         .          .     16:	for i := 0; i < b.N; i++ {
         .          .     17:		output.Reset()
         .   194.51MB     18:		algOne(in, find, repl, &output)
         .          .     19:	}
         .          .     20:}
(pprof)
```

## Compiler Reporting

### Build the program

```commandline
go build -gcflags "-m -m"
```

### Output

```commandline
# github.com/williammunozr/learning-go/language_mechanics/memory_profiling
./main.go:8:6: can inline main with cost 0 as: func() {  }
./main.go:12:6: cannot inline algOne: function too complex: cost 698 exceeds budget 80
./main.go:14:26: inlining call to bytes.NewReader
./main.go:24:26: inlining call to io.ReadFull
./main.go:31:27: inlining call to io.ReadFull
./main.go:38:19: inlining call to bytes.Compare
./main.go:42:28: inlining call to io.ReadFull
./main.go:58:6: can inline assembleInputStream with cost 3 as: func() []byte { return ([]byte)("abcelvisaElvisabcelviseelvisaelvisaabeeeelvise l v i saa bb e l v i saa elvi\nselvielviselvielvielviselvi1elvielviselvis") }
./main.go:20:13: make([]byte, size) escapes to heap:
./main.go:20:13:   flow: {heap} = &{storage for make([]byte, size)}:
./main.go:20:13:     from make([]byte, size) (non-constant size) at ./main.go:20:13
./main.go:14:26: &bytes.Reader{...} escapes to heap:
./main.go:14:26:   flow: ~R0 = &{storage for &bytes.Reader{...}}:
./main.go:14:26:     from &bytes.Reader{...} (spill) at ./main.go:14:26
./main.go:14:26:     from ~R0 = &bytes.Reader{...} (assign-pair) at ./main.go:14:26
./main.go:14:26:   flow: input = ~R0:
./main.go:14:26:     from input := ~R0 (assign) at ./main.go:14:8
./main.go:14:26:   flow: io.r = input:
./main.go:14:26:     from input (interface-converted) at ./main.go:42:29
./main.go:14:26:     from io.r, io.buf := input, buf[:end] (assign-pair) at ./main.go:42:28
./main.go:14:26:   flow: {heap} = io.r:
./main.go:14:26:     from io.ReadAtLeast(io.r, io.buf, len(io.buf)) (call parameter) at ./main.go:42:28
./main.go:12:13: parameter data leaks to {storage for &bytes.Reader{...}} with derefs=0:
./main.go:12:13:   flow: bytes.b = data:
./main.go:12:13:     from bytes.b := data (assign-pair) at ./main.go:14:26
./main.go:12:13:   flow: {storage for &bytes.Reader{...}} = bytes.b:
./main.go:12:13:     from bytes.Reader{...} (struct literal element) at ./main.go:14:26
./main.go:12:13: leaking param: data
./main.go:12:26: find does not escape
./main.go:12:39: repl does not escape
./main.go:12:52: output does not escape
./main.go:14:26: &bytes.Reader{...} escapes to heap
./main.go:20:13: make([]byte, size) escapes to heap
./main.go:59:16: ([]byte)("abcelvisaElvisabcelviseelvisaelvisaabeeeelvise l v i saa bb e l v i saa elvi\nselvielviselvielvielviselvi1elvielviselvis") escapes to heap:
./main.go:59:16:   flow: ~r0 = &{storage for ([]byte)("abcelvisaElvisabcelviseelvisaelvisaabeeeelvise l v i saa bb e l v i saa elvi\nselvielviselvielvielviselvi1elvielviselvis")}:
./main.go:59:16:     from ([]byte)("abcelvisaElvisabcelviseelvisaelvisaabeeeelvise l v i saa bb e l v i saa elvi\nselvielviselvielvielviselvi1elvielviselvis") (spill) at ./main.go:59:16
./main.go:59:16:     from return ([]byte)("abcelvisaElvisabcelviseelvisaelvisaabeeeelvise l v i saa bb e l v i saa elvi\nselvielviselvielvielviselvi1elvielviselvis") (return) at ./main.go:59:2
./main.go:59:16: ([]byte)("abcelvisaElvisabcelviseelvisaelvisaabeeeelvise l v i saa bb e l v i saa elvi\nselvielviselvielvielviselvi1elvielviselvis") escapes to heap
```

### Lines to check

```commandline
./main.go:20:13: make([]byte, size) escapes to heap:
./main.go:14:26: &bytes.Reader{...} escapes to heap
./main.go:20:13: make([]byte, size) escapes to heap
```
## Profiling AlgorithmTwo

```commandline
go test -run none -bench AlgorithmTwo -benchtime 3s -benchmem -memprofile memtwo.out
```

Output:

```commandline
goos: linux
goarch: amd64
pkg: github.com/williammunozr/learning-go/language_mechanics/memory_profiling
cpu: 11th Gen Intel(R) Core(TM) i5-1135G7 @ 2.40GHz
BenchmarkAlgorithmTwo-8   	 3681998	       978.6 ns/op	       5 B/op	       1 allocs/op
PASS
ok  	github.com/williammunozr/learning-go/language_mechanics/memory_profiling	4.596s
```

### Profiling AlgorithmTwo

```commandline
go tool pprof -alloc_space memcpu.test memtwo.out
ile: memory_profiling.test
Type: alloc_space
Time: Mar 8, 2023 at 11:41am (-05)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) list algTwo
Total: 29MB
ROUTINE ======================== github.com/williammunozr/learning-go/language_mechanics/memory_profiling.algTwo in /home/william/go/src/github.com/williammunozr/learning-go/language_mechanics/memory_profiling/main.go
   28.50MB    28.50MB (flat, cum) 98.27% of Total
         .          .     62:func algTwo(data []byte, find []byte, repl []byte, output *bytes.Buffer) {
         .          .     63:	// use a bytes buffer to provide a stream to process.
         .          .     64:	input := bytes.NewBuffer(data)
         .          .     65:
         .          .     66:	// the number of bytes we are looking for.
         .          .     67:	size := len(find)
         .          .     68:
         .          .     69:	// declare the buffers we need to process the stream.
   28.50MB    28.50MB     70:	buf := make([]byte, size)
         .          .     71:	end := size - 1
         .          .     72:
         .          .     73:	// read in an initial number of bytes we need to get started.
         .          .     74:	if n, err := input.Read(buf[:end]); err != nil || n < end {
         .          .     75:		output.Write(buf[:n])
(pprof)
```

### Build the code

```commandline
go build -gcflags "-m -m" 
```

Output:

```commandline
# github.com/williammunozr/learning-go/language_mechanics/memory_profiling
./main.go:8:6: can inline main with cost 0 as: func() {  }
./main.go:12:6: cannot inline algOne: function too complex: cost 698 exceeds budget 80
./main.go:14:26: inlining call to bytes.NewReader
./main.go:24:26: inlining call to io.ReadFull
./main.go:31:27: inlining call to io.ReadFull
./main.go:38:19: inlining call to bytes.Compare
./main.go:42:28: inlining call to io.ReadFull
./main.go:58:6: can inline assembleInputStream with cost 3 as: func() []byte { return ([]byte)("abcelvisaElvisabcelviseelvisaelvisaabeeeelvise l v i saa bb e l v i saa elvi\nselvielviselvielvielviselvi1elvielviselvis") }
./main.go:62:6: cannot inline algTwo: function too complex: cost 684 exceeds budget 80
./main.go:64:26: inlining call to bytes.NewBuffer
./main.go:74:25: inlining call to bytes.(*Buffer).Read
./main.go:81:26: inlining call to bytes.(*Buffer).Read
./main.go:88:19: inlining call to bytes.Compare
./main.go:92:27: inlining call to bytes.(*Buffer).Read
./main.go:74:25: inlining call to bytes.(*Buffer).empty
./main.go:74:25: inlining call to bytes.(*Buffer).Reset
./main.go:81:26: inlining call to bytes.(*Buffer).empty
./main.go:81:26: inlining call to bytes.(*Buffer).Reset
./main.go:92:27: inlining call to bytes.(*Buffer).empty
./main.go:92:27: inlining call to bytes.(*Buffer).Reset
./main.go:20:13: make([]byte, size) escapes to heap:
./main.go:20:13:   flow: {heap} = &{storage for make([]byte, size)}:
./main.go:20:13:     from make([]byte, size) (non-constant size) at ./main.go:20:13
./main.go:14:26: &bytes.Reader{...} escapes to heap:
./main.go:14:26:   flow: ~R0 = &{storage for &bytes.Reader{...}}:
./main.go:14:26:     from &bytes.Reader{...} (spill) at ./main.go:14:26
./main.go:14:26:     from ~R0 = &bytes.Reader{...} (assign-pair) at ./main.go:14:26
./main.go:14:26:   flow: input = ~R0:
./main.go:14:26:     from input := ~R0 (assign) at ./main.go:14:8
./main.go:14:26:   flow: io.r = input:
./main.go:14:26:     from input (interface-converted) at ./main.go:42:29
./main.go:14:26:     from io.r, io.buf := input, buf[:end] (assign-pair) at ./main.go:42:28
./main.go:14:26:   flow: {heap} = io.r:
./main.go:14:26:     from io.ReadAtLeast(io.r, io.buf, len(io.buf)) (call parameter) at ./main.go:42:28
./main.go:12:13: parameter data leaks to {storage for &bytes.Reader{...}} with derefs=0:
./main.go:12:13:   flow: bytes.b = data:
./main.go:12:13:     from bytes.b := data (assign-pair) at ./main.go:14:26
./main.go:12:13:   flow: {storage for &bytes.Reader{...}} = bytes.b:
./main.go:12:13:     from bytes.Reader{...} (struct literal element) at ./main.go:14:26
./main.go:12:13: leaking param: data
./main.go:12:26: find does not escape
./main.go:12:39: repl does not escape
./main.go:12:52: output does not escape
./main.go:14:26: &bytes.Reader{...} escapes to heap
./main.go:20:13: make([]byte, size) escapes to heap
./main.go:59:16: ([]byte)("abcelvisaElvisabcelviseelvisaelvisaabeeeelvise l v i saa bb e l v i saa elvi\nselvielviselvielvielviselvi1elvielviselvis") escapes to heap:
./main.go:59:16:   flow: ~r0 = &{storage for ([]byte)("abcelvisaElvisabcelviseelvisaelvisaabeeeelvise l v i saa bb e l v i saa elvi\nselvielviselvielvielviselvi1elvielviselvis")}:
./main.go:59:16:     from ([]byte)("abcelvisaElvisabcelviseelvisaelvisaabeeeelvise l v i saa bb e l v i saa elvi\nselvielviselvielvielviselvi1elvielviselvis") (spill) at ./main.go:59:16
./main.go:59:16:     from return ([]byte)("abcelvisaElvisabcelviseelvisaelvisaabeeeelvise l v i saa bb e l v i saa elvi\nselvielviselvielvielviselvi1elvielviselvis") (return) at ./main.go:59:2
./main.go:59:16: ([]byte)("abcelvisaElvisabcelviseelvisaelvisaabeeeelvise l v i saa bb e l v i saa elvi\nselvielviselvielvielviselvi1elvielviselvis") escapes to heap
./main.go:74:25: algTwo ignoring self-assignment in bytes.b.buf = bytes.b.buf[:0]
./main.go:81:26: algTwo ignoring self-assignment in bytes.b.buf = bytes.b.buf[:0]
./main.go:92:27: algTwo ignoring self-assignment in bytes.b.buf = bytes.b.buf[:0]
./main.go:70:13: make([]byte, size) escapes to heap:
./main.go:70:13:   flow: {heap} = &{storage for make([]byte, size)}:
./main.go:70:13:     from make([]byte, size) (non-constant size) at ./main.go:70:13
./main.go:62:13: data does not escape
./main.go:62:26: find does not escape
./main.go:62:39: repl does not escape
./main.go:62:52: output does not escape
./main.go:64:26: &bytes.Buffer{...} does not escape
./main.go:70:13: make([]byte, size) escapes to heap
```

### Analysis

During the build process we can see the following line is escaping:

```commandline
./main.go:70:13: make([]byte, size) escapes to heap:
./main.go:70:13:   flow: {heap} = &{storage for make([]byte, size)}:
./main.go:70:13:     from make([]byte, size) (non-constant size) at ./main.go:70:13
```

This is happening because values can only be placed on the stack if the compiler knows the size of the value at compile time.

### Temporarily hard code the size

```commandline
buf := make([]byte, 5)
```

```commandline
Compile the program again, and we should see the following:

./main.go:71:13: make([]byte, 5) does not escape
```
