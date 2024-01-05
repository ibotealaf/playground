import helper

def main():
	helper.createTable()

	while True:
		helper.instructions()
		try:
			command = int(input("Enter command: "))
			if command == 0:
				products = helper.allProducts()
				if len(products) == 0:
					print("No products yet")
				for product in products:
					print(" Name || Price ")
					print(f"{product[0]} || {product[1]}")
	
			if command == 1:
				product_name = input("Search for product: ")
				price = helper.getProductPrice(product_name)
				if price == None:
						print(f"Could not find product with {product_name}")
						continue

				print(f"The price for {product_name} is {price[0]}")

			if command == 2:
				product_name = input("Enter product name: ")
				product_price = float(input("Enter product price: "))
				newProduct = helper.createProduct(product_name, product_price)
				if newProduct != None:
						print("New product created")

			if command == 3:
				print("Session ended")
				break
			else:
				continue
		except:
			print("Only integers accepted")


if __name__ == "__main__":
	main()
