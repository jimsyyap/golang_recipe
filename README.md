# my golang_recipe

## TODO

- seek security analyst + golang

### notes
rand.Seed deprecated? Do...
Source := rand.NewSource(time.Now().UnixNano())
r := rand.New(source)
number := r.Intn(10)
