### Test Env

golang 1.12.x

dell server 48 cpu core, hz 2.5

### pipeline mode

```
pipeline cost 53.839017ms
pipeline cost 87.617389ms
pipeline cost 35.041264ms
pipeline cost 30.389709ms
pipeline cost 36.155615ms
pipeline cost 34.633911ms
pipeline cost 44.634442ms
pipeline cost 43.291504ms
pipeline cost 38.112881ms
pipeline cost 34.843634ms
pipeline cost 79.052806ms
pipeline cost 33.663211ms
pipeline cost 38.634923ms
pipeline cost 47.739686ms
```

### concurrent mode (go func)

```
concurrent cost 12.272914ms
concurrent cost 23.300305ms
concurrent cost 21.565337ms
concurrent cost 14.084461ms
concurrent cost 27.099352ms
concurrent cost 25.383744ms
concurrent cost 10.514287ms
concurrent cost 23.86968ms
concurrent cost 24.406822ms
concurrent cost 34.802821ms
concurrent cost 39.235525ms
concurrent cost 28.664666ms
concurrent cost 25.674713ms
concurrent cost 33.523613ms
```

### go pool mode

```
go pool cost 12.883377ms
go pool cost 11.292233ms
go pool cost 10.899465ms
go pool cost 10.747805ms
go pool cost 8.306871ms
go pool cost 10.613021ms
go pool cost 13.103071ms
go pool cost 16.094125ms
go pool cost 12.562816ms
go pool cost 18.157812ms
go pool cost 26.127534ms
go pool cost 10.951803ms
go pool cost 12.208376ms
go pool cost 28.038126ms
go pool cost 17.979167ms
```
