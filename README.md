# Ninja tech challenge 2025

This API provides backend support for a marketplace focused on travelers and backpackers, including user management, product catalog, orders, and more.

This service operates behind an API Gateway to manage routing and authentication [(API Gateway repository link)](https://github.com/caiomarinello/ninjaGateway).

---

## Table of Contents

1. [Product Management Endpoints](#product-management-endpoints)
2. [Order Management Endpoints](#order-management-endpoints)
3. [Review Management Endpoints](#review-management-endpoints)
4. [Tip/Story Exchange Endpoints](#tipstory-exchange-endpoints)
5. [Authentication & Authorization Endpoints](#authentication--authorization-endpoints)
6. [User Management Endpoints](#user-management-endpoints)

---

## Product Management Endpoints

| **Method** | **Endpoint**    | **Description**                                | **Access** | **Status**     |
| ---------- | --------------- | ---------------------------------------------- | ---------- | -------------- |
| GET        | `/products`     | Retrieve a list of all products.               | User       | 🟢 Implemented |
| GET        | `/product/{id}` | Retrieve detailed information about a product. | User       | 🟢 Implemented |
| POST       | `/product`      | Add a new product.                             | Admin      | 🟢 Implemented |
| PUT        | `/product/{id}` | Update an existing product.                    | Admin      | 🟢 Implemented |
| DELETE     | `/product/{id}` | Delete a product.                              | Admin      | 🟢 Implemented |

---

## Order Management Endpoints

| **Method** | **Endpoint**  | **Description**                                                                      | **Access** | **Status**     |
| ---------- | ------------- | ------------------------------------------------------------------------------------ | ---------- | -------------- |
| POST       | `/checkout`   | Finalize the cart, saves the order in the database and send user confirmation email. | User       | 🟢 Implemented |
| GET        | `/order/{id}` | Retrieve a specific order’s details.                                                 | User       | 🟢 Implemented |
| GET        | `/orders`     | List all orders for the authenticated user.                                          | User       | 🟢 Implemented |

---

## Review Management Endpoints

| **Method** | **Endpoint**            | **Description**                          | **Access** | **Status**         |
| ---------- | ----------------------- | ---------------------------------------- | ---------- | ------------------ |
| POST       | `/review`               | Create a new review for a product.       | User       | 🔴 Not Implemented |
| GET        | `/reviews/{product_id}` | Retrieve reviews for a specific product. | User       | 🔴 Not Implemented |
| PUT        | `/review/{id}`          | Update an existing review.               | User       | 🔴 Not Implemented |
| DELETE     | `/review/{id}`          | Delete a review.                         | User/Admin | 🔴 Not Implemented |

---

## Tip/Story Exchange Endpoints

| **Method** | **Endpoint** | **Description**                                  | **Access** | **Status**         |
| ---------- | ------------ | ------------------------------------------------ | ---------- | ------------------ |
| GET        | `/tips`      | Retrieve a list of all travel tips/stories.      | User       | 🔴 Not Implemented |
| GET        | `/tip/{id}`  | Retrieve detailed information about a tip/story. | User       | 🔴 Not Implemented |
| POST       | `/tip`       | Post a new tip or story.                         | User       | 🔴 Not Implemented |
| PUT        | `/tip/{id}`  | Update an existing tip/story.                    | User       | 🔴 Not Implemented |
| DELETE     | `/tip/{id}`  | Delete a tip/story.                              | User/Admin | 🔴 Not Implemented |

---

## User Management Endpoints

| **Method** | **Endpoint**    | **Description**                            | **Access** | **Status**         |
| ---------- | --------------- | ------------------------------------------ | ---------- | ------------------ |
| GET        | `/user/profile` | Retrieve the authenticated user’s profile. | User       | 🔴 Not Implemented |
| PUT        | `/user/profile` | Update the user’s profile.                 | User       | 🔴 Not Implemented |

---

## Authentication & Authorization Endpoints

This functionality is handled by the API Gateway. Refer to the [API Gateway Documentation](https://github.com/caiomarinello/ninjaGateway) for details.

---
