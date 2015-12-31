# THE MNIST DATABASE of handwritten digits (img split)

Split the "the mnist database of handwritten digits" in several 28x28
pictures (and lables)

## Split images

```sh
go run split.go -type image -in train-images-idx3-ubyte -out img-train/
```

## Split labels


```sh
go run split.go -type label -in train-labels-idx1-ubyte -out img-train/
```

