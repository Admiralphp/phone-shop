// src/app/modules/products/product-form/product-form.component.ts
import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { ProductService } from '../../../core/services/product.service';
import { Product } from '../../../core/models/product.model';

@Component({
  selector: 'app-product-form',
  templateUrl: './product-form.component.html',
  styleUrls: ['./product-form.component.scss']
})
export class ProductFormComponent implements OnInit {
  productForm: FormGroup;
  isEdit = false;
  productId: string | null = null;
  isLoading = false;
  error = '';

  constructor(
    private fb: FormBuilder,
    private route: ActivatedRoute,
    private router: Router,
    private productService: ProductService
  ) {
    this.productForm = this.fb.group({
      name: ['', [Validators.required]],
      description: ['', [Validators.required]],
      price: [0, [Validators.required, Validators.min(0)]],
      category: ['', [Validators.required]],
      stockQuantity: [0, [Validators.required, Validators.min(0)]],
      imageUrl: ['', [Validators.required]]
    });
  }

  ngOnInit(): void {
    this.productId = this.route.snapshot.paramMap.get('id');
    
    if (this.productId) {
      this.isEdit = true;
      this.loadProduct(this.productId);
    }
  }

  loadProduct(id: string): void {
    this.isLoading = true;
    this.productService.getProduct(id).subscribe({
      next: (product) => {
        this.productForm.patchValue({
          name: product.name,
          description: product.description,
          price: product.price,
          category: product.category,
          stockQuantity: product.stockQuantity,
          imageUrl: product.imageUrl
        });
        this.isLoading = false;
      },
      error: (err) => {
        this.error = 'Failed to load product';
        this.isLoading = false;
      }
    });
  }

  onSubmit(): void {
    if (this.productForm.invalid) {
      return;
    }

    this.isLoading = true;
    
    if (this.isEdit && this.productId) {
      this.productService.updateProduct(this.productId, this.productForm.value).subscribe({
        next: () => {
          this.router.navigate(['/products']);
        },
        error: (err) => {
          this.error = 'Failed to update product';
          this.isLoading = false;
        }
      });
    } else {
      this.productService.createProduct(this.productForm.value).subscribe({
        next: () => {
          this.router.navigate(['/products']);
        },
        error: (err) => {
          this.error = 'Failed to create product';
          this.isLoading = false;
        }
      });
    }
  }
}
