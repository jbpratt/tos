package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/jbpratt/tos/internal/models"
	"github.com/jbpratt/tos/internal/pb"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	// db driver
	_ "github.com/mattn/go-sqlite3"
)

// MenuDB is everything that interacts with the database
// involving the menu
type MenuDB interface {
	SeedMenu(context.Context) error
	CreateMenuItem(context.Context, *pb.Item) (int64, error)
	DeleteMenuItem(context.Context, int64) error
	UpdateMenuItem(context.Context, *pb.Item) error
	// CreateMenuItemOption() error
	GetMenu(context.Context) (*pb.Menu, error)
}

// MenuService the the abstraction for the MenuDB
type MenuService interface {
	MenuDB
}

var _ MenuDB = (*menuDB)(nil)

type menuService struct {
	MenuDB
}

type menuDB struct {
	sync.RWMutex
	db *sql.DB
}

// NewMenuService creates a menu service for interacting with the database
func NewMenuService(db *sql.DB) (MenuService, error) {
	return &menuService{&menuDB{db: db}}, nil
}

func (m *menuDB) SeedMenu(ctx context.Context) error {
	m.Lock()
	defer m.Unlock()

	menu, err := loadStaticMenu()
	if err != nil {
		return fmt.Errorf("failed to load static menu: %w", err)
	}

	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_ = menu
	/*
		for _, category := range menu.GetCategories() {
			c := &models.ItemKind{Name: category.GetName()}
			if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
				tx.Rollback()
				return err
			}
			for _, item := range category.GetItems() {
				i := &models.Item{Name: item.GetName(), Price: int64(item.GetPrice()), KindID: c.ID.Int64}
				if err = i.Insert(ctx, tx, boil.Infer()); err != nil {
					tx.Rollback()
					return err
				}
				for _, option := range item.GetOptions() {
					res, err := tx.Exec(
						"INSERT INTO options (name, price, selected) VALUES (?,?,?)",
						option.GetName(), option.GetPrice(), option.GetSelected())
					if err != nil {
						tx.Rollback()
						return err
					}
					optionid, _ := res.LastInsertId()
					_, err = tx.Exec(
						"INSERT INTO item_options (item_id, option_id) VALUES (?,?)",
						itemid, optionid)
					if err != nil {
						tx.Rollback()
						return err
					}
				}
			}
		}
	*/

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

// TODO: stop using fmt.Sprintf to format queries
func (m *menuDB) GetMenu(ctx context.Context) (*pb.Menu, error) {
	m.RLock()
	defer m.RUnlock()

	var categories []*pb.Category
	menu := &pb.Menu{
		Categories: categories,
	}
	/*
		if err := m.db.Select(&menu.Categories, "SELECT * from categories"); err != nil {
			return nil, err
		}
		for _, category := range menu.GetCategories() {
			if err := m.db.Select(&category.Items,
				fmt.Sprintf("SELECT * FROM items WHERE category_id = %v", category.GetId())); err != nil {
				return nil, err
			}
			for _, item := range category.GetItems() {
				if err := m.db.Select(&item.Options, fmt.Sprintf(
					`
					SELECT name,price,selected,options.id
					FROM options JOIN item_options as io ON options.id = io.option_id
					WHERE item_id = %d`, item.GetId())); err != nil {
					return nil, err
				}
			}
		}
	*/
	return menu, nil
}

// need to reload
func (m *menuDB) CreateMenuItem(ctx context.Context, item *pb.Item) (int64, error) {
	m.Lock()
	defer m.Unlock()

	i := &models.Item{
		Name:   item.GetName(),
		Price:  item.GetPrice(),
		KindID: item.GetCategoryID(),
	}
	if err := i.InsertG(ctx, boil.Infer()); err != nil {
		return 0, err
	}

	return i.ID.Int64, nil
}

func (m *menuDB) UpdateMenuItem(ctx context.Context, item *pb.Item) error {
	m.Lock()
	defer m.Unlock()

	itm, err := models.FindItemG(ctx, null.Int64From(item.Id))
	if err != nil {
		return fmt.Errorf("failed to find item: %w", err)
	}

	itm.Name = item.Name
	itm.Price = item.Price

	_, err = itm.UpdateG(ctx, boil.Infer())
	if err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}

	return nil
}

func (m *menuDB) DeleteMenuItem(ctx context.Context, id int64) error {
	m.Lock()
	defer m.Unlock()

	aff, err := models.Items(qm.Where("id=?", id)).DeleteAllG(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete item: %w", err)
	}

	if aff == 0 {
		return errors.New("ID not found")
	}

	return nil
}

func loadStaticMenu() (*pb.Menu, error) {
	var menu *pb.Menu = &pb.Menu{}
	if err := json.Unmarshal([]byte(staticMenu), menu); err != nil {
		return nil, err
	}
	return menu, nil
}

const staticMenu = `
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
}`
