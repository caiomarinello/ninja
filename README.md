# Ninja tech challenge 2025

This API provides backend support for a marketplace focused on travelers and backpackers, including user management, product catalog, orders, and more.

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

| **Method** | **Endpoint**     | **Description**                                | **Access** | **Status**         |
| ---------- | ---------------- | ---------------------------------------------- | ---------- | ------------------ |
| GET        | `/products`      | Retrieve a list of all products.               | User       | 🔴 Not Implemented |
| GET        | `/products/{id}` | Retrieve detailed information about a product. | User       | 🔴 Not Implemented |
| POST       | `/products`      | Add a new product.                             | Admin      | 🔴 Not Implemented |
| PUT        | `/products/{id}` | Update an existing product.                    | Admin      | 🔴 Not Implemented |
| DELETE     | `/products/{id}` | Delete a product.                              | Admin      | 🔴 Not Implemented |

---

## Order Management Endpoints

| **Method** | **Endpoint**          | **Description**                             | **Access** | **Status**         |
| ---------- | --------------------- | ------------------------------------------- | ---------- | ------------------ |
| POST       | `/checkout`           | Finalize the cart and create an order.      | User       | 🔴 Not Implemented |
| GET        | `/orders/{id}`        | Retrieve a specific order’s details.        | User       | 🔴 Not Implemented |
| GET        | `/orders`             | List all orders for the authenticated user. | User       | 🔴 Not Implemented |
| PUT        | `/orders/{id}/status` | Update the status of an order.              | Admin      | 🔴 Not Implemented |

---

## Review Management Endpoints

| **Method** | **Endpoint**            | **Description**                          | **Access** | **Status**         |
| ---------- | ----------------------- | ---------------------------------------- | ---------- | ------------------ |
| POST       | `/reviews`              | Create a new review for a product.       | User       | 🔴 Not Implemented |
| GET        | `/reviews/{product_id}` | Retrieve reviews for a specific product. | User       | 🔴 Not Implemented |
| PUT        | `/reviews/{id}`         | Update an existing review.               | User       | 🔴 Not Implemented |
| DELETE     | `/reviews/{id}`         | Delete a review.                         | User/Admin | 🔴 Not Implemented |

---

## Tip/Story Exchange Endpoints

| **Method** | **Endpoint** | **Description**                                  | **Access** | **Status**         |
| ---------- | ------------ | ------------------------------------------------ | ---------- | ------------------ |
| GET        | `/tips`      | Retrieve a list of all travel tips/stories.      | User       | 🔴 Not Implemented |
| GET        | `/tips/{id}` | Retrieve detailed information about a tip/story. | User       | 🔴 Not Implemented |
| POST       | `/tips`      | Post a new tip or story.                         | User       | 🔴 Not Implemented |
| PUT        | `/tips/{id}` | Update an existing tip/story.                    | User       | 🔴 Not Implemented |
| DELETE     | `/tips/{id}` | Delete a tip/story.                              | User/Admin | 🔴 Not Implemented |

---

## Authentication & Authorization

| **Method** | **Endpoint**            | **Description**                                 | **Access** | **Status**         |
| ---------- | ----------------------- | ----------------------------------------------- | ---------- | ------------------ |
| POST       | `/auth/register`        | Register a new user.                            | User       | 🔴 Not Implemented |
| POST       | `/auth/login`           | Authenticate a user and return a token.         | User       | 🔴 Not Implemented |
| POST       | `/auth/logout`          | Log out the authenticated user.                 | User       | 🔴 Not Implemented |
| POST       | `/auth/forgot-password` | Send a password reset link to the user’s email. | Public     | 🔴 Not Implemented |

---

## User Management

| **Method** | **Endpoint**     | **Description**                            | **Access** | **Status**         |
| ---------- | ---------------- | ------------------------------------------ | ---------- | ------------------ |
| GET        | `/users/profile` | Retrieve the authenticated user’s profile. | User       | 🔴 Not Implemented |
| PUT        | `/users/profile` | Update the user’s profile.                 | User       | 🔴 Not Implemented |
