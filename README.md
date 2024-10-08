# Gestimate CLI

`gestimate` is a Command Line Interface (CLI) tool built in Go that performs 3-point estimation (PERT) based on provided best, likely, and worst-case dates. It calculates and displays confidence intervals for 68%, 90%, and 95% probabilities.

## NOTICE

This is a prototype, it probably has a lot of bugs and may be inaccurate. Please see the LICENSE for additional information about warranties.

## Features

- Takes in best, likely, and worst case dates as inputs.
- Calculates PERT-based estimates.
- Outputs confidence intervals for 68%, 90%, and 95%.
- Can be run natively or inside a Docker container.

## Requirements

- [Go](https://golang.org/doc/install) (see go.mod for the minimum version)
- [Docker](https://docs.docker.com/get-docker/) (for containerized usage)
- Make (optional, but recommended for build automation)

## Quick Start

1. **Clone the repository**:

    ```bash
    git clone https://github.com/SudoBrendan/gestimate
    cd gestimate
    ```

2. **Build the CLI**:

    ```bash
    make build
    ```

3. **Run the CLI**:

    ```bash
    ./gestimate --best 2024/01/01 --likely 2024/02/01 --worst 2024/03/01
    ```

## Usage

### Command Line Options:

- `--best`, `-b`: Best case date (format: YYYY/MM/DD) [required].
- `--likely`, `-l`: Most likely case date (format: YYYY/MM/DD) [required].
- `--worst`, `-w`: Worst case date (format: YYYY/MM/DD) [required].

### Example:

```bash
./gestimate --best 2024/01/01 --likely 2024/02/01 --worst 2024/03/01
```

which  will output

```
Confidence Interval Table
------------------------
68% Confidence: 2024/02/10
90% Confidence: 2024/02/16
95% Confidence: 2024/02/19
```

## Docker Usage

You can build and run the application inside a Docker container:

```
# build
make docker-build

# run
make docker-run
```