# Proposal: Add `simd` package to standard library

## Background

After doing a little bit of research on previous issues mentioning adding SIMD to Go, I found the following proposals:

[Go 2: Add implicit single program multiple data](https://github.com/golang/go/issues/58610)

This proposal relies on the `ispmd` keyword which does not follow the compatibility guarantees for Go 1.

[add package for using SIMD instructions](https://github.com/golang/go/issues/53171)

This proposal is closer to what I am proposing, however, there are multiple packages per CPU architecture, and I'm providing a more complete explanation on how to achieve adding SIMD to the Go standard library.

[crypto: replace assembly implementations with internal instrinsics](https://github.com/golang/go/issues/64634)

This proposal is more specific to the crypto package but in essence this is close to what I'm proposing here.

## Goals

The main goal of this proposal is to provide an alternative approach to designing a `simd` package for the Go standard library. As of right now, there is no consensus on what the API should look like and this proposal intends to drive the discussion further.

This proposal mainly includes two things:

- Adding a new kind of build tag that lets the user specify which SIMD ISA to use at compile time.
- Using compiler intrinsics to generate inline SIMD instructions in the code.

### Build Tag

For the first point, I was thinking about optional build tags like the following:

```go
//go:simd sse2
```
```go
//go:simd neon
```
```go
//go:simd avx512
```

etc.

As mentioned, these build tags are optional. If not specified, we would resolve to the appropriate SIMD ISA available on the current OS and ARCH. However, I still think that these build tags are needed for deeper optimization and platform-specific operations. If we know that some instruction is more performant or only available on a certain architecture, we should be able to enforce using it manually.

Finally, having the optional build tag would let the developer choose, at compile time, which SIMD ISA to target and thus cross compile. We could write something similar to:

    $ go build -simd neon ...

With this, we could take advantage of platform-specific features and know at compile time the size of the vector registers (e.g., 128, 256, or 512 bits). This would help us make better decisions for optimizations on the compiler side.

### Compiler Intrinsics

The next crucial step would be to create a portable SIMD package that would rely on the compiler to generate SIMD instructions through compiler intrinsics. I demonstrated that this is feasible with a [POC](https://github.com/Clement-Jean/simd-go-POC). As of right now, it looks like the following (you can see more examples [here](https://github.com/Clement-Jean/simd-go-POC/blob/main/examples), including UTF8 string validation):

```go
package main

import (
    "fmt"
    "simd"
)

func main() {
    a := simd.Uint8x16{...}
    b := simd.Uint8x16{...}
    c := simd.AddU8x16(a, b)

    fmt.Printf("%v\n", c)
}
```

And here, the `AddU8x16` gets lowered down to a `vector add` instruction after SSA lowering.

### Notes

- We can provide functions like `AddU8x16`, `AddU8x32`, etc. without changing the generics implementation. Other implementations like [Highway](https://github.com/google/highway/blob/87ab8b81c9b11d8e28c9ebbd593bef7c39f7a61d/hwy/ops/arm_neon-inl.h#L801), [Rust std::simd](https://doc.rust-lang.org/std/simd/prelude/struct.Simd.html), and [Zig @Vector](https://ziglang.org/documentation/master/#Vectors) rely on generics for the API. In Go, we do not have non-type parameters in generics; thus we cannot have something like `Simd[uint8, 16]`.

- Also, we do not have a compile time `Sizeof` which could help us have:
    ```go
    type Simd[T SupportedSimdTypes] = [VectorRegisterSize / SizeofInBits(T)]T
    ```

## Philosophy

It is important to understand that this proposal is not describing an abstraction of SIMD features. Such an abstraction could create noticeable performance difference between ISAs. That's why we are trying to avoid it. This means that **if an operation is not available on an ISA, we simply don't provide it**. Each intrinsic should only have 1 underlying instruction, not a sequence of instructions.

## Challenges to Overcome

- Under the hood, the current POC, works with pointers on arrays. The main reason is that fixed-size arrays are not currently SSAable. But because we work with these pointers on arrays which are stored in general purpose registers, the performance is not great (allocations required and load of non-contiguous memory) and it requires us to do the LD/ST dance for each function in the simd package.

I believe we would need some kind of type aliases like the following:

```go
type Int8x16 [16]int8
type Uint8x16 [16]uint8
//...
```

These types should not be indexable and only instantiable through init functions like `Splat8x16`, `Load8x16`, etc. The compiler would then promote these special types to vector registers. This would remove all the LD/ST dance and memory allocation that I have in my POC and thus make everything a lot faster.

In the end the previous code snippet could look like this:

```go
package main

import (
    "fmt"
    "simd"
)

func main() {
    a := simd.SplatU8x16(1)
    b := simd.LoadU8x16([16]uint8{...})
    c := simd.AddU8x16(a, b)

    fmt.Printf("%v\n", c)
}
```

- A lot of instructions on arm (and I suppose on other ISAs) are missing. The current POC is using constants (`WORD $0x...`) to encode these.

I believe we should avoid having to use constants when defining intrinsics. We could implement the missing instructions along the way with the implementation of intrinsics.

- Naming these intrinsics is not always easy. For example, NEON has instructions called `VMIN` and `UMINV`. The former returns a vector of min elements, and the latter reduce to the minimum in a given vector. As we don't have function overloads, we will need to find a way to name them appropriately.

I believe we should try to make the horizontal operations more verbose (e.g. `ReduceMin8x16`) and promote the vertical ones (e.g. `Min8x16`). For the example of `VMIN` and `UMINV`, the latter does not even seem to exist in SSE2 whereas the first one does.

There are other operations that have "conflicting" names. For example `shift right` and `shift left` have both `logical` and `arithmetic` shifts. For these cases, I believe we could just call them `LogicalShiftRight`, `ArithmeticShiftRight`, ... I agree that this is verbose but it makes it clear what is happening.

- The current POC did not implement the concept of Masks. This is an important concept but also a tricky one to implement without proper compiler support. After discussion with [Jan Wassenberg](https://github.com/jan-wassenberg) (author of [Highway](https://github.com/google/highway)), I realized that some platforms do not treat masks in the same way. Here is a summary:

   - on NEON, SSE4, and AVX2: 1 bit per bit of the vector.
   - on SVE: 1 bit per byte of vector (variable vector size).
   - on AVX-512, and RVV: 1 bit per lane.
   - AVX-512 has a separate register file for masks.

We could have other types aliases, like the one made for vectors. These types would have different shapes and be loaded differently depending on the platform.

- Some operations need parameters that are known at compile time and within a certain range. For example `SSHR` (shift right) on NEON takes an immediate `n` that need to be restricted between 1 and 8 (see [vshr_n_s8](https://developer.arm.com/architectures/instruction-sets/intrinsics/vshr_n_s8)). I ran into some problems where during the build of the compiler, the `n` would resolve to 0 (default value of int passed as param) and it would crash the program.

I believe we could have some way to check at compile time these values are within a certain range. Like a [static_assert](https://en.cppreference.com/w/cpp/language/static_assert) in C++ or checks on the AST.

## Why is it Important?

Without SIMD, we are missing on a lot of potential optimizations. Here is a non-exhaustive list of concrete things that could improve performance in daily life scenarios:

- [simdjson](https://github.com/simdjson/simdjson)
- [Decoding Billions of Integers Per Second Through Vectorization](https://people.csail.mit.edu/jshun/6886-s19/lectures/lecture19-1.pdf)
- [Vectorized and performance-portable Quicksort](https://opensource.googleblog.com/2022/06/Vectorized%20and%20performance%20portable%20Quicksort.html)
- [Introduction to Hyperscan](https://www.intel.com/content/www/us/en/developer/articles/technical/introduction-to-hyperscan.html)

Furthermore, it would make these currently existing packages more portable and maintainable:

- [simdjson-go](https://github.com/minio/simdjson-go)
- [sha256-simd](https://github.com/minio/sha256-simd)
- [md5-simd](https://github.com/minio/md5-simd)

There are obviously way more applications of SIMD but I am just trying to say that this is useful in practical scenarios.
