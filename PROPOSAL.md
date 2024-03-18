# Proposal: Add `simd` package to standard library

## Background

After doing a little bit of research on previous issues mentionning adding SIMD to Go, I found the following:

[Go 2: Add implicit single program multiple data](https://github.com/golang/go/issues/58610)

This proposal relies on the `ispmd` keyword which does not follow the compatibility guarantees for Go 1.

[add package for using SIMD instructions](https://github.com/golang/go/issues/53171)

This proposal is closer to what I am proposing however there are multiple packages per CPU architecture, and I'm providing a more complete explanation on how to achieve adding SIMD to Go standard library.

[crypto: replace assembly implementations with internal instrinsics](https://github.com/golang/go/issues/64634)

This proposal is more specific to the crypto package but in essence this is close to what I'm proposing here.

## Goals

The main goal of this proposal is to provide an alternative approach to designing a `simd` package for the Go standard library. As of right now, there is no concensus on what the API should look like.

This proposal mainly propose two things:

-   Adding a new kind of build tag that let user specify which SIMD ISA to use at compile time.
-   Using compiler intrinsics to generate inline SIMD instructions in the code.

### Build Tag

For the first point, I was thinking about something like the following:

```go
// +simd sse2
```
``` go
// +simd neon
```
```go
// +simd avx512
```

etc.

This would let the developer choose at compile time which SIMD ISA to target. We could write something similar to:

    $ go build -simd neon

And as such, we could take advantage of platform specific features and know at compile time the size of the vector registers (e.g. 128 or 256 bits). This would help us make better decisions for optimizations.

### Compiler Intrinsics

The next crucial step would be to create a portable SIMD package that would rely on the compiler to generate SIMD instruction through compiler intrinsics. I demonstrated that this is feasible with a [POC](https://github.com/Clement-Jean/simd-go-POC) (only working on ARM64 NEON for now). As of right now, it looks like the following (you can see more examples [here](https://github.com/Clement-Jean/simd-go-POC/blob/main/tests/arithm_test.go)):

```go
package main

import (
    "fmt"
    "simd"
)

func main() {
    a := &[16]uint8{...}
    b := &[16]uint8{...}
    c := simd.AddU8x16(a, b)

    fmt.Printf("%v\n", c)
}
```

And the `AddU8x16` gets lowered down to a `VADD` instruction after SSA lowering.

There are a few things to note:

-   We do not need to have type aliases. This bloats the code and we can simply rely on the good old fixed-size arrays.
-   We can provide functions like `AddU8x16`, `AddU8x32`, etc. without changing the generics implementation. Other implementations like [Highway](https://github.com/google/highway/blob/87ab8b81c9b11d8e28c9ebbd593bef7c39f7a61d/hwy/ops/arm_neon-inl.h#L801), [Rust std::simd](https://doc.rust-lang.org/std/simd/prelude/struct.Simd.html), and [Zig @Vector](https://ziglang.org/documentation/master/#Vectors) rely on generics for the API. In Go, we do not have non-type parameter in generics, thus we cannot have something like `Simd[uint8, 16]`.
-   We also do not have a compile time `Sizeof` which could have help us have:
    ```go
    type Simd[T SupportedSimdTypes] = [VectorRegisterSize / SizeofInBits(T)]T
    ```

## Challenges to Overcome

-   You might have noticed that the previous code snippet contained pointer on arrays and not arrays. This is because fixed-size arrays are not SSAable for now. I do not know much about why but @randall77 taught me this when I was working on the POC.
-   Because we work with pointers on arrays, the performance is not great (allocations and load of non contiguous memory) and it requires us to do the LD/ST dance for each function in the simd package.

## Scope

This proposal focuses on adding the following set of intrinsics:

- Splat
- Masking
- Arithmetic (wrapping and saturating)
- Bit (and, or, xor, shl, shr)
- Reduce (any, all, eq, neq)

more can be added later but we are trying to be realistic. We can discuss what is to be included.

## Why is it Important?

We are missing on a lot of optimizations. Here is a non exhaustive list of concrete things that could improve performance in daily life scenarios:

-   [simdjson](https://github.com/simdjson/simdjson)
-   [Decoding Billions of Integers Per Second Through Vectorization](https://people.csail.mit.edu/jshun/6886-s19/lectures/lecture19-1.pdf)
-   [Vectorized and performance-portable Quicksort](https://opensource.googleblog.com/2022/06/Vectorized%20and%20performance%20portable%20Quicksort.html)
-   [Introduction to Hyperscan](https://www.intel.com/content/www/us/en/developer/articles/technical/introduction-to-hyperscan.html)

Furthermore, it would make these currently existing packages more portable and maintainable:

-   [simdjson-go](https://github.com/minio/simdjson-go)
-   [sha256-simd](https://github.com/minio/sha256-simd)
-   [md5-simd](https://github.com/minio/md5-simd)

The are obviously way more applications of SIMD. I am just trying to say that this is useful in practical scenarios, not only in theory.
