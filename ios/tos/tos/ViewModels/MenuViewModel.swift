import Foundation
import Combine

final class MenuViewModel: ChannelViewModel, ObservableObject, Identifiable {
    private var menuClient: Tospb_MenuServiceClient?
    @Published private(set) var menu: Tospb_Menu? = nil

    override init() {
        super.init()
        menuClient = Tospb_MenuServiceClient(channel: super.client)
        getMenu()
    }

    func getMenu() {
        #if DEBUG
        menu = loadMenu()
        #else
        let request = Tospb_Empty()
        let call = menuClient!.getMenu(request)
        let response = try? call.response.wait()
        menu = response!
        #endif
    }

    func createItem(_ item: Tospb_Item) {
        do {
            _ = try menuClient!.createMenuItem(item).response.wait()
        } catch {
            print("createMenuItem failed: \(error)")
            return
        }
    }

    func deleteItem(_ itemID: Int64) {
        let req: Tospb_DeleteMenuItemRequest = .with {
            $0.id = itemID
        }
        do {
            _ = try menuClient!.deleteMenuItem(req).response.wait()
        } catch {
            print("deleteMenuItem failed: \(error)")
        }
    }

    func updateItem(_ item: Tospb_Item) {
        do {
            _ = try menuClient!.updateMenuItem(item).response.wait()
        } catch {
            print("updateMenuItem failed: \(error)")
        }
    }
}

func loadMenu() -> Tospb_Menu {
    guard let menu = try? Tospb_Menu(jsonUTF8Data: Data(jsonMenu.utf8)) else {
        fatalError("failed to decode menu")
    }
    return menu
}

