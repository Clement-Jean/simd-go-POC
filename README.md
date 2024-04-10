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
go/bin/go run [-tags neon|sse2|sse4.1] .
```

## Architecture

- patches -> patching existing files
- overlays -> new dirs+files

## Supported Intrinsics

legend:
  - :no_entry_sign: not available
  - :white_check_mark: implemented
  - :x: not implemented

| Intrinsic            | NEON               | SSE2               | SSE4.1             |
|----------------------|--------------------|--------------------|--------------------|
| `Add8x16`            | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| `AddU8x16`           | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| `Sub8x16`            | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| `SubU8x16`           | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| `SaturatingAdd8x16`  | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| `SaturatingAddU8x16` | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| `SaturatingSub8x16`  | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| `SaturatingSubU8x16` | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| `And8x16`            | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| `AndU8x16`           | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| `Or8x16`             | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| `OrU8x16`            | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| `Xor8x16`            | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| `Xor8x16`            | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| `Max8x16`            | :white_check_mark: | :no_entry_sign:    | :white_check_mark: |
| `MaxU8x16`           | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| `Min8x16`            | :white_check_mark: | :no_entry_sign:    | :white_check_mark: |
| `MinU8x16`           | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| `ReduceMax8x16`      | :white_check_mark: | :no_entry_sign:    | :white_check_mark: |
| `ReduceMaxU8x16`     | :white_check_mark: | :no_entry_sign:    | :white_check_mark: |
| `ReduceMin8x16`      | :white_check_mark: | :no_entry_sign:    | :white_check_mark: |
| `ReduceMinU8x16`     | :white_check_mark: | :no_entry_sign:    | :white_check_mark: |

