# code-generator
> Generate hundreds of thousands of codes in a few seconds

[![Build Status](https://travis-ci.org/ecrmnn/code-generator.svg?branch=master)](https://travis-ci.org/ecrmnn/code-generator)

### Installation
```bash
go get github.com/ecrmnn/code-generator
```

### Usage
Generate 3 codes with pattern: 4 letters + dash + 4 numbers ``llll-nnnn``
```bash
cd $GOPATH/src/github.com/ecrmnn/code-generator
go run main.go -pattern llll-nnnn -length 3
# Generated 3 of 3 codes
# Done. Codes stored in codes.txt
# Execution time 402.662µs
nano codes.txt
# All codes stored in codes.txt
```

### Sort of related
- [passlawl](https://github.com/ecrmnn/passlawl) - CLI for generating passwords

### License
MIT © [Daniel Eckermann](http://danieleckermann.com)