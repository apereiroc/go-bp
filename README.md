# Go blueprint

Example

```bash
project=~/my_new_fancy_project
mkdir -p $project
go-bp file makefile $project 
```

Fill project info and see the result

```bash
cat $project/Makefile
```

See available commands with

```bash
go-bp file -h
```
