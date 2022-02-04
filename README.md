# fita-backend-master
Backend Job Application for Fita


## Command List
### Run Build
```
make build
```

### Run Test
```
make test
```

### Run Server
```
make run
```

### Exec Server Binary file
```
make exec
```

### Wire Generate
```
make wire
```

### GraphQL Generate
```
make gql
```


## Dummy List:
1. clientID = 99999
2. clientID = 88888
3. clientID = 77777
4. clientID = 99998



## Schema
### Get All Items
``` 
query {
  items {
    sku
    name
    price
    inventoryQty
  }
}
```

### Add Cart
```
mutation {
  addCart(input: {clientID: 99999, sku: "43N23P", qty: 1}) {
    clientID
    items {
      sku
      name
      qty
      price
    }
  }
}
```


### Checkout
```
mutation {
  checkout(input: {clientID: 99999}) {
    items {
      sku
      name
      qty
      price
    }
    totalAmount
  }
}
```