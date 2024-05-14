# simd-go-POC

This repository is an attempt at adding SIMD to Go through compiler intrinsics.

## Init

```
git submodule update --init --recursive
```

## Build

```
pushd go && ../build.sh && popd
```

## Run

```
go/bin/go run [-tags neon|sse2|sse4.1] ./examples/{NAME}.go
```

## Architecture

- patches -> patching existing files
- overlays -> new dirs+files

## Supported Intrinsics

legend:
  - :no_entry_sign: not available
  - :white_check_mark: implemented
  - :x: not implemented (maybe unavailable)

| Intrinsic            | NEON               | SSE2               | SSE4.1             | SSSE3              |
|----------------------|--------------------|--------------------|--------------------|--------------------|
| `Add8x16`            | :white_check_mark: | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `AddU8x16`           | :white_check_mark: | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `Sub8x16`            | :white_check_mark: | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `SubU8x16`           | :white_check_mark: | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `SaturatingAdd8x16`  | :white_check_mark: | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `SaturatingAddU8x16` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `SaturatingSub8x16`  | :white_check_mark: | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `SaturatingSubU8x16` | :white_check_mark: | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `And8x16`            | :white_check_mark: | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `AndU8x16`           | :white_check_mark: | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `And16x8`            | :x:                | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `AndU16x8`           | :x:                | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `Or8x16`             | :white_check_mark: | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `OrU8x16`            | :white_check_mark: | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `Xor8x16`            | :white_check_mark: | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `Xor8x16`            | :white_check_mark: | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `ShiftLeftU8x16`     | :white_check_mark: | :no_entry_sign:    | :no_entry_sign:    | :no_entry_sign:    |
| `ShiftLeft8x16`      | :white_check_mark: | :no_entry_sign:    | :no_entry_sign:    | :no_entry_sign:    |
| `ShiftRightU8x16`    | :white_check_mark: | :no_entry_sign:    | :no_entry_sign:    | :no_entry_sign:    |
| `ShiftRight8x16`     | :white_check_mark: | :no_entry_sign:    | :no_entry_sign:    | :no_entry_sign:    |
| `ShiftRightU16x8`    | :x:                | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `ShiftRight16x8`     | :x:                | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `Max8x16`            | :white_check_mark: | :no_entry_sign:    | :white_check_mark: | :no_entry_sign:    |
| `MaxU8x16`           | :white_check_mark: | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `Min8x16`            | :white_check_mark: | :no_entry_sign:    | :white_check_mark: | :no_entry_sign:    |
| `MinU8x16`           | :white_check_mark: | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `ReduceMax8x16`      | :white_check_mark: | :no_entry_sign:    | :no_entry_sign:    | :no_entry_sign:    |
| `ReduceMaxU8x16`     | :white_check_mark: | :no_entry_sign:    | :no_entry_sign:    | :no_entry_sign:    |
| `ReduceMin8x16`      | :white_check_mark: | :no_entry_sign:    | :no_entry_sign:    | :no_entry_sign:    |
| `ReduceMinU8x16`     | :white_check_mark: | :x:                | :x:                | :no_entry_sign:    |
| `ExtractU8x16`       | :white_check_mark: | :x:                | :x:                | :white_check_mark: |
| `Extract8x16`        | :white_check_mark: | :x:                | :x:                | :white_check_mark: |
| `LookupU8x16`        | :white_check_mark: | :x:                | :x:                | :white_check_mark: |
| `Lookup8x16`         | :white_check_mark: | :x:                | :x:                | :white_check_mark: |
| `AllZerosU8x16`      | :x:                | :x:                | :white_check_mark: | :no_entry_sign:    |
| `AllZeros8x16`       | :x:                | :x:                | :white_check_mark: | :no_entry_sign:    |
| `MovMaskByteU8x16`   | :no_entry_sign:    | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `MovMaskByte8x16`    | :no_entry_sign:    | :white_check_mark: | :white_check_mark: | :no_entry_sign:    |
| `SplatU8x16`         | :white_check_mark: | :x:                | :x:                | :no_entry_sign:    |
| `Splat8x16`          | :white_check_mark: | :x:                | :x:                | :no_entry_sign:    |

