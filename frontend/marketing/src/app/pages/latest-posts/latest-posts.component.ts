import { Component, OnInit, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { CardComponent } from '../../components/card/card.component';
import { InfiniteScrollComponent } from '../../components/infinite-scroll/infinite-scroll.component';
import { LatestPostsStore } from './latest-posts.store';

@Component({
  selector: 'app-latest-posts',
  standalone: true,
  imports: [CommonModule, CardComponent, InfiniteScrollComponent],
  templateUrl: './latest-posts.component.html',
})
export class LatestPostsComponent implements OnInit {
  protected readonly store = inject(LatestPostsStore);

  ngOnInit() {
    this.store.loadMore();
  }

  onScroll() {
    this.store.loadMore();
  }
}
