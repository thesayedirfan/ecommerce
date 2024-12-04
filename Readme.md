# Uniblox Assigment

## Thought process
- a user can add items to the cart he can add mutiple items it the user adds the same item again we increase just the quantity of the item
- when the user hit the specified threshold (2 in this case) of the item we give him the coupon code to apply to the checkout
- when the user checkouts with discount code we give him 10 percent off and then we give him the code after nth purchase only
- when the user checkouts without the discount code then he can request the admin to generate the coupon for the next order
- when the user can then apply the coupon code for it and get the discount
- admin can generate the coupon code for the use only if he has met the above threshold
- admin also have the api to stats of the total orders


### Installation
1. have latest version of go installed
    *  verify the installation of go

    ```bash
    go version
    ```
2. Clone the repository 
    ```bash
    git clone https://github.com/thesayedirfan/ecommerce.git

    cd ecommerce
    ```
3. Start The Project

    - start the server
    ```bash
    make run

4. Run the tests
    ```bash
    go test ./...
    ```

## API Documentation
### API Endpoints

| **Category** | **Endpoint Name**   | **Method** | **URL**                    | **Request Body**                                                                                                                                               | **Description** |
|--------------|---------------------|------------|----------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------|-----------------|
| Admin        | admin-stats         | GET        | `{{URL}}admin/stats`       | N/A                                                                                                                                                           |         allows admin to get the stats of total orders,discount, discount codes    |
| Admin        | discount-code       | POST       | `{{URL}}admin/discount/generate` | ```{  "user_id": "1"}```                                                                                                                           |        allows user to get the discount if the threshold id met         |
| Cart         | AddItemToCart       | POST       | `{{URL}}cart/add`          | ```{  "product_id": "2",  "user_id": "1",  "name": "Bottle 2",  "price": 10,  "quantity": 1,  "Price": 10}```                            |             allows user to add the item in cart    |
| Cart         | Checkout            | POST       | `{{URL}}cart/checkout`     | ```{  "user_id": "1"}```                                                                                  |           allows use to checkout with or without discout code      |

### Notes

- **URL Variables:** The `{{URL}}` variable is used as a placeholder in all endpoints.
