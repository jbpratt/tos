# Mookies' BBQ TOS
<p align="center">
  <img style="float: right;" src="assets/logo.png" alt="Mookies logo" />
</p>

## Goal
Implement a ticket order system in Go

### Requirements

##### Client 1: Front
* Ability to requests a slice of items, displays them
* Ability to add item(s) to a 'cart' to build up an order
* Ability to send the order to the server

##### Client 2: Kitchen
* Ability to receive orders through server push
* Ability to mark orders as complete, send order to server

##### Server:
* Ability to send slice of items
* Ability to receive new order
* Ability to push new order to client
* Ability to receive completed order
* Validate orders from the server as they are built in the 'cart'

#### Mockups

##### Client 1: Front
![front](https://i.vgy.me/zyYZjo.png)

##### Client 2: Kitchen
![kitchen](https://i.vgy.me/nuYZ5k.png)
