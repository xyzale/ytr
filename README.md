# YTR

Ytr is a simple automation tool which allows to execute a list of commands by typing only one
instruction.

The commands list must be organised in a .yaml file called Ytrfile.yaml. Here an example:

```
commands:
  example:
    - pwd
    - ls
  createfile:
    - echo "build" > foo.txt
    - cat foo.txt
```

Once cloned the repository, execute `go build` to create an executable file, then move it to 
your bin folder such as `/usr/local/bin`.

In order to run a list of commands, just type `ytr <command-name>`. Following the example, 
the command-name could be `example` or `createfile`.
