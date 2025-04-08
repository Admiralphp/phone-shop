// src/app/core/services/analytics.service.ts
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { environment } from '../../../environments/environment';

export interface SalesSummary {
  totalSales: number;
  dailySales: number;
  weeklySales: number;
  monthlySales: number;
  topProducts: Array<{
    productId: string;
    productName: string;
    quantity: number;
    revenue: number;
  }>;
}

export interface InventorySummary {
  totalProducts: number;
  lowStockProducts: number;
  outOfStockProducts: number;
  productsNeedingRestock: Array<{
    productId: string;
    productName: string;
    currentStock: number;
  }>;
}

export interface UserActivitySummary {
  totalUsers: number;
  newUsers: number;
  activeUsers: number;
  userGrowthRate: number;
}

@Injectable({
  providedIn: 'root'
})
export class AnalyticsService {
  private readonly API_URL = `${environment.apiUrl}/analytics`;

  constructor(private http: HttpClient) {}

  getSalesSummary(period: 'day' | 'week' | 'month' | 'year' = 'month'): Observable<SalesSummary> {
    return this.http.get<SalesSummary>(`${this.API_URL}/sales?period=${period}`);
  }

  getInventorySummary(): Observable<InventorySummary> {
    return this.http.get<InventorySummary>(`${this.API_URL}/inventory`);
  }

  getUserActivitySummary(period: 'day' | 'week' | 'month' | 'year' = 'month'): Observable<UserActivitySummary> {
    return this.http.get<UserActivitySummary>(`${this.API_URL}/user-activity?period=${period}`);
  }
}