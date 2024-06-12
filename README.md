# Packing Manager

A tool written in go to split an order down in the most efficient way

## How to use

Set the query parameter `quantity` to be the amount you want to manage

## Hosting

The app is hosted with digital ocean and is accessible with:

https://walrus-app-qbp8c.ondigitalocean.app/

To get stock samples you can use:

https://walrus-app-qbp8c.ondigitalocean.app/pack?quantity=399&packs=100,75,50,20,10

setting quantity for number of stock items, and if *optionally* set packs with comma seperated values
packs must be passed in size order


| Items Ordered | Correct Number of Packs           | Incorrect Number of Packs                        | link                                                            |
|---------------|-----------------------------------|-------------------------------------------------|-----------------------------------------------------------------|
| 1             | 1 x 250                           | 1 x 500 – more items than necessary              | https://walrus-app-qbp8c.ondigitalocean.app/pack?quantity=1     |
| 250           | 1 x 250                           | 1 x 500 – more items than necessary              | https://walrus-app-qbp8c.ondigitalocean.app/pack?quantity=250   |
| 251           | 1 x 500                           | 2 x 250 – more packs than necessary              | https://walrus-app-qbp8c.ondigitalocean.app/pack?quantity=251   |
| 501           | 1 x 500, 1 x 250                  | 1 x 1000 – more items than necessary, 3 x 250 – more packs than necessary | https://walrus-app-qbp8c.ondigitalocean.app/pack?quantity=501   |
| 12001         | 2 x 5000, 1 x 2000, 1 x 250       | 3 x 5000 – more items than necessary             | https://walrus-app-qbp8c.ondigitalocean.app/pack?quantity=12001 |

