// src/app/core/models/product.model.ts
export interface Product {
  id: string;
  name: string;
  description: string;
  price: number;
  category: string;
  stockQuantity: number;
  imageUrl: string;
  createdAt: Date;
  updatedAt: Date;
}
