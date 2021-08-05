

## Lessons learnt
Day 4
* Object.values
* Object.values(Enum)

```
enum Enum {
    FOO,
    BAR
}

console.log(Object.values(Enum));
console.log(Object.keys(Enum));

```
output:
```
[ 'FOO', 'BAR', 0, 1 ]
[ '0', '1', 'FOO', 'BAR' ]
```

If the enum values are the same as the keys. Both `Object.values`/`Object.keys` return the same arrays
```
enum Enum {
    FOO = "FOO",
    BAR = "BAR"
}

console.log(Object.values(Enum));
console.log(Object.keys(Enum));
```
output:
```
[ 'FOO', 'BAR' ]
[ 'FOO', 'BAR' ]
```
