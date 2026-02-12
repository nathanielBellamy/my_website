import { CommonModule } from '@angular/common';
import { Component, input, output, TemplateRef } from '@angular/core';

@Component({
  selector: 'app-infinite-scroll',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './infinite-scroll.component.html',
  styleUrls: ['./infinite-scroll.component.css'],
})
export class InfiniteScrollComponent {
  itemTemplate = input.required<TemplateRef<any>>();
  items = input<any[]>([]);
  loading = input<boolean>(false);
  allLoaded = input<boolean>(false);
  error = input<string | null>(null);
  scrolled = output<void>();

  onScroll(event: Event) {
    const element = event.target as HTMLElement;
    if (element.scrollHeight - element.scrollTop <= element.clientHeight + 100) {
      if (!this.loading() && !this.allLoaded()) {
        this.scrolled.emit();
      }
    }
  }
}
