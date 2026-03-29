import { CommonModule } from '@angular/common';
import { AfterViewInit, Component, ElementRef, OnDestroy, OnInit, ViewChild, effect, input, output, TemplateRef } from '@angular/core';

@Component({
  selector: 'app-infinite-scroll',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './infinite-scroll.component.html',
  styleUrls: ['./infinite-scroll.component.css'],
})
export class InfiniteScrollComponent implements OnInit, AfterViewInit, OnDestroy {
  itemTemplate = input.required<TemplateRef<any>>();
  items = input<any[]>([]);
  loading = input<boolean>(false);
  allLoaded = input<boolean>(false);
  error = input<string | null>(null);
  scrolled = output<void>();

  @ViewChild('sentinel') sentinel!: ElementRef<HTMLElement>;
  private observer?: IntersectionObserver;
  private isIntersecting = false;
  
  private retryCount = 0;
  private maxRetries = 5;
  private retryTimeoutId?: any;

  constructor() {
    effect(() => {
      const isLoading = this.loading();
      const isAllLoaded = this.allLoaded();
      const hasError = !!this.error();

      if (isLoading || isAllLoaded) {
        if (this.retryTimeoutId) {
          clearTimeout(this.retryTimeoutId);
          this.retryTimeoutId = undefined;
        }
        return;
      }

      if (this.isIntersecting) {
        if (hasError) {
          if (this.retryCount < this.maxRetries) {
            const backoffTime = Math.pow(2, this.retryCount) * 1000;
            this.retryCount++;
            
            this.retryTimeoutId = setTimeout(() => {
              this.scrolled.emit();
            }, backoffTime);
          }
        } else {
          // Success or initial state, reset retries
          this.retryCount = 0;
          this.scrolled.emit();
        }
      }
    });
  }

  ngOnInit() {
    this.observer = new IntersectionObserver((entries) => {
      this.isIntersecting = entries[0].isIntersecting;
      
      if (!this.isIntersecting) {
        if (this.retryTimeoutId) {
          clearTimeout(this.retryTimeoutId);
          this.retryTimeoutId = undefined;
        }
      } else if (!this.loading() && !this.allLoaded()) {
        // We just scrolled into view and we aren't loading, emit immediately.
        // We reset the retry count assuming a new manual scroll attempt represents user intent
        this.retryCount = 0;
        this.scrolled.emit();
      }
    });
  }

  ngAfterViewInit() {
    if (this.sentinel) {
      this.observer?.observe(this.sentinel.nativeElement);
    }
  }

  ngOnDestroy() {
    if (this.retryTimeoutId) {
      clearTimeout(this.retryTimeoutId);
    }
    this.observer?.disconnect();
  }
}
