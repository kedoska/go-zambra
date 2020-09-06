# go-zambra

A minimal Go application to execute tasks.

### Execute single command

```bash
go run main.go run node -e "console.log(1)"
```

### Execute Yaml workflow

```yaml
name: 'test flow'

steps:
  - name: node
    args:
      - -e
      - "console.log('1st')"
  - name: node
    args:
      - -e
      - "console.log('second')"

```

```bash
go run main.go execute ./test.yaml 
```