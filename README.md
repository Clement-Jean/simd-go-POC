# simd-go-POC

This repository is an attempt at adding SIMD to Go through compiler intrinsics.

## Supported Intrinsics

| Intrinsic            | NEON               | SSE2 |
|----------------------|--------------------|------|
| `Add8x16`            | :white_check_mark: | :white_check_mark:  |
| `AddU8x16`           | :white_check_mark: | :white_check_mark:  |
| `Sub8x16`            | :white_check_mark: | :white_check_mark:  |
| `SubU8x16`           | :white_check_mark: | :white_check_mark:  |
| `SaturatingAdd8x16`  | :white_check_mark: | :white_check_mark:  |
| `SaturatingAddU8x16` | :white_check_mark: | :white_check_mark:  |
| `SaturatingSub8x16`  | :white_check_mark: | :white_check_mark:  |
| `SaturatingSubU8x16` | :white_check_mark: | :white_check_mark:  |
| `And8x16`            | :white_check_mark: | :white_check_mark:  |
| `AndU8x16`           | :white_check_mark: | :white_check_mark:  |
| `Or8x16`             | :white_check_mark: | :white_check_mark:  |
| `OrU8x16`            | :white_check_mark: | :white_check_mark:  |
| `Xor8x16`            | :white_check_mark: | :white_check_mark:  |
| `Xor8x16`            | :white_check_mark: | :white_check_mark:  |
| `Max8x16`            | :white_check_mark: | :white_check_mark:  |
| `MaxU8x16`           | :white_check_mark: | :white_check_mark:  |
| `Min8x16`            | :white_check_mark: | :white_check_mark:  |
| `MinU8x16`           | :white_check_mark: | :white_check_mark:  |
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

## SSA trace

From root, execute:

```
GOSSAFUNC=AddU8x16 go/bin/go run .
```

you will get a file called `ssa.html`, just open that in your browser
