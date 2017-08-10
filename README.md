# Datetime

Go的time.Time的增强类型

## How To Use：

```go
dt := datetime.New()
fmt.Println(dt) //yyyy-MM-dd HH:mm:ss
```

## get、set `yyyy-MM-dd HH:mm:ss`

```go
//get
dtStr := dt.String()

//set
err := dt.SetString("yyyy-MM-dd HH:mm:ss")
```

## get、set `time.Time`

```go
//get
t := dt.Time()

//set
dt.SetTime(time.Now())
```

## get、set `Unix timestamp`

```go
//get
u := dt.Unix()

//set
dt.SetUnix(int64(123))
```

## for `JSON`

```go
type st struct {
    DT Datetime
}
s := &st{}

//Marshal
j,e := json.Marshal(s)
fmt.Println(string(j)) // {"DT": "yyyy-MM-dd HH:mm:ss"}

//Unmarshal
e := json.UnMarshal([]byte(`{"DT": "yyyy-MM-dd HH:mm:ss"}`), s) // it's ok
```