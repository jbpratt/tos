import SwiftUI

struct ContentView: View {
    var body: some View {
        Menu()
    }
}

struct Menu: View {
    var menu: Tospb_Menu
    init() {
        self.menu = loadMenu()
    }

    var body: some View {
        ScrollView {
            VStack(alignment: .leading){
                ForEach(menu.categories, id: \.self) { cat in
                    Category(category: cat)
                }.padding()
            }
        }
    }
}

struct Category: View {
    var category: Tospb_Category

    var body: some View {
        VStack {
            Text(category.name).font(.largeTitle)
            ForEach(category.items, id: \.self) { item in
                HStack {
                    Text(item.name)
                    Spacer()
                    Text(String(format: "%.2f", item.price / 100))
                }
            }
        }
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
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
