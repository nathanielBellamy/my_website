import { Component, OnInit, inject } from '@angular/core';
import { AboutStore } from './about.store';
import { CommonModule } from '@angular/common';
import { InfiniteScrollComponent } from '../../components/infinite-scroll/infinite-scroll.component';
import { CardComponent } from '../../components/card/card.component';

@Component({
  selector: 'app-about',
  standalone: true,
  imports: [CommonModule, InfiniteScrollComponent, CardComponent],
  templateUrl: './about.component.html',
})
export class AboutComponent implements OnInit {
  protected readonly store = inject(AboutStore);

  ngOnInit() {
    this.store.loadMore();
  }

  onScroll() {
    this.store.loadMore();
  }
}
