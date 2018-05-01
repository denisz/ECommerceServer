package models


type Price int

const MaxPrice Price = 1<<31 - 1
const MinPrice Price = -1 << 31
