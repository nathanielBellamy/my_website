import { Component, OnInit, inject } from '@angular/core';
import { GrooveJrStore } from './groove-jr.store';
import { CommonModule } from '@angular/common';
import { InfiniteScrollComponent } from '../../components/infinite-scroll/infinite-scroll.component';
import { CardComponent } from '../../components/card/card.component';
import { ScrollFadeInDirective } from '../../directives/scroll-fade-in.directive';

@Component({
  selector: 'app-groove-jr',
  standalone: true,
  imports: [CommonModule, InfiniteScrollComponent, CardComponent, ScrollFadeInDirective],
  templateUrl: './groove-jr.component.html',
})
export class GrooveJrComponent implements OnInit {
  protected readonly store = inject(GrooveJrStore);

  ngOnInit() {
    this.store.loadMore();
  }

  onScroll() {
    this.store.loadMore();
  }
}
