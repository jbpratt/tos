import {Order, Item, Option} from "./tos_pb"

export default [
  {
    "id": 1,
    "items": [
      {
        "id": 1,
        "name": "LG Smoked Pulled Pork",
        "options": [
          {
            "name": "pickles",
            "selected": true
          }
        ],
        "orderItemID": 1,
        "price": 495
      }
    ],
    "name": "Majora",
    "status": "active",
    "time_ordered": "2020-10-12 13:01:59",
    "total": 495
  },
  {
    "id": 2,
    "items": [
      {
        "id": 1,
        "name": "LG Smoked Pulled Pork",
        "options": [
          {
            "name": "pickles",
            "selected": true
          }
        ],
        "orderItemID": 2,
        "price": 495
      }
    ],
    "name": "Majora",
    "status": "active",
    "time_ordered": "2020-10-12 13:02:02",
    "total": 495
  },
  {
    "id": 3,
    "items": [
      {
        "id": 1,
        "name": "LG Smoked Pulled Pork",
        "options": [
          {
            "name": "pickles",
            "selected": true
          }
        ],
        "orderItemID": 3,
        "price": 495
      }
    ],
    "name": "Majora",
    "status": "active",
    "time_ordered": "2020-10-12 13:02:04",
    "total": 495
  },
  {
    "id": 4,
    "items": [
      {
        "id": 1,
        "name": "LG Smoked Pulled Pork",
        "options": [
          {
            "name": "pickles",
            "selected": true
          }
        ],
        "orderItemID": 4,
        "price": 495
      }
    ],
    "name": "Majora",
    "status": "active",
    "time_ordered": "2020-10-12 13:02:12",
    "total": 495
  },
  {
    "id": 5,
    "items": [
      {
        "id": 1,
        "name": "LG Smoked Pulled Pork",
        "options": [
          {
            "name": "pickles",
            "selected": true
          }
        ],
        "orderItemID": 5,
        "price": 495
      }
    ],
    "name": "Majora",
    "status": "active",
    "time_ordered": "2020-10-12 13:02:57",
    "total": 495
  },
  {
    "id": 6,
    "items": [
      {
        "id": 1,
        "name": "LG Smoked Pulled Pork",
        "options": [
          {
            "name": "pickles",
            "selected": true
          }
        ],
        "orderItemID": 6,
        "price": 495
      }
    ],
    "name": "Majora",
    "status": "active",
    "time_ordered": "2020-10-12 13:02:59",
    "total": 495
  },
  {
    "id": 7,
    "items": [
      {
        "id": 1,
        "name": "LG Smoked Pulled Pork",
        "options": [
          {
            "name": "pickles",
            "selected": true
          }
        ],
        "orderItemID": 7,
        "price": 495
      }
    ],
    "name": "Majora",
    "status": "active",
    "time_ordered": "2020-10-12 13:03:01",
    "total": 495
  },
  {
    "id": 8,
    "items": [
      {
        "id": 1,
        "name": "LG Smoked Pulled Pork",
        "options": [
          {
            "name": "pickles",
            "selected": true
          }
        ],
        "orderItemID": 8,
        "price": 495
      }
    ],
    "name": "Majora",
    "status": "active",
    "time_ordered": "2020-10-12 13:03:03",
    "total": 495
  },
  {
    "id": 9,
    "items": [
      {
        "id": 1,
        "name": "LG Smoked Pulled Pork",
        "options": [
          {
            "name": "pickles",
            "selected": true
          }
        ],
        "orderItemID": 9,
        "price": 495
      }
    ],
    "name": "Majora",
    "status": "active",
    "time_ordered": "2020-10-12 13:03:09",
    "total": 495
  },
  {
    "id": 10,
    "items": [
      {
        "id": 1,
        "name": "LG Smoked Pulled Pork",
        "options": [
          {
            "name": "pickles",
            "selected": true
          }
        ],
        "orderItemID": 10,
        "price": 495
      }
    ],
    "name": "Majora",
    "status": "active",
    "time_ordered": "2020-10-12 13:03:14",
    "total": 495
  }
].map(ord => {
  var order = new Order();
  order.setId(ord.id);
  order.setName(ord.name);
  order.setStatus(ord.status);
  order.setTotal(ord.total);
  order.setTimeOrdered(ord.time_ordered);
  order.setItemsList(ord.items.map(itm => {
    var item = new Item();
    item.setName(itm.name);
    item.setId(itm.id);
    item.setPrice(itm.price);
    item.setOptionsList(itm.options.map(opt => {
      var option = new Option();
      option.setName(opt.name);
      option.setSelected(opt.selected);
      return option;
    }));
    return item;
  }));
  return order;
}) as Order[];
