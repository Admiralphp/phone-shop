// src/app/modules/dashboard/dashboard.component.ts
import { Component, OnInit } from '@angular/core';
import { AnalyticsService, SalesSummary, InventorySummary, UserActivitySummary } from '../../core/services/analytics.service';

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.scss']
})
export class DashboardComponent implements OnInit {
  salesSummary: SalesSummary | null = null;
  inventorySummary: InventorySummary | null = null;
  userActivity: UserActivitySummary | null = null;
  isLoading = true;
  error = '';

  constructor(private analyticsService: AnalyticsService) {}

  ngOnInit(): void {
    this.loadDashboardData();
  }

  loadDashboardData(): void {
    this.isLoading = true;
    
    // Use forkJoin or similar to combine these requests if desired
    this.analyticsService.getSalesSummary().subscribe({
      next: (data) => {
        this.salesSummary = data;
        this.checkLoadingComplete();
      },
      error: (err) => {
        this.error = 'Failed to load sales data';
        this.isLoading = false;
      }
    });

    this.analyticsService.getInventorySummary().subscribe({
      next: (data) => {
        this.inventorySummary = data;
        this.checkLoadingComplete();
      },
      error: (err) => {
        this.error = 'Failed to load inventory data';
        this.isLoading = false;
      }
    });

    this.analyticsService.getUserActivitySummary().subscribe({
      next: (data) => {
        this.userActivity = data;
        this.checkLoadingComplete();
      },
      error: (err) => {
        this.error = 'Failed to load user activity data';
        this.isLoading = false;
      }
    });
  }

  private checkLoadingComplete(): void {
    if (this.salesSummary && this.inventorySummary && this.userActivity) {
      this.isLoading = false;
    }
  }
}