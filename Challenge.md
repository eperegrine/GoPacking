# Software Engineering Challenge
Imagine for a moment that one of our product lines ships in various pack sizes:
- 250 Items
- 500 Items
- 1000 Items
- 2000 Items
- 5000 Items
- 
Our customers can order any number of these items through our website, but they
will always only be given complete packs.

1. Only whole packs can be sent. Packs cannot be broken open.
2. Within the constraints of Rule 1 above, send out no more items than
necessary to fulfil the order.
3. Within the constraints of Rules 1 & 2 above, send out as few packs as
possible to fulfil each order.
So, for example:

| Items Ordered | Correct Number of Packs           | Incorrect Number of Packs                        |
|---------------|-----------------------------------|-------------------------------------------------|
| 1             | 1 x 250                           | 1 x 500 – more items than necessary              |
| 250           | 1 x 250                           | 1 x 500 – more items than necessary              |
| 251           | 1 x 500                           | 2 x 250 – more packs than necessary              |
| 501           | 1 x 500, 1 x 250                  | 1 x 1000 – more items than necessary, 3 x 250 – more packs than necessary |
| 12001         | 2 x 5000, 1 x 2000, 1 x 250       | 3 x 5000 – more items than necessary             |


Write an application that can calculate the number of packs we need to ship to the
customer.