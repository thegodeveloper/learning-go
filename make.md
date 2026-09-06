# Make Commands

Default target: `run` (runs when `make` is called with no target).

## Targets

### Format

Formats all Go source files in the project.

```
make fmt
```

### Vet

Formats code, then checks it for suspicious constructs.

```
make vet
```

### Run

Formats, vets, and starts the API.

```
make run
```

### Clean

Removes build artifacts.

```
make clean
```

### Default

Same as `make run`.

```
make
```
