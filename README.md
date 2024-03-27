# simd-go-POC

This repository is an attempt at adding SIMD to Go through compiler intrinsics.

:warning: It only works on ARM64 for now :warning:

## Supported Intrinsics

- `Add8x16` & `AddU8x16`
- `Sub8x16` & `SubU8x16`
- `SaturatingAdd8x16` & `SaturatingAddU8x16`
- `SaturatingSub8x16` & `SaturatingSubU8x16`
- `And8x16` & `AndU8x16`
- `Or8x16` & `OrU8x16`
- `Xor8x16` & `XorU8x16`
- `Max8x16` & `MaxU8x16`
- `Min8x16` & `MinU8x16`
- `ReduceMax8x16` & `ReduceMaxU8x16`
- `ReduceMin8x16` & `ReduceMinU8x16`

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
