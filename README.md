# simd-go-POC

This repository is an attempt at adding SIMD to go through compiler intrinsics.

## Build

```
pushd go && ../build.sh && popd
```

## SSA trace

From root, execute:

```
GOSSAFUNC=AddU8x16 go/bin/go run .
```

you will get a file called ssa.html, just open that in you browser