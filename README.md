# s3-bucket-wordlist-generator

This tool creates a wordlist, based on a supplied prefix, optional deliminator and suffix wordlist, where each word in the new wordlist is a concatenation of the prefix, deliniator, and suffix. The created wordlist can then be fed into an s3 enumeration tool, such as gobuster, to scan for s3 buckets.

The included suffix wordlist, list.txt, is taken from [koaj](https://github.com/koaj/aws-s3-bucket-wordlist.git), and is a wordlist based on most common aws s3 bucket names.

## Build and Install
```sh
git clone https://github.com/sidsherrill1/s3-bucket-wordlist-generator.git
cd s3-bucket-wordlist-generator
go build concat_words.go
```

## Usage
```sh
# Produce new_words.txt
./concat_words -prefix=pre -delimiter=_ -wordlist=words.txt -output=new_words.txt
# Pass new_words.txt to gobuster
gobuster s3 -w output.txt
```
