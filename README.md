# Path Traversal False-Positive Pattern (Go)

This project intentionally demonstrates a control-flow pattern that can trigger a taint-analysis false positive:

- Source: `--path` CLI flag (`inputPath`)
- Guard: `if inputPath != "" { return inputPath, nil }`
- Sink: `os.ReadFile(...)` appears later in the same function

The key detail: the sink executes only when `inputPath == ""`.
So there is no runtime path where a non-empty attacker-controlled `--path` value reaches `os.ReadFile`.

## Run

```bash
cd path-traversal-fp-pattern
go run . --path "../../../../etc/passwd"
```

Expected output:

```text
resolved path: ../../../../etc/passwd
```

No file read is performed in this case because the function returns at the guard.

Now run with an empty path:

```bash
go run .
```

Expected output:

```text
resolved path: ./safe-default/location
```

In this case, `os.ReadFile("fixtures/default_path.txt")` is executed, but no attacker-controlled path is involved.

## Why this is useful

It reproduces the **same branch-exclusivity pattern** behind reports like:

- "tainted CLI flag reaches `os.ReadFile`"

even though:

- non-empty input takes an early return path
- sink path is reachable only for empty input
