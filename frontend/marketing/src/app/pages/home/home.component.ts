import { Component, OnInit, inject } from '@angular/core';
import { CardComponent } from '../../components/card/card.component';
import { HomeStore } from './home.store';
import { CommonModule } from '@angular/common';
import { InfiniteScrollComponent } from '../../components/infinite-scroll/infinite-scroll.component';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [CardComponent, CommonModule, InfiniteScrollComponent],
  templateUrl: './home.component.html',
})
export class HomeComponent implements OnInit {
  protected readonly store = inject(HomeStore);

  ngOnInit() {
    this.store.loadMore();
  }

  onScroll() {
    this.store.loadMore();
  }
}