let jsonMenu = """
{
 "categories": [
  {
   "name": "Sandwiches",
   "id": 1,
   "items": [
    {
     "name": "LG Smoked Pulled Pork",
     "price": 495,
     "options": [
      {
       "name": "pickles",
       "selected": true
      },
      {
       "name": "bbq sauce",
       "selected": true
      },
      {
       "name": "white sauce"
      },
      {
       "name": "ketchup"
      },
      {
       "name": "mayo"
      },
      {
       "name": "coleslaw",
       "price": 25
      },
      {
       "name": "tomato",
       "price": 25
      },
      {
       "name": "lettuce",
       "price": 25
      },
      {
       "name": "cheese",
       "price": 25
      },
      {
       "name": "onion",
       "price": 25
      }
     ]
    },
    {
     "name": "RG Smoked Pulled Pork",
     "price": 395,
     "options": [
      {
       "name": "pickles",
       "selected": true
      },
      {
       "name": "bbq sauce",
       "selected": true
      },
      {
       "name": "white sauce"
      },
      {
       "name": "ketchup"
      },
      {
       "name": "mayo"
      },
      {
       "name": "coleslaw",
       "price": 25
      },
      {
       "name": "tomato",
       "price": 25
      },
      {
       "name": "lettuce",
       "price": 25
      },
      {
       "name": "cheese",
       "price": 25
      },
      {
       "name": "onion",
       "price": 25
      }
     ]
    },
    {
     "name": "LG Smoked Chicken Breast",
     "price": 495,
     "options": [
      {
       "name": "pickles"
      },
      {
       "name": "bbq sauce"
      },
      {
       "name": "white sauce"
      },
      {
       "name": "ketchup"
      },
      {
       "name": "mayo",
       "selected": true
      },
      {
       "name": "coleslaw",
       "price": 25
      },
      {
       "name": "tomato",
       "price": 25,
       "selected": true
      },
      {
       "name": "lettuce",
       "price": 25,
       "selected": true
      },
      {
       "name": "cheese",
       "price": 25
      },
      {
       "name": "onion",
       "price": 25
      }
     ]
    },
    {
     "name": "RG Smoked Chicken Breast",
     "price": 395,
     "options": [
      {
       "name": "pickles"
      },
      {
       "name": "bbq sauce"
      },
      {
       "name": "white sauce"
      },
      {
       "name": "ketchup",
       "selected": true
      },
      {
       "name": "mayo",
       "selected": true
      },
      {
       "name": "coleslaw",
       "price": 25
      },
      {
       "name": "tomato",
       "price": 25,
       "selected": true
      },
      {
       "name": "lettuce",
       "price": 25,
       "selected": true
      },
      {
       "name": "cheese",
       "price": 25
      },
      {
       "name": "onion",
       "price": 25
      }
     ]
    },
    {
     "name": "'The Molly'",
     "price": 395,
     "options": [
      {
       "name": "pickles",
       "selected": true
      },
      {
       "name": "bbq sauce",
       "selected": true
      },
      {
       "name": "white sauce",
       "selected": true
      },
      {
       "name": "ketchup"
      },
      {
       "name": "mayo"
      },
      {
       "name": "coleslaw",
       "price": 25,
       "selected": true
      },
      {
       "name": "tomato",
       "price": 25
      },
      {
       "name": "lettuce",
       "price": 25
      },
      {
       "name": "cheese",
       "price": 25
      },
      {
       "name": "onion",
       "price": 25
      }
     ]
    },
    {
     "name": "LG Hamburger",
     "price": 495,
     "options": [
      {
       "name": "pickles",
       "selected": true
      },
      {
       "name": "bbq sauce",
       "selected": true
      },
      {
       "name": "white sauce"
      },
      {
       "name": "ketchup"
      },
      {
       "name": "mayo"
      },
      {
       "name": "coleslaw",
       "price": 25
      },
      {
       "name": "tomato",
       "price": 25
      },
      {
       "name": "lettuce",
       "price": 25
      }
     ]
    },
    {
     "name": "Hamburger",
     "price": 395,
     "options": [
      {
       "name": "pickles",
       "selected": true
      },
      {
       "name": "bbq sauce",
       "selected": true
      },
      {
       "name": "white sauce"
      },
      {
       "name": "ketchup"
      },
      {
       "name": "mayo"
      },
      {
       "name": "coleslaw",
       "price": 25
      },
      {
       "name": "tomato",
       "price": 25
      },
      {
       "name": "lettuce",
       "price": 25
      }
     ]
    },
    {
     "name": "LG Cheeseburger",
     "price": 550,
     "options": [
      {
       "name": "pickles",
       "selected": true
      },
      {
       "name": "bbq sauce",
       "selected": true
      },
      {
       "name": "white sauce"
      },
      {
       "name": "ketchup"
      },
      {
       "name": "mayo"
      },
      {
       "name": "coleslaw",
       "price": 25
      },
      {
       "name": "tomato",
       "price": 25
      },
      {
       "name": "lettuce",
       "price": 25
      }
     ]
    },
    {
     "name": "Cheeseburger",
     "price": 425,
     "options": [
      {
       "name": "pickles",
       "selected": true
      },
      {
       "name": "bbq sauce",
       "selected": true
      },
      {
       "name": "white sauce"
      },
      {
       "name": "ketchup"
      },
      {
       "name": "mayo"
      },
      {
       "name": "coleslaw",
       "price": 25
      },
      {
       "name": "tomato",
       "price": 25
      },
      {
       "name": "lettuce",
       "price": 25
      }
     ]
    },
    {
     "name": "Grilled Cheese",
     "price": 300,
     "options": [
      {
       "name": "pickles",
       "selected": true
      },
      {
       "name": "bbq sauce",
       "selected": true
      },
      {
       "name": "white sauce"
      },
      {
       "name": "ketchup"
      },
      {
       "name": "mayo"
      },
      {
       "name": "coleslaw",
       "price": 25
      },
      {
       "name": "tomato",
       "price": 25
      },
      {
       "name": "lettuce",
       "price": 25
      }
     ]
    },
    {
     "name": "Pulled Pork Melt",
     "price": 395,
     "options": [
      {
       "name": "pickles",
       "selected": true
      },
      {
       "name": "bbq sauce",
       "selected": true
      },
      {
       "name": "white sauce"
      },
      {
       "name": "ketchup"
      },
      {
       "name": "mayo"
      },
      {
       "name": "coleslaw",
       "price": 25
      },
      {
       "name": "tomato",
       "price": 25
      },
      {
       "name": "lettuce",
       "price": 25
      }
     ]
    }
   ]
  },
  {
   "name": "Plates",
   "id": 2,
   "items": [
    {
     "name": "Smoked Pulled Pork",
     "price": 990,
     "options": [
      {
       "name": "potato salad"
      },
      {
       "name": "baked beans"
      },
      {
       "name": "coleslaw"
      },
      {
       "name": "chips"
      },
      {
       "name": "fries"
      },
      {
       "name": "baked potato",
       "price": 100
      }
     ]
    },
    {
     "name": "RG Rib",
     "price": 995,
     "options": [
      {
       "name": "potato salad"
      },
      {
       "name": "baked beans"
      },
      {
       "name": "coleslaw"
      },
      {
       "name": "chips"
      },
      {
       "name": "fries"
      },
      {
       "name": "baked potato",
       "price": 100
      }
     ]
    },
    {
     "name": "½ Smoked Chicken",
     "price": 995,
     "options": [
      {
       "name": "potato salad"
      },
      {
       "name": "baked beans"
      },
      {
       "name": "coleslaw"
      },
      {
       "name": "chips"
      },
      {
       "name": "fries"
      },
      {
       "name": "baked potato",
       "price": 100
      }
     ]
    },
    {
     "name": "Smoked Chicken Breast",
     "price": 725,
     "options": [
      {
       "name": "potato salad"
      },
      {
       "name": "baked beans"
      },
      {
       "name": "coleslaw"
      },
      {
       "name": "chips"
      },
      {
       "name": "fries"
      },
      {
       "name": "baked potato",
       "price": 100
      }
     ]
    },
    {
     "name": "Smoked Wing (8 wings)",
     "price": 999,
     "options": [
      {
       "name": "potato salad"
      },
      {
       "name": "baked beans"
      },
      {
       "name": "coleslaw"
      },
      {
       "name": "chips"
      },
      {
       "name": "fries"
      },
      {
       "name": "baked potato",
       "price": 100
      }
     ]
    },
    {
     "name": "BBQ Loaded Nachos",
     "price": 875,
     "options": [
      {
       "name": "cheddar cheese",
       "selected": true
      },
      {
       "name": "green onions",
       "selected": true
      },
      {
       "name": "jalapeños",
       "selected": true
      },
      {
       "name": "bbq sauce"
      },
      {
       "name": "white sauce"
      }
     ]
    },
    {
     "name": "Chicken Loaded Nachos",
     "price": 875,
     "options": [
      {
       "name": "cheddar cheese",
       "selected": true
      },
      {
       "name": "green onions",
       "selected": true
      },
      {
       "name": "jalapeños",
       "selected": true
      },
      {
       "name": "bbq sauce"
      },
      {
       "name": "white sauce"
      }
     ]
    }
   ]
  },
  {
   "name": "Baskets",
   "id": 3,
   "items": [
    {
     "name": "Smoked Wing",
     "price": 500,
     "options": [
      {
       "name": "potato salad"
      },
      {
       "name": "baked beans"
      },
      {
       "name": "coleslaw"
      },
      {
       "name": "chips"
      },
      {
       "name": "fries"
      },
      {
       "name": "baked potato",
       "price": 100
      }
     ]
    },
    {
     "name": "Rib",
     "price": 500,
     "options": [
      {
       "name": "potato salad"
      },
      {
       "name": "baked beans"
      },
      {
       "name": "coleslaw"
      },
      {
       "name": "chips"
      },
      {
       "name": "fries"
      },
      {
       "name": "baked potato",
       "price": 100
      }
     ]
    }
   ]
  },
  {
   "name": "Potatoes",
   "id": 4,
   "items": [
    {
     "name": "Loaded Pork",
     "price": 899,
     "options": [
      {
       "name": "cheese",
       "selected": true
      },
      {
       "name": "bacon bits",
       "selected": true
      },
      {
       "name": "green onions",
       "selected": true
      },
      {
       "name": "butter",
       "selected": true
      },
      {
       "name": "sour cream",
       "selected": true
      }
     ]
    },
    {
     "name": "Loaded Chicken",
     "price": 899,
     "options": [
      {
       "name": "cheese",
       "selected": true
      },
      {
       "name": "bacon bits",
       "selected": true
      },
      {
       "name": "green onions",
       "selected": true
      },
      {
       "name": "butter",
       "selected": true
      },
      {
       "name": "sour cream",
       "selected": true
      }
     ]
    },
    {
     "name": "Loaded Potato (no meat)",
     "price": 699,
     "options": [
      {
       "name": "cheese",
       "selected": true
      },
      {
       "name": "bacon bits",
       "selected": true
      },
      {
       "name": "green onions",
       "selected": true
      },
      {
       "name": "butter",
       "selected": true
      },
      {
       "name": "sour cream",
       "selected": true
      }
     ]
    },
    {
     "name": "(Pork) Smothered and Covered Fries",
     "price": 899,
     "options": [
      {
       "name": "cheese",
       "selected": true
      },
      {
       "name": "bacon bits",
       "selected": true
      },
      {
       "name": "green onions",
       "selected": true
      },
      {
       "name": "butter",
       "selected": true
      },
      {
       "name": "sour cream",
       "selected": true
      }
     ]
    },
    {
     "name": "(Chicken) Smothered and Covered Fries",
     "price": 899,
     "options": [
      {
       "name": "cheese",
       "selected": true
      },
      {
       "name": "bacon bits",
       "selected": true
      },
      {
       "name": "green onions",
       "selected": true
      },
      {
       "name": "butter",
       "selected": true
      },
      {
       "name": "sour cream",
       "selected": true
      }
     ]
    }
   ]
  },
  {
   "name": "Sides",
   "id": 5,
   "items": [
    {
     "name": "SM Cole Slaw",
     "price": 300
    },
    {
     "name": "LG Cole Slaw",
     "price": 600
    },
    {
     "name": "SM Baked Beans",
     "price": 300
    },
    {
     "name": "LG Baked Beans",
     "price": 600
    },
    {
     "name": "SM Potato Salad",
     "price": 300
    },
    {
     "name": "LG Potato Salad",
     "price": 600
    },
    {
     "name": "Plain Chips",
     "price": 100
    },
    {
     "name": "Fries",
     "price": 300
    }
   ]
  },
  {
   "name": "Drinks",
   "id": 6,
   "items": [
    {
     "name": "Canned Drink",
     "price": 100
    },
    {
     "name": "Bottled Water",
     "price": 150
    }
   ]
  },
  {
   "name": "Desserts",
   "id": 7,
   "items": [
    {
     "name": "Oreo Dream",
     "price": 350
    },
    {
     "name": "Lemon Delight",
     "price": 350
    },
    {
     "name": "Strawberry Pizza",
     "price": 350
    },
    {
     "name": "Whole Dessert",
     "price": 3000
    },
    {
     "name": "SM Banana Pudding",
     "price": 350
    },
    {
     "name": "LG Banana Pudding",
     "price": 700
    }
   ]
  },
  {
   "name": "Sauces",
   "id": 8,
   "items": [
    {
     "name": "Extra Ranch",
     "price": 50
    },
    {
     "name": "Extra White",
     "price": 50
    },
    {
     "name": "Extra Buffalo Wing",
     "price": 50
    },
    {
     "name": "Extra BBQ",
     "price": 50
    }
   ]
  }
 ]
}
"""
