# simd-go-POC

This repository is an attempt at adding SIMD to Go through compiler intrinsics.

## Supported Intrinsics

| Intrinsic            | NEON               | SSE2 |
|----------------------|--------------------|------|
| `Add8x16`            | :white_check_mark: | :x:  |
| `AddU8x16`           | :white_check_mark: | :x:  |
| `Sub8x16`            | :white_check_mark: | :x:  |
| `SubU8x16`           | :white_check_mark: | :x:  |
| `SaturatingAdd8x16`  | :white_check_mark: | :x:  |
| `SaturatingAddU8x16` | :white_check_mark: | :x:  |
| `SaturatingSub8x16`  | :white_check_mark: | :x:  |
| `SaturatingSubU8x16` | :white_check_mark: | :x:  |
| `And8x16`            | :white_check_mark: | :x:  |
| `AndU8x16`           | :white_check_mark: | :x:  |
| `Or8x16`             | :white_check_mark: | :x:  |
| `OrU8x16`            | :white_check_mark: | :x:  |
| `Xor8x16`            | :white_check_mark: | :x:  |
| `Xor8x16`            | :white_check_mark: | :x:  |
| `Max8x16`            | :white_check_mark: | :x:  |
| `MaxU8x16`           | :white_check_mark: | :x:  |
| `Min8x16`            | :white_check_mark: | :x:  |
| `MinU8x16`           | :white_check_mark: | :x:  |
| `ReduceMax8x16`      | :white_check_mark: | :x:  |
| `ReduceMaxU8x16`     | :white_check_mark: | :x:  |
| `ReduceMin8x16`      | :white_check_mark: | :x:  |
| `ReduceMinU8x16`     | :white_check_mark: | :x:  |

## Architecture

- patches -> patching existing files
- overlays -> new dirs+files

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
go/bin/go run .
```
