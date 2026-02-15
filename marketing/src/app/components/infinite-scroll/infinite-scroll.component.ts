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

  constructor() {
    effect(() => {
      if (!this.loading() && !this.allLoaded() && this.isIntersecting) {
        this.scrolled.emit();
      }
    });
  }

  ngOnInit() {
    this.observer = new IntersectionObserver((entries) => {
      this.isIntersecting = entries[0].isIntersecting;
      if (this.isIntersecting && !this.loading() && !this.allLoaded()) {
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
    this.observer?.disconnect();
  }
}
