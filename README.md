# Locality-Sensing Hashing (LSH)

- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
    - [Importing package](#importing-package)
    - [Creating an LSH Instance](#creating-an-lsh-Instance)
    - [Adding Vectors](#adding-vectors)
    - [Querying by Vector ID](#querying-by-vector-id)
    - [Querying by Vector](#querying-by-vector)

## Introduction

At its core, Locality-Sensitive Hashing is a technique that allows for the efficient approximate nearest neighbor search in high-dimensional spaces. Unlike traditional methods that require exhaustive comparisons, LSH employs hashing functions to map similar items to the same bucket with high probability, thereby dramatically reducing the search space.

LSH operates under the premise of locality sensitivity, which implies that similar items are likely to remain close to each other in the hashed space. By exploiting this property, LSH enables rapid retrieval of approximate nearest neighbors, making it an invaluable tool in scenarios where exact matches are not necessary.

## Installation

To use this library, you need to have Go installed. You can install the library using the following command:

```bash
go get github.com/agtabesh/lsh
```

## Usage

### Importing package

Import the library in your Go code:

```go
import (
    "context"
    "fmt"

    "github.com/agtabesh/lsh"
    "github.com/agtabesh/lsh/types"
)
```

To create an instance of LSH, you need to specify the configuration parameters using the `LSHConfig` struct and call the `NewLSH` function:

### Creating an LSH Instance

```go
config := lsh.LSHConfig{
    SignatureSize:     128,
    BandSize:          64,
}

hashFamily := lsh.NewXXHASH64HashFamily(config.SignatureSize)
similarityMeasure := lsh.NewHammingSimilarity()
store := lsh.NewInMemoryStore()

instance, err := NewLSH(config, hashFamily, similarityMeasure, store)
if err != nil {
    // Handle error
}
```

### Adding Vectors
You can add vectors to the LSH service using the `Add` method:

```go
ctx := context.Background()
vectors := []types.Vector{
    {"feat1": 1, "feat2": 1, "feat3": 1},
    {"feat1": 1, "feat4": 1, "feat5": 1},
    {"feat1": 1, "feat2": 1, "feat6": 1, "feat7": 1},
}
for i, vector := range vectors {
    vectorID := types.VectorID(fmt.Sprint(i))
    err := instance.Add(ctx, vectorID, vector)
    if err != nil {
        // Handle error
    }
}
```

> [!NOTE]  
> Vectors can be in any dimension or length.

### Querying by Vector ID

You can perform a query using a vector ID using the `QueryByVectorID` method:

```go
ctx := context.Background()
vectorID := types.VectorID("0")
count := 5
similarVectorsID, err := instance.QueryByVectorID(ctx, vectorID, count)
if err != nil {
    // Handle error
}
fmt.Println("similarVectors", similarVectorsID)
```

### Querying by Vector

You can perform a query using a vector using the `QueryByVector` method:

```go
ctx := context.Background()
vector := types.Vector{"feat1": 1, "feat2": 1, "feat3": 1}
count := 5
similarVectorsID, err := lsh.QueryByVector(ctx, vector, count)
if err != nil {
    // Handle error
}
fmt.Println("similarVectors", similarVectorsID)
```


