	package Methods

	import (
		"Test_task/DB_conn"
		"github.com/gin-gonic/gin"
		_ "github.com/lib/pq"
	)

	type Item struct {
		ID int
		SKU string
		Name string
		Type string
		Price string
	}


	// возможно изменить price на другой тип данных
	//сделать обработчик ошибок
	// нормально ли возвращается id возвращать id
	func AddItem(c * gin.Context) {
		SKU := c.Query("sku")
		name := c.Query("name")
		itype := c.Query("type")
		price := c.Query("price")

		//addItem
		insrt, err := DB_conn.Connect().Query("INSERT INTO items(sku, name, type, price) values ($1, $2, $3, $4)",SKU, name, itype, price)
		if err != nil {
			panic(err)
		}
		defer insrt.Close()

		//get id
		slct, err := DB_conn.Connect().Query("SELECT id FROM items WHERE sku=$1",SKU)
		if err != nil {
			panic(err)
		}
		slct.Next()
		var id int
		err = slct.Scan(&id)
		if err != nil {
			panic(err)
		}
		c.JSON(200, gin.H{"id": id})
	}

	func EditItem(c *gin.Context){
		id := c.Query("id")
		check, err:= DB_conn.Connect().Query("SELECT id FROM iems WHERE id=?", id)
		if err != nil{
			c.JSON(404, gin.H{"massage": "Query error"})
		}
		defer check.Close()

		check.Next()
		var itm Item
		err = check.Scan(&itm.ID)
		if err != nil {
			c.JSON(404, gin.H{"massage":"no such id"})
		}else {
			name := c.Query("name")
			category := c.Query("category")
			price := c.Query("price")

			upd, err := DB_conn.Connect().Query("UPDATE items SET name=?, category=?, price=? WHERE id=?", name, category, price, id)
			if err != nil {
				c.JSON(404, gin.H{"massage": "Query error"})
			}
			defer upd.Close()
		}
	}

	//ready || delete by sku
	func DeleteItem(c *gin.Context){
		id := c.Query("id")
		//does id exist
		get_id, err := DB_conn.Connect().Query("Select * FROM items WHERE id=$1", id )
		if err != nil {
			c.JSON(404, gin.H{"massage":"err"})
		}
		defer get_id.Close()
		get_id.Next()
		err = get_id.Scan(&id)
		if err != nil {
			c.JSON(404, gin.H{"massage":"no such id"})
		}else {
			del, err := DB_conn.Connect().Query("DELETE FROM items WHERE id = $1", id)
			if err != nil{
				massage := "query error"
				c.JSON(404, gin.H{
					"massage": massage,
				})
			}else{
				c.JSON(200, gin.H{
					"massage": "deleted",
				})
			}
			defer del.Close()
		}

	}

	//добавить ИЛИ SKU
	func GetItem(c *gin.Context){
		id := c.Query("id")
		f, err := DB_conn.Connect().Query("SELECT * FROM items WHERE id = $1", id)
		if err != nil{
			massage := "id type error"
			c.JSON(404, gin.H{"massage": massage})
		}
		defer f.Close()

		f.Next()
		itm := new(Item)
		err = f.Scan(&itm.ID, &itm.SKU, &itm.Name, &itm.Type, &itm.Price)
		if err != nil {
			msg := "item with id " + id + " not found"
			c.JSON(404, gin.H{
				"message": msg,
			})
		}else {
			c.JSON(200, gin.H{
				"id":    itm.ID,
				"sku":   itm.SKU,
				"name":  itm.Name,
				"type":  itm.Type,
				"price": itm.Price,
			})
		}
	}

	//а как выводить частями
	func GetAllItems(c *gin.Context){
		get, err := DB_conn.Connect().Query("select * from items")
		if err != nil {
			panic(err)
		}

		defer get.Close()
		for get.Next() {
			var itm Item
			err = get.Scan(&itm.ID, &itm.SKU, &itm.Name, &itm.Type, &itm.Price)
			if err != nil {
				panic(err)
			}
			c.JSON(200,gin.H{"id":itm.ID, "sku": itm.SKU, "name": itm.Name, "type":itm.Type, "price": itm.Price})
		}
	}