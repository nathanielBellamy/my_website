import { Component, OnInit, inject } from '@angular/core';
import { BlogStore } from './blog.store';
import { CommonModule } from '@angular/common';
import { InfiniteScrollComponent } from '../../components/infinite-scroll/infinite-scroll.component';
import { CardComponent } from '../../components/card/card.component';

@Component({
  selector: 'app-blog',
  standalone: true,
  imports: [CommonModule, InfiniteScrollComponent, CardComponent],
  templateUrl: './blog.component.html',
})
export class BlogComponent implements OnInit {
  protected readonly store = inject(BlogStore);

  ngOnInit() {
    this.store.loadMore();
  }

  onScroll() {
    this.store.loadMore();
  }
}
