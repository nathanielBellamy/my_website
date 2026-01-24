import { CommonModule } from '@angular/common';
import { Component, EventEmitter, Input, Output, TemplateRef } from '@angular/core';

@Component({
  selector: 'app-infinite-scroll',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './infinite-scroll.component.html',
  styleUrls: ['./infinite-scroll.component.css'],
})
export class InfiniteScrollComponent {
  @Input() itemTemplate!: TemplateRef<any>;
  @Input() items: any[] = [];
  @Input() loading = false;
  @Input() allLoaded = false;
  @Input() error: string | null = null;
  @Output() scrolled = new EventEmitter<void>();

  onScroll(event: Event) {
    const element = event.target as HTMLElement;
    if (element.scrollHeight - element.scrollTop <= element.clientHeight + 100) {
      if (!this.loading && !this.allLoaded) {
        this.scrolled.emit();
      }
    }
  }
}
